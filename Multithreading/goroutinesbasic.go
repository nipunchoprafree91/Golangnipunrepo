package Multithreading

import (
	"fmt"
	"sync"
	// "time"
	"unicode"
)

func Basicgoroutines() {
	wg := new(sync.WaitGroup)

	wg.Add(20)
	go goroutines1(wg, 20)
	wg.Add(20)
	go goroutines2(wg, 't')

	wg.Wait()

}

func goroutines1(wg *sync.WaitGroup, num int) {

	for i := 0; i <= num; i++ {

		fmt.Println("In Goroutine1")
		fmt.Println("value of i : ", i)
		wg.Done()
	}
}

func goroutines2(wg *sync.WaitGroup, str rune) {

	for r := 'a'; r <= str; r++ {

		fmt.Println("In Goroutine2")
		R := unicode.ToUpper(r)
		fmt.Printf("%c%c\n", r, R)
		wg.Done()
	}
	// wrong as this will get called once and
	// rest 19 goroutines will wait endlessly
	// wg.Done()
}
