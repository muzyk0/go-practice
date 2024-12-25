package main

import (
	"flag"
	"fmt"
	"os"
)

// go run main.go -thumb -w 1024 -file "./in/img0001.jpg" -dest "/home/user/images"
// func main() {
// 	// первый аргумент — имя запущенного файла
// 	fmt.Printf("Command: %v\n", os.Args[0])
// 	// выведем остальные параметры
// 	for i, v := range os.Args[1:] {
// 		fmt.Println(i+1, v)
// 	}
// }

// go run . -file /home/user/input/img001.png
// func main() {
// 	// указываем имя флага, значение по умолчанию и описание
// 	imgFile := flag.String("file", "", "input image file")
// 	// делаем разбор командной строки
// 	flag.Parse()
// 	fmt.Println("Image file:", *imgFile)
// }

// go run . -thumb -w 1024 -file "./in/img0001.jpg" -dest "/home/user/images"
// func main() {
// 	imgFile := flag.String("file", "", "input image file")
// 	destDir := flag.String("dest", "./output", "destination folder")
// 	width := flag.Int("w", 1024, "width of the image")
// 	isThumb := flag.Bool("thumb", false, "create thumb")

// 	// разбор командной строки
// 	flag.Parse()
// 	fmt.Println("Image file:", *imgFile)
// 	fmt.Println("Destination folder:", *destDir)
// 	fmt.Println("Width:", *width)
// 	fmt.Println("Thumbs:", *isThumb)
// }

// var (
// 	width *int
// 	thumb *bool
// )

// func init() {
// 	// используем init-функцию
// 	width = flag.Int("width", 1024, "width of the image")
// 	thumb = flag.Bool("thumb", false, "create thumb")
// }

// // или сразу используем глобальные переменные
// // var (
// //     width = flag.Int("width", 1024, "width of the image")
// //     thumb = flag.Bool("thumb", false, "create thumb")
// // )

// func main() {

// 	// разбор командной строки
// 	flag.Parse()

// 	fmt.Println("Width:", *width)
// 	fmt.Println("Thumbs:", *thumb)
// }
// **********************************************************************
// go run . "./in/img0001.jpg" "./in/img0002.jpg"
// Out:
/*
Image file (0):
%!(EXTRA string=./in/img0001.jpg)Image file (1):
%!(EXTRA string=./in/img0002.jpg)Destination folder: ./output
Width: 1024
Thumbs: false
*/
// func main() {
// 	destDir := flag.String("dest", "./output", "destination folder")
// 	width := flag.Int("w", 1024, "width of the image")
// 	isThumb := flag.Bool("thumb", false, "create thumb")

//		flag.Parse()
//		// получаем список файлов
//		for i, v := range flag.Args() {
//			fmt.Printf("Image file (%d):\r\n", i, v)
//		}
//		fmt.Println("Destination folder:", *destDir)
//		fmt.Println("Width:", *width)
//		fmt.Println("Thumbs:", *isThumb)
//	}
//
// **********************************************************************
// Наборы флагов
// func main() {
// 	// декларируем наборы флагов для подкоманд
// 	cnvFlags := flag.NewFlagSet("cnv", flag.ExitOnError)
// 	filterFlags := flag.NewFlagSet("filter", flag.ExitOnError)
// 	// декларируем флаги набора cnvFlags
// 	destDir := cnvFlags.String("dest", "./output", "destination folder")
// 	width := cnvFlags.Int("w", 1024, "width of the image")
// 	isThumb := cnvFlags.Bool("thumb", false, "create thumb")

//		// флаги набора filterFlags
//		isGray := filterFlags.Bool("gray", false, "convert to grayscale")
//		isSepia := filterFlags.Bool("sepia", false, "convert to sepia")
//		// проверяем, задана ли подкоманда
//		// os.Arg[0] имя команды
//		// os.Arg[1] имя подкоманды
//		if len(os.Args) < 2 {
//			fmt.Println("set or get subcommand required")
//			os.Exit(1)
//		}
//		// в зависимости от переданной подкоманды
//		// делаем парсинг флагов соответствующего набора,
//		// передаём функции FlagSet.Parse() аргументы командной строки
//		// os.Args[2:] содержит все аргументы,
//		// следующие за os.Args[1], за именем подкоманды
//		switch os.Args[1] {
//		case "cnv":
//			cnvFlags.Parse(os.Args[2:])
//		case "filter":
//			filterFlags.Parse(os.Args[2:])
//		default:
//			// PrintDefaults выводит параметры командной строки
//			flag.PrintDefaults()
//			os.Exit(1)
//		}
//		// проверяем, какой набор флагов использовался,
//		// то есть какая подкоманда была передана,
//		// функция FlagSet.Parsed() возвращает false, если
//		// парсинг флагов набора не проводился
//		if cnvFlags.Parsed() {
//			// логика для img cnv
//		}
//		if filterFlags.Parsed() {
//			// логика для img filter
//		}
//	}
//
// **********************************************************************
// Пользовательские правила обработки
// func main() {
// 	// готовим переменную для аргументов
// 	var effects []string
// 	// декларируем функцию-обработчик
// 	flag.Func("effects", "Rotation and mirror", func(flagValue string) error {
// 		// разбиваем значение флага на слайс строк через запятую
// 		// и заливаем в переменную
// 		effects = strings.Split(flagValue, ",")
// 		return nil
// 	})
// 	// запускаем парсинг
// 	flag.Parse()
// }
// **********************************************************************
// Интерфейс пользовательской обработки

// type NetAddress struct {
// 	Host string
// 	Port int
// }

// // допишите код реализации методов интерфейса
// // ...

// func (n *NetAddress) String() string {
// 	return n.Host + ":" + strconv.Itoa(n.Port)
// }

// func (n *NetAddress) Set(flagValue string) error {
// 	args := strings.Split(flagValue, ":")
// 	if len(args) != 2 {
// 		return errors.New(`need address in a format "host:port"`)
// 	}

// 	port, err := strconv.Atoi(args[1])
// 	if err != nil {
// 		return err
// 	}

// 	n.Host = args[0]
// 	n.Port = port

// 	return nil
// }

// func main() {
// 	addr := new(NetAddress)
// 	// если интерфейс не реализован,
// 	// здесь будет ошибка компиляции
// 	_ = flag.Value(addr)
// 	// проверка реализации
// 	flag.Var(addr, "addr", "Net address host:port")
// 	flag.Parse()
// 	fmt.Println(addr.Host)
// 	fmt.Println(addr.Port)
// }

// **********************************************************************

var version = "0.0.1"

func main() {
	// допишите код
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "App version: %v\nUsage of %s:\n", version, os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()
}
