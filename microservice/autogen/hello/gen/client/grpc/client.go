package hello_clientgrpc

import (
	context "context"

        "github.com/go-kit/kit/log"
        "google.golang.org/grpc"
        grpctransport "github.com/go-kit/kit/transport/grpc"
        "github.com/go-kit/kit/endpoint"
        jwt "github.com/go-kit/kit/auth/jwt"

        pb "codelab/microservice/protogen/pb"
        endpoints "codelab/microservice/protogen/hello/gen/endpoints"
)



func New(conn *grpc.ClientConn, logger log.Logger) pb.HelloServiceServer {
        
		
			var callEndpoint endpoint.Endpoint
			{
				callEndpoint = grpctransport.NewClient(
					conn,
					"hello.HelloService",
					"Call",
					EncodeCallRequest,
					DecodeCallResponse,
					pb.CallResponse{},
					append([]grpctransport.ClientOption{}, grpctransport.ClientBefore(jwt.FromGRPCContext()))...,
				).Endpoint()
			}
		
        
		
        
		
        

        return &endpoints.Endpoints {
                
			
				CallEndpoint: callEndpoint,
			
                
			
                
			
                
        }
}


	
		func EncodeCallRequest(_ context.Context, request interface{}) (interface{}, error) {
			req := request.(*pb.CallRequest)
			return req, nil
		}

		func DecodeCallResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
			response := grpcResponse.(*pb.CallResponse)
			return response, nil
		}
	

	

	

