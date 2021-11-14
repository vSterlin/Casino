package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/lib/pq"
	"github.com/vSterlin/casino/internal/config"
	"github.com/vSterlin/casino/internal/server"
)

func main() {

	if err := config.Load(); err != nil {
		fmt.Println(err.Error())
	}
	db, err := sql.Open("postgres", "user=v dbname=casino sslmode=disable")
	if err != nil {
		fmt.Println(err.Error())
	}

	db.Ping()

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatalln("please provide a valid PORT number")
	}

	s := server.NewServer(port, db)
	s.Init()
	defer s.Shutdown()

}
