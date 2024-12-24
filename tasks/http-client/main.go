package main

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
)

// Native net/http
// var bearer = "Bearer <Token>"
// func main() {
//     // создаём новый запрос
//     req, err := http.NewRequest("GET", "https://yandex.ru", nil)
//     if err != nil {
//         log.Println(err)
//         return
//     }

//     // добавляем авторизацию
//     req.Header.Add("Authorization", bearer)

//     // создаём клиент
//     client := &http.Client{}
//     resp, err := client.Do(req)
//     if err != nil {
//         log.Println("Error on response.\n[ERROR] -", err)
//         return
//     }
//     defer resp.Body.Close()

//     if resp.StatusCode != http.StatusOK {
//         log.Println("Bad status code on response: ", resp.StatusCode)
//         return
//     }

//     body, err := io.ReadAll(resp.Body)
//     // продолжаем работу
//     fmt.Println(body)
// }

// retry v2 client

// MyApiError — описание ошибки при неверном запросе.
// type MyApiError struct {
// 	Code      int       `json:"code"`
// 	Message   string    `json:"message"`
// 	Timestamp time.Time `json:"timestamp"`
// }

// // Post — модель, описание основного объекта.
// type Post struct {
// 	UserID int    `json:"userId"`
// 	ID     int    `json:"id"`
// 	Title  string `json:"title"`
// 	Text   string `json:"text"`
// }

// func main() {
// 	client := resty.New()

// 	client.
// 	// устанавливаем количество повторений
// 	SetRetryCount(3).
// 	// длительность ожидания между попытками
// 	SetRetryWaitTime(30 * time.Second).
// 	// длительность максимального ожидания
// 	SetRetryMaxWaitTime(90 * time.Second)

// 	var responseErr MyApiError
// 	var post Post

// 	_, err := client.R().
// 		SetError(&responseErr).
// 		SetResult(&post).
// 		SetPathParams(map[string]string{
// 			"postID": "2",
// 		}).
// 		Get("https://jsonplaceholder.typicode.com/posts/{postID}")

// 	if err != nil {
// 		fmt.Println(responseErr)
// 		panic(err)
// 		return
// 	}

// 	fmt.Println(post)
// }

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func main() {
	var users []User
	url := "https://jsonplaceholder.typicode.com/users"

	client := resty.New()
	// если выбрали resty, используйте SetResult(&users)
	// для получения результата сразу в виде массива
	// ...

	_, err := client.R().SetResult(&users).Get(url)
	if err != nil {
		panic(err)
	}

	// var usersStr string
	// for _, v := range users {
	// 	usersStr += fmt.Sprintf("%v, ", v.Username)
	// }

	// fmt.Print(usersStr)

	//or

	var out []string
	for _, v := range users {
		out = append(out, v.Username)
	}

	fmt.Println(strings.Join(out, ", "))
}
