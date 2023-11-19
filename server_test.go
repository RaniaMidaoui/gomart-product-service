package test

import (
	"context"
	"log"
	"net"
	"testing"

	"github.com/RaniaMidaoui/goMart-product-service/pkg/db"
	"github.com/RaniaMidaoui/goMart-product-service/pkg/pb"
	"github.com/RaniaMidaoui/goMart-product-service/pkg/services"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

func server(ctx context.Context) (pb.ProductServiceClient, func(), error) {
	lis := bufconn.Listen(1024 * 1024)

	s := grpc.NewServer()

	h := db.Mock()

	ss := services.Server{
		H: h,
	}

	pb.RegisterProductServiceServer(s, &ss)

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()

	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
		return lis.Dial()
	}), grpc.WithInsecure())

	if err != nil {
		return nil, nil, err
	}

	closer := func() {
		err := lis.Close()
		if err != nil {
			log.Printf("error closing listener: %v", err)
		}
		s.Stop()
	}

	return pb.NewProductServiceClient(conn), closer, nil
}

func TestFindProduct(t *testing.T) {
	ctx := context.Background()

	c, closer, err := server(ctx)

	if err != nil {
		t.Fatalf("failed to start test server: %v", err)
	}

	defer closer()

	type expectation struct {
		status int64
	}

	tests := map[string]struct {
		req  *pb.FindOneRequest
		want expectation
	}{
		"success": {
			req: &pb.FindOneRequest{
				Id: 1,
			},
			want: expectation{
				status: 200,
			},
		},
		"not found": {
			req: &pb.FindOneRequest{
				Id: 999,
			},
			want: expectation{
				status: 404,
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			res, err := c.FindOne(ctx, tc.req)

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if res.GetStatus() != tc.want.status {
				t.Errorf("expected %v, got %v", tc.want.status, res.GetStatus())
			}
		})
	}

}
