package orderservice

import (
	"context"
	"fmt"
	"github.com/iam-benyamin/hellofresh/contract/goproto/productproto"
	"github.com/iam-benyamin/hellofresh/contract/goproto/userproto"
	"github.com/iam-benyamin/hellofresh/entity/orderentity"
	"github.com/iam-benyamin/hellofresh/param/orderparam"
	"github.com/iam-benyamin/hellofresh/param/productparam"
	"github.com/iam-benyamin/hellofresh/param/userparam"
	"github.com/iam-benyamin/hellofresh/pkg/errmsg"
	"github.com/iam-benyamin/hellofresh/pkg/richerror"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

func (s Service) CreateNewOrder(ctx context.Context, req orderparam.CreateOrderRequest) error {
	const op = "orderservice.CreateNewOrder"

	// TODO: what about call user microservice with the gateway authenticate and authorization
	// user at that level and append user info to request
	p, err := fetchProduct(ctx, req.ProductCode)
	if err != nil {
		fmt.Println("FetchDataFromProduct", err)
		return richerror.New(op)
	}

	// TODO: have a cache layer (i.e. redis) save the products at the cache and don't call
	// another service every time (reduce load and increasing latency) for and we can listen
	// to the queue the broadcast product updated
	u, err := fetchUserProfile(ctx, req.UserID)
	if err != nil {
		return richerror.New(op).WithErr(err)
	}

	err = s.repo.SaveOrder(ctx, orderparam.SaveOrder{
		UserID:           u.ID,
		ProductCode:      p.Code,
		CustomerFullName: fmt.Sprintf("%s %s", u.FirstName, u.LastName),
		ProductName:      p.Name,
		TotalAmount:      p.Price,
	})
	if err != nil {
		return richerror.New(op).WithErr(err).WithMessage(errmsg.ErrorMsgCantCreateItem).WithKind(richerror.KindUnexpected)
	}

	// TODO: publish message
	if err = s.broker.PublishCreatedOrder(ctx, orderparam.Message{
		Producer: "",
		SentAt:   "",
		Type:     "",
		Payload: orderparam.Payload{
			Order: orderentity.Order{
				ID:               "",
				UserID:           "",
				ProductCode:      "",
				CustomerFullName: "",
				ProductName:      "",
				TotalAmount:      0,
				CreatedAt:        time.Time{},
			},
		},
	}, "created_order"); err != nil {
		return richerror.New(op).WithErr(err).WithMessage(errmsg.ErrorMsgSomeThingWentWrong).
			WithKind(richerror.KindUnexpected)
	}

	return nil
}

func fetchUserProfile(ctx context.Context, userID string) (userparam.ProfileResponse, error) {
	const op = "orderservice.fetchUserProfile"

	//TODO: technical dept - WithInsecure and WithBlock and Dial is deprecated
	conn, err := grpc.Dial("localhost:8086", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return userparam.ProfileResponse{}, richerror.New(op).WithErr(err).WithKind(richerror.KindNotFound).
			WithMessage(errmsg.ErrorMsgNetwork)
	}
	defer func(conn *grpc.ClientConn) {
		_ = conn.Close()
	}(conn)

	client := userproto.NewUserServiceClient(conn)

	req := &userproto.ProfileRequest{UserId: userID}
	res, err := client.Profile(ctx, req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.NotFound {
			return userparam.ProfileResponse{}, richerror.New(op).WithErr(err).WithKind(richerror.KindNotFound).
				WithMessage(errmsg.ErrorMsgNotFound)
		}

		return userparam.ProfileResponse{}, richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected).
			WithMessage(errmsg.ErrorMsgNetwork)
	}

	return userparam.ProfileResponse{
		ID:        res.Id,
		FirstName: res.FirstName,
		LastName:  res.LastName,
	}, nil
}

func fetchProduct(ctx context.Context, ProductCode string) (productparam.ProductResponse, error) {
	const op = "orderservice.fetchDataFromProduct"

	conn, err := grpc.Dial("localhost:8087", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		// TODO: handel the error with retry and | or circuit Breaker pattern
		return productparam.ProductResponse{}, richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected).
			WithMessage(errmsg.ErrorMsgNetwork)
	}
	defer func(conn *grpc.ClientConn) {
		_ = conn.Close()
	}(conn)

	client := productproto.NewProductServiceClient(conn)

	// Contact the server and print out its response.
	req := &productproto.ProductRequest{ProductCode: ProductCode}
	res, err := client.Product(ctx, req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.NotFound {
			return productparam.ProductResponse{}, richerror.New(op).WithErr(err).WithKind(richerror.KindNotFound).
				WithMessage(errmsg.ErrorMsgNotFound)
		}

		return productparam.ProductResponse{}, richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected).
			WithMessage(errmsg.ErrorMsgNetwork)
	}

	return productparam.ProductResponse{
		ID:    res.Id,
		Name:  res.Name,
		Code:  res.Code,
		Price: res.Price,
	}, nil
}
