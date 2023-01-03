package miniblog

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewMinBlogCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "miniblog",
		Short:        "a good go practical project",
		Long:         "a good go practical project, used to create user with basic information.",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any argumnets, got %q", cmd.CommandPath(), args)
				}
			}
			return nil
		},
	}
	return cmd
}

func run() error {
	fmt.Println("Hello MiniBlog!!")
	return nil
}
