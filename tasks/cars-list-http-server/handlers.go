package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

var cars = map[string]string{
	"id1": "Renault Logan",
	"id2": "Renault Duster",
	"id3": "BMW X6",
	"id4": "BMW M5",
	"id5": "VW Passat",
	"id6": "VW Jetta",
	"id7": "Audi A4",
	"id8": "Audi Q7",
}

// carsListFunc — вспомогательная функция для вывода всех машин.
func carsListFunc() []string {
	var list []string
	for _, c := range cars {
		list = append(list, c)
	}
	return list
}

// carFunc — вспомогательная функция для вывода определённой машины.
func carFunc(id string) string {
	if c, ok := cars[id]; ok {
		return c
	}
	return "unknown identifier " + id
}

func carsHandle(rw http.ResponseWriter, r *http.Request) {
	carsList := carsListFunc()
	io.WriteString(rw, strings.Join(carsList, ", "))
}

func carHandle(rw http.ResponseWriter, r *http.Request) {
	// net/http
	// carID := r.URL.Query().Get("id")
	carID := chi.URLParam(r, "id")
	// при запросе "/car" вернётся ошибка 404,
	// поэтому не нужно проверять id на пустоту
	// 	http.Error(rw, "carID param is missed", http.StatusBadRequest)
	// 	return
	// }
	rw.Write([]byte(carFunc(carID)))
}

func brandHandle(rw http.ResponseWriter, r *http.Request) {
	list := make([]string, 0)

	brand := strings.ToLower(chi.URLParam(r, "brand"))

	for _, c := range cars {
		bm := strings.Split(strings.ToLower(c), " ")
		if bm[0] == brand {
			list = append(list, c)
		}
	}

	io.WriteString(rw, strings.Join(list, ", "))
}

func modelHandle(rw http.ResponseWriter, r *http.Request) {
	car := fmt.Sprintf("%s %s", chi.URLParam(r, "brand"), chi.URLParam(r, "model"))

	for _, c := range cars {
		if strings.EqualFold(c, car) {
			io.WriteString(rw, c)
			return
		}
	}

	http.Error(rw, "unknown model: "+car, http.StatusNotFound)

}
