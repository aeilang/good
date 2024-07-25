package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/aeilang/backend/internal/server"
	_ "github.com/lib/pq"
)

func main() {
	connStr := "postgres://lang:password@localhost:5432/test_db?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil || db.Ping() != nil {
		panic("db connecting failed!")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	queries, err := store.Prepare(ctx, db)
	if err != nil {
		panic("create queries failed")
	}

	serv := server.New(queries)

	r := http.NewServeMux()
	r.HandleFunc("GET /users", serv.HandleGetUsers)
	r.HandleFunc("GET /user/{id}", serv.HandleGetUser)
	r.HandleFunc("DELETE /user/{id}", serv.HandleDeleteUser)
	r.HandleFunc("POST /user", serv.HandleCreateUser)
	r.HandleFunc("PUT /user", serv.HandleUpdateUser)

	log.Println("listen to port 8888")
	if err := http.ListenAndServe(":8888", r); err != nil {
		log.Println(err)
	}
}
