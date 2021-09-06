package main

import (
	"context"
	trfmain "facebook-connector/pkg/rule/TRF_Main"
	util "facebook-connector/pkg/util"
)

import (
	_ "facebook-connector/docs"
)
import (
	"facebook-connector/pkg/middleware"
)
import (
	"facebook-connector/pkg/controller"
	"facebook-connector/pkg/xiLogger"
	"facebook-connector/sidecar"
	"log"
	"net"
)

func main() {

	cntxt := util.CustomContext{}
	cntxt.AppGoContext = context.Background()
	config := make(map[string]interface{})
	trfmain.TRF_Main(&cntxt, config)

	RunServer()
}

func RunServer() {
	port := middleware.GetGrpcPort()
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := middleware.GetGrpcServerConfigs()

	xiLogger.Log.Info("\n Server listening on port %v \n", ":"+port)
	sidecar.RegisterServiceServer(s, &controller.Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
