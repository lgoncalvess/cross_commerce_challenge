package main

import (
	"fmt"

	"github.com/labstack/echo"
)

//Inicializando o servidor e chamando a unica rota
func main() {
	server := echo.New()
	Router(server)

	port := "3000"

	address := fmt.Sprintf("%s:%s", "0.0.0.0", port)
	fmt.Println(address)
	server.Start(address)
}
