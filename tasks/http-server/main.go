package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{
		Timeout: time.Second * 60,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println(req.URL)
			return nil
		},
	}
	response, err := client.Get("http://ya.ru")

	if err != nil {
		fmt.Println(err)
		return
	}

	// io.Discard выступает в качестве приёмника ненужных данных
	_, err = io.Copy(io.Discard, response.Body)
	response.Body.Close()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(response)
}

// func main() {
// 	client := &http.Client{
// 		Timeout: time.Second * 60,
// 		CheckRedirect: func(req *http.Request, via []*http.Request) error {
// 			fmt.Println(req.URL)
// 			return nil
// 		},
// 	}

// 	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080", nil)

// 	req.Header.Set(`MyHeader`, "Hello")
// 	req.Header.Add(`MyHeader`, "Привет")

// 	if err != nil {
// 		panic(err)
// 	}
// 	response, err := client.Do(req)
// 	if err != nil {
// 		panic(err)
// 	}
// 	io.Copy(os.Stdout, response.Body) // вывод ответа в консоль
// 	response.Body.Close()
// }

// json
// func main() {
// 	client := &http.Client{}

// 	var body = []byte(`{"message":"Hello"}`)
// 	request, err := http.NewRequest(http.MethodPost, "http://localhost:8080", bytes.NewBuffer(body))
// 	if err != nil {
// 		// обрабатываем ошибку
// 	}
// 	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
// 	_, err = client.Do(request)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// "multipart/form-data"
// func main() {
// 	client := &http.Client{}

// 	file, _ := os.Open("main.go") // открываем файл
// 	defer file.Close()            // не забываем закрыть
// 	body := &bytes.Buffer{}       // создаём буфер
// 	// на основе буфера конструируем multipart.Writer из пакета mime/multipart
// 	writer := multipart.NewWriter(body)
// 	// готовим форму для отправки файла на сервер
// 	part, err := writer.CreateFormFile("uploadfile", filename)
// 	if err != nil {
// 		// обрабатываем ошибку
// 	}
// 	// копируем файл в форму
// 	// multipart.Writer отформатирует данные и запишет в предоставленный буфер
// 	_, err = io.Copy(part, file)
// 	if err != nil {
// 		// обрабатываем ошибку
// 	}
// 	writer.Close()

// 	// пишем запрос
// 	request, err := http.NewRequest(http.MethodPost, url, body)
// 	if err != nil {
// 		// обрабатываем ошибку
// 	}
// 	// добавляем заголовок запроса
// 	request.Header.Set("Content-Type", writer.FormDataContentType())
// 	response, err := client.Do(request)
// }

// "application/x-www-form-urlencoded"
// func main() {
// 	client := &http.Client{}

// 	// готовим контейнер для данных
// 	// используем тип url.Values из пакета net/url
// 	data := url.Values{}
// 	// устанавливаем данные
// 	data.Set("key1", "value1")
// 	data.Set("key2", "value2")
// 	// пишем запрос
// 	request, err := http.NewRequest(http.MethodPost, "http://localhost:8080", strings.NewReader(data.Encode()))
// 	if err != nil {
// 		// обрабатываем ошибку
// 	}
// 	// устанавливаем заголовки
// 	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
// 	request.Header.Set("Content-Length", strconv.Itoa(len(data.Encode())))
// 	response, err := client.Do(request)
// }
