package handlers

import (
	"cats/internal/domain/dto"
	"cats/internal/domain/types"
	"cats/internal/domain/usecase"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"net/http"
)

type CatsHandlers struct {
	uc        *usecase.CatsUseCase
	validator *validator.Validate
}

func NewCatsHandlers(uc *usecase.CatsUseCase, validator *validator.Validate) *CatsHandlers {
	return &CatsHandlers{
		uc:        uc,
		validator: validator,
	}
}

// GetCats godoc
// @Summary      Get cats
// @Tags         cats
// @Accept       json
// @Produce      json
// @Success      200  {array}  entity.Cat
// @Failure      400  {object}  rest.Error
// @Failure      404  {object}  rest.Error
// @Failure      500  {object}  rest.Error
// @Router       /cats [get]
func (h *CatsHandlers) GetCats(c *fiber.Ctx) error {
	res, err := h.uc.List(c.Context())
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(res)
}

// GetCat godoc
// @Summary      Get cat by ID
// @Tags         cats
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Cat ID"
// @Success      200  {object}  entity.Cat
// @Failure      400  {object}  rest.Error
// @Failure      404  {object}  rest.Error
// @Failure      500  {object}  rest.Error
// @Router       /cats/{id} [get]
func (h *CatsHandlers) GetCat(c *fiber.Ctx) error {
	id := c.Params("id")
	uid, err := uuid.Parse(id)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	cat, err := h.uc.FindByID(c.Context(), uid)
	if err != nil {
		if errors.Is(err, types.ErrCatNotFound) {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(cat)
}

// CreateCat godoc
// @Summary      Create cat
// @Tags         cats
// @Accept       json
// @Produce      json
// @Param		dto.CreateCatDTO	body	dto.CreateCatDTO	true	"Create cat request"
// @Success      200  {object}  entity.Cat
// @Failure      400  {object}  rest.Error
// @Failure      404  {object}  rest.Error
// @Failure      500  {object}  rest.Error
// @Router       /cats [post]
func (h *CatsHandlers) CreateCat(c *fiber.Ctx) error {
	p := &dto.CreateCatDTO{}
	if err := c.BodyParser(p); err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "cannot parse body: "+err.Error())
	}

	if err := h.validator.Struct(p); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	cat, err := h.uc.Create(c.Context(), p)
	if err != nil {
		if errors.Is(err, types.ErrBreedNotValid) {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(cat)
}

// UpdateCat godoc
// @Summary      Update cat by ID
// @Tags         cats
// @Accept       json
// @Produce      json
// @Param		dto.UpdateCatDTO	body	dto.UpdateCatDTO	true	"Update cat request"
// @Param        id   path      string  true  "Cat ID"
// @Success      200  {object}  entity.Cat
// @Failure      400  {object}  rest.Error
// @Failure      404  {object}  rest.Error
// @Failure      500  {object}  rest.Error
// @Router       /cats/{id} [put]
func (h *CatsHandlers) UpdateCat(c *fiber.Ctx) error {
	id := c.Params("id")
	uid, err := uuid.Parse(id)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	p := &dto.UpdateCatDTO{}
	if err := c.BodyParser(p); err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "cannot parse body: "+err.Error())
	}
	if err := h.validator.Struct(p); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	cat, err := h.uc.Update(c.Context(), uid, p)
	if err != nil {
		if errors.Is(err, types.ErrCatNotFound) {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(cat)
}

// DeleteCat godoc
// @Summary      Delete cat by ID
// @Tags         cats
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Cat ID"
// @Success      204 ""
// @Failure      400  {object}  rest.Error
// @Failure      404  {object}  rest.Error
// @Failure      500  {object}  rest.Error
// @Router       /cats/{id} [delete]
func (h *CatsHandlers) DeleteCat(c *fiber.Ctx) error {
	id := c.Params("id")
	uid, err := uuid.Parse(id)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	err = h.uc.Delete(c.Context(), uid)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.SendStatus(http.StatusNoContent)
}

func (h *CatsHandlers) Routes(r fiber.Router) {
	r.Get("", h.GetCats).
		Post("", h.CreateCat).
		Get("/:id", h.GetCat).
		Put("/:id", h.UpdateCat).
		Delete("/:id", h.DeleteCat)
}
