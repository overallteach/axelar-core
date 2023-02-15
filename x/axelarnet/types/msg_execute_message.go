package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/axelarnetwork/axelar-core/utils"
)

// NewExecuteMessage creates a message of type ExecuteMessageRequest
func NewExecuteMessage(sender sdk.AccAddress, id string, payload []byte) *ExecuteMessageRequest {
	return &ExecuteMessageRequest{
		Sender:  sender,
		ID:      id,
		Payload: payload,
	}
}

// Route returns the route for this message
func (m ExecuteMessageRequest) Route() string {
	return RouterKey
}

// Type returns the type of the message
func (m ExecuteMessageRequest) Type() string {
	return "ExecuteMessage"
}

// ValidateBasic executes a stateless message validation
func (m ExecuteMessageRequest) ValidateBasic() error {
	if err := sdk.VerifyAddressFormat(m.Sender); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, sdkerrors.Wrap(err, "sender").Error())
	}

	if err := utils.ValidateString(m.ID); err != nil {
		return sdkerrors.Wrap(err, "invalid general message id")
	}

	if len(m.Payload) == 0 {
		return fmt.Errorf("invalid payload")
	}

	return nil
}

// GetSignBytes returns the message bytes that need to be signed
func (m ExecuteMessageRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&m)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the set of signers for this message
func (m ExecuteMessageRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Sender}
}