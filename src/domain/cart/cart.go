package cart

import (
	"errors"
	"github.com/bitlogic/go-startup/src/domain/customer"
	"github.com/bitlogic/go-startup/src/domain/product"
	"github.com/google/uuid"
)

type Cart struct {
	id         uuid.UUID
	customerId uuid.UUID
	items      map[uuid.UUID]item
}

func NewCart(customer *customer.Customer) (*Cart, error) {
	if customer == nil {
		return nil, errors.New("no customer provided")
	}

	return &Cart{
		id:         uuid.New(),
		customerId: customer.GetId(),
		items:      map[uuid.UUID]item{},
	}, nil
}

func (c Cart) Size() int {
	var cartSize int
	for _, v := range c.items {
		cartSize += v.quantity
	}

	return cartSize
}

func (c *Cart) AddItem(product *product.Product, quantity int) (item, error) {
	if product == nil {
		return item{}, errors.New("invalid product")
	}

	if quantity < 1 {
		return item{}, errors.New("invalid quantity")
	}

	cartItem, err := c.findItem(product.GetId())
	if err != nil {
		return c.addNewItem(product, quantity), nil
	}

	return c.updateItemQuantity(cartItem, quantity), nil
}

func (c Cart) findItem(productId uuid.UUID) (item, error) {
	cartItem, found := c.items[productId]
	if !found {
		return item{}, errors.New("item not found")
	}

	return cartItem, nil
}

func (c *Cart) addNewItem(product *product.Product, quantity int) item {
	item := newItem(product, quantity)

	c.items[product.GetId()] = item
	return item
}

func (c *Cart) updateItemQuantity(cartItem item, quantityToAdd int) item {
	cartItem.quantity += quantityToAdd
	c.items[cartItem.productId] = cartItem

	return cartItem
}

func newItem(product *product.Product, quantity int) item {
	return item{
		productId: product.GetId(),
		price:     product.GetPrice(),
		quantity:  quantity,
	}
}