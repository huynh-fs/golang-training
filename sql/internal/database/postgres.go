package database

import (
	"database/sql"
	"log"

	"github.com/huynh-fs/sql/internal/config"
	_ "github.com/lib/pq" 
)

func ConnectDB(cfg config.Config) *sql.DB {
	db, err := sql.Open("postgres", cfg.GetDSN())
	if err != nil {
		log.Fatalf("Lỗi khi mở kết nối CSDL: %v", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("Lỗi khi kiểm tra kết nối CSDL: %v", err)
	}
	log.Println("Kết nối CSDL thành công!")
	return db
}