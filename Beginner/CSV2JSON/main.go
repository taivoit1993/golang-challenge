package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

// convert csv string to json string
func ConvertCSV2JSON(csv string) string {
	jsonString := "["
	var title []string
	//loop multiple line csv string
	//first line always title
	for pos, line := range strings.Split(strings.TrimSuffix(csv, "\n"), "\n") {
		//title of csv => key of json
		if pos == 0 {
			title = strings.Split(line, ",")
		} else {
			temp := strings.Split(line, ",")
			length := len(temp)
			jsonString += "{"
			for i, v := range temp {
				jsonString += fmt.Sprint(title[i], ":", v)
				if i < length-1 {
					jsonString += ","
				}

			}
			jsonString += "},"
		}
	}
	jsonString = jsonString[:len(jsonString)-1] + "]"
	return jsonString
}
func main() {
	var csv string = "\"Id\",\"UserName\"\n\"1\",\"Sam Smith\"\n\"2\",\"Fred Frankly\"\n\"1\",\"Zachary Zupers\""
	var temp interface{}
	jsonString := ConvertCSV2JSON(csv)
	fmt.Println("Json String: ", jsonString)
	//Parse String To Json
	err := json.Unmarshal([]byte(jsonString), &temp)
	if err == nil {
		fmt.Println(temp)
	} else {
		fmt.Println("Can't Parse Csv To Json", err)
	}

}
