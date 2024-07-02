package main

type Application struct {
}

func NewApplication() *Application {
	return &Application{}
}

func (a *Application) handleRequest(url string, method string) (int, string) {
	if url == "/app/status" && method == "GET" {
		return 200, "OK"
	}

	if url == "/app/users" && method == "POST" {
		return 201, "User created"
	}

	return 404, "Not found"
}
