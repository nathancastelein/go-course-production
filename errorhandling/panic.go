package errorhandling

import "fmt"

func Panic(recoverMe bool) {
	if recoverMe {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in f", r)
			}
		}()
	}
	fmt.Println([]int{}[0])
}
