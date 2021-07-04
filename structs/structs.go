package libstructs

import (
	"database/sql"
	fmt "fmt"
	"strconv"
	"strings"

	"github.com/mitchellh/mapstructure"
)

// MergeRow merge sql.Rows with a struct
func MergeRow(dst interface{}, rows *sql.Rows) {
	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		fmt.Printf("libstructs: %v\n", err.Error())
	}

	// Get column type
	columnsType, err := rows.ColumnTypes()
	if err != nil {
		fmt.Printf("libstructs: %v\n", err.Error())
	}

	// Make a slice for the values
	values := make([]sql.RawBytes, len(columns))

	// rows.Scan wants '[]interface{}' as an argument, so we must copy the
	// references into such a slice
	// See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// get RawBytes from data
	err = rows.Scan(scanArgs...)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Now do something with the data.
	// Here we just print each column as a string.
	m := make(map[string]interface{})
	// var value string
	for i, col := range values {
		// Here we can check if the value is nil (NULL value)
		if col != nil {
			switch value := strings.ToLower(columnsType[i].DatabaseTypeName()); value {
			case "varchar":
				m[columns[i]] = string(col)
			case "decimal":
				s, _ := strconv.ParseFloat(string(col), 64)
				m[columns[i]] = s
			case "text":
				m[columns[i]] = string(col)
			case "tinyint":
				n, _ := strconv.Atoi(string(col))
				if n > 0 {
					m[columns[i]] = true
				} else {
					m[columns[i]] = false
				}
			default:
				if strings.Contains(value, "int") {
					m[columns[i]], err = strconv.Atoi(string(col))
				} else {
					m[columns[i]], err = strconv.Atoi(string(col))
					if err != nil {
						m[columns[i]] = string(col)
					}
				}
			}
		}
	}

	config := &mapstructure.DecoderConfig{
		TagName: "db",
		Result:  &dst,
	}
	decoder, _ := mapstructure.NewDecoder(config)
	err = decoder.Decode(m)
	if err != nil {
		fmt.Printf("libstructs: %v\n", err.Error())
	}
}
