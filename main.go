package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/orololuwa/reimagined-robot/handlers"
)

func run()(*sql.DB, *chi.Mux, error){
	dbHost := "localhost"
	dbPort := "5432"
	dbName := "bookings"
	dbUser := "orololuwa"
	dbPassword := ""
	dbSSL := "disable"

	// Connecto to DB
	log.Println("Connecting to dabase")
	dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s", dbHost, dbPort, dbName, dbUser, dbPassword, dbSSL)

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal("Cannot conect to database: Dying!", err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	log.Println("Connected to database")

	handlers.NewHandler(db)
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Post("/user", handlers.Repo.CreateAUser)
	router.Get("/user/{id}", handlers.Repo.GetAUser)
	return db, router, nil
}


const portNumber = ":8080"
func main(){
	db, route, err := run()
	if (err != nil){
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))

	srv := &http.Server{
		Addr: portNumber,
		Handler: route,
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

