# gostruct2html

Convert any GO structure to HTML representation using reflection.

Example:
```
package main
import "fmt"

type Struct1 struct {
	field1 string
	field2 map[string]string
	field3 []string
	field5 Struct3
}

type Struct3 struct {
	subfield1 string
	subfield2 int
}

type Struct2 struct {
	struct1 Struct1
	field4  string
}

var test_struct = Struct2{
	struct1: Struct1{
		field1: "test",
		field2: map[string]string{
			"host_id":     "h.host_id",
			"site_id":     "s.site_id",
			"platform_id": "p.platform_id",
			"agent_id":    "a.agent_id",
			"subnet_id":   "b.subnet_id",
			"master_id":   "e.master_id",
		},
		field3: []string{"this", "is", "a", "test"},
		field5: Struct3{subfield1: "sub1 value", subfield2: 10},
	},
	field4: "This is field4",
}

func main() {
  fmt.Println(Struct2html(test_struct, 10, ""))
}

```

go run ./*.go > test_struct.html

![struct](https://user-images.githubusercontent.com/31523474/32586566-76accc10-c4c0-11e7-9615-e0f5361e491e.png)
