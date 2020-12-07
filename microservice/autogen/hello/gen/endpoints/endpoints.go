package hello_endpoints



import (
	context "context"
	"fmt"

	oldcontext "golang.org/x/net/context"
        pb "..\..\Workspace\github.com\liguangsheng\codelab\microservice/hello/gen/pb"
	"github.com/go-kit/kit/endpoint"
)

var _ = endpoint.Chain
var _ = fmt.Errorf
var _ = context.Background

type StreamEndpoint func(server interface{}, req interface{}) (err error)

type Endpoints struct {
	
		
			CallEndpoint endpoint.Endpoint
		
	
		
			StreamEndpoint StreamEndpoint
		
	
		
			PingPongEndpoint StreamEndpoint
		
	
}


	
		
			func (e *Endpoints)Call(ctx oldcontext.Context, in *pb.Request) (*pb.Response, error) {
				out, err := e.CallEndpoint(ctx, in)
				if err != nil {
					return &pb.Response{ErrMsg: err.Error()}, err
				}
				return out.(*pb.Response), err
			}
		
	

	
		
			func (e *Endpoints)Stream(in *pb.StreamRequest, server pb.HelloService_StreamServer) error {
				return fmt.Errorf("not implemented")
			}
		
	

	
		
			func (e *Endpoints)PingPong(server pb.HelloService_PingPongServer) error {
				return fmt.Errorf("not implemented")
			}
		
	



	
		func MakeCallEndpoint(svc pb.HelloServiceServer) endpoint.Endpoint {
			return func(ctx context.Context, request interface{}) (interface{}, error) {
				req := request.(*pb.Request)
				rep, err := svc.Call(ctx, req)
				if err != nil {
					return &pb.Response{ErrMsg: err.Error()}, err
				}
				return rep, nil
			}
		}
	

	
		func MakeStreamEndpoint(svc pb.HelloServiceServer) StreamEndpoint {
			return func(server interface{}, request interface{}) error {
				
				return svc.Stream(request.(*pb.StreamRequest), server.(pb.HelloService_StreamServer))
				
			}
		}
	

	
		func MakePingPongEndpoint(svc pb.HelloServiceServer) StreamEndpoint {
			return func(server interface{}, request interface{}) error {
				
				return svc.PingPong(server.(pb.HelloService_PingPongServer))
				
			}
		}
	


func MakeEndpoints(svc pb.HelloServiceServer) Endpoints {
	return Endpoints{
		
			CallEndpoint: MakeCallEndpoint(svc),
		
			StreamEndpoint: MakeStreamEndpoint(svc),
		
			PingPongEndpoint: MakePingPongEndpoint(svc),
		
	}
}
