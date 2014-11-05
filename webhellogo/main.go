package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/hoisie/web"
)

const CounterFile = "/data/counter"

func main() {
	os.Mkdir("/data", os.ModeDir|0755)
	web.Get("/", func() string {
		msg := fmt.Sprintf("Hello Go言語 %d!", readUpdatedCounter()) // (Hello GoLanguage)
		fmt.Println(msg)
		return msg
	})
	web.Run(":8080")
}

func readUpdatedCounter() int {
	store, _ := ioutil.ReadFile(CounterFile)
	fmt.Println(string(store))
	var i = 0
	fmt.Sscanf(string(store), "%d", &i)
	i++
	fmt.Println(i)
	store = []byte(fmt.Sprintf("%d", i))
	fmt.Println(string(store))
	ioutil.WriteFile(CounterFile, store, 0755)
	return i
}
