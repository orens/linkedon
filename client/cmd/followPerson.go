package cmd

import (
	"context"
	"fmt"
	"strconv"

	"github.com/orens/linkedon/linkedon"
	"github.com/spf13/cobra"
)

var followPersonCmd = &cobra.Command{
	Use:   "follow-person <followerId> <followeeId>",
	Short: "Follow a person",
	Long:  "Follow a person in the social network",
	RunE: func(cmd *cobra.Command, args []string) error {
		followerId, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("failed to convert followerId to int: %v", err)
		}
		followeeId, err := strconv.Atoi(args[1])
		if err != nil {
			return fmt.Errorf("failed to convert followeeId to int: %v", err)
		}
		_, err = grpcClient.FollowPerson(context.Background(), &linkedon.FollowPersonRequest{FollowerId: int32(followerId), FolloweeId: int32(followeeId)})
		return err
	},
}

func init() {
	rootCmd.AddCommand(followPersonCmd)
}
