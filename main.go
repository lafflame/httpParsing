package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("What should we parse today?\n1. YouTube\n2. VK.com\n")
	var choice int
	fmt.Scan(&choice)
	switch choice {
	case 1:
		youtube()
	case 2:
		vk()
	}
}

func youtube() {
	getForm("https://youtue.com")
}

func vk() {
	getForm("https://vk.com")
}

func getForm(address string) {
	resp, err := http.Get(address)
	if err != nil {
		fmt.Println("При подключении к серверу произошла ошибка:" + err.Error())
	}
	defer resp.Body.Close()

	statusCode := resp.StatusCode

	fmt.Println(statusCode)
}
