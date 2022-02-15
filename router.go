package main

import (
	"ccs/controller"
	"net/http"
	"path"

	"github.com/labstack/echo"
)

func Router(server *echo.Echo) {
	server.GET(path.Join("/"), handleGet)
}

//Função responsável por rotear todo o ciclo de extract, transform e retornar o resultado
func handleGet(context echo.Context) error {
	unorderedNumbers := controller.Extract()
	orderedNumbers := controller.MergeSort(unorderedNumbers)
	return context.JSON(http.StatusOK, orderedNumbers)
}
