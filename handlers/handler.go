package handlers

type Handler struct{}

// New creates a new HTTP handler
func New() *Handler {
	return &Handler{}
}
