package http

import (
	http "net/http"

	influxdb "github.com/influxdata/influxdb"
	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"
)

// CheckBackend are all services a checkhandler requires
type CheckBackend struct {
	influxdb.HTTPErrorHandler
	Logger *zap.Logger
}

// NewCheckBackend returns a new checkbackend
func NewCheckBackend(b *APIBackend) *CheckBackend {
	return &CheckBackend{
		Logger: b.Logger.With(zap.String("handler", "check")),
	}
}

// CheckHandler responds to /api/v2/checks requests
type CheckHandler struct {
	*httprouter.Router
	influxdb.HTTPErrorHandler
	Logger *zap.Logger
}

const (
	checksPath   = "/api/v2/checks"
	checksIDPath = "/api/v2/checks/:id"
)

// NewCheckHandler returns a new checkhandler
func NewCheckHandler(b *CheckBackend) *CheckHandler {
	h := &CheckHandler{
		Router: NewRouter(b.HTTPErrorHandler),
		Logger: zap.NewNop(),
	}

	h.HandlerFunc("GET", checksPath, h.handleGetChecks)
	h.HandlerFunc("POST", checksPath, h.handleCreateCheck)

	h.HandlerFunc("GET", checksIDPath, h.handleGetCheck)
	h.HandlerFunc("PATCH", checksIDPath, h.handleUpdateCheck)
	h.HandlerFunc("DELETE", checksIDPath, h.handleDeleteCheck)

	return h
}

func (h *CheckHandler) handleGetChecks(w http.ResponseWriter, r *http.Request) {

}

func (h *CheckHandler) handleCreateCheck(w http.ResponseWriter, r *http.Request) {

}

func (h *CheckHandler) handleGetCheck(w http.ResponseWriter, r *http.Request) {

}

func (h *CheckHandler) handleUpdateCheck(w http.ResponseWriter, r *http.Request) {

}

func (h *CheckHandler) handleDeleteCheck(w http.ResponseWriter, r *http.Request) {

}
