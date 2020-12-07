package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/gorilla/handlers"
	"google.golang.org/grpc"

	session_endpoints "codelab/microservice/autogen/hello/gen/endpoints"
	session_pb "codelab/microservice/autogen/hello/gen/pb"
	session_grpctransport "codelab/microservice/autogen/hello/gen/transports/grpc"
	session_httptransport "codelab/microservice/autogen/hello/gen/transports/http"
)

func main() {
	mux := http.NewServeMux()
	errc := make(chan error)
	s := grpc.NewServer()
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stdout)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	// initialize services
	{
		svc := session_svc.New()
		endpoints := session_endpoints.MakeEndpoints(svc)
		srv := session_grpctransport.MakeGRPCServer(endpoints)
		session_pb.RegisterSessionServiceServer(s, srv)
		session_httptransport.RegisterHandlers(svc, mux, endpoints)
	}
	{
		svc := sprint_svc.New()
		endpoints := sprint_endpoints.MakeEndpoints(svc)
		srv := sprint_grpctransport.MakeGRPCServer(endpoints)
		sprint_pb.RegisterSprintServiceServer(s, srv)
		sprint_httptransport.RegisterHandlers(svc, mux, endpoints)
	}
	{
		svc := user_svc.New()
		endpoints := user_endpoints.MakeEndpoints(svc)
		srv := user_grpctransport.MakeGRPCServer(endpoints)
		user_pb.RegisterUserServiceServer(s, srv)
		user_httptransport.RegisterHandlers(svc, mux, endpoints)
	}

	// start servers
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		logger := log.With(logger, "transport", "HTTP")
		logger.Log("addr", ":8000")
		errc <- http.ListenAndServe(":8000", handlers.LoggingHandler(os.Stderr, mux))
	}()

	go func() {
		logger := log.With(logger, "transport", "gRPC")
		ln, err := net.Listen("tcp", ":9000")
		if err != nil {
			errc <- err
			return
		}
		logger.Log("addr", ":9000")
		errc <- s.Serve(ln)
	}()

	logger.Log("exit", <-errc)
}
