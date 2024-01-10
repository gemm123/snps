package service

import (
	"github.com/google/uuid"
	"log"
	"synapsis/internal/model"
	"synapsis/internal/repository"
	"time"
)

type cartService struct {
	cartRepository    repository.CartRepository
	productRepository repository.ProductRepository
}

type CartService interface {
	AddProductToCart(cartRequest model.CartRequest, idUser uuid.UUID) error
}

func NewCartService(cartRepository repository.CartRepository, productRepository repository.ProductRepository) *cartService {
	return &cartService{
		cartRepository:    cartRepository,
		productRepository: productRepository,
	}
}

func (s *cartService) AddProductToCart(cartRequest model.CartRequest, idUser uuid.UUID) error {
	cart, err := s.cartRepository.GetCartByIdUser(idUser)
	if err != nil {
		log.Println("error: " + err.Error())
		return err
	}

	if cart.Total != 0 {
		cartProducts, err := s.cartRepository.GetAllCartProductByIdCart(cart.Id)
		if err != nil {
			log.Println("error: " + err.Error())
			return err
		}

		for _, cp := range cartProducts {
			if cp.IdProduct == cartRequest.IdProduct {
				cp.Quantity = cp.Quantity + cartRequest.Quantity
				cp.Total = cp.Total + cartRequest.Total
				cp.UpdatedAt = time.Now()

				if err := s.cartRepository.UpdateCartProduct(cp); err != nil {
					log.Println("error: " + err.Error())
					return err
				}

				cart.Total = cart.Total + cartRequest.Total
				cart.UpdatedAt = time.Now()

				if err := s.cartRepository.UpdateCart(cart); err != nil {
					log.Println("error: " + err.Error())
					return err
				}

				return nil
			}
		}

		newCartProduct := model.CartProduct{
			Id:        uuid.New(),
			IdCart:    cart.Id,
			IdProduct: cartRequest.IdProduct,
			Quantity:  cartRequest.Quantity,
			Total:     cartRequest.Total,
		}

		if err := s.cartRepository.InsertCartProduct(newCartProduct); err != nil {
			log.Println("error: " + err.Error())
			return err
		}

		cart.Total = cart.Total + cartRequest.Total
		cart.UpdatedAt = time.Now()

		if err := s.cartRepository.UpdateCart(cart); err != nil {
			log.Println("error: " + err.Error())
			return err
		}

		return nil
	}

	newCart := model.Cart{
		Id:     uuid.New(),
		IdUser: idUser,
		Total:  cartRequest.Total,
	}

	if err := s.cartRepository.InsertCart(newCart); err != nil {
		log.Println("error: " + err.Error())
		return err
	}

	cartProduct := model.CartProduct{
		Id:        uuid.New(),
		IdCart:    newCart.Id,
		IdProduct: cartRequest.IdProduct,
		Quantity:  cartRequest.Quantity,
		Total:     cartRequest.Total,
	}

	if err := s.cartRepository.InsertCartProduct(cartProduct); err != nil {
		log.Println("error: " + err.Error())
		return err
	}

	return nil
}
