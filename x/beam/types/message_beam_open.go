package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgOpenBeam{}

// NewMsgOpenBeam Build a open beam message based on parameters
func NewMsgOpenBeam(creator string, amount int64, secret string) *MsgOpenBeam {
	return &MsgOpenBeam{
		Creator: creator,
		Amount:  amount,
		Secret:  secret,
	}
}

// Route dunno
func (msg MsgOpenBeam) Route() string {
	return RouterKey
}

// Type Return the message type
func (msg MsgOpenBeam) Type() string {
	return "OpenBeam"
}

// GetSigners Return the list of signers for the given message
func (msg *MsgOpenBeam) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

// GetSignBytes Return the generated bytes from the signature
func (msg *MsgOpenBeam) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic Validate the message payload before dispatching to the local kv store
func (msg *MsgOpenBeam) ValidateBasic() error {
	// Ensure the address is correct and that we are able to acquire it
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid creator address (%s)", err)
	}

	// Validate the amount
	if msg.Amount <= 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "Invalid amount: must be greater than 0")
	}

	// Validate the secret
	// TODO: implement
	return nil
}
