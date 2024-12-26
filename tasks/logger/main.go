package main

import (
	"go.uber.org/zap"
)

//func main() {
//	// создаём файл info.log и обрабатываем ошибку, если что-то пошло не так
//	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
//	if err != nil {
//		log.Fatal(err)
//	}
//	// откладываем закрытие файла
//	defer file.Close()
//
//	// устанавливаем назначение вывода в файл info.log
//	log.SetOutput(file)

// добавляем флаги для вывода даты, времени, имени файла
// log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
//	log.Print("Logging to a file in Go!")
//}

//func main() {
//	flog, err := os.OpenFile(`server.log`, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer flog.Close()
//	mylog := log.New(flog, `serv `, log.LstdFlags|log.Lshortfile)
//	mylog.Println(`Start server`)
//	mylog.Println(`Finish server`)
//
//}

//func main() {
//	var buf bytes.Buffer
//	// допишите код
//	// 1) создайте переменную типа *log.Logger
//	// 2) запишите в неё нужные строки
//
//	logger := log.New(&buf, "mylog: ", 0)
//
//	logger.Print("Hello, world!")
//	logger.Print("Goodbye")
//
//	fmt.Print(&buf)
//	// должна вывести
//	// mylog: Hello, world!
//	// mylog: Goodbye
//}

// logrus
//func main() {
//	// создаём файл info.log и обрабатываем ошибку
//	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// откладываем закрытие файла
//	defer file.Close()
//
//	// устанавливаем вывод логов в файл
//	log.SetOutput(file)
//	// устанавливаем вывод логов в формате JSON
//	log.SetFormatter(&log.JSONFormatter{})
//	// устанавливаем уровень предупреждений
//	log.SetLevel(log.WarnLevel)
//
//	// определяем стандартные поля JSON
//	log.WithFields(log.Fields{
//		"genre": "metal",
//		"name":  "Rammstein",
//	}).Info("Немецкая метал-группа, образованная в январе 1994 года в Берлине.")
//
//	log.WithFields(log.Fields{
//		"omg":  true,
//		"name": "Garbage",
//	}).Warn("В 2021 году вышел новый альбом No Gods No Masters.")
//
//	log.WithFields(log.Fields{
//		"omg":  true,
//		"name": "Linkin Park",
//	}).Fatal("Группа Linkin Park взяла паузу после смерти вокалиста Честера Беннингтона 20 июля 2017 года.")
//}

// zap
func main() {
	// добавляем предустановленный логер NewDevelopment
	logger, err := zap.NewDevelopment()
	if err != nil {
		// вызываем панику, если ошибка
		panic("cannot initialize zap")
	}
	// это нужно добавить, если логер буферизован
	// в данном случае не буферизован, но привычка хорошая
	defer logger.Sync()

	// для примера берём простой URL
	const url = "http://example.com"

	// делаем логер SugaredLogger
	sugar := logger.Sugar()

	// выводим сообщение уровня Info с парой "url": url в виде JSON, это SugaredLogger
	sugar.Infow(
		"Failed to fetch URL",
		"url", url,
	)

	// выводим сообщение уровня Info, но со строкой URL, это тоже SugaredLogger
	sugar.Infof("Failed to fetch URL: %s", url)
	// выводим сообщение уровня Error со строкой URL, и это SugaredLogger
	sugar.Errorf("Failed to fetch URL: %s", url)

	// переводим в обычный Logger
	plain := sugar.Desugar()

	// выводим сообщение уровня Info обычного регистратора (не SugaredLogger)
	plain.Info("Hello, Go!")
	// также уровня Warn (не SugaredLogger)
	plain.Warn("Simple warning")
	// и уровня Error, но добавляем строго типизированное поле "url" (не SugaredLogger)
	plain.Error("Failed to fetch URL", zap.String("url", url))
}
