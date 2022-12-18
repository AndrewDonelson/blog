package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateImage = "create_image"

var _ sdk.Msg = &MsgCreateImage{}

func NewMsgCreateImage(creator string, ipfsurl string, tags string, views string, likes string, dislikes string) *MsgCreateImage {
	return &MsgCreateImage{
		Creator:  creator,
		Ipfsurl:  ipfsurl,
		Tags:     tags,
		Views:    views,
		Likes:    likes,
		Dislikes: dislikes,
	}
}

func (msg *MsgCreateImage) Route() string {
	return RouterKey
}

func (msg *MsgCreateImage) Type() string {
	return TypeMsgCreateImage
}

func (msg *MsgCreateImage) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateImage) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateImage) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
