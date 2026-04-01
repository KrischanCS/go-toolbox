package optional_test

import (
	"encoding/json"
	"encoding/xml"
	"fmt"

	"github.com/KrischanCS/go-toolbox/optional"
)

func ExampleOptional_Get() {
	optPresent := optional.Of("value")
	optEmpty := optional.Empty[string]()

	value, ok := optPresent.Get()
	fmt.Printf("value 1:%s\n", value)
	fmt.Printf("ok 1:%t\n", ok)

	fmt.Println()

	value, ok = optEmpty.Get()
	fmt.Printf("value 2:%s\n", value)
	fmt.Printf("ok 2:%t\n", ok)

	// Output:
	// value 1:value
	// ok 1:true
	//
	// value 2:
	// ok 2:false
}

func ExampleOptional_Or() {
	optPresent := optional.Of("value")
	value := optPresent.Or("fallback")
	fmt.Printf("value 1: %s\n", value)

	optEmpty := optional.Empty[string]()
	value = optEmpty.Or("fallback")
	fmt.Printf("value 2: %s\n", value)

	// Output:
	//
	// value 1: value
	// value 2: fallback
}

func ExampleOptional_MarshalJSON() {
	type Test struct {
		Value1 optional.Optional[string] `json:"value1"`
		Value2 optional.Optional[string] `json:"value2"`
		Value3 optional.Optional[string] `json:"value3,omitempty"` // omitempty will be kept, event if optional is empty
		Value4 optional.Optional[string] `json:"value4,omitzero"`  // omitzero will correctly omit an empty optional
	}

	t := Test{
		Value1: optional.Of("value1"),
		Value2: optional.Empty[string](),
		// Not setting Value3 & Value4, zero values are empty optionals,
		// equivalent as setting Value2 to empty explicitly
	}

	jsonBytes, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonBytes))

	// Output:
	// {
	//   "value1": "value1",
	//   "value2": null,
	//   "value3": null
	// }
}

func ExampleOptional_UnmarshalJSON() {
	type Test struct {
		Value1 optional.Optional[string] `json:"value1"`
		Value2 optional.Optional[string] `json:"value2"`
		Value3 optional.Optional[string] `json:"value3"`
	}

	jsonBytes := []byte(`{"value1":"value1","value2":null}`)

	var t Test
	if err := json.Unmarshal(jsonBytes, &t); err != nil {
		panic(err)
	}

	fmt.Println(t.Value1.String())
	fmt.Println(t.Value2.String())
	fmt.Println(t.Value3.String())

	// Output:
	// (Optional[string]: value1)
	// (Optional[string] <empty>)
	// (Optional[string] <empty>)
}

func ExampleOptional_MarshalXML() {
	type Test struct {
		Value1 optional.Optional[string] `xml:"value1"`
		Value2 optional.Optional[string] `xml:"value2"`           // empty existing tags are considered existing empty strings
		Value3 optional.Optional[string] `xml:"value3,omitempty"` // omitempty will be kept, event if optional is empty
	}

	t := Test{
		Value1: optional.Of("value1"),
		Value2: optional.Empty[string](),
		// Not setting Value3 & Value4, zero values are empty optionals, equivalent as setting Value2 to empty explicitly
	}

	xmlBytes, err := xml.MarshalIndent(t, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(xmlBytes))

	// Output:
	// <Test>
	//   <value1>value1</value1>
	// </Test>
}

func ExampleOptional_UnmarshalXML() {
	type Test struct {
		Value1 optional.Optional[string] `xml:"value1"`
		Value2 optional.Optional[string] `xml:"value2"`
		Value3 optional.Optional[string] `xml:"value3"`
	}

	xmlBytes := []byte(`<Test><value1>value1</value1><value2></value2></Test>`)

	var t Test

	err := xml.Unmarshal(xmlBytes, &t)
	if err != nil {
		panic(err)
	}

	fmt.Println(t.Value1.String())
	fmt.Println(t.Value2.String()) // Empty existing tag will be considered existing empty string
	fmt.Println(t.Value3.String())

	// Output:
	// (Optional[string]: value1)
	// (Optional[string]: )
	// (Optional[string] <empty>)
}

func ExampleOptional_String() {
	optPresent := optional.Of("value")
	optEmpty := optional.Empty[string]()

	fmt.Println(optPresent.String())
	fmt.Println(optEmpty.String())

	// Output:
	// (Optional[string]: value)
	// (Optional[string] <empty>)
}
