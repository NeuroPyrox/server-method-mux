package methodmux

import (
	"fmt"
	"net/http"
)

type Mux struct {
	Getter,
	Header,
	Poster,
	Putter,
	Patcher,
	Deleter,
	Connecter,
	Optioner,
	Tracer http.HandlerFunc
}

func (m *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if method := m.GetMethod(r.Method); method != nil {
		method(w, r)
	} else {
		methodNotAllowed(w, r)
	}
}

func (m *Mux) GetMethod(method string) http.HandlerFunc {
	switch method {
	case http.MethodGet:
		return m.Getter
	case http.MethodHead:
		return m.Header
	case http.MethodPost:
		return m.Poster
	case http.MethodPut:
		return m.Putter
	case http.MethodPatch:
		return m.Patcher
	case http.MethodDelete:
		return m.Deleter
	case http.MethodConnect:
		return m.Connecter
	case http.MethodOptions:
		return m.Optioner
	case http.MethodTrace:
		return m.Tracer
	default:
		return nil
	}
}

func methodNotAllowed(w http.ResponseWriter, r *http.Request) {
	err := fmt.Sprintf("Method not allowed: %v", r.Method)
	http.Error(w, err, http.StatusMethodNotAllowed)
}
