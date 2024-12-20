package main

import "net/http"

// Redirect
//func main() {
//	fs := http.FileServer(http.Dir("./static"))
//
//	http.HandleFunc("/json/", func(w http.ResponseWriter, r *http.Request) {
//		http.ServeFile(w, r, "./static/json.json")
//	})
//
//	err := http.ListenAndServe(":8080", http.StripPrefix("/assets/", fs))
//	if err != nil {
//		panic(err)
//	}
//}

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir(".."))
	mux.Handle("/golang/", http.StripPrefix("/golang/", fs))
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./main.go")
	})

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
