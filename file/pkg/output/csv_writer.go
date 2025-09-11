// File: internal/pkg/output/csv_writer.go

package output

import (
	"github.com/huynh-fs/file/internal/model"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const outputFileName = "bingo_result.csv"

func WriteToCSV(data *model.ResultData) error {
	file, err := os.Create(outputFileName)
	if err != nil {
		return fmt.Errorf("không thể tạo file: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, row := range data.InitialTicket {
		if err := writeRow(writer, row[:]); err != nil { return err }
	}

	calledStrs := make([]string, len(data.CalledNumbers))
	for i, num := range data.CalledNumbers {
		calledStrs[i] = strconv.Itoa(num)
	}
	if err := writer.Write([]string{strings.Join(calledStrs, " ")}); err != nil {
		return err
	}

	if err := writer.Write([]string{data.WinLine}); err != nil {
		return err
	}

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