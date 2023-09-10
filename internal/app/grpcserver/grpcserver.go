package grpcserver

import (
	"context"
	"github.com/exclide/movie-service/api/proto/stringer"
)

type GrpcServer struct {
	stringer.UnimplementedReverserServer
}

func (s *GrpcServer) Reverse(ctx context.Context, in *stringer.StringRequest) (*stringer.ReverseResponse, error) {
	runes := []rune(in.Str)
	l, r := 0, len(runes)-1
	for l < r {
		runes[l], runes[r] = runes[r], runes[l]
		l++
		r--
	}

	return &stringer.ReverseResponse{Str: string(runes)}, nil
}

func (s *GrpcServer) Counter(ctx context.Context, in *stringer.StringRequest) (*stringer.CountResponse, error) {
	resp := int32(len(in.Str))
	return &stringer.CountResponse{Cnt: resp}, nil
}
