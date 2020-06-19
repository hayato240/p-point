package controllers

import (
	"errors"
	"reflect"
	"strconv"

	"github.com/hayato240/p-point/domain"
	"github.com/hayato240/p-point/interfaces/database"
	"github.com/hayato240/p-point/usecase"
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

	identifier, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, err)
		return
	}

	user, err := controller.Interactor.Show(identifier)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, user)
}

func (controller *UserController) Points(c Context) {
	u := domain.User{}
	c.Bind(&u)

	identifier, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, err)
		return
	}

	u.ID = identifier

	if reflect.TypeOf(u.Amount).Kind() != reflect.Int {
		err = errors.New("type of Amount is not int")
		c.JSON(400, err)
		return
	}

	user, err := controller.Interactor.Points(u)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, user)

}
