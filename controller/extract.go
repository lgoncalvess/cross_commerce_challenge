package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

//Utilizando rotinas sincronizadas para as chamadas a API externa
func Extract() []float64 {
	//WaitGroup está sendo utilizado para sincronizar as Go Routines abaixo
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	var dataGroup1 []float64
	go func() {
		dataGroup1 = callCrossCommerceApi(1)
		waitGroup.Done()
	}()

	var dataGroup2 []float64
	go func() {
		dataGroup2 = callCrossCommerceApi(2)
		waitGroup.Done()
	}()
	waitGroup.Wait()

	unorderedNumbers := append(dataGroup1, dataGroup2...)
	return unorderedNumbers
}

//função resposável para realizar a chamada externa
func callCrossCommerceApi(pageNumber int) (totalNumbers []float64) {
	var mapBody = map[string][]float64{}
	for {
		endpoint := "http://challenge.dienekes.com.br/api/numbers?page=" + strconv.FormatInt(int64(pageNumber), 10)
		response, err := http.Get(endpoint)

		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}

		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		json.Unmarshal([]byte(responseData), &mapBody)

		numbers := mapBody["numbers"]
		if len(numbers) == 0 {
			break
		} else {
			totalNumbers = append(totalNumbers, numbers...)
			pageNumber += 2
		}
	}
	return totalNumbers
}
