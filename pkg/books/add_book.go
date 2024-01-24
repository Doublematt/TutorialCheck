package books

import (
	"github.com/Doublematt/TutorialCheck/pkg/common/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AddBookRequestBody struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

func (h handler) AddBook(ctx *gin.Context) {
	body := AddBookRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var book models.Book

	book.Title = body.Title
	book.Author = body.Author
	book.Description = body.Description

	if result := h.DB.Create(&book); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusCreated, &book)
}
