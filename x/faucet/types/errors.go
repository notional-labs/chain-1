package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrMintTooOften = sdkerrors.Register(ModuleName, 1100, "Please wait 24 hours between each claim")
)
