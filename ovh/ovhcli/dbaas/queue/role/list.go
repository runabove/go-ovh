package role

import (
	"github.com/ovhlabs/go-sdk/ovh"
	"github.com/ovhlabs/go-sdk/ovh/ovhcli/common"

	"github.com/spf13/cobra"
)

var withDetails bool

func init() {
	cmdList.PersistentFlags().BoolVarP(&withDetails, "withDetails", "", false, "Display details")
}

var cmdList = &cobra.Command{
	Use:   "list",
	Short: "List all roles on a service: ovhcli dbaas queue role list (--name=AppName | <--id=appID>)",
	Run: func(cmd *cobra.Command, args []string) {

		client, err := ovh.NewDefaultClient()
		common.Check(err)

		if name != "" {
			app, errInfo := client.DBaasQueueAppInfoByName(name)
			common.Check(errInfo)
			id = app.ID
		}

		apps, err := client.DBaasQueueRoleList(id, withDetails)
		common.Check(err)

		common.FormatOutputDef(apps)
	},
}
