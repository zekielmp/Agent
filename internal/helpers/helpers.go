package helpers

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		return err
	}
	return nil

}

func ReadJSON(w http.ResponseWriter, r http.Request, data any) error {
	r.Body = http.MaxBytesReader(w, r.Body, 1<<24)

	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		_ = WriteJSON(w, http.StatusBadRequest, "invalid credentials")
		return err
	}

	return nil
}
