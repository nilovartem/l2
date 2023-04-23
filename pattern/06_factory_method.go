package pattern

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/patrickmn/go-cache"
	yaml "gopkg.in/yaml.v2"
)

/*
	Реализовать паттерн «фабричный метод».

Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.

	https://en.wikipedia.org/wiki/Factory_method_pattern
*/
type method string

const (
	CACHE method = "cache"
	YAML  method = "yaml"
)

type User struct {
	Id   int    `yaml:"id"`
	Name string `yaml:"username"`
}
type Keeper interface {
	//Put()
	Load(*User)
	Unload() *User
	Name() string
}

// Данные в кеше
type Cache struct {
	*cache.Cache
}

func (c *Cache) Unload() *User {
	if c.Cache == nil {
		return nil
	}
	user := User{}
	id, _ := c.Get("Id")
	user.Id = id.(int)
	name, _ := c.Get("Name")
	user.Name = name.(string)
	return &user
}
func (c *Cache) Load(user *User) {
	if c.Cache == nil {
		c.Cache = cache.New(5*time.Minute, 10*time.Minute)
	}
	c.Cache.Set("Id", user.Id, cache.NoExpiration)
	c.Cache.Set("Name", user.Name, cache.NoExpiration)
}
func (c *Cache) Name() string {
	return "Cache"
}

// Сохраняем данные в файле
type File struct {
}

func (file *File) Read() []byte {
	path, _ := filepath.Abs("../L2/" + "user.yaml")
	contents, _ := os.ReadFile(path)
	return contents
}
func (file *File) Unload() *User {
	user := User{}
	_ = yaml.Unmarshal(file.Read(), &user)
	fmt.Println("File unload is OK")
	return &user
}
func (file *File) Load(user *User) {
	path, _ := filepath.Abs("../L2/" + "user.yaml")
	data, _ := yaml.Marshal(user)
	_ = ioutil.WriteFile(path, data, 0)
}
func (file *File) Name() string {
	return "YAML File"
}

type Memorizer interface {
	Memory(action method) Keeper
}
type Memory struct{}

func (m *Memory) Memory(action method) Keeper {
	var k Keeper
	switch action {
	case CACHE:
		k = &Cache{}
	case YAML:
		k = &File{}
	}
	return k
}

// Пользователь может выбирать, с каким хранилищем работать
func RunFactory() {
	/*
		user := &User{Id: 1, Name: "Artem"}
		var keeper Keeper
		m := Memory{}
	*/
	//Работа с файлом
	/*
		keeper = m.Memory(YAML)
		keeper.Load(user)
		user = nil
		user = keeper.Unload()
		fmt.Println(user.Id)
		fmt.Println(user.Name)
	*/
	//Работа с кэшем
	/*
		keeper = m.Memory(CACHE)
		keeper.Load(user)
		user = nil
		user = keeper.Unload()
		fmt.Println(user.Id)
		fmt.Println(user.Name)
	*/

}
