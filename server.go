package main

import "github.com/anmho/inference/gen/inference/v1/proto/inferencev1connect"



var _ inferencev1connect.GreetServiceHandler = (*Service)(nil)
type Service struct {
}

