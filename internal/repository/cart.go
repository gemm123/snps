package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"synapsis/internal/model"
)

type cartRepository struct {
	DB *gorm.DB
}

type CartRepository interface {
	InsertCart(cart model.Cart) error
	InsertCartProduct(cartProduct model.CartProduct) error
	GetCartByIdUser(idUser uuid.UUID) (model.Cart, error)
	UpdateCart(cart model.Cart) error
	UpdateCartProduct(ProductCart model.CartProduct) error
	GetAllCartProductByIdCart(idCart uuid.UUID) ([]model.CartProduct, error)
	DeleteCartProduct(cartProduct model.CartProduct) error
	GetCartProduct(idCart, idProduct uuid.UUID) (model.CartProduct, error)
	DeleteCart(cart model.Cart) error
}

func NewCartRepository(DB *gorm.DB) *cartRepository {
	return &cartRepository{
		DB: DB,
	}
}

func (r *cartRepository) InsertCart(cart model.Cart) error {
	err := r.DB.Table("carts").Create(&cart).Error
	return err
}

func (r *cartRepository) InsertCartProduct(cartProduct model.CartProduct) error {
	err := r.DB.Table("cart_products").Create(&cartProduct).Error
	return err
}

func (r *cartRepository) GetCartByIdUser(idUser uuid.UUID) (model.Cart, error) {
	var cart model.Cart
	err := r.DB.Table("carts").Where("id_user = ?", idUser).Find(&cart).Error
	return cart, err
}

func (r *cartRepository) UpdateCart(cart model.Cart) error {
	err := r.DB.Save(&cart).Error
	return err
}

func (r *cartRepository) UpdateCartProduct(ProductCart model.CartProduct) error {
	err := r.DB.Save(&ProductCart).Error
	return err
}

func (r *cartRepository) GetAllCartProductByIdCart(idCart uuid.UUID) ([]model.CartProduct, error) {
	var cartProduct []model.CartProduct
	err := r.DB.Table("cart_products").Where("id_cart = ?", idCart).Find(&cartProduct).Error
	return cartProduct, err
}

func (r *cartRepository) GetCartProduct(idCart, idProduct uuid.UUID) (model.CartProduct, error) {
	var cartProduct model.CartProduct
	err := r.DB.Table("cart_products").Where("id_cart = ? AND id_product = ?", idCart, idProduct).Find(&cartProduct).Error
	return cartProduct, err
}

func (r *cartRepository) DeleteCartProduct(cartProduct model.CartProduct) error {
	existingRecord := r.DB.Table("cart_products").Where("id = ?", cartProduct.Id).First(&model.CartProduct{})
	if existingRecord.Error != nil {
		return gorm.ErrRecordNotFound
	}

	err := r.DB.Table("cart_products").Where("id = ?", cartProduct.Id).Delete(&cartProduct).Error
	return err
}

func (r *cartRepository) DeleteCart(cart model.Cart) error {
	existingRecord := r.DB.Table("carts").Where("id = ?", cart.Id).First(&model.Cart{})
	if existingRecord.Error != nil {
		return gorm.ErrRecordNotFound
	}

	err := r.DB.Table("carts").Where("id = ?", cart.Id).Delete(&cart).Error
	return err
}
