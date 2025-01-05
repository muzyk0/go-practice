package main

import (
	"fmt"
	"github.com/mailru/easyjson"
	myjson "tasks/os/easyjson"
)

//type MyType struct {
//	User      string    `json:"user,omitempty" example:"Bob"`
//	CreatedAt time.Time `json:"created_at"`
//}
//
//const (
//	targetField = "User" // имя поля, о котором нужно получить информацию
//	targetTag   = "json" // тег, значение которого нужно получить
//)
//
//func main() {
//
//	obj := MyType{}
//
//	// получаем Go-описание типа
//	objType := reflect.TypeOf(obj)
//
//	// ищем поле по имени
//	field, ok := objType.FieldByName(targetField)
//	if !ok {
//		panic(fmt.Errorf("field (%s): not found", targetField))
//	}
//
//	// ищем тег по имени
//	tagValue, ok := field.Tag.Lookup(targetTag)
//	if !ok {
//		panic(fmt.Errorf("tag (%s) for field (%s): not found", targetTag, targetField))
//	}
//
//	fmt.Printf("Значение тега %s поля %s: %s\n", targetTag, targetField, tagValue)
//	fmt.Printf("Теги поля %s: %s\n", targetField, field.Tag)
//}

func main() {
	balance := myjson.AccountBalance{
		AccountIdHash: []byte{0x10, 0x20, 0x0A, 0x0B},
		Amounts: []myjson.CurrencyAmount{
			{Amount: 1000000, Decimals: 2, Symbol: "RUB"},
			{Amount: 2510, Decimals: 2, Symbol: "USD"},
		},
		IsBlocked: true,
	}

	// преобразуем значение переменной balance в JSON-формат
	out, err := easyjson.Marshal(balance)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))
}
