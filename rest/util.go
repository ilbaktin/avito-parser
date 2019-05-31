package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func setJsonHeadersHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func writeObjectToResp(w http.ResponseWriter, iFace interface{}) {
	jsonBytes, err := json.MarshalIndent(iFace, "", "\t")
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("%s: %v", INTERNAL_ERROR, err)))
		return
	}
	w.WriteHeader(200)
	w.Write(jsonBytes)
}