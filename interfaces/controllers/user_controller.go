package controllers

import (
	"github.com/hayato240/p-point/domain"
	"github.com/hayato240/p-point/interfaces/database"
	"github.com/hayato240/p-point/usecase"
	"strconv"
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
	// validate := validator.New()
	u := domain.User{}
	c.Bind(&u)

	// if err := validate.Struct(c); err != nil {
	// 	log.Fatal(err)
	// }

	identifier, err := strconv.Atoi(c.Param("id"))

	user, err := controller.Interactor.Show(identifier)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, user)
}

func (controller *UserController) PointUp(c Context) {
	u := domain.User{}
	c.Bind(&u)

	user, err := controller.Interactor.PointUp(u)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, user)

}
