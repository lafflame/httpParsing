package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("What should we parse today?\n1. YouTube\n2. VK.com\n3. Custom URL")
	var choice int
	fmt.Scan(&choice)
	switch choice {
	case 1:
		getForm("https://youtube.com")
	case 2:
		getForm("https://vk.com")
	case 3:
		var customUrl string
		fmt.Scan(&customUrl)
		getForm(customUrl)
	}
}

func getForm(address string) {
	resp, err := http.Get(address)
	if err != nil {
		fmt.Println("При подключении к серверу произошла ошибка")
		return
	}
	defer resp.Body.Close()

	statusCode := resp.StatusCode

	codes(statusCode)
	fmt.Println(resp.StatusCode)
	fmt.Printf("The server is %s, code:[%d]", resp.Status, resp.StatusCode)
}

// Расшифровка кодов
func codes(statusCode int) {
	switch statusCode {
	case http.StatusOK:
		fmt.Print("OK, the server is working, code: ")
	case http.StatusNotFound:
		fmt.Print("NotFound, code:")
	case http.StatusForbidden:
		fmt.Print("Forbidden, code:")
	default:
		fmt.Print("Unknown error, code:")
	}
}
