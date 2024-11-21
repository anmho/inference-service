package main

import (
	"connectrpc.com/connect"
	"context"
	"fmt"
	pb "github.com/anmho/inference/gen/protos/v1"
	"github.com/anmho/inference/gen/protos/v1/inferencev1connect"
	"github.com/spf13/cobra"
	"log"
	"net/http"
)

var client inferencev1connect.InferenceServiceClient

var completionsCmd = &cobra.Command{
	Use:   "completions",
	Short: "Completions commands",
}

var getCompletionsCmd = &cobra.Command{
	Use:   "get",
	Short: "Get completions",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Getting completions!")
		ctx := context.Background()
		msg := &pb.CompletionsRequest{
			Prompt: "What is your name?",
		}
		request := &connect.Request[pb.CompletionsRequest]{Msg: msg}

		response, err := client.GetCompletions(ctx, request)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(response)
	},
}

var streamCompletionsCmd = &cobra.Command{
	Use:   "stream",
	Short: "Stream completions",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("streaming completions!")
	},
}

var rootCmd = &cobra.Command{
	Use:   "completions-cli",
	Short: "A completions cli",
}

const (
	BaseURL = "http://localhost:8080"
)

func main() {
	fmt.Println("hello")
	client = inferencev1connect.NewInferenceServiceClient(http.DefaultClient, BaseURL)

	completionsCmd.AddCommand(getCompletionsCmd)
	completionsCmd.AddCommand(streamCompletionsCmd)

	rootCmd.AddCommand(completionsCmd)
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
	}
}
