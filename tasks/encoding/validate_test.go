package main

import (
	"reflect"
	"strconv"
	"strings"
	"testing"
)

// User используется для тестирования.
type User struct {
	Nick string
	Age  int `limit:"18"`
	Rate int `limit:"0,100"`
}

// Str2Int конвертирует строку в int.
func Str2Int(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return v
}

// Validate проверяет min и max для int c тегом limit.
func Validate(obj interface{}) bool {
	vobj := reflect.ValueOf(obj)
	objType := vobj.Type() // получаем описание типа

	// перебираем все поля структуры
	for i := 0; i < objType.NumField(); i++ {
		// берём значение текущего поля и проверяем, что это int
		if v, ok := vobj.Field(i).Interface().(int); ok {
			// подсказка: тег limit надо искать в поле objType.Field(i)
			// objType.Field(i).Tag.Lookup или objType.Field(i).Tag.Get

			if limitTag, ok := objType.Field(i).Tag.Lookup("limit"); ok {
				minMaxValues := strings.Split(limitTag, ",")

				if v < Str2Int(minMaxValues[0]) {
					return false
				}

				if len(minMaxValues) >= 2 && v > Str2Int(minMaxValues[1]) {
					return false
				}
			}
		}
	}
	return true
}

func TestValidate(t *testing.T) {
	var table = []struct {
		name string
		age  int
		rate int
		want bool
	}{
		{"admin", 20, 88, true},
		{"su", 45, 10, true},
		{"", 16, 5, false},
		{"usr", 24, -2, false},
		{"john", 18, 0, true},
		{"usr2", 30, 200, false},
	}
	for _, v := range table {
		if Validate(User{v.name, v.age, v.rate}) != v.want {
			t.Errorf("Не прошла проверку запись %s", v.name)
		}
	}
}
