/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/orens/linkedon/linkedon"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var grpcClient linkedon.LinkedonServiceClient

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "Linkedon client",
	Short: "Linkedon client",
	Long:  `Linkedon client to manage your social network`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if grpcClient != nil {
			return nil
		}
		conn, err := grpc.NewClient("localhost:8044", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("failed to connect: %v", err)
			panic(err)
		}
		cobra.OnFinalize(func() {
			conn.Close()
		})
		grpcClient = linkedon.NewLinkedonServiceClient(conn)
		return nil
	}}

func Execute() error {
	return rootCmd.Execute()
}

var verbose bool

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
}
