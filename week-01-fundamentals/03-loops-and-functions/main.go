package main

import (

	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)
for {

	oldZ := z

	z -= (z*z - x) / (2*z)


	if math.Abs(z - oldZ) < 0.000001{
	break
}

}
 return z
}

func main() {
     fmt.Println(Sqrt(2))
}     

