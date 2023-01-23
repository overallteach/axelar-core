package axelarnet

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	porttypes "github.com/cosmos/ibc-go/v3/modules/core/05-port/types"
	ibcexported "github.com/cosmos/ibc-go/v3/modules/core/exported"

	"github.com/axelarnetwork/axelar-core/x/axelarnet/keeper"
	"github.com/axelarnetwork/axelar-core/x/axelarnet/types"
	nexus "github.com/axelarnetwork/axelar-core/x/nexus/exported"
	"github.com/axelarnetwork/utils/funcs"
)

// RateLimiter implements an ICS4Wrapper middleware to rate limit IBC transfers
type RateLimiter struct {
	keeper  keeper.Keeper
	channel porttypes.ICS4Wrapper
	nexus   types.Nexus
}

// NewRateLimiter returns a new RateLimiter
func NewRateLimiter(keeper keeper.Keeper, channel porttypes.ICS4Wrapper, nexus types.Nexus) RateLimiter {
	return RateLimiter{
		keeper:  keeper,
		channel: channel,
		nexus:   nexus,
	}
}

// RateLimitPacket implements rate limiting of IBC packets
// - If the IBC channel that the packet is sent on is a registered chain, check the activation status.
// - If the packet is an ICS-20 coin transfer, apply rate limiting on (chain, base denom) pair.
// - If the rate limit is exceeded, an error is returned.
// Incoming direction is used for tokens incoming to Axelar (unlocked from IBC escrow/minted as an IBC denom).
// Outgoing direction is used for tokens going out from Axelar (locked in the IBC escrow/burned as an IBC denom).
func (r RateLimiter) RateLimitPacket(ctx sdk.Context, packet ibcexported.PacketI, direction nexus.TransferDirection, ibcPath string) error {
	chainName, ok := r.keeper.GetChainNameByIBCPath(ctx, ibcPath)
	if !ok {
		// If the IBC channel is not registered as a chain, skip rate limiting
		return nil
	}
	chain := funcs.MustOk(r.nexus.GetChain(ctx, chainName))

	if !r.nexus.IsChainActivated(ctx, chain) {
		return fmt.Errorf("chain %s registered for IBC path %s is deactivated", chain.Name, ibcPath)
	}

	token, err := parseTokenFromPacket(packet)
	if err != nil {
		return err
	}

	if err := r.nexus.RateLimitTransfer(ctx, chain.Name, token, direction); err != nil {
		return err
	}

	return nil
}

// SendPacket implements the ICS4 Wrapper interface
func (r RateLimiter) SendPacket(ctx sdk.Context, chanCap *capabilitytypes.Capability, packet ibcexported.PacketI) error {
	if err := r.channel.SendPacket(ctx, chanCap, packet); err != nil {
		return err
	}

	// Cross-chain transfers using IBC have already been tracked by EnqueueTransfer, so skip those
	if _, found := r.keeper.GetSeqIDMapping(ctx, packet.GetSourcePort(), packet.GetSourceChannel(), packet.GetSequence()); found {
		return nil
	}

	return r.RateLimitPacket(ctx, packet, nexus.Outgoing, types.NewIBCPath(packet.GetSourcePort(), packet.GetSourceChannel()))
}

// WriteAcknowledgement implements the ICS4 Wrapper interface
func (r RateLimiter) WriteAcknowledgement(ctx sdk.Context, chanCap *capabilitytypes.Capability, packet ibcexported.PacketI, ack ibcexported.Acknowledgement) error {
	return r.channel.WriteAcknowledgement(ctx, chanCap, packet, ack)
}

func parseTokenFromPacket(packet ibcexported.PacketI) (sdk.Coin, error) {
	// Only ICS-20 packets are expected in the middleware
	var data ibctransfertypes.FungibleTokenPacketData
	if err := ibctransfertypes.ModuleCdc.UnmarshalJSON(packet.GetData(), &data); err != nil {
		return sdk.Coin{}, sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "cannot unmarshal ICS-20 transfer packet data: %s", err.Error())
	}

	asset := data.Denom
	// If the asset being transferred is an IBC denom originating on the destination chain,
	// then the full denom in the IBC transfer contains the IBC channel to the destination chain as a prefix, `transfer/channel-x/asset`.
	// e.g. For IBC channel Axelar channel-0 <-> channel-1 Cosmoshub,
	// For an asset `uusdc` originating on Axelar, the full denom when sending it from Cosmoshub -> Axelar will be `transfer/channel-1/uusdc`
	// Similarly, for an asset `uatom` originating on Cosmoshub, the full denom when sending it from Axelar -> Cosmoshub will be `transfer/channel-0/uatom`
	// So, if the source channel of the packet is a prefix of the denom being transferred, we remove it to use the remaining denom as the asset for the rate limit.
	if ibctransfertypes.ReceiverChainIsSource(packet.GetSourcePort(), packet.GetSourceChannel(), data.Denom) {
		ibcTransferPrefix := ibctransfertypes.GetDenomPrefix(packet.GetSourcePort(), packet.GetSourceChannel())
		asset = data.Denom[len(ibcTransferPrefix):]
	}

	// parse the transfer amount
	transferAmount, ok := sdk.NewIntFromString(data.Amount)
	if !ok {
		return sdk.Coin{}, sdkerrors.Wrapf(ibctransfertypes.ErrInvalidAmount, "unable to parse transfer amount (%s) into sdk.Int", data.Amount)
	}
	token := sdk.Coin{Denom: asset, Amount: transferAmount}
	if err := token.Validate(); err != nil {
		return sdk.Coin{}, err
	}

	return token, nil
}