package cmd

import (
	"context"
	"fmt"

	"github.com/orens/linkedon/linkedon"
	"github.com/spf13/cobra"
)

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset the social network",
	Long:  "Reset the social network",
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := grpcClient.Reset(context.Background(), &linkedon.ResetRequest{})
		if err != nil {
			return fmt.Errorf("failed to reset: %v", err)
		}
		fmt.Printf("reset: done\n")
		return err
	},
}

func init() {
	rootCmd.AddCommand(resetCmd)
}
