package inference

import (
	"connectrpc.com/connect"
	"context"
	pb "github.com/anmho/inference/gen/protos/v1"
	"github.com/anmho/inference/gen/protos/v1/inferencev1connect"
	"github.com/google/generative-ai-go/genai"
	"github.com/openai/openai-go"
)

var _ inferencev1connect.InferenceServiceHandler = (*Service)(nil)

type Service struct {
	openaiClient   *openai.Client
	googleAIClient *genai.Client
}

func NewService(openaiClient *openai.Client, googleAIClient *genai.Client) *Service {
	return &Service{}
}
func (s Service) GetCompletions(ctx context.Context, request *connect.Request[pb.CompletionsRequest]) (*connect.Response[pb.GetCompletionsResponse], error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) StreamCompletions(ctx context.Context, c *connect.Request[pb.CompletionsRequest], c2 *connect.ServerStream[pb.StreamCompletionsResponse]) error {
	//TODO implement me
	panic("implement me")
}
