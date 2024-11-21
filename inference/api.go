package inference

import (
	"connectrpc.com/connect"
	"context"
	v1 "github.com/anmho/inference/gen/protos/v1"
	"github.com/anmho/inference/gen/protos/v1/inferencev1connect"
	"github.com/openai/openai-go"
)

var _ inferencev1connect.InferenceServiceHandler = (*Service)(nil)

type Service struct {
}

func NewService(openaiClient *openai.Client) *Service {
	return &Service{}
}
func (s Service) GetCompletions(ctx context.Context, c *connect.Request[v1.CompletionsRequest]) (*connect.Response[v1.GetCompletionsResponse], error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) StreamCompletions(ctx context.Context, c *connect.Request[v1.CompletionsRequest], c2 *connect.ServerStream[v1.StreamCompletionsResponse]) error {
	//TODO implement me
	panic("implement me")
}
