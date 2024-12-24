package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// func main() {
// 	// определяем хендлер, который выводит все машины
// 	http.HandleFunc("/cars", carsHandle)
// 	// определяем хендлер, который выводит определённую машину
// 	http.HandleFunc("/car", carHandle)

// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(TimerTrace)
	// или
	// r.Use(middleware.RealIP, middleware.Logger, middleware.Recoverer)

	// r.Route("/cars", func(r chi.Router) {
	// 	r.Get("/", carsHandle)
	// 	r.Route("/{brand}", func(r chi.Router) {
	// 		r.Get("/", brandHandle)
	// 		r.Route("/{model}", func(r chi.Router) {
	// 			r.Get("/", modelHandle)

	// 		})
	// 	})
	// })

	// r.Get("/", func(rw http.ResponseWriter, r *http.Request) {
	// 	rw.Write([]byte("chi"))
	// })
	// r.Get("/item/{id}", func(rw http.ResponseWriter, r *http.Request) {
	// 	// получаем значение URL-параметра id
	// 	id := chi.URLParam(r, "id")
	// 	io.WriteString(rw, fmt.Sprintf("item = %s", id))
	// })

	// // определяем хендлер, который выводит все машины
	// r.Get("/cars", carsHandle)
	// // определяем хендлер, который выводит определённую машину
	// r.Get("/car/{id}", carHandle)

	//  можно объединять обработчики запросов в древовидную структуру методом Route()
	// r.Get("/cars", carsHandle)                  // GET /cars
	// r.Get("/cars/{brand}", brandHandle)         // GET /cars/renault
	// r.Get("/cars/{brand}/{model}", modelHandle) // GET /cars/renault/duster

	// // то же самое можно описать, используя Route
	// r.Route("/cars", func(r chi.Router) {
	// 	r.Get("/", carsHandle) // GET /cars
	// 	// Route можно вкладывать один в другой
	// 	r.Route("/{brand}", func(r chi.Router) {
	// 		r.Get("/", brandHandle)        // GET /cars/renault
	// 		r.Get("/{model}", modelHandle) // GET /cars/renault/duster
	// 	})
	// })

	// Ещё Route() позволяет группировать разные HTTP-методы для одного и того же запроса:
	// r.Post("/car", newCar)           // POST /car
	// r.Get("/car/{id}", getCar)       // GET /car/1234
	// r.Put("/car/{id}", updateCar)    // PUT /car/1234
	// r.Delete("/car/{id}", deleteCar) // DELETE /car/1234

	// // то же самое, используя Router
	// r.Route("/car", func(r chi.Router) {
	// 	r.Post("/", newCar) // POST /car
	// 	r.Route("/{id}", func(r chi.Router) {
	// 		r.Get("/", getCar)       // GET /car/1234
	// 		r.Put("/", updateCar)    // PUT /car/1234
	// 		r.Delete("/", deleteCar) // DELETE /car/1234
	// 	})
	// })

	// r передаётся как http.Handler
	http.ListenAndServe(":8080", CarRouter())
}
