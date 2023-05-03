Что выведет программа? Объяснить вывод программы.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()
	return c
}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4 ,6, 8)
	c := merge(a, b )
	for v := range c {
		fmt.Println(v)
	}
}
```

Ответ:
```
1
2
4
6
3
5
7
8
0

Это пример вывода, потому что является бесконечным и программа завершится лишь принудительно.
Порядок вывода чисел будет случайным и зависит от планировщика. Что касается 0, то это стандартное "нулевое" значение для типа int. Оно выводится, потому что если мы делаем range по закрытому каналу, мы всегда будем получать стандартное "нулевое" значение в зависимости от типа.

```
