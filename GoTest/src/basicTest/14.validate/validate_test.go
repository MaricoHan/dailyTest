package validate

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

// Binding Binding
// 判断结构体的必要字段是否存在
// 在官方支持omitempty关键字的
// 基础上增加required关键字
// tag: json
// key: omitempty/required
func Binding(i interface{}) error {
	elems := reflect.ValueOf(i).Elem()
	rtypeof := elems.Type()
	if rtypeof.Kind() != reflect.Struct {
		return errors.New("Type of i is not struct")
	}

	for i := 0; i < elems.NumField(); i++ {
		vfield := elems.Field(i)
		tfield := rtypeof.Field(i)

		var tag string
		tags := tfield.Tag.Get("json")
		if index := strings.Index(tags, ","); index != -1 {
			tag = tags[index+1:]
		}

		if tag == "required" {
			if vfield.Kind() == reflect.Ptr {
				if vfield.IsNil() {
					return errors.New("Field '" + tfield.Name + "' is required")
				}
			} else if vfield.Kind() == reflect.String {
				if vfield.IsZero() {
					return errors.New("Field '" + tfield.Name + "' is required")
				}
			} else if vfield.Kind() == reflect.Interface {
				if vfield.IsNil() {
					return errors.New("Field '" + tfield.Name + "' is required")
				}
			} else if vfield.Kind() == reflect.Slice {
				if vfield.IsNil() {
					return errors.New("Field '" + tfield.Name + "' is required")
				}
			} else {
				return errors.New("Field type error")
			}
		}
	}
	return nil
}

type Student struct {
	Name string      `json:"name,required"`
	Age  *int        `json:"age,required"`
	Info interface{} `json:"info,required"`
}

func TestValidate(t *testing.T) {
	strs := []string{
		`{"name":"Bob","age":10}`,
		"{\"name\":\"Bob\",\"info\":\"info aaaaaaa\"}",
		"{\"age\":10,\"info\":\"info aaaaaaa\"}",
		"{\"name\":\"Bob\",\"age\":10,\"info\":\"info aaaaaaa\"}",
	}

	for i := 0; i < len(strs); i++ {
		s := &Student{}
		if err := json.Unmarshal([]byte(strs[i]), s); err != nil {
			continue
		}
		fmt.Println(*s)

		if err := Binding(s); err != nil {
			fmt.Println(strs[i], "lack of fileds, errMsg:", err.Error())
			continue
		}
		fmt.Println(strs[i], "perfect")
	}
}

type person struct {
	Name string `json:"name"`
	Age  string `json:"age,omitempty"` // omitempty ：转为json字符串时，如果该value为空，则忽略该key
}
type person1 struct {
	Name string   `json:"name"`
	Age  []string `json:"age,omitempty"`
}

func TestOmitEmpty(t *testing.T) {
	j := string(`{"name":"han"}`)

	p1 := person1{}
	json.Unmarshal([]byte(j), &p1)
	fmt.Println("p1",p1)
	bz, _ := json.Marshal(p1)
	s:=string(bz)
	fmt.Println("json字符串：",s)

	p2 := person1{}
	json.Unmarshal([]byte(s), &p2)
	fmt.Println(p2)

	fmt.Println(reflect.ValueOf(p2.Age).IsZero())
}
func TestOmitEmpty1(t *testing.T) {
	j := string(`{"name":"han","age":[]}`)

	p1 := person1{}
	json.Unmarshal([]byte(j), &p1)
	fmt.Println("p1",p1)
	bz, _ := json.Marshal(p1)
	s:=string(bz)
	fmt.Println("json字符串：",s)

	p2 := person1{}
	json.Unmarshal([]byte(s), &p2)
	fmt.Println(p2)

	fmt.Println(reflect.ValueOf(p2.Age).IsZero())
}
