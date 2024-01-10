package service

import (
	"errors"
	"github.com/google/uuid"
	"log"
	"synapsis/internal/model"
	"synapsis/internal/repository"
)

type orderService struct {
	orderRepository   repository.OrderRepository
	cartRepository    repository.CartRepository
	productRepository repository.ProductRepository
}

type OrderService interface {
	AddOrderProduct(orderProducts []model.OrderProductRequest, idUser uuid.UUID) error
	GetAllOrderProduct(idUser uuid.UUID) ([]model.OrderResponse, error)
	PayOrder(idOrder uuid.UUID) error
}

func NewOrderService(
	orderRepository repository.OrderRepository,
	cartRepository repository.CartRepository,
	productRepository repository.ProductRepository,
) *orderService {
	return &orderService{
		orderRepository:   orderRepository,
		cartRepository:    cartRepository,
		productRepository: productRepository,
	}
}

func (s *orderService) AddOrderProduct(orderProducts []model.OrderProductRequest, idUser uuid.UUID) error {
	cart, err := s.cartRepository.GetCartByIdUser(idUser)
	if err != nil {
		log.Println("error: " + err.Error())
		return err
	}
	if cart.Total == 0 {
		log.Println("error: cart not found")
		return errors.New("cart not found")
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

		product, err := s.productRepository.GetProductById(op.IdProduct)
		if err != nil {
			log.Println("error: " + err.Error())
			return err
		}

		product.Stok = product.Stok - orderProduct.Quantity
		if err := s.productRepository.UpdateProduct(product); err != nil {
			log.Println("error: " + err.Error())
			return err
		}
	}

	return nil
}

func (s *orderService) GetAllOrderProduct(idUser uuid.UUID) ([]model.OrderResponse, error) {
	var orderResponses []model.OrderResponse

	orders, err := s.orderRepository.GetOrdersByIdUser(idUser)
	if err != nil {
		log.Println("error: " + err.Error())
		return orderResponses, err
	}

	for _, order := range orders {
		orderResponse := model.OrderResponse{
			IdOrder:    order.Id,
			TotalPrice: order.Total,
			Status:     order.Status,
		}

		orderProducts, err := s.orderRepository.GetAllOrderProductByIdOrder(order.Id)
		if err != nil {
			log.Println("error: " + err.Error())
			return orderResponses, err
		}

		for _, orderProduct := range orderProducts {
			product, err := s.productRepository.GetProductById(orderProduct.IdProduct)
			if err != nil {
				log.Println("error: " + err.Error())
				return orderResponses, err
			}

			op := model.OrderProductResponse{
				IdProduct: uuid.New(),
				Name:      product.Name,
				Price:     product.Price,
				Quantity:  orderProduct.Quantity,
				Total:     orderProduct.Total,
			}

			orderResponse.OrderProduct = append(orderResponse.OrderProduct, op)
		}

		orderResponses = append(orderResponses, orderResponse)
	}

	return orderResponses, nil
}

func (s *orderService) PayOrder(idOrder uuid.UUID) error {
	order, err := s.orderRepository.GetOrderById(idOrder)
	if err != nil {
		log.Println("error: " + err.Error())
		return err
	}

	order.Status = "Paid"

	if err := s.orderRepository.UpdateOrder(order); err != nil {
		log.Println("error: " + err.Error())
		return err
	}

	return nil
}
