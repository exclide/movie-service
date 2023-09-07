package grpcserver

import (
	"context"
	"github.com/exclide/movie-service/pkg/pb"
)

type GrpcServer struct {
	pb.UnimplementedReverserServer
}

func (s *GrpcServer) Reverse(ctx context.Context, in *pb.StringRequest) (*pb.ReverseResponse, error) {
	runes := []rune(in.Str)
	l, r := 0, len(runes)-1
	for l < r {
		runes[l], runes[r] = runes[r], runes[l]
		l++
		r--
	}

	return &pb.ReverseResponse{Str: string(runes)}, nil
}

func (s *GrpcServer) Counter(ctx context.Context, in *pb.StringRequest) (*pb.CountResponse, error) {
	resp := int32(len(in.Str))
	return &pb.CountResponse{Cnt: resp}, nil
}
