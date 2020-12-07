package hello_grpctransport



import (
	context "context"
        "fmt"

        oldcontext "golang.org/x/net/context"
	grpctransport "github.com/go-kit/kit/transport/grpc"

        pb "..\..\Workspace\github.com\liguangsheng\codelab\microservice/hello/gen/pb"
        endpoints "..\..\Workspace\github.com\liguangsheng\codelab\microservice/hello/gen/endpoints"
)

// avoid import errors
var _ = fmt.Errorf

func MakeGRPCServer(endpoints endpoints.Endpoints) pb.HelloServiceServer {
     	var options []grpctransport.ServerOption
        _ = options
	return &grpcServer{
		
			
				call: grpctransport.NewServer(
					endpoints.CallEndpoint,
					decodeRequest,
					encodeCallResponse,
					options...,
				),
			
                
			
				stream: &server{
					e: endpoints.StreamEndpoint,
				},
			
                
			
				pingpong: &server{
					e: endpoints.PingPongEndpoint,
				},
			
                
	}
}

type grpcServer struct {
	
		
			call grpctransport.Handler
		
	
		
			stream streamHandler
		
	
		
			pingpong streamHandler
		
	
}


	
		func (s *grpcServer) Call(ctx oldcontext.Context, req *pb.Request) (*pb.Response, error) {
		_, rep, err := s.call.ServeGRPC(ctx, req)
			if err != nil {
				return nil, err
			}
			return rep.(*pb.Response), nil
		}

		func encodeCallResponse(ctx context.Context, response interface{}) (interface{}, error) {
			resp := response.(*pb.Response)
			return resp, nil
		}
	

	
		func (s *grpcServer) Stream(req *pb.StreamRequest, server pb.HelloService_StreamServer) error {
		        return s.stream.Do(server, req)
		}
	

	
		func (s *grpcServer) PingPong(server pb.HelloService_PingPongServer) error {
		        return s.pingpong.Do(server, nil)
		}
	


func decodeRequest(ctx context.Context, grpcReq interface{}) (interface{}, error) {
	return grpcReq, nil
}

type streamHandler interface{
	Do(server interface{}, req interface{}) (err error)
}

type server struct {
	e endpoints.StreamEndpoint
}

func (s server) Do(server interface{}, req interface{}) (err error) {
	if err := s.e(server, req); err != nil {
		return err
	}
	return nil
}
