package service_test

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"gorm.io/gorm"

	"github.com/huynh-fs/gin-api/internal/model"
	"github.com/huynh-fs/gin-api/pkg/config"
	"github.com/huynh-fs/gin-api/pkg/database"
)

// Các biến toàn cục cho cả package test
var (
	testDB *gorm.DB // Dùng cho integration tests và benchmarks
	testCfg *config.Config
)

// TestMain là hàm setup và teardown duy nhất cho toàn bộ package service.
func TestMain(m *testing.M) {
	// --- SETUP ---
	// Tải file .env.test. Nếu không có file này, các integration test/benchmark sẽ thất bại.
	if err := godotenv.Load("../../.env.test"); err != nil {
		log.Println("WARNING: .env.test file not found. Integration tests/benchmarks will fail.")
	}
	
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config for tests: %v", err)
	}
	testCfg = cfg

	// Cố gắng kết nối DB. Nếu thất bại, chỉ in cảnh báo và vẫn tiếp tục.
	// Điều này cho phép các unit test/benchmark (dùng mock) vẫn có thể chạy.
	db, err := database.Connect(cfg.DSN())
	if err != nil {
		log.Printf("WARNING: Could not connect to test database: %v. Integration tests/benchmarks will fail.", err)
	} else {
		testDB = db
		// Chỉ migrate và dọn dẹp nếu kết nối thành công
		db.AutoMigrate(&model.Todo{}, &model.User{}, &model.RefreshToken{})
		defer func() {
			log.Println("Cleaning up test database...")
			db.Migrator().DropTable(&model.Todo{}, &model.User{}, &model.RefreshToken{})
		}()
	}
	
	// --- Chạy tất cả các test và benchmark ---
	exitCode := m.Run()

	// --- TEARDOWN ---
	// Teardown được thực hiện bởi 'defer' ở trên.
	os.Exit(exitCode)
}