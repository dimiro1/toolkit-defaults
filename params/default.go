package params

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Default struct{}

func (Default) ByName(r *http.Request, name string) string {
	vars := mux.Vars(r)
	if len(vars) == 0 {
		return ""
	}

	return vars[name]
}


func New() Default {
	return Default{}
}