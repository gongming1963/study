str2 := "{\"foo\":{\"baz\": [1,2,3]}}"

var y map[string]interface{}
json.Unmarshal([]byte(str2), &y)

fmt.Println("%v", y)
