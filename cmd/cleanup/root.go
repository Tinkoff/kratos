package cleanup

import (
	"github.com/spf13/cobra"
)

func NewCleanupCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "cleanup",
		Short: "Various cleanup helpers",
	}
}

func RegisterCommandRecursive(parent *cobra.Command) {
	c := NewCleanupCmd()
	parent.AddCommand(c)
	c.AddCommand(NewCleanupSQLCmd())
}
