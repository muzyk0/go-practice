package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func mainPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body := fmt.Sprintf("Method: %s\r\n", r.Method)
	body += "Header ===============\r\n"
	for k, v := range r.Header {
		body += fmt.Sprintf("%s: %v\r\n", k, v)
	}

	// Эквивалентно коду получения параметров ниже
	//err := r.ParseForm()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//body += "Query parameters ===============\r\n"
	//for k, v := range r.Form {
	//	body += fmt.Sprintf("%s: %v\r\n", k, v)
	//}

	body += "Query parameters ===============\r\n"
	for k, v := range r.URL.Query() {
		body += fmt.Sprintf("%s: %v\r\n", k, v)
	}

	w.WriteHeader(http.StatusOK)

	if _, err := w.Write([]byte(body)); err != nil {
		fmt.Printf("main page error: %v", err)
	}

}

func apiPage(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hello from the API!"))
}

type Subj struct {
	Product string `json:"name"`
	Price   int    `json:"price"`
}

func JSONHandler(w http.ResponseWriter, req *http.Request) {
	// собираем данные
	//subj := Subj{"Milk",  30}
	subj := Subj{Product: "Milk", Price: 30}
	// кодируем в JSON
	resp, err := json.Marshal(subj)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	// устанавливаем заголовок Content-Type
	// для передачи клиенту информации, кодированной в JSON
	w.Header().Set("content-type", "application/json")
	// устанавливаем код 200
	w.WriteHeader(http.StatusOK)
	// пишем тело ответа
	w.Write(resp)
}

func WriteHandle(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "1")
	fmt.Fprint(w, "2")
	w.Write([]byte("3"))
	// Выведется 123
}

func rootHandle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Привет"))
}

//func main() {
//	//mux := http.NewServeMux()
//
//	//mux.HandleFunc("/", mainPage)
//	//mux.HandleFunc("/api/", apiPage)
//	//mux.HandleFunc("/json/", JSONHandler)
//	//mux.HandleFunc("/", LoginHandler)
//	//mux.HandleFunc("/", WriteHandle)
//
//	// Middlewares
//	//http.Handle("/", middleware(http.HandlerFunc(rootHandle)))
//	//http.Handle("/", Conveyor(http.HandlerFunc(rootHandle), middleware, middleware, middleware))
//
//	//err := http.ListenAndServe(":8080", mux)
//	//
//	//if err != nil {
//	//	panic(err)
//	//}
//}

func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://kagi.com", http.StatusMovedPermanently)
}

// Redirect
//func main() {
//
//	http.HandleFunc(("/search/"), redirect)
//	http.Handle("/yandex/", http.NotFoundHandler())
//	http.Handle("/google/", http.RedirectHandler("https://google.com", http.StatusMovedPermanently))
//
//	log.Fatal(http.ListenAndServe(":8080", nil))
//
//}
