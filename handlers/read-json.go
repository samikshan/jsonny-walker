package handlers

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/samikshan/upgraded-umbrella/jsonny-walker/data"
)

func (h *Handler) GetPathsData(c echo.Context) error {
	type req struct {
		K         int
		Threshold float64
	}

	r := &req{}
	if err := c.Bind(r); err != nil {
		return &echo.HTTPError{
			Code:    http.StatusUnprocessableEntity,
			Message: "failed to process request get paths",
		}
	}

	paths := data.GetPaths(r.K, r.Threshold)

	return c.JSON(http.StatusOK, paths)
}
