package databases

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func DB() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Ошибка загрузки .env файла: %v", err)
	}

	st := "host=" + os.Getenv("DB_HOST") +
		" port=" + os.Getenv("DB_PORT") +
		" user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASS") +
		" dbname=" + os.Getenv("DB_NAME") +
		" sslmode=disable"

	db, err := sql.Open("postgres", st)
	if err != nil {
		log.Fatalf("Ошибка открытия соединения: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Не удалось подключиться к БД: %v", err)
	}

	return db
}
