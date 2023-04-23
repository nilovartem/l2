package pattern

import (
	"fmt"
	"os"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

/*
	Реализовать паттерн «цепочка вызовов».

Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.

	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

type Object struct {
	Username string `yaml:"username"`
}
type validator interface {
	validate(object *Object, params interface{})
}
type ExistValidator struct{ next validator }

func (v *ExistValidator) validate(object *Object, filename string) {
	path, _ := filepath.Abs("../L2/" + filename)
	if _, err := os.Stat(path); err == nil {
		fmt.Println("ExistValidator is OK")
		v.next.validate(object, filename)
	} else {
		fmt.Println(err)
		fmt.Printf("[Exist] Can't locate file, aborting\n")
		return
	}
}

type ContentValidator struct{ next validator }

func (v *ContentValidator) validate(object *Object, param interface{}) {
	contents, err := os.ReadFile(param.(string))
	if err == nil {
		fmt.Println("ContentValidator is OK")
		v.next.validate(object, contents)
	} else {
		fmt.Printf("[Content] Can't read file, aborting\n")
		return
	}
}

type JsonValidator struct{}

func (v *JsonValidator) validate(object *Object, param interface{}) {
	err := yaml.Unmarshal(param.([]byte), object)

	if err != nil {
		fmt.Printf("[Json] Can't unmarshal file, aborting")
		return
	}
	fmt.Println("JsonValidator is OK")
}
func RunChain() {
	object := Object{}
	fmt.Printf("Object before file validation : %+v\n", object)
	jsonValidator := JsonValidator{}
	contentValidator := ContentValidator{&jsonValidator}
	existValidator := ExistValidator{&contentValidator}
	existValidator.validate(&object, "chain.yaml")
	fmt.Printf("Object after file validation : %+v", object)
}
