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

type MissionsHandlers struct {
	uc        *usecase.MissionsUseCase
	validator *validator.Validate
}

func NewMissionsHandlers(uc *usecase.MissionsUseCase, validator *validator.Validate) *MissionsHandlers {
	return &MissionsHandlers{
		uc:        uc,
		validator: validator,
	}
}

// GetMissions godoc
// @Summary      Get missions
// @Tags         mission
// @Accept       json
// @Produce      json
// @Success      200  {array}  entity.Mission
// @Failure      400  {object}  rest.Error
// @Failure      404  {object}  rest.Error
// @Failure      500  {object}  rest.Error
// @Router       /missions  [get]
func (h *MissionsHandlers) GetMissions(ctx *fiber.Ctx) error {
	missions, err := h.uc.List(ctx.Context())
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(missions)
}

// CreateMission godoc
// @Summary      Create mission
// @Tags         mission
// @Accept       json
// @Produce      json
// @Param        dto.CreateMissionDTO body dto.CreateMissionDTO true "Create mission request"
// @Success      200  {object}  entity.Mission
// @Failure      400  {object}  rest.Error
// @Failure      404  {object}  rest.Error
// @Failure      500  {object}  rest.Error
// @Router       /missions  [post]
func (h *MissionsHandlers) CreateMission(ctx *fiber.Ctx) error {
	p := &dto.CreateMissionDTO{}
	if err := ctx.BodyParser(p); err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "cannot parse body: "+err.Error())
	}
	if err := h.validator.Struct(p); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	miss, err := h.uc.Create(ctx.Context(), p)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(miss)
}

// GetMission godoc
// @Summary      Get mission by ID
// @Tags         mission
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Mission ID"
// @Success      200  {object}  entity.Mission
// @Failure      400  {object}  rest.Error
// @Failure      404  {object}  rest.Error
// @Failure      500  {object}  rest.Error
// @Router       /missions/{id}  [get]
func (h *MissionsHandlers) GetMission(ctx *fiber.Ctx) error {
	id := ctx.Params("missionID")
	uid, err := uuid.Parse(id)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	mission, err := h.uc.Get(ctx.Context(), uid)
	if err != nil {
		if errors.Is(err, types.ErrMissionNotFound) {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(mission)
}

// UpdateMission godoc
// @Summary      Update mission by ID
// @Tags         mission
// @Accept       json
// @Produce      json
// @Param        dto.UpdateMissionDTO body dto.UpdateMissionDTO true "Update mission request"
// @Param        id   path      string  true  "Mission ID"
// @Success      200  {object}  entity.Mission
// @Failure      400  {object}  rest.Error
// @Failure      404  {object}  rest.Error
// @Failure      500  {object}  rest.Error
// @Router       /missions/{id} [put]
func (h *MissionsHandlers) UpdateMission(ctx *fiber.Ctx) error {
	id := ctx.Params("missionID")
	uid, err := uuid.Parse(id)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	p := &dto.UpdateMissionDTO{}
	if err := ctx.BodyParser(p); err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "cannot parse body: "+err.Error())
	}
	if err := h.validator.Struct(p); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	miss, err := h.uc.Update(ctx.Context(), uid, p)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(miss)
}

// DeleteMission godoc
// @Summary      Delete mission by ID
// @Tags         mission
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Mission ID"
// @Success      204 ""
// @Failure      400  {object}  rest.Error
// @Failure      404  {object}  rest.Error
// @Failure      500  {object}  rest.Error
// @Router       /missions/{id} [delete]
func (h *MissionsHandlers) DeleteMission(ctx *fiber.Ctx) error {
	id := ctx.Params("missionID")
	uid, err := uuid.Parse(id)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	if err := h.uc.Delete(ctx.Context(), uid); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return ctx.SendStatus(http.StatusNoContent)
}

// AddTarget godoc
// @Summary      Create target
// @Tags         mission
// @Accept       json
// @Produce      json
// @Param        dto.CreateTargetDTO body dto.CreateTargetDTO true "Create target request"
// @Param        id   path      string  true  "Mission ID"
// @Success      200  {object}  entity.Mission
// @Failure      400  {object}  rest.Error
// @Failure      404  {object}  rest.Error
// @Failure      500  {object}  rest.Error
// @Router       /missions/{id}/targets [post]
func (h *MissionsHandlers) AddTarget(ctx *fiber.Ctx) error {
	id := ctx.Params("missionID")
	uid, err := uuid.Parse(id)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	p := &dto.CreateTargetDTO{}
	if err := ctx.BodyParser(p); err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "cannot parse body: "+err.Error())
	}
	if err := h.validator.Struct(p); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	miss, err := h.uc.AddTarget(ctx.Context(), uid, p)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(miss)

}

// UpdateTarget godoc
// @Summary      Update target
// @Tags         mission
// @Accept       json
// @Produce      json
// @Param        dto.UpdateTargetDTO body dto.UpdateTargetDTO true "Update target request"
// @Param        id   path      string  true  "Mission ID"
// @Param        tid   path      string  true  "Target Name"
// @Success      200  {object}  entity.Mission
// @Failure      400  {object}  rest.Error
// @Failure      404  {object}  rest.Error
// @Failure      500  {object}  rest.Error
// @Router       /missions/{id}/targets/{tid} [put]
func (h *MissionsHandlers) UpdateTarget(ctx *fiber.Ctx) error {
	id := ctx.Params("missionID")
	uid, err := uuid.Parse(id)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	targetName := ctx.Params("targetName")

	p := &dto.UpdateTargetDTO{}
	if err := ctx.BodyParser(p); err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "cannot parse body: "+err.Error())
	}

	miss, err := h.uc.UpdateTarget(ctx.Context(), uid, targetName, p)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(miss)
}

// DeleteTarget godoc
// @Summary      Delete target
// @Tags         mission
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Mission ID"
// @Param        tid   path      string  true  "Target Name"
// @Success      200  {object}  entity.Mission
// @Failure      400  {object}  rest.Error
// @Failure      404  {object}  rest.Error
// @Failure      500  {object}  rest.Error
// @Router       /missions/{id}/targets/{tid} [delete]
func (h *MissionsHandlers) DeleteTarget(ctx *fiber.Ctx) error {
	id := ctx.Params("missionID")
	uid, err := uuid.Parse(id)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	targetName := ctx.Params("targetName")

	miss, err := h.uc.DeleteTarget(ctx.Context(), uid, targetName)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(miss)
}

func (h *MissionsHandlers) Routes(r fiber.Router) {
	r.Get("", h.GetMissions).
		Post("", h.CreateMission).
		Get("/:missionID", h.GetMission).
		Put("/:missionID", h.UpdateMission).
		Delete("/:missionID", h.DeleteMission).
		Post("/:missionID/targets", h.AddTarget).
		Put("/:missionID/targets/:targetName", h.UpdateTarget).
		Delete("/:missionID/targets/:targetName", h.DeleteTarget)

}
