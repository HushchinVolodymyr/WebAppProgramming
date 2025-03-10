package main

import (
	"example.com/m/Controllers"
	"example.com/m/Middleware"
	"net/http"
)

type RouterService struct {
	mux *http.ServeMux
}

func NewRouterService() *RouterService {
	mux := http.NewServeMux()
	return &RouterService{mux: mux}
}

func (rs *RouterService) InitializeRoutes() {
	//Echo Controller Routes
	rs.mux.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			Controllers.GetEchoController(w, r)
		case http.MethodPost:
			Controllers.PostEchoController(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Task1 Controller Routes Part 1 (Handle POST request, other methods not allowed)
	rs.mux.HandleFunc("/task1/part1", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			Controllers.PostTask1Part1Controller(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Task1 Controller Routes Part 2 (Handle POST request, other methods not allowed)
	rs.mux.HandleFunc("/task1/part2", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			Controllers.PostTask1Part2Controller(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Task2 Controller Routes (Handle POST request, other methods not allowed)
	rs.mux.HandleFunc("/task2", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			Controllers.PostTask2Controller(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Task3 Controller Routes (Handle POST request, other methods not allowed)
	rs.mux.HandleFunc("/task3", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			Controllers.PostTask3Controller(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Task4 Controller Routes (Handle POST request, other methods not allowed)
	rs.mux.HandleFunc("/task4", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			Controllers.PostTask4Controller(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Task5 Controller Routes (Handle POST request, other methods not allowed)
	rs.mux.HandleFunc("/task5", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			Controllers.PostTask5Controller(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Task6 Controller Routes (Handle POST request, other methods not allowed)
	rs.mux.HandleFunc("/task6", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			Controllers.PostTask6Controller(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}

// GetHandler returns the handler for the RouterService
func (rs *RouterService) GetHandler() http.Handler {
	return Middleware.LoggingMiddleware(rs.mux)
}
