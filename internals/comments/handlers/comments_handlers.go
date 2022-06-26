package handlers

import (
	"back-usm/internals/comments/core/domain"
	"back-usm/internals/comments/core/ports"

	"github.com/gofiber/fiber/v2"
)

type CommentHandlers struct {
	commentServices ports.CommentServices
}

func NewCommentHandlers(commentServices ports.CommentServices) *CommentHandlers {
	return &CommentHandlers{
		commentServices: commentServices,
	}
}

func (h *CommentHandlers) CreateComment(ctx *fiber.Ctx) error {
	comment := domain.Comment{}
	if err := ctx.BodyParser(&comment); err != nil {
		return err
	}

	if err := h.commentServices.CreateComment(&comment); err != nil {
		return err
	}

	return ctx.JSON(comment)
}

func (h *CommentHandlers) FindAllComments(ctx *fiber.Ctx) error {
	comments, err := h.commentServices.FindAllComments(ctx.Params("productId"))
	if err != nil {
		return err
	}

	return ctx.JSON(comments)
}

func (h *CommentHandlers) DeleteComment(ctx *fiber.Ctx) error {
	if err := h.commentServices.DeleteComment(ctx.Params("id")); err != nil {
		return err
	}

	return ctx.JSON(map[string]string{"message": "Comment deleted"})
}
