package main 

import "fmt"

func PrintRange(start, end int) string {
	return PrintIntervals(start, end, 1)
}

func PrintIntervals(start, end, interval int) string {
	ret := ""
	for i := start; i <= end; i += interval {
		ret += fmt.Sprintf("%d ", i)
	}
	return ret[:len(ret)-1]
}
