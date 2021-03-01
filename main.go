package main

import (
	"context"
	"ecommerce-auth/db/postgresql"
	"ecommerce-auth/handlers"
	"flag"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"net/http"
	"os"
)


func main() {
	addr := flag.String("addr", ":8000", "HTTP network address")
	flag.Parse()
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	pool, err := pgxpool.Connect(context.Background(), "user=postgres password=agahan02 host=localhost port=5433 dbname=GoLangExamples sslmode=disable pool_max_conns=50")
	if err != nil {
		log.Fatalf("Unable to connection to database: %v\n", err)
	}
	defer pool.Close()

	app := &application{
		errorLog,
		infoLog,
		&postgresql.UserModel{Pool: pool},
	}
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(), // Call the new app.routes() method
	}

	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	snippets *postgresql.UserModel
}

func (a application) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/signin", handlers.Signin)
	mux.HandleFunc("/welcome", handlers.Welcome)
	mux.HandleFunc("/refresh", handlers.Refresh)

	return mux
}