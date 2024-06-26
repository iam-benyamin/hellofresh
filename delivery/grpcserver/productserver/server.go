package productserver

import (
	"context"
	"fmt"
	"github.com/iam-benyamin/hellofresh/contract/goproto/productproto"
	"github.com/iam-benyamin/hellofresh/param/productparam"
	"github.com/iam-benyamin/hellofresh/pkg/richerror"
	"github.com/iam-benyamin/hellofresh/service/productservice"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"sync"
)

type ProductServer struct {
	productproto.UnimplementedProductServiceServer
	svc productservice.Service
}

func New(svc productservice.Service) ProductServer {
	return ProductServer{
		UnimplementedProductServiceServer: productproto.UnimplementedProductServiceServer{},
		svc:                               svc,
	}
}

func (p ProductServer) Product(ctx context.Context, req *productproto.ProductRequest) (*productproto.ProductResponse, error) {
	const op = "productserver.Product"

	product, err := p.svc.Product(ctx, productparam.ProductRequest{ProductCode: req.ProductCode})
	if err != nil {
		e := richerror.New(op).WithErr(err)

		if e.Kind() == richerror.KindNotFound {
			return nil, status.Errorf(codes.NotFound, e.Message())
		}

		return nil, status.Errorf(codes.Unknown, e.Message())
	}

	return &productproto.ProductResponse{
		Id:    product.ID,
		Name:  product.Name,
		Code:  product.Code,
		Price: product.Price,
	}, nil
}

func (p ProductServer) Start(done <-chan bool, wg *sync.WaitGroup) {
	address := fmt.Sprintf(":%d", 8087)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()

	productproto.RegisterProductServiceServer(grpcServer, p)

	go func() {
		fmt.Println("Starting grpc server on " + address)
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatal("couldn't start product grpc server")
		}
	}()

	go func() {
		wg.Add(1)
		defer wg.Done()
		<-done
		grpcServer.GracefulStop()
		fmt.Println("grpc product server shutdown gracefully")
	}()
}
