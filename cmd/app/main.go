package main

import (
	"backend-survey-app/internal/storage"
	"backend-survey-app/internal/transport/http/router"
	"backend-survey-app/pkg/databases"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	db := databases.DB()
	defer db.Close()

	storage.GetDB(db)

	fmt.Println("listening localhost:2025")
	http.ListenAndServe("localhost:2025", router.NewRouterCompl())
}
