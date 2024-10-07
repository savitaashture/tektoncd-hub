// Code generated by goa v3.19.1, DO NOT EDIT.
//
// admin HTTP server
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/design

package server

import (
	"context"
	"net/http"

	admin "github.com/tektoncd/hub/api/gen/admin"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
	"goa.design/plugins/v3/cors"
)

// Server lists the admin service endpoint HTTP handlers.
type Server struct {
	Mounts        []*MountPoint
	UpdateAgent   http.Handler
	RefreshConfig http.Handler
	CORS          http.Handler
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the admin service endpoints using the
// provided encoder and decoder. The handlers are mounted on the given mux
// using the HTTP verb and path defined in the design. errhandler is called
// whenever a response fails to be encoded. formatter is used to format errors
// returned by the service methods prior to encoding. Both errhandler and
// formatter are optional and can be nil.
func New(
	e *admin.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"UpdateAgent", "PUT", "/system/user/agent"},
			{"RefreshConfig", "POST", "/system/config/refresh"},
			{"CORS", "OPTIONS", "/system/user/agent"},
			{"CORS", "OPTIONS", "/system/config/refresh"},
		},
		UpdateAgent:   NewUpdateAgentHandler(e.UpdateAgent, mux, decoder, encoder, errhandler, formatter),
		RefreshConfig: NewRefreshConfigHandler(e.RefreshConfig, mux, decoder, encoder, errhandler, formatter),
		CORS:          NewCORSHandler(),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "admin" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.UpdateAgent = m(s.UpdateAgent)
	s.RefreshConfig = m(s.RefreshConfig)
	s.CORS = m(s.CORS)
}

// MethodNames returns the methods served.
func (s *Server) MethodNames() []string { return admin.MethodNames[:] }

// Mount configures the mux to serve the admin endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountUpdateAgentHandler(mux, h.UpdateAgent)
	MountRefreshConfigHandler(mux, h.RefreshConfig)
	MountCORSHandler(mux, h.CORS)
}

// Mount configures the mux to serve the admin endpoints.
func (s *Server) Mount(mux goahttp.Muxer) {
	Mount(mux, s)
}

// MountUpdateAgentHandler configures the mux to serve the "admin" service
// "UpdateAgent" endpoint.
func MountUpdateAgentHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleAdminOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("PUT", "/system/user/agent", f)
}

// NewUpdateAgentHandler creates a HTTP handler which loads the HTTP request
// and calls the "admin" service "UpdateAgent" endpoint.
func NewUpdateAgentHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeUpdateAgentRequest(mux, decoder)
		encodeResponse = EncodeUpdateAgentResponse(encoder)
		encodeError    = EncodeUpdateAgentError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "UpdateAgent")
		ctx = context.WithValue(ctx, goa.ServiceKey, "admin")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountRefreshConfigHandler configures the mux to serve the "admin" service
// "RefreshConfig" endpoint.
func MountRefreshConfigHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleAdminOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/system/config/refresh", f)
}

// NewRefreshConfigHandler creates a HTTP handler which loads the HTTP request
// and calls the "admin" service "RefreshConfig" endpoint.
func NewRefreshConfigHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeRefreshConfigRequest(mux, decoder)
		encodeResponse = EncodeRefreshConfigResponse(encoder)
		encodeError    = EncodeRefreshConfigError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "RefreshConfig")
		ctx = context.WithValue(ctx, goa.ServiceKey, "admin")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountCORSHandler configures the mux to serve the CORS endpoints for the
// service admin.
func MountCORSHandler(mux goahttp.Muxer, h http.Handler) {
	h = HandleAdminOrigin(h)
	mux.Handle("OPTIONS", "/system/user/agent", h.ServeHTTP)
	mux.Handle("OPTIONS", "/system/config/refresh", h.ServeHTTP)
}

// NewCORSHandler creates a HTTP handler which returns a simple 204 response.
func NewCORSHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(204)
	})
}

// HandleAdminOrigin applies the CORS response headers corresponding to the
// origin for the service admin.
func HandleAdminOrigin(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			h.ServeHTTP(w, r)
			return
		}
		if cors.MatchOrigin(origin, "*") {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
			if acrm := r.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT")
				w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
				w.WriteHeader(204)
				return
			}
			h.ServeHTTP(w, r)
			return
		}
		h.ServeHTTP(w, r)
		return
	})
}