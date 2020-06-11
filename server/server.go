package server

import (
	"github.com/labstack/echo"
	"github.com/samikshan/upgraded-umbrella/jsonny-walker/handlers"
)

// Server is an HTTP server
type Server struct {
	E *echo.Echo
	H *handlers.Handler
}

// New returns a new http server instance
func New() *Server {
	sv := &Server{
		E: echo.New(),
		H: handlers.New(),
	}
	sv.routes()
	return sv
}

func (sv *Server) routes() {
	// Routes
	sv.E.POST("/addJSONObject", sv.H.AddJSONObject)
	sv.E.POST("/getPathsData", sv.H.GetPathsData)
}
