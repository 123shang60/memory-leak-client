package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

func main() {
	i := 1
	for {
		url := "http://127.0.0.1:8080/tesk"
		method := "POST"

		payload := strings.NewReader("{\n    \"username\" : \"11\",\n    \"password\" : \"33333\"\n}")

		client := &http.Client {
		}
		req, err := http.NewRequest(method, url, payload)

		if err != nil {
			fmt.Println(err)
		}
		req.Header.Add("Content-Type", "application/json")

		res, err := client.Do(req)
		res.Body.Close()
		client.Do(req)
		time.Sleep(time.Second * 5)

		fmt.Println("第",i ,"次执行成功！")
		i += 1
	}
}