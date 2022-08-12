package rooter

import (
	"www/src/controller"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Init() *echo.Echo {
	// インスタンス
	e := echo.New()

	// ミドルウェア
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルーティング
	e.GET("/books", controller.Index)
	e.GET("/books/:id", controller.Show)
	e.POST("/books", controller.Create)
	e.PUT("/books/:id", controller.Update)
	e.DELETE("/books/:id", controller.Delete)

	return e
}
