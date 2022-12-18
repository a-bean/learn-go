package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func Add(a, b int) int {
	total := a + b
	return total
}
func main() {

	http.HandleFunc("/add", func(writer http.ResponseWriter, request *http.Request) {
		_ = request.ParseForm()
		fmt.Println("request:", request)
		a, _ := strconv.Atoi(request.Form["a"][0])
		b, _ := strconv.Atoi(request.Form["b"][0])
		fmt.Println(Add(a, b))
		writer.Header().Set("Content-Type", "application/json")
		data, _ := json.Marshal(map[string]int{
			"data": Add(a, b),
		})
		_, _ = writer.Write(data)
	})
	http.ListenAndServe(":8003", nil)
}
