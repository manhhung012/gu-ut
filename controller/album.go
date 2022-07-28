package controller

import (
	models "go.ut/hello/models"
	services "go.ut/hello/services"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
)

type AlbumController struct {
	Ctx     iris.Context
	Service services.AlbumService
	Session *sessions.Session
}

var registerBookStaticView = mvc.View{
	Name: "book_home.html",
	Data: iris.Map{"Title": "Book Home Page"},
}

func (c *AlbumController) Get() (results []models.Album) {
	return c.Service.GetAll()
}
