package cmd

import (
	"context"
	"fmt"
	"strconv"

	"github.com/orens/linkedon/linkedon"
	"github.com/spf13/cobra"
)

var createPersonCmd = &cobra.Command{
	Use:   "create-person <name> <age> <id>",
	Short: "Create a new person",
	Long:  "Create a new person in the social network",
	RunE: func(cmd *cobra.Command, args []string) error {
		age, err := strconv.Atoi(args[1])
		if err != nil {
			return fmt.Errorf("failed to convert age to int: %v", err)
		}
		id, err := strconv.Atoi(args[2])
		if err != nil {
			return fmt.Errorf("failed to convert id to int: %v", err)
		}
		person := &linkedon.Person{
			Name: args[0],
			Age:  int32(age),
			Id:   int32(id),
		}
		_, err = grpcClient.CreatePerson(context.Background(), &linkedon.CreatePersonRequest{Person: person})
		if err != nil {
			return fmt.Errorf("failed to create person: %v", err)
		}
		fmt.Printf("created person: %v\n", person)
		return err
	},
}

func init() {
	rootCmd.AddCommand(createPersonCmd)
}
