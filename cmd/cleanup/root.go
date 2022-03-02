package cleanup

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ory/kratos/driver"
	"github.com/ory/kratos/driver/config"
	"github.com/ory/x/cmdx"
	"github.com/ory/x/configx"
	"github.com/spf13/cobra"
)

func NewCleanupCmd() *cobra.Command {
	c := &cobra.Command{
		Use:   "cleanup",
		Short: "Various cleanup helpers",
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
		Args: cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			dsn := os.Getenv("DSN")
			if !strings.HasPrefix(dsn, "postgres://") {
				fmt.Println("cleanup: DSN currently support postgres only")
				os.Exit(1)
			}
			d := driver.New(
				cmd.Context(),
				cmd.ErrOrStderr(),
				configx.WithFlags(cmd.Flags()),
				configx.WithValues(map[string]interface{}{
					config.ViperKeyDSN: dsn,
				}),
				configx.SkipValidation(),
			)
			err := d.Persister().CleanupDatabase(cmd.Context())
			cmdx.Must(err, "An error occurred while cleanup expired data: %s", err)
			fmt.Println("Successfully cleanup!")
			os.Exit(0)
		},
	}

	c.Flags().Int("limit", 1000, "Define how many records are deleted. (default 1000)")
	c.Flags().Int("batch-size", 100, "Define how many records are deleted with each iteration. (default 100)")
	c.Flags().Duration("keep-if-younger", 2160*time.Hour, "Keep database records that are younger than a specified duration e.g. 3000h.")
	c.Flags().Bool("cleanup-sessions", false, "If set then cleaning up expired sessions")
	c.Flags().Bool("cleanup-continuity-containers", false, "If set then cleaning up expired continuity containers")
	c.Flags().Bool("cleanup-login-flows", false, "If set then cleaning up expired login flows")
	c.Flags().Bool("cleanup-recovery-flows", false, "If set then cleaning up expired recovery flows")
	c.Flags().Bool("cleanup-registration-flows", false, "If set then cleaning up expired registation flows")
	c.Flags().Bool("cleanup-settings-flows", false, "If set then cleaning up expired settings flows")
	c.Flags().Bool("cleanup-verification-flows", false, "If set then cleaning up expired verification flows")

	return c
}

func RegisterCommandRecursive(parent *cobra.Command) {
	parent.AddCommand(NewCleanupCmd())
}
