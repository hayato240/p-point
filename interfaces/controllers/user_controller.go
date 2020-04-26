package controllers

import (
	"github.com/p-point/domain"
	"github.com/p-point/interfaces/database"
	"github.com/p-point/usecase"
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
	u := domain.User{}
	c.Bind(&u)
	user, err := controller.Interactor.Add(u)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(201, user)
}

func (controller *UserController) Show(c Context) {
	u := domain.User{}
	c.Bind(&u)
	user, err := controller.Interactor.Show(u)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, user)
}
