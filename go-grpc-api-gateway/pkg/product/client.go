package product

import (
	"fmt"

	"github.com/datvvan/go-grpc-api-gateway/pkg/config"
	"github.com/datvvan/go-grpc-api-gateway/pkg/product/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.ProductServiceClient
}

func InitServiceClient(c *config.Config) pb.ProductServiceClient {
	cc, err := grpc.Dial(c.ProductSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}
	return pb.NewProductServiceClient(cc)
}
