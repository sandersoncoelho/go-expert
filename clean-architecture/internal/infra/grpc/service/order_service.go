package service

import (
	"context"

	"github.com/devfullcycle/20-CleanArch/internal/infra/grpc/pb"
	"github.com/devfullcycle/20-CleanArch/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase usecase.CreateOrderUseCase
}

func NewOrderService(createOrderUseCase usecase.CreateOrderUseCase) *OrderService {
	return &OrderService{
		CreateOrderUseCase: createOrderUseCase,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	dto := usecase.OrderInputDTO{
		ID:    in.Id,
		Price: float64(in.Price),
		Tax:   float64(in.Tax),
	}
	output, err := s.CreateOrderUseCase.Create(dto)
	if err != nil {
		return nil, err
	}
	return &pb.CreateOrderResponse{
		Id:         output.ID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}

func (s *OrderService) ListOrders(ctx context.Context, in *pb.Blank) (*pb.OrderList, error) {
	ordersDTO, err := s.CreateOrderUseCase.List()
	if err != nil {
		return nil, err
	}

	var ordersResponse []*pb.CreateOrderResponse

	for _, order := range ordersDTO {
		orderResponse := &pb.CreateOrderResponse{
			Id:         order.ID,
			Price:      float32(order.Price),
			Tax: 		float32(order.Tax),
			FinalPrice: float32(order.FinalPrice),
		}

		ordersResponse = append(ordersResponse, orderResponse)
	}

	return &pb.OrderList{Orders: ordersResponse}, nil
}
