package workfile

import (
	"strconv"
	"strings"

	"dms/src/structs"
)

func Scan_Table_Links(filepath string, line_number int) (structs.Hash_Table, int) {
	line, err := Read_Line_Fromfile(filepath, line_number)
	if err != nil {
		return structs.Hash_Table{}, -1
	}
	startIndex := strings.Index(line, "[") + 1
	endIndex := strings.Index(line, "}")
	sizeIndex := strings.Index(line, ",")
	size := line[startIndex:sizeIndex]
	val := line[sizeIndex+len(size)+2 : endIndex]
	val = strings.ReplaceAll(val, ")", "")
	val = strings.ReplaceAll(val, "(", "")
	arr := strings.Split(val, ",")
	size_int, _ := strconv.Atoi(size)
	hash_table := structs.NewHashTable(size_int)
	for i := 0; i < len(arr); i++ {
		if arr[i] != "" {
			key := strings.ReplaceAll(arr[i], " ", "")
			i++
			value := strings.ReplaceAll(arr[i], " ", "")
			hash_table.Hset(key, value)
		}
	}
	if len(arr) == 0 {
		return *hash_table, 0
	} else {
		return *hash_table, len(arr) / 2
	}
}
