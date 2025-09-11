package handler

import (
	"github.com/huynh-fs/file/internal/model"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const outputFileName = "bingo_result.csv"

func WriteResult(data *model.ResultData) error {
	file, err := os.Create(outputFileName)
	if err != nil {
		return fmt.Errorf("không thể tạo file: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// ghi tấm vé ban đầu
	for _, row := range data.InitialTicket {
		if err := writeRow(writer, row[:]); err != nil { return err }
	}

	// ghi các số đã gọi
	calledStrs := make([]string, len(data.CalledNumbers))
	for i, num := range data.CalledNumbers {
		calledStrs[i] = strconv.Itoa(num)
	}
	if err := writer.Write([]string{strings.Join(calledStrs, " ")}); err != nil {
		return err
	}

	// ghi dòng thắng
	if err := writer.Write([]string{data.WinLine}); err != nil {
		return err
	}

	// ghi tấm vé cuối cùng
	for _, row := range data.FinalTicket {
		if err := writeRow(writer, row[:]); err != nil { return err }
	}

	fmt.Printf("Đã ghi kết quả ra file '%s'\n", outputFileName)
	return writer.Error()
}

func writeRow(writer *csv.Writer, row []int) error {
	record := make([]string, len(row))
	for i, val := range row {
		record[i] = strconv.Itoa(val)
	}
	return writer.Write(record)
}