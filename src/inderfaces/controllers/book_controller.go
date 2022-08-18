package controllers

import (
	"strconv"
	"www/src/domain"
	"www/src/inderfaces/database"
	"www/src/usecase"

	"github.com/labstack/echo"
)

type BookController struct {
	Interactor usecase.BookInteractor
}

func NewBookController(sqlHandler database.SqlHandler) *BookController {
	return &BookController{
		Interactor: usecase.BookInteractor{
			BookRepository: &database.BookRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *BookController) Show(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	book, err := controller.Interactor.BookById(id)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, book)
	return
}

func (controller *BookController) Index(c echo.Context) (err error) {
	books, err := controller.Interactor.Books()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, books)
	return
}

func (controller *BookController) Create(c echo.Context) (err error) {
	b := domain.Book{}
	c.Bind(&b)
	book, err := controller.Interactor.Add(b)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, book)
	return
}

func (controller *BookController) Save(c echo.Context) (err error) {
	b := domain.Book{}
	c.Bind(&b)
	book, err := controller.Interactor.Update(b)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, book)
	return
}

func (controller *BookController) Delete(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	book := domain.Book{
		Id: id,
	}
	err = controller.Interactor.DeletebyId(book)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, book)
	return
}
