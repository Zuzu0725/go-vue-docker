package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"www/src/model"

	"github.com/labstack/echo"
)

func Index(c echo.Context) error {
	books := model.Books{}
	books.GetBooks()

	return c.JSON(http.StatusOK, books)
}

func Show(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	book := model.Book{}
	book.GetBook(id)

	return c.JSON(http.StatusOK, book)
}

func Create(c echo.Context) error {
	book := new(model.Book)
	if err := c.Bind(book); err != nil {
		fmt.Println("データの取得に失敗しました")
		return err
	}
	book.CreateBook()

	return c.JSON(http.StatusOK, book)
}

func Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	b := new(model.Book)
	if err := c.Bind(b); err != nil {
		fmt.Println("データの取得に失敗しました")
		return err
	}

	book := model.Book{
		Id:      id,
		Title:   b.Title,
		Author:  b.Author,
		Summary: b.Summary,
	}
	book.UpdateBook()

	return c.JSON(http.StatusOK, book)
}

func Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	book := model.Book{}
	book.DeleteBook(id)

	return c.JSON(http.StatusOK, book)
}
