package workfile

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

func Read_Line_Fromfile(path string, line_num int) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", errors.New("-->file not found")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
		if lineCount == line_num {
			return scanner.Text(), nil
		}
	}

	if err := scanner.Err(); err != nil {
		return "", errors.New("-->file not found")
	}

	return "", errors.New("-->line not found")
}

func WriteLineFromFile(path string, lineNum int, data string) error {
	lines, err := ReadLinesFromFile(path)
	if err != nil {
		return err
	}

	if lineNum < 1 || lineNum > len(lines) {
		return errors.New("-->table not found")
	}

	lines[lineNum-1] = data

	return WriteLinesToFile(path, lines)
}

func ReadLinesFromFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func WriteLinesToFile(path string, lines []string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	return writer.Flush()
}

func Search_Number_Table(line string, name_table string) int {
	startIndex := strings.Index(line, name_table) + len(name_table) + 2
	if startIndex > len(line) {
		return 0
	}
	endIndex := strings.Index(line[startIndex:], "}")
	valueStr := line[startIndex : startIndex+endIndex]
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return 0
	}
	return value
}

func Search_Table(filepath string, command string, name_table string) (error, string, int) {
	if command == "post" || command == "get"{
		line, err := Read_Line_Fromfile(filepath, 1)
		if err != nil {
			return errors.New("-->file not found"), "", -1
		}
		number_line := Search_Number_Table(line, "links")
		if number_line == 0 {
			return errors.New("-->table not found"), "", -1
		}
		return nil, "link", number_line
	}
	return errors.New("-->incorrect command"), "", -1
}
