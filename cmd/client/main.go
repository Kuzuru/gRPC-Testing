package main

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"grpctesting/pkg/orderservice"
)

func main() {
	cwt, _ := context.WithTimeout(context.Background(), time.Second*5)

	conn, err := grpc.DialContext(cwt, "localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	os := orderservice.NewOrdersAPIClient(conn)

	newOrder := &orderservice.Order{
		Name:     "iPhone 12 Pro MAX",
		Price:    "149000",
		Category: "smartphone",
	}

	or, err := os.CreateOrderRequest(cwt, newOrder)
	if err != nil {
		panic(err)
	}

	if or.IsSuccess {
		fmt.Println("Order created")
	}

	orders, err := os.GetOrders(cwt, &orderservice.Filter{})
	if err != nil {
		panic(err)
	}

	for i := range orders.List {
		order := orders.List[i]

		fmt.Printf("Name: %v \nPrice: %v \nCategory: %v\n\n", order.Name, order.Price, order.Category)
	}
}
