package main

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"log"
)

// Native pars envs
//func main() {
//	if res, exist := os.LookupEnv("USERNAME"); exist {
//		fmt.Printf("Username env is: %s, %b", res, exist)
//	}
//	u := os.Getenv("USERNAME")
//	fmt.Println(u)
//
//	envList := os.Environ()
//	// выводим первые пять элементов
//	for i := 0; i < 5 && i < len(envList); i++ {
//		fmt.Println(envList[i])
//	}
//}

// With package github.com/caarlos0/env/v6
//type Config struct {
//	Files []string `env:"FILES" envSeparator:":"`
//	Home  string   `env:"HOME"`
//	// required требует, чтобы переменная TASK_DURATION была определена
//	TaskDuration time.Duration `env:"TASK_DURATION,required"`
//}
//
//func main() {
//	var cfg Config
//	err := env.Parse(&cfg)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	log.Println(cfg)
//}

type Config struct {
	User string `env:"USERNAME"` // укажите тег env
}

func main() {
	var cfg Config
	// допишите код

	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Current user is %s\n", cfg.User)
}
