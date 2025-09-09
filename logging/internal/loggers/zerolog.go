package loggers

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func DemonstrateZerolog() {
	fmt.Println("--- Bắt đầu minh họa Zerolog ---")
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	log.Info().
		Str("service", "my-app").
		Int("user_id", 12345).
		Msg("Người dùng đã đăng nhập thành công")

	log.Warn().
		Str("component", "database").
		Str("error_code", "DB_CONN_TIMEOUT").
		Msg("Kết nối đến database bị timeout")

	fmt.Println("Output của Zerolog là structured log (JSON), rất phù hợp cho máy móc phân tích.")
	fmt.Println("--- Kết thúc minh họa Zerolog ---")
}