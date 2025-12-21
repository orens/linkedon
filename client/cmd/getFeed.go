package cmd

import (
	"context"
	"fmt"
	"strconv"

	"github.com/orens/linkedon/linkedon"
	"github.com/spf13/cobra"
)

var getFeedCmd = &cobra.Command{
	Use:   "get-feed <personId>",
	Short: "Get the feed for a person",
	Long:  "Get the feed for a person in the social network",
	RunE: func(cmd *cobra.Command, args []string) error {
		personId, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("failed to convert personId to int: %v", err)
		}
		feed, err := grpcClient.GetFeed(context.Background(), &linkedon.GetFeedRequest{PersonId: int32(personId)})
		if err != nil {
			return fmt.Errorf("failed to get feed: %v", err)
		}
		for _, post := range feed.Posts {
			fmt.Printf("--------------------------------\n")
			fmt.Printf("%d: Posted by %s: %s\n", post.PostId, post.AuthorName, post.Content)
		}
		fmt.Printf("---------- END OF FEED --------------------\n")
		return err
	},
}

func init() {
	rootCmd.AddCommand(getFeedCmd)
}
