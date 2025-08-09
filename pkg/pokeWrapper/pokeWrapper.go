package pokewrapper

import (
	"fmt"
	"io"
	"net/http"
)

func getType(name string) ([2]string, error) {
	var result [2]string
	return result, nil

}

func Foo() {
	res, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	data, _ := io.ReadAll(res.Body)
	fmt.Println(string(data))
}
