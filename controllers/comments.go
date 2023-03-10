package controllers

import (
	"net/http"

	"github.com/arshamalh/blogo/databases"
	"github.com/arshamalh/blogo/models"
	"github.com/arshamalh/blogo/tools"
	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
)

type commentController struct {
	db databases.Database
}

func NewCommentController(db databases.Database) *commentController {
	return &commentController{
		db: db,
	}
}

func (cc *commentController) CreateComment(ctx echo.Context) error {
	var comment models.Comment
	user_id, _ := tools.ExtractUserID(ctx)
	if err := ctx.Bind(&comment); err != nil {
		return ctx.JSON(http.StatusBadRequest, gin.H{"error": "unable to parse comment"})
	}
	comment.UserID = user_id
	if err := cc.db.AddComment(&comment); err != nil {
		return ctx.JSON(http.StatusInternalServerError, gin.H{"error": "unable to add comment"})
	}

	return ctx.JSON(http.StatusOK, gin.H{"comment": comment})
}