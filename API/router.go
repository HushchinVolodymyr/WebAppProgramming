<<<<<<< HEAD
package main
=======
<<<<<<< HEAD
package main
=======
package Services
>>>>>>> c3690e10a22f1be29bd25231a34a1d186b2e7f10
>>>>>>> 5092a91f6602365a02fcd88947a41393efc7501c

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
<<<<<<< HEAD
=======
<<<<<<< HEAD
>>>>>>> 5092a91f6602365a02fcd88947a41393efc7501c

	// Task1 Controller Routes Part 2 (Handle POST request, other methods not allowed)
	rs.mux.HandleFunc("/task1/part2", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			Controllers.PostTask1Part2Controller(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
<<<<<<< HEAD
=======
=======
>>>>>>> c3690e10a22f1be29bd25231a34a1d186b2e7f10
>>>>>>> 5092a91f6602365a02fcd88947a41393efc7501c
}

// GetHandler returns the handler for the RouterService
func (rs *RouterService) GetHandler() http.Handler {
	return Middleware.LoggingMiddleware(rs.mux)
}
