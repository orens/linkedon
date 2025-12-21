package cmd

import (
	"context"
	"fmt"
	"strconv"

	"github.com/orens/linkedon/linkedon"
	"github.com/spf13/cobra"
)

var postCmd = &cobra.Command{
	Use:   "post <personId> <postId> <content>",
	Short: "Post a new post",
	Long:  "Post a new post in the social network",
	RunE: func(cmd *cobra.Command, args []string) error {
		personId, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("failed to convert personId to int: %v", err)
		}
		postId, err := strconv.Atoi(args[1])
		if err != nil {
			return fmt.Errorf("failed to convert postId to int: %v", err)
		}
		_, err = grpcClient.Post(context.Background(), &linkedon.PostRequest{PersonId: int32(personId), PostId: int32(postId), Content: args[2]})
		if err != nil {
			return fmt.Errorf("failed to post: %v", err)
		}
		fmt.Printf("posted: %d posted %s\n", personId, args[1])
		return err
	},
}

func init() {
	rootCmd.AddCommand(postCmd)
}
