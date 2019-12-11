package flexible_reflect

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestDeepEqual(t *testing.T) {
	a := map[int]string{1: "one", 2: "two", 3: "three"}
	b := map[int]string{1: "one", 2: "two", 3: "three"}
	t.Log("a == b ?", reflect.DeepEqual(a, b))

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	s3 := []int{2, 1, 3}

	t.Log("s1 == s2 ?", reflect.DeepEqual(s1, s2))
	t.Log("s1 == s3 ?", reflect.DeepEqual(s1, s3))

	c1 := Customer{"1", "Mike", 40}
	c2 := Customer{"1", "Mike", 40}
	fmt.Println(c1 == c2)
	fmt.Println(reflect.DeepEqual(c1, c2))
}

type Employee struct {
	EmployeeID string
	Name       string `"format":"normal"`
	Age        int
}

func (e *Employee) UpdateAge(newVal int) {
	e.Age = newVal
}

type Customer struct {
	CookieID string
	Name     string
	Age      int
}

func fillBySetting(st interface{}, settings map[string]interface{}) error {

	if reflect.TypeOf(st).Kind() != reflect.Ptr {
		return errors.New("the first para should be a pointer to the struct type")
	}

	if reflect.TypeOf(st).Elem().Kind() != reflect.Struct {
		return errors.New("the first param should be a pointer to the struct type")
	}

	if settings == nil {
		return errors.New("setting is nil")
	}

	var (
		field reflect.StructField
		ok    bool
	)

	for k, v := range settings {
		if field, ok = (reflect.ValueOf(st)).Elem().Type().FieldByName(k); !ok {
			continue
		}

		if field.Type == reflect.TypeOf(v) {
			vstr := reflect.ValueOf(st)
			vstr = vstr.Elem()
			vstr.FieldByName(k).Set(reflect.ValueOf(k))
		}

	}
	return nil
}

func TestFillNameAndAge(t *testing.T) {
	setting := map[string]interface{}{"Name": "Mike", "Age": 30}
	e := Employee{}
	if err := fillBySetting(&e, setting); err != nil {
		t.Fatal(err)
	}
	t.Log(e)
	c := new(Customer)
	if err := fillBySetting(c, setting); err != nil {
		t.Fatal(err)
	}
	t.Log(*c)
}
