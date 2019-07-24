package grpc

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/grpc-ecosystem/go-grpc-prometheus"

	"github.com/gol4ng/skeleton/internal/service"
	"github.com/gol4ng/skeleton/pkg/my_package/data_transformer"
	protos "github.com/gol4ng/skeleton/pkg/my_package/protos"
	"github.com/gol4ng/skeleton/pkg/my_package/repository"
)

type Server struct {
	server   *grpc.Server
	GRPCPort string

	documentRepository repository.DocumentRepository
}

func NewServer(container *service.Container) *Server {
	s := &Server{
		server: grpc.NewServer(
			grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
			grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
		),
		GRPCPort:           container.Cfg.GRPCPort,
		documentRepository: container.GetDocumentRepository(),
	}
	protos.RegisterMyPackageServer(s.server, s)
	reflection.Register(s.server)
	grpc_prometheus.EnableHandlingTimeHistogram()
	grpc_prometheus.Register(s.server)

	return s
}

func (s *Server) Start() error {
	lis, err := net.Listen("tcp", ":"+s.GRPCPort)
	if err != nil {
		return err
	}

	if err := s.server.Serve(lis); err != nil {
		return err
	}

	return nil
}

//@todo create serverInterface
func (s *Server) Stop() error {
	s.server.GracefulStop()
	return nil
}

func (s *Server) Store(ctx context.Context, req *protos.StoreRequest) (*protos.StoreResponse, error) {
	return &protos.StoreResponse{}, nil
}

func (s *Server) Find(req *protos.FindRequest, findSrv protos.MyPackage_FindServer) error {
	ctx := findSrv.Context()

	documents, err := s.documentRepository.Find(ctx)
	if err != nil {
		return err
	}
	if len(documents) == 0 {
		return findSrv.Send(nil)
	}
	for _, document := range documents {
		select {
		case <-ctx.Done():
			log.Printf("\tclient close connection before EOF: %s\n", ctx.Err())
			return ctx.Err()
		default:
			resp := &protos.FindResponse{
				Document: data_transformer.TransformDocument(document),
			}

			if err := findSrv.Send(resp); err != nil {
				return err
			}
		}
	}
	return nil
}
