package api

import (
	"context"
	"github.com/fi-GarO/CTC-2022/05/pkg"
	"github.com/fi-GarO/CTC-2022/05/pkg/store"
)

type Server struct {
	st store.Store
}

var _ = ApiServer(&Server{})

func NewServer(st store.Store) *Server {
	return &Server{st: st}
}

func (s *Server) Get(ctx context.Context, request *GetRequest) (*GetResponse, error) {
	kv, err := s.st.Get(ctx, request.Key)
	if err != nil {
		return nil, pkg.ToGrpcError(err)
	}

	return &GetResponse{Value: kv}, nil
}

func (s *Server) Put(ctx context.Context, request *PutRequest) (*PutResponse, error) {
	return &PutResponse{}, s.st.Put(ctx, request.Key, request.Value)
}

func (s *Server) Delete(ctx context.Context, request *DeleteRequest) (*DeleteResponse, error) {
	_, err := s.st.Delete(ctx, request.Key)
	if err != nil {
		return nil, pkg.ToGrpcError(err)
	}

	return &DeleteResponse{}, nil
}

func (s *Server) mustEmbedUnimplementedApiServer() {
}
