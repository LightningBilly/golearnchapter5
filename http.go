package main

import (
	"net/http"
	"fmt"
	"io"
	"os"
)

func main() {
	resp, err:=http.Head("http://www.baidu.com")

	if err!=nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}