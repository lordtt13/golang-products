package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct {}

func main(){
	resp, err := http.Get("https://github.com/lordtt13")
	if err != nil{
		fmt.Println("Error: ",err)
		os.Exit(1)
	}

	lw := logWriter{}

	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error){
	fmt.Println(string(bs))
	fmt.Println("Wrote ",len(bs))

	return len(bs), nil
}