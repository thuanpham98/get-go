package main

import (
	"fmt"

	numbergoodpair "github.com/thuanpham98/get-go/number-good-pair"
)

func main() {
	var ret = numbergoodpair.NumberGoodPair([1,1,3,4,4,3,5,6]);
	fmt.Println(ret);

}