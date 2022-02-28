package cleanup

import (
	"time"

	"github.com/spf13/cobra"

	"github.com/ory/kratos/cmd/cliclient"
	"github.com/ory/x/configx"
)

func NewCleanupSQLCmd() *cobra.Command {
	c := &cobra.Command{
		Use:   "cleanup",
		Short: "Cleanup sql database from expired flows and sessions",
		Long: `Run this command as frequently as you need.
It is recommended to run this command close to the SQL instance (e.g. same subnet) instead of over the public internet.
This decreases risk of failure and decreases time required.
You can read in the database URL, for example:
	export DSN=...
	kratos cleanup
### optional params ###
	--limit
	--batch-size
    --keep-if-younger
	--cleanup-sessions
	--cleanup-continuity-containers
	--cleanup-login-flows
	--cleanup-recovery-flows
	--cleanup-registration-flows
	--cleanup-settings-flows
	--cleanup-verification-flows
### WARNING ###
Before running this command on an existing database, create a back up!
`,
		Run: func(cmd *cobra.Command, args []string) {
			cliclient.NewCleanupHandler().CleanupSQL(cmd, args)
		},
	}
	configx.RegisterFlags(c.PersistentFlags())

	c.Flags().IntP("limit", "l", 1000, "Define how many records are deleted. (default 1000)")
	c.Flags().IntP("batch-size", "b", 100, "Define how many records are deleted with each iteration. (default 100)")
	c.Flags().DurationP("keep-if-younger", "k", 2160*time.Hour, "Keep database records that are younger than a specified duration e.g. 3000h.")
	c.Flags().BoolP("cleanup-sessions", "s", false, "If set then cleaning up expired sessions")
	c.Flags().BoolP("cleanup-continuity-containers", "c", false, "If set then cleaning up expired continuity containers")
	c.Flags().BoolP("cleanup-login-flows", "i", false, "If set then cleaning up expired login flows")
	c.Flags().BoolP("cleanup-recovery-flows", "r", false, "If set then cleaning up expired recovery flows")
	c.Flags().BoolP("cleanup-registration-flows", "g", false, "If set then cleaning up expired registation flows")
	c.Flags().BoolP("cleanup-settings-flows", "t", false, "If set then cleaning up expired settings flows")
	c.Flags().BoolP("cleanup-verification-flows", "v", false, "If set then cleaning up expired verification flows")

	return c
}
