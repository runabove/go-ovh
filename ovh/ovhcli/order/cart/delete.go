package cart

import (
	"github.com/runabove/go-sdk/ovh"
	"github.com/runabove/go-sdk/ovh/ovhcli/common"

	"github.com/spf13/cobra"
)

var cmdCartDelete = &cobra.Command{
	Use:   "delete <cartID>",
	Short: "Delete cart",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			common.WrongUsage(cmd)
		}
		cartID := args[0]

		client, err := ovh.NewDefaultClient()
		common.Check(err)

		err = client.OrderDeleteCart(cartID)
		common.Check(err)
	},
}
