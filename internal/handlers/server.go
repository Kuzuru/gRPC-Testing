package handlers

import (
	"context"

	"grpctesting/pkg/orderservice"
)

type Server struct {
	Orders []*orderservice.Order
	orderservice.UnimplementedOrdersAPIServer
}

func (s *Server) CreateOrderRequest(_ context.Context, order *orderservice.Order) (*orderservice.CreateOrderResponse, error) {
	s.Orders = append(s.Orders, order)

	return &orderservice.CreateOrderResponse{IsSuccess: true}, nil
}

func (s *Server) GetOrders(_ context.Context, _ *orderservice.Filter) (*orderservice.GetOrdersRequest, error) {
	return &orderservice.GetOrdersRequest{List: s.Orders}, nil
}
