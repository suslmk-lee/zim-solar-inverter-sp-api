package main

import (
	"github.com/rs/cors"
	"log"
	"net/http"
	"zim-solar-inverter-sp-api/config"
	"zim-solar-inverter-sp-api/handler"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/data", handler.DataHandler)

	// config에서 AllowedOrigins 불러오기
	origins := config.GetAllowedOrigins()
	if origins == nil {
		log.Fatal("AllowedOrigins not defined in config")
	}

	// CORS 설정 추가
	c := cors.New(cors.Options{
		AllowedOrigins:   origins,
		AllowCredentials: true,
	})
	handlerWithCors := c.Handler(mux)

	log.Println("Starting server on :8080")
	http.ListenAndServe(":8080", handlerWithCors)
}
