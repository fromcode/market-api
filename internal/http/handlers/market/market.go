package market

import "net/http"

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) { //mengatur handler pada url "/"
		w.Write([]byte("Selamat datang di Market API"))
	}
}
