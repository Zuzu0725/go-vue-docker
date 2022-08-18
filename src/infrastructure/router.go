package infrastructure

import (
	"www/src/inderfaces/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Init() {
	// インスタンス作成
	e := echo.New()

	bookController := controllers.NewBookController(NewSqlHandler())

	// ミドルウェア
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルーティング
	e.GET("/books", func(c echo.Context) error { return bookController.Index(c) })
	e.GET("/books/:id", func(c echo.Context) error { return bookController.Show(c) })
	e.POST("/books", func(c echo.Context) error { return bookController.Create(c) })
	e.PUT("/books/:id", func(c echo.Context) error { return bookController.Save(c) })
	e.DELETE("/books/:id", func(c echo.Context) error { return bookController.Delete(c) })

	// サーバー起動
	e.Logger.Fatal(e.Start(":8080"))
}
