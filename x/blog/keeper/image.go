package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"blog/x/blog/types"
)

func (k Keeper) GetImageCount(ctx sdk.Context) uint64 {
	// Get the store using storeKey (which is "blog") and ImageCountKey (which is "Image/count/")
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.ImageCountKey))

	// Convert the ImageCountKey to bytes
	byteKey := []byte(types.ImageCountKey)

	// Get the value of the count
	bz := store.Get(byteKey)

	// Return zero if the count value is not found (for example, it's the first post)
	if bz == nil {
		return 0
	}

	// Convert the count into a uint64
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetImageCount(ctx sdk.Context, count uint64) {
	// Get the store using storeKey (which is "blog") and ImageCountKey (which is "Image/count/")
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.ImageCountKey))

	// Convert the ImageCountKey to bytes
	byteKey := []byte(types.ImageCountKey)

	// Convert count from uint64 to string and get bytes
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)

	// Set the value of Image/count/ to count
	store.Set(byteKey, bz)
}

func (k Keeper) AppendImage(ctx sdk.Context, post types.Image) uint64 {
	// Get the current number of posts in the store
	count := k.GetImageCount(ctx)

	// Assign an ID to the post based on the number of posts in the store
	post.Id = count

	// Get the store
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.ImageKey))

	// Convert the post ID into bytes
	byteKey := make([]byte, 8)
	binary.BigEndian.PutUint64(byteKey, post.Id)

	// Marshal the post into bytes
	appendedValue := k.cdc.MustMarshal(&post)

	// Insert the post bytes using post ID as a key
	store.Set(byteKey, appendedValue)

	// Update the post count
	k.SetImageCount(ctx, count+1)
	return count
}
