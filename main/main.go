package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/rs/cors"
	"instagram-clone-server/handler"
	"instagram-clone-server/middleware"
	"instagram-clone-server/service"
	"log"
	"net/http"
	"time"
)

func main() {
	db, err := sqlx.Connect("postgres", "postgresql://admin:admin@localhost:5432/instagram?sslmode=disable")
	if err != nil {
		panic (err)
	}
	api := service.NewAPI(db)

	r := mux.NewRouter()
	r.HandleFunc("/posts", handler.AllPosts(api)).Methods("GET")
	r.HandleFunc("/posts", handler.CreatePost(api)).Methods("POST")
	r.HandleFunc("/auth/sign-up", handler.SignUp(api)).Methods("POST")
	r.HandleFunc("/auth/sign-in", handler.SignIn(api)).Methods("POST")

	r.Use(middleware.LoggerMiddleware)
	r.Use(mux.CORSMethodMiddleware(r))


	// For dev only - Set up CORS so our client can consume the api
	corsWrapper := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "PATCH", "PUT"},
		AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"},
	})

	srv := &http.Server{
		Handler: corsWrapper.Handler(r),
		Addr: "127.0.0.1:4000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}
	fmt.Println("Listening at:", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}