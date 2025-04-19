package handlers

import (
	"net/http"
	"strconv"

	"traning/models"

	"github.com/labstack/echo/v4"
)

func (h *Handler) CreateExercise(c echo.Context) error {
	var exercise models.Exercise
	if err := c.Bind(&exercise); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	if err := h.services.CreateExercise(&exercise); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, exercise)
}

func (h *Handler) UpdateExercise(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var exercise models.Exercise
	if err := c.Bind(&exercise); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	exercise.ID = uint(id)

	if err := h.services.UpdateExercise(&exercise); err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, exercise)
}

func (h *Handler) DeleteExercise(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.services.DeleteExercise(uint(id)); err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Exercise not found"})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "Exercise deleted"})
}

func (h *Handler) GetExercises(c echo.Context) error {
	exercises, err := h.services.GetAllExercises()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, exercises)
}

func (h *Handler) CreateExerciseLog(c echo.Context) error {
	exerciseID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid exercise ID"})
	}

	userID, ok := c.Get("user_id").(uint)
	if !ok {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid user"})
	}

	log, err := h.services.CreateExerciseLog(uint(exerciseID), userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, log)
}

func (h *Handler) GetExerciseLog(c echo.Context) error {
	logID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid log ID"})
	}

	userID, ok := c.Get("user_id").(uint)
	if !ok {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid user"})
	}

	log, timesWithPrev, err := h.services.GetExerciseLog(uint(logID), userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Log not found"})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"log":   log,
		"times": timesWithPrev,
	})
}
