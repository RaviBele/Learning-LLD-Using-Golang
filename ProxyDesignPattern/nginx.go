package main

type Nginx struct {
	application            *Application
	maxAllowedRequestCount int
	rateLimiter            map[string]int
}

func NewNginxServer() *Nginx {
	return &Nginx{
		application:            NewApplication(),
		maxAllowedRequestCount: 2,
		rateLimiter:            make(map[string]int),
	}
}

func (n *Nginx) handleRequest(url, method string) (int, string) {
	allowed := n.checkRateLimiter(url)

	if !allowed {
		return 403, "Not allowed"
	}

	return n.application.handleRequest(url, method)
}

func (n *Nginx) checkRateLimiter(url string) bool {
	if _, ok := n.rateLimiter[url]; !ok {
		n.rateLimiter[url] = 1
	} else {
		n.rateLimiter[url] += 1
	}

	if n.rateLimiter[url] > n.maxAllowedRequestCount {
		return false
	}

	return true
}
