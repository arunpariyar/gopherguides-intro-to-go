package demo

import "fmt"

func Add(a int, b int)(int, error){
	if a == 0 || b == 0 {
		return 0, fmt.Errorf("exp non values got a:%d, d:%d", a, b)
	}

	return a+b, nil
}