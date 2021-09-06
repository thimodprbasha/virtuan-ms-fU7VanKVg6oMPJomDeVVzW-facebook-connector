package controller

import (
	api "facebook-connector/pkg/api"
	sidecar "facebook-connector/sidecar"
)

import (
	"context"
)

type Server struct {
	sidecar.UnimplementedServiceServer
}

func (s *Server) Sidecar(ctx context.Context, request *sidecar.SidecarRequest) (*sidecar.ResponseDto, error) {

	result, err := api.Sidecar(request.RequestDto, ctx)

	return result, err
}
