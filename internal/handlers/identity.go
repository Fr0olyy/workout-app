package handlers

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) UserIdentity(c echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		header := c.Request().Header.Get(authorizationHeader)

		if header == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "User unauthorized"})
		}

		headerParts := strings.Split(header, " ")

		if len(headerParts) != 2 {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid auth header"})
		}

		userID, err := h.services.Authorization.ParseToken(headerParts[1])

		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Could not parse token"})
		}

		c.Set(userCtx, userID)
		return c.NoContent(http.StatusNoContent)

	}
}
