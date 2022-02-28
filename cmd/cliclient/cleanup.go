package cliclient

import (
	"fmt"
	"os"

	"github.com/ory/kratos/driver"
	"github.com/ory/x/cmdx"
	"github.com/ory/x/configx"
	"github.com/spf13/cobra"
)

type CleanupHandler struct{}

func NewCleanupHandler() *CleanupHandler {
	return &CleanupHandler{}
}

func (h *CleanupHandler) CleanupSQL(cmd *cobra.Command, args []string) {
	var d driver.Registry

	d = driver.NewWithoutInit(
		cmd.Context(),
		cmd.ErrOrStderr(),
		configx.WithFlags(cmd.Flags()),
		configx.SkipValidation())
	if len(d.Config(cmd.Context()).DSN()) == 0 {
		fmt.Println(cmd.UsageString())
		fmt.Println("")
		fmt.Println("Environment variable DSN must be set")
		os.Exit(1)
		return
	}

	err := d.Init(cmd.Context(), driver.SkipNetworkInit)
	cmdx.Must(err, "An error occurred initializing cleanup: %s", err)

	err = d.Persister().CleanupDatabase(cmd.Context())
	cmdx.Must(err, "An error occurred while cleanup expired data: %s", err)
	fmt.Println("Successfully cleanup!")
}
