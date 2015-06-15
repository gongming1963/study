// Go offers built-in support for JSON encoding and
// decoding, including to and from built-in and custom
// data types.

package main

import "encoding/json"
import "fmt"
import "os"
import "reflect"

// We'll use these two structs to demonstrate encoding and
// decoding of custom types below.
type Response1 struct {
	Page   int
	Fruits []string
}
type Response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func test1() {

	// First we'll look at encoding basic data types to
	// JSON strings. Here are some examples for atomic
	// values.
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	// And here are some for slices and maps, which encode
	// to JSON arrays and objects as you'd expect.
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	// The JSON package can automatically encode your
	// custom data types. It will only include exported
	// fields in the encoded output and will by default
	// use those names as the JSON keys.
	res1D := &Response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	// You can use tags on struct field declarations
	// to customize the encoded JSON key names. Check the
	// definition of `Response2` above to see an example
	// of such tags.
	res2D := &Response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	// Now let's look at decoding JSON data into Go
	// values. Here's an example for a generic data
	// structure.
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	// We need to provide a variable where the JSON
	// package can put the decoded data. This
	// `map[string]interface{}` will hold a map of strings
	// to arbitrary data types.
	var dat map[string]interface{}

	// Here's the actual decoding, and a check for
	// associated errors.
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	// In order to use the values in the decoded map,
	// we'll need to cast them to their appropriate type.
	// For example here we cast the value in `num` to
	// the expected `float64` type.
	num := dat["num"].(float64)
	fmt.Println(num)

	// Accessing nested data requires a series of
	// casts.
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	// We can also decode JSON into custom data types.
	// This has the advantages of adding additional
	// type-safety to our programs and eliminating the
	// need for type assertions when accessing the decoded
	// data.
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := &Response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	// In the examples above we always used bytes and
	// strings as intermediates between the data and
	// JSON representation on standard out. We can also
	// stream JSON encodings directly to `os.Writer`s like
	// `os.Stdout` or even HTTP response bodies.
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}

func test2() {
	str2 := "{\"foo\":{\"baz\": [1,2,3]}}"

	var y map[string]interface{}
	json.Unmarshal([]byte(str2), &y)

	for key, value := range y {
		fmt.Println("Key:", key, "Value:", value)
		fmt.Println(reflect.TypeOf(value))
		//for k1, v1 := range value {
		//	fmt.Println("Key:", k1, "Value:", v1)
		//}
	}

	fmt.Println("%v", y)
}

func test3() {
	str1 := `[{"北京":29},{"重庆":24},{"昆明":26},{"哈尔滨":5},{"天津":17},{"成都":27},{"太原":8},{"深圳":30},{"青岛":6},{"佛山":14},{"苏州":13},           {"杭州":23},{"郑州":19},{"福州":9},{"无锡":8},{"长沙":9},{"沈阳":8},{"广州":30},{"温州":10},{"大连":4},{"宁波":9},{"贵阳":16},{"南宁":3},      {"上海":30},{"西安":9},{"长春":5},{"南京":38},{"常州":4},{"东莞":27},{"武汉":26},{"石家庄":3},{"厦门":8},{"合肥":8},{"济南":13}]`
	//str1 := `{"北京":29}`
	var y []map[string]int
	json.Unmarshal([]byte(str1), &y)

	//for key, value := range y {
	//	fmt.Println("Key:", key, "Value:", value)
	//	fmt.Println(reflect.TypeOf(value))
	//}
	fmt.Printf("%+v", y)

}

func main() {
	test3()
}
