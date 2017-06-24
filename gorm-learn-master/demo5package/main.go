package demo5package

import (
	"fmt"

	"github.com/gorm-learn-master/demo5package/math"
)

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := math.Average(xs)
	fmt.Println(avg)
}
