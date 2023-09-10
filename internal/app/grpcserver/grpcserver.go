package grpcserver

import (
	"context"
	"github.com/exclide/movie-service/api/proto"
)

type GrpcServer struct {
	proto.UnimplementedReverserServer
}

func (s *GrpcServer) Reverse(ctx context.Context, in *proto.StringRequest) (*proto.ReverseResponse, error) {
	runes := []rune(in.Str)
	l, r := 0, len(runes)-1
	for l < r {
		runes[l], runes[r] = runes[r], runes[l]
		l++
		r--
	}

	return &proto.ReverseResponse{Str: string(runes)}, nil
}

func (s *GrpcServer) Counter(ctx context.Context, in *proto.StringRequest) (*proto.CountResponse, error) {
	resp := int32(len(in.Str))
	return &proto.CountResponse{Cnt: resp}, nil
}
