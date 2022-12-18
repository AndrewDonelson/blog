package keeper

import (
	"context"

	"blog/x/blog/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateImage(goCtx context.Context, msg *types.MsgCreateImage) (*types.MsgCreateImageResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Create variable of type Image
	var image = types.Image{
		Creator:  msg.Creator,
		Ipfsurl:  msg.Ipfsurl,
		Tags:     msg.Tags,
		Views:    msg.Views,
		Likes:    msg.Likes,
		Dislikes: msg.Dislikes,
	}

	// Add a post to the store and get back the ID
	id := k.AppendImage(ctx, image)

	return &types.MsgCreateImageResponse{Id: id}, nil
}
