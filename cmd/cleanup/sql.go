package cleanup

import (
	"github.com/spf13/cobra"

	"github.com/ory/kratos/cmd/cliclient"
	"github.com/ory/x/configx"
)

func NewCleanupSQLCmd() *cobra.Command {
	c := &cobra.Command{
		Use:   "sql <database-url>",
		Short: "Cleanup sql database from expired flows and sessions",
		Long: `Run this command as frequently as you need.
It is recommended to run this command close to the SQL instance (e.g. same subnet) instead of over the public internet.
This decreases risk of failure and decreases time required.
You can read in the database URL using the -e flag, for example:
	export DSN=...
	kratos cleanup sql -e
### WARNING ###
Before running this command on an existing database, create a back up!
`,
		Run: func(cmd *cobra.Command, args []string) {
			cliclient.NewCleanupHandler().CleanupSQL(cmd, args)
		},
	}

	configx.RegisterFlags(c.PersistentFlags())
	c.Flags().BoolP("read-from-env", "e", false, "If set, reads the database connection string from the environment variable DSN or config file key dsn.")
	c.Flags().BoolP("yes", "y", false, "If set all confirmation requests are accepted without user interaction.")
	return c
}
