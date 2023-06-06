package messages

import (
  "fmt"
  "strconv"

  "strings"
  "unicode/utf8"
  "encoding/json"
)

func MessageConvertSQL (message string) string{
  sqlMessage := "Check"
  fmt.Println(sqlMessage)
  return sqlMessage

}
func trimLastChar(s string) string {
    r, size := utf8.DecodeLastRuneInString(s)
    if r == utf8.RuneError && (size == 0 || size == 1) {
        size = 0
    }
    return s[:len(s)-size]
}


func ConvertToJSONAndKeyValue(message string) (string, map[string]int, error) {
  //fmt.Println("ok", message)
	message = trimLastChar(message)
 // fmt.Println("2 ok", message)
  message = strings.ReplaceAll(message, " ","")
  parts := strings.Split(message, ":")

	if len(parts)%2 != 0 {
		return "", nil, fmt.Errorf("invalid message format")
	}

	jsonData := make(map[string]interface{})
	keyValue := make(map[string]int)

	for i := 0; i < len(parts); i += 2 {
		key := parts[i]
		value := parts[i+1]

    values, _ := strconv.Atoi(value)

		jsonData[key] = values
		keyValue[key] = values

	}

	jsonBytes, err := json.Marshal(jsonData)
	if err != nil {
		return "", nil, err
	}

	jsonStr := string(jsonBytes)
  return jsonStr, keyValue, nil
}


func MessageConvertNoSQL (message string) string{
  noSqlMessage := "check"
  fmt.Println(noSqlMessage) 
  return noSqlMessage
}
