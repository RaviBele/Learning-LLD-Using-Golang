package main

import "fmt"

func main() {

	server := NewNginxServer()

	appStatusURL := "/app/status"
	createUserURL := "/app/users"

	status, response := server.handleRequest(appStatusURL, "GET")
	fmt.Println("status: ", status, "response: ", response)

	status, response = server.handleRequest(appStatusURL, "GET")
	fmt.Println("status: ", status, "response: ", response)

	status, response = server.handleRequest(appStatusURL, "GET")
	fmt.Println("status: ", status, "response: ", response)

	status, response = server.handleRequest(createUserURL, "POST")
	fmt.Println("status: ", status, "response: ", response)

	status, response = server.handleRequest(createUserURL, "GET")
	fmt.Println("status: ", status, "response: ", response)
}
