package routes

import (
	"context"
	"net/http"

	"github.com/datvvan/go-grpc-api-gateway/pkg/order/pb"
	"github.com/gin-gonic/gin"
)

type CreateOrderRequestBody struct {
	ProductID int64 `json:"productID"`
	Quantity  int64 `json:"quantity"`
}

func CreateOrder(ctx *gin.Context, c pb.OrderServiceClient) {
	body := CreateOrderRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userId, _ := ctx.Get("userID")
	if userId == nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": "none user"})
	}

	res, err := c.CreateOrder(context.Background(), &pb.CreateOrderRequest{
		ProductID: body.ProductID,
		Quantity:  body.Quantity,
		UserID:    userId.(int64),
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
