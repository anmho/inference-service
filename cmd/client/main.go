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
		ctx := context.Background()
		msg := &pb.CompletionsRequest{
			Prompt: "What is your name?",
		}
		request := &connect.Request[pb.CompletionsRequest]{Msg: msg}
		fmt.Println("Getting completions!")
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

func main() {
	client = inferencev1connect.NewInferenceServiceClient(http.DefaultClient, "localhost:8080")

	completionsCmd.AddCommand(getCompletionsCmd)
	completionsCmd.AddCommand(streamCompletionsCmd)

}
