package service

import (
	"github.com/google/uuid"
	"log"
	"synapsis/internal/model"
	"synapsis/internal/repository"
)

type orderService struct {
	orderRepository repository.OrderRepository
	cartRepository  repository.CartRepository
}

type OrderService interface {
	AddOrderProduct(orderProducts []model.OrderProductRequest, idUser uuid.UUID) error
}

func NewOrderService(orderRepository repository.OrderRepository, cartRepository repository.CartRepository) *orderService {
	return &orderService{
		orderRepository: orderRepository,
		cartRepository:  cartRepository,
	}
}

func (s *orderService) AddOrderProduct(orderProducts []model.OrderProductRequest, idUser uuid.UUID) error {
	cart, err := s.cartRepository.GetCartByIdUser(idUser)
	if err != nil {
		log.Println("error: " + err.Error())
		return err
	}

	order := model.Order{
		Id:     uuid.New(),
		IdUser: idUser,
		Total:  cart.Total,
		Status: "Unpaid",
	}

	if err := s.orderRepository.InsertOrder(order); err != nil {
		log.Println("error: " + err.Error())
		return err
	}

	for _, op := range orderProducts {
		orderProduct := model.OrderProduct{
			Id:        uuid.New(),
			IdOrder:   order.Id,
			IdProduct: op.IdProduct,
			Quantity:  op.Quantity,
			Total:     op.Total,
		}

		if err := s.orderRepository.InsertOrderProduct(orderProduct); err != nil {
			log.Println("error: " + err.Error())
			return err
		}
	}

	return nil
}
