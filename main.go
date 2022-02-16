package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

//Inicializando o servidor e chamando a unica rota
func main() {
	server := echo.New()
	Router(server)

	godotenv.Load()
	port := os.Getenv("PORT")

	address := fmt.Sprintf("%s:%s", "0.0.0.0", port)
	fmt.Println(address)
	server.Start(address)
}
