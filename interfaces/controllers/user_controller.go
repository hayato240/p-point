package controllers

import (
	"errors"
	"github.com/p-point/usecase"
	"github.com/p-point/interfaces/database"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

func NewUserController(sqlHandler database.SqlHandler) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *UserController) Create(c Context) {
	c.JSON(201, errors.New("Not Implement"))
}

func (controller *UserController) Show(c Context) {
	c.JSON(200, errors.New("Not Implement"))
}