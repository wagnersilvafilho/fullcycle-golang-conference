package auth

import "net/http"

type Authenticated interface {
	GetID() int64
	GetName() string
}

// define func to authenticate
type authenticateFunc func(string, string) (Authenticated, error)

type handler struct {
	authenticate authenticateFunc
}

func HandleAuth(fn authenticateFunc) func(http.ResponseWriter, *http.Request) {
	h := handler{fn}

	return h.auth
}
