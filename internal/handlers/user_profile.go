package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetUserProfile(c echo.Context) error {
	userID := c.Get("userID").(uint)

	profile, err := h.services.GetUserProfile(userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "User not found"})
	}

	return c.JSON(http.StatusOK, profile)
}
