package main

import (
	"fmt"
	"reflect"
	"strings"
)

var space = "                                    "

const t1 = "class=box cellspacing=0 cellpadding=3 width=100%"
const t2 = "class=bottom width=7%"
const t20 = "class=bottom width=5% "
const t3 = "class=bottom1"

var colors = []string{"FFFFF8", "#F8F8FF", "FFF8F8", "F8FFFF", "FFF8FF", "F8FFF8",
	"FFFFF8", "#F8F8FF", "FFF8F8", "F8FFFF", "FFF8FF", "F8FFF8",
	"FFFFF8", "#F8F8FF", "FFF8F8", "F8FFFF", "FFF8FF", "F8FFF8"}

// Struct2html convert any go object to html structure.
func Struct2html(data interface{}, max int, skip string) string {

	var out = `
	<style>
	td.box {
		border: 1px solid #E0E0E0;
	}
	table.box {
    border-left: 1px solid #E0E0E0;
		border-right: 1px solid #E0E0E0;
		border-top: 1px solid #E0E0E0;
  }
	td.bottom1 {

		border-bottom: 1px solid #E0E0E0;
		border-leftc: 1px solid #E0E0E0;
	}
	td.bottom {
		vertical-align: top;
    text-align: right;
		border-bottom: 1px solid #E0E0E0;
		border-leftc: 1px solid #E0E0E0;
	}
	</style>
	`
	o, _ := visit(reflect.ValueOf(data), 0, max, skip)
	out += o
	return out
}

func visit(v reflect.Value, level int, max int, skip string) (string, string) {
	out := ""
	pointerV := v
	pointer := false
	for {
		if pointerV.Kind() == reflect.Interface {
			pointerV = pointerV.Elem()
		}
		if pointerV.Kind() == reflect.Ptr {
			pointer = true
			v = reflect.Indirect(pointerV)
		}
		if pointer {
			pointerV = v
		}
		pointer = false
		switch pointerV.Kind() {
		case reflect.Ptr, reflect.Interface:
			continue
		}
		break
	}
	//originalV := v
	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	k := v.Kind()
	spacer := ""

	if level > max {
		return "", ""
	}
	elType := ""
	if k == reflect.Slice || k == reflect.Array {
		out1, type1 := visitArray(v, level, max, spacer, skip)
		elType = type1
		out += out1

	} else if k == reflect.Struct {
		out1, type1 := visitStruct(v, level, max, spacer, skip)
		elType = type1
		out += out1

	} else if k == reflect.Map {
		out1, type1 := visitMap(v, level, max, spacer, skip)
		elType = type1
		out += out1
	} else {
		out += fmt.Sprintf("%v", v)
	}
	return out, elType
}

func visitMap(v reflect.Value, level, max int, spacer, skip string) (string, string) {
	out := ""
	len := len(v.MapKeys())
	title := false
	if len == 1 && level == 0 {
		title = true
	}
	if !title {
		out += spacer + fmt.Sprintf("<table bgcolor=%s %s>\n", colors[level], t1)
	}
	for _, k := range v.MapKeys() {
		kv := v.MapIndex(k)
		out1, type1 := visit(kv, level+1, max, skip)
		if title {
			out += fmt.Sprintf("<big><b>%s%s</b></big>%s", k, type1, out1)
		} else {
			out += spacer + fmt.Sprintf("<tr><td %s><font color=gray>%v%s:</td><td %s>%s</td>\n",
				t2, k, type1, t3, out1)
		}
	}
	out += spacer + "</table>\n"
	return out, "{}"
}

func visitStruct(v reflect.Value, level, max int, spacer, skip string) (string, string) {
	out := ""
	out += spacer + fmt.Sprintf("<table bgcolor=%s %s>\n", colors[level], t1)
	vt := v.Type()

	for i := 0; i < vt.NumField(); i += 1 {
		sf := vt.Field(i)
		f := v.FieldByIndex([]int{i})
		stype := sf.Type.String()
		six := strings.LastIndex(stype, ".")
		if six > 0 {
			stype = stype[six+1:]
		}
		stype = strings.Replace(stype, "string", "str", -1)
		//fmt.Println("name=", sf.Name, sf.Type, skip)
		if skip == "" || strings.Index(","+skip+",", ","+sf.Name+",") < 0 {
			out1, type1 := visit(f, level+1, max, skip)
			out += spacer + fmt.Sprintf("<tr><td %s><font color=gray>%s%s:<br>%s</td><td %s>%v</td>\n",
				t2, sf.Name, type1, stype, t3, out1)
		}
	}
	out += spacer + "</table>\n"
	return out, ""
}

func visitArray(v reflect.Value, level, max int, spacer, skip string) (string, string) {
	elType := fmt.Sprintf("[%d]", v.Len())
	var out = ""
	if level == max {
		out += fmt.Sprintf("Count: %d", v.Len())
	} else {
		out += spacer + fmt.Sprintf("<table bgcolor=%s %s>\n", colors[level], t1)
		for i := 0; i < v.Len(); i++ {
			elem := v.Index(i)
			out1, _ := visit(elem, level+1, max, skip)
			out += spacer + fmt.Sprintf("<tr><td %s>#%d:</td><td %s>%s</td>\n",
				t20, i, t3, out1)
		}
		out += spacer + "\n</table>\n"
	}
	return out, elType
}
