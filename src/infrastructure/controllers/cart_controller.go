package controllers

import (
	"errors"

	"github.com/bitlogic/go-startup/src/application"
	"github.com/labstack/echo/v4"
)

type CartService interface {
	CreateNewCart(application.CreateCartCommand) (application.CartDto, error)
}

type CartController struct {
	cartService CartService
}

func NewCartController(service CartService) (*CartController, error) {
	if service == nil {
		return nil, errors.New("cart service was nil")
	}

	return &CartController{
		cartService: service,
	}, nil
}

func (cc *CartController) CreateNewCart(c echo.Context) error {
	var command application.CreateCartCommand
	if err := c.Bind(&command); err != nil {
		return err
	}

	if err := c.Validate(command); err != nil {
		return err
	}

	cartDto, err := cc.cartService.CreateNewCart(command)
	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}

	return c.JSON(201, cartDto)
}
