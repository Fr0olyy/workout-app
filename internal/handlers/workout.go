package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type WorkoutController struct {
	DB *gorm.DB
}

func (h *Handler) GetWorkouts(c echo.Context) error {
	workouts, err := h.services.GetAllWorkouts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, workouts)
}

func (h *Handler) GetWorkout(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid ID"})
	}

	workout, minutes, err := h.services.GetWorkout(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"workout": workout,
		"minutes": minutes,
	})
}

func (h *Handler) CreateWorkout(c echo.Context) error {
	var body struct {
		Name        string `json:"name"`
		ExerciseIDs []uint `json:"exerciseIds"`
	}
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}

	workout, err := h.services.CreateWorkout(body.Name, body.ExerciseIDs)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, workout)
}

func (h *Handler) UpdateWorkout(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid ID"})
	}

	var body struct {
		Name        string `json:"name"`
		ExerciseIDs []uint `json:"exerciseIds"`
	}
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}

	workout, err := h.services.UpdateWorkout(uint(id), body.Name, body.ExerciseIDs)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, workout)
}

func (h *Handler) DeleteWorkout(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid ID"})
	}

	if err := h.services.DeleteWorkout(uint(id)); err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "Workout deleted!"})
}
