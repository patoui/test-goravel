package controllers

import (
	"goravel/app/models"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

type UserController struct {
	//Dependent services
}

func NewUserController() *UserController {
	return &UserController{
		//Inject services
	}
}

func (r *UserController) Show(ctx http.Context) http.Response {
	id := ctx.Request().RouteInt("id")
	var user models.User
	facades.Orm().Query().Where("id = ?", id).First(&user)
	return ctx.Response().Success().Json(http.Json{
		"name": user.Name,
		"email": user.Email,
	})
}
