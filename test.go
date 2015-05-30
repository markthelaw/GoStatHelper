// test
package main

import (
	"fmt"
	"github.com/markthelaw/GoStatHelper/StatUtil"
)


func main() {
	fmt.Println("Hello World!")
	
	
	
	fmt.Println(StatUtil.ErfInv(0.1))
	fmt.Println(StatUtil.InverserNormalCDF(2))
	fmt.Println(StatUtil.InverseNormalDist(0.5, 1, 2))
	fmt.Println(StatUtil.InverseNormalDist(0.4, 1, 4))
	fmt.Println(StatUtil.InverseNormalDist(0.1, 0.1, 0.12))
	
}
