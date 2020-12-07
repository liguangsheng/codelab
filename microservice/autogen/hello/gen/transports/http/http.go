package hello_httptransport



import (
       "log"
	"net/http"
	"encoding/json"
	"context"

        pb "..\..\Workspace\github.com\liguangsheng\codelab\microservice/hello/gen/pb"
        gokit_endpoint "github.com/go-kit/kit/endpoint"
        httptransport "github.com/go-kit/kit/transport/http"
        endpoints "..\..\Workspace\github.com\liguangsheng\codelab\microservice/hello/gen/endpoints"
)

var _ = log.Printf
var _ = gokit_endpoint.Chain
var _ = httptransport.NewClient



	
		func MakeCallHandler(svc pb.HelloServiceServer, endpoint gokit_endpoint.Endpoint) *httptransport.Server {
			return httptransport.NewServer(
				endpoint,
				decodeCallRequest,
				encodeResponse,
				[]httptransport.ServerOption{}...,
			)
		}

		func decodeCallRequest(ctx context.Context, r *http.Request) (interface{}, error) {
			var req pb.Request
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				return nil, err
			}
			return &req, nil
		}
	

	

	


func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func RegisterHandlers(svc pb.HelloServiceServer, mux *http.ServeMux, endpoints endpoints.Endpoints) error {
	
		
			log.Println("new HTTP endpoint: \"/Call\" (service=Hello)")
			mux.Handle("/Call", MakeCallHandler(svc, endpoints.CallEndpoint))
		
	
		
	
		
	
	return nil
}
