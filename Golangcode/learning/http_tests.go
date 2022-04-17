package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

)

func oneTimeOp() {
	fmt.Println("one time op start")
	time.Sleep(3 * time.Second)
	fmt.Println("one time op started")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("http handler start")
		var once sync.Once
		once.Do(oneTimeOp)
		fmt.Println("http handler end")
		w.Write([]byte("done!"))
	})
	log.Fatal(http.ListenAndServe(":8089", nil))
}
