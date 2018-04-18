package conn

import (
	"database/sql"
	"fmt"
	"log"
	// pq
	_ "github.com/lib/pq"
)

// DatabaseConfig :: configure database connection
type DatabaseConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Driver   string `json:"driver"`
	Database string `json:"database"`
}

// DB .. save connection
var DB *sql.DB

// InitDB ..
func InitDB(cfg DatabaseConfig) {
	log.Println("Initalizing database")

	connString := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		cfg.UserName, cfg.Password, cfg.Host, cfg.Database)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatalln(err)
		return
	}
	DB = db

	log.Println("Database succesfully connected")
}
