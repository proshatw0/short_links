package workfile

import (
	"errors"
	"strconv"

	"dms/src/checkURL"
)

func Processing_Request(filepath string, commands [2]string) (error, string) {
	table_name := "links"
	err, mod, number_line := Search_Table(filepath, commands[0], table_name)
	if err != nil || number_line == -1 {
		return err, ""
	}
	if mod == "" {
		return errors.New("unknown command"), ""
	}
	if number_line != -1 {
		if commands[1] == "" {
			return errors.New("Example request: `<command>\n<link>`"), ""

		}
		links, len := Scan_Table_Links(filepath, number_line)
		if links.Size <= 0 || len == -1 {
			return errors.New("table not found"), ""
		}
		switch commands[0] {
		case "post":
			if !checkURL.CheckURL(commands[1]) {
				return errors.New("original link is not available"), ""
			}
			base := "http://10.241.125.222/"
			base += strconv.Itoa(len + 1)
			err := links.Hset(base, commands[1])
			if err != nil {
				return err, ""
			}
			err = Print_Table_Hash_Table(filepath, number_line, "links", links)
			if err != nil {
				return err, ""
			}
			return nil, base
		case "get":
			value, err := links.Hget(commands[1])
			if err != nil {
				return err, ""
			}
			return nil, value
		}
	} else {
		return errors.New("table not found"), ""
	}
	return errors.New("table not found"), ""
}
