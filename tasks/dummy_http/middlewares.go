package main

import "net/http"

type Middleware func(http.Handler) http.Handler

func Conveyor(h http.Handler, middlewares ...Middleware) http.Handler {
	for _, middleware := range middlewares {
		h = middleware(h)
	}
	return h
}

// middleware принимает параметром Handler и возвращает тоже Handler.
func middleware(next http.Handler) http.Handler {
	// получаем Handler приведением типа http.HandlerFunc
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// здесь пишем логику обработки
		// например, разрешаем запросы cross-domain
		// w.Header().Set("Access-Control-Allow-Origin", "*")
		// ...
		// замыкание: используем ServeHTTP следующего хендлера
		next.ServeHTTP(w, r)
	})
}
