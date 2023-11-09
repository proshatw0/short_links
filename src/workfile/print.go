package workfile

import (
	"strconv"

	"dms/src/structs"
)

func Print_Table_Hash_Table(filepath string, line_number int, name_table string, ht structs.Hash_Table) error {
	out := name_table
	out += ": [" + strconv.Itoa(ht.Size) + ", {"
	for index := 0; index < ht.Size; index++ {
		currentNode := ht.Table[index].Head
		for currentNode != nil {
			out += "(" + currentNode.Data.Key + ", " + currentNode.Data.Value + "), "
			currentNode = currentNode.Next
		}
	}
	if len(out) == len(name_table)+len(strconv.Itoa(ht.Size))+6 {
		out += "}]"
	} else {
		out = out[0:len(out)-2] + "}]"
	}
	err := WriteLineFromFile(filepath, line_number, out)
	if err != nil {
		return err
	}
	return nil
}
