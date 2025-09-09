package output

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ResultData struct {
	InitialTicket [5][5]int
	CalledNumbers []int
	Winline string
	FinalTicket [5][5]int
}

func WriteResultToCSV(filename string, data ResultData) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("không thể tạo file: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()
	
	// in vé bingo ban đầu
	for _, row := range data.InitialTicket {
		record := make([]string, len(row))
		for i, val := range row {
			record[i] = strconv.Itoa(val)
		}
		if err := writer.Write(record); err != nil {
			return err
		}
	}

	calledNumsStr := make([]string, len(data.CalledNumbers))
	for i, num := range data.CalledNumbers {
		calledNumsStr[i] = strconv.Itoa(num)
	}

	// in danh sách số đã gọi
	if err := writer.Write([]string{strings.Join(calledNumsStr, " ")}); err != nil {
		return err
	}

	// in dòng bingo thắng
	if err := writer.Write([]string{"BINGO theo " + data.Winline}); err != nil {
		return err
	}

	// in vé bingo cuối cùng
	for _, row := range data.FinalTicket {
		record := make([]string, len(row))
		for i, val := range row {
			record[i] = strconv.Itoa(val)
		}
		if err := writer.Write(record); err != nil {
			return err
		}
	}

	return nil
}