package Datatypes

import (
	"fmt"
	"time"
)

type freqmapinfo struct {
	number int
	freq   int
}

func Maps_freq_calculator(arr []int) []freqmapinfo {

	freqmaparray := []freqmapinfo{}

	fmt.Println("Time now is : ", time.Now())

	freq_map := make(map[int]int)
	for _, num := range arr {
		freq_map[num] = freq_map[num] + 1
		freqmaparray = append(freqmaparray, freqmapinfo{num, freq_map[num]})

	}
	return freqmaparray

}
