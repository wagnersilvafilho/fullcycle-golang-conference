package main

import (
	"fmt"
	"strconv"
)

func main() {
	total, err := convertAndSum("1", "2", "5", "Q")
	fmt.Println(total, err)
}

func convertAndSum(s ...string) (total int, err error) {
	for _, v := range s {
		n, e := strconv.Atoi(v)
		if e != nil {
			total = 0
			err = fmt.Errorf("valor '%s' é inválido.", v)
			break
		}
		total += n
	}
	return
}
