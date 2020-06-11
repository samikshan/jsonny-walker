package handlers

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/samikshan/upgraded-umbrella/jsonny-walker/data"
)

func (h *Handler) AddJSONObject(c echo.Context) error {
	type req struct {
		Data map[string]interface{}
	}

	r := &req{}
	if err := c.Bind(r); err != nil {
		return &echo.HTTPError{
			Code:    http.StatusUnprocessableEntity,
			Message: "failed to process new json object request",
		}
	}

	data.JSONData["nObjects"] = data.JSONData["nObjects"].(int) + 1
	data.JSONData["components"] = data.ProcessJSONInput(
		r.Data,
		data.JSONData["components"].(map[string]interface{}),
		"",
	)

	log.Println(data.JSONData["components"])

	return c.NoContent(http.StatusOK)
}
