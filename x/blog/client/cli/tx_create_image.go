package cli

import (
	"strconv"

	"blog/x/blog/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCreateImage() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-image [ipfsurl] [tags] [views] [likes] [dislikes]",
		Short: "Broadcast message createImage",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argIpfsurl := args[0]
			argTags := args[1]
			argViews := args[2]
			argLikes := args[3]
			argDislikes := args[4]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateImage(
				clientCtx.GetFromAddress().String(),
				argIpfsurl,
				argTags,
				argViews,
				argLikes,
				argDislikes,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
