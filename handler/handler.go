package handler

import (
	"encoding/json"
	"net/http"
	"zim-solar-inverter-sp-api/data"
)

func DataHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data := data.GenerateData()
	json.NewEncoder(w).Encode(data)
}
