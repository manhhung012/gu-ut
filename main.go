package main

import (
	controller "go.ut/hello/controller"
	models "go.ut/hello/models"
	repos "go.ut/hello/repository"
	service "go.ut/hello/services"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")

	//masterpage
	tmpl := iris.HTML("./templates", ".html").Layout("masterpage.html").Reload(true)
	app.RegisterView(tmpl)

	app.StaticContent("/static","./static", []byte(""))

	//routes
	app.Get("/", homeHandler)

	// **** BOOKS (MySQL)
	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/recordings"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Album{})

	albumRepo := repos.NewAlbumRepository(db)
	albumService := service.NewAlbumService(albumRepo)
	albums := mvc.New(app.Party("/album"))
	albums.Register(albumService)
	albums.Handle(new(controller.AlbumController))

	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("Message", ctx.Values().GetStringDefault("message", "The page you're looking for doesn't exist"))
		ctx.View("error.html")
	})

	app.Run(
		iris.Addr(":8080"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)

}
