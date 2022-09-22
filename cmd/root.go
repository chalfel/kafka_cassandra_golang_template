package cmd

import (
	"context"

	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	ctx := context.Background()

	rootCmd := &cobra.Command{
		Use:   "app",
		Short: "go boilerplate",
	}

	s := NewServerCmd(ctx)

	rootCmd.AddCommand(s)

	return rootCmd
}
