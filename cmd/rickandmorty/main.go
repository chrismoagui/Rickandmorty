package main

import (
	"fmt"
	"gomysql/internal/data"
	"gomysql/internal/server"
	"log"
	"os"
	"os/signal"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	//log.Fatal(http.ListenAndServe(":8080", nil))
	fmt.Println(("API RICK AND MORTY"))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	serv, err := server.New(port)
	if err != nil {
		log.Fatal(err)
	}

	//conexion a la base de datos
	d := data.New()
	if err := d.DB.Ping(); err != nil {
		log.Fatal(err)
	}

	go serv.Start()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	serv.Close()
	data.Close()
}
