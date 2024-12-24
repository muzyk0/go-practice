package main

import (
	"net/http"
	"net/http/cookiejar"
)

// func main() {
// 	client := &http.Client{}
// 	req, err := http.NewRequest(http.MethodGet, `http://localhost:8080`, nil)
// 	if err != nil {
// 		panic(err)
// 	}
// 	req.AddCookie(&http.Cookie{
// 		Name:  "ID",
// 		Value: "3675",
// 	})
// 	req.AddCookie(&http.Cookie{
// 		Name:   "Token",
// 		Value:  "TEST_TOKEN",
// 		MaxAge: 360,
// 	})
// 	_, err = client.Do(req)
// if err != nil {
// 	panic(err)
// }
// }

func main() {
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}
	client := &http.Client{
		Jar: jar,
	}

	req, err := http.NewRequest(http.MethodGet, `http://localhost:8080`, nil)
	if err != nil {
		panic(err)
	}
	cookie := &http.Cookie{
		Name:   "Token",
		Value:  "TEST_TOKEN",
		MaxAge: 300,
	}
	req.AddCookie(cookie)
	_, err = client.Do(req)
	if err != nil {
		panic(err)
	}
}
