package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
	"encoding/json"
	"strconv"
	"sync"
	"path"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)


func main() {
	server := echo.New()
	server.GET(path.Join("/"), handleGet)

	godotenv.Load()
	port := os.Getenv("PORT")

	address := fmt.Sprintf("%s:%s", "0.0.0.0", port)
	fmt.Println(address)
	server.Start(address)
}

func handleGet(context echo.Context) error {
	numbers := run()
	return context.JSON(http.StatusOK,numbers)
}

func run() []float64{
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)
	
	var dataGroup1 []float64
	go func() {
		dataGroup1 = extract(1)
		waitGroup.Done()
	}()

	var dataGroup2 []float64
	go func() {
		dataGroup2 = extract(2)
		waitGroup.Done()
	}()
	waitGroup.Wait()

	unorderedNumbers := append(dataGroup1, dataGroup2...)
	orderedNumbers := mergeSort(unorderedNumbers)
	return orderedNumbers
}

func extract(pageNumber int) (totalNumbers []float64){
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

		json.Unmarshal([]byte(responseData),&mapBody)
		
		numbers := mapBody["numbers"]
		if len(numbers) == 0{
			break
		} else {
			totalNumbers = append(totalNumbers,numbers...)
			pageNumber+=2
		}
	}
	return totalNumbers
}

func mergeSort(items []float64) []float64 {
    if len(items) < 2 {
        return items
    }
    first := mergeSort(items[:len(items)/2])
    second := mergeSort(items[len(items)/2:])
    return merge(first, second)
}

func merge(a []float64, b []float64) []float64 {
    final := []float64{}
    i := 0
    j := 0
    for i < len(a) && j < len(b) {
        if a[i] < b[j] {
            final = append(final, a[i])
            i++
        } else {
            final = append(final, b[j])
            j++
        }
    }
    for ; i < len(a); i++ {
        final = append(final, a[i])
    }
    for ; j < len(b); j++ {
        final = append(final, b[j])
    }
    return final
}
