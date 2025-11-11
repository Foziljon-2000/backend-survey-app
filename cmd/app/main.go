package main

import (
	"backend-survey-app/pkg/databases"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	db := databases.DB()
	defer db.Close()

	if err := db.Ping(); err != nil {
		panic("не удалось подключиться: " + err.Error())
	}
	fmt.Println("Подключение успешно!")
}
