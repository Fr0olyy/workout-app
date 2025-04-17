package handlers

import (
	"net/http"
	"traning/models"

	"github.com/labstack/echo/v4"
)

func (h *Handler) SignUp(c echo.Context) error {
	var input models.User
	if err := c.Bind(&input); err != nil {

		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {

		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "invalid request"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) SignIn(c echo.Context) error {
	var input models.User
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Could not login"})
	}

	token, err := h.services.Authorization.GenerateToken(input.Email, input.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not generate token"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
