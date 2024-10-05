package main

import (
	"fmt"
	"os"

	l "linear/functions"
)

func main() {
	v := l.ReadFile(os.Args[1])
	v1 := l.ConvertData(v)
	sumx, sumy, sumxy, sumx2, sumy2 := l.Sums(v1)
	b, a := l.Linear_Regression_Line(sumx, sumy, sumxy, sumx2, sumy2, v1)
	r := l.Pearson_Correlation_Coefficient(sumx, sumy, sumxy, sumx2, sumy2, v1)
	fmt.Printf("Linear Regression Line: y = %.6fx + %.6f\n", b, a)
	fmt.Printf("Pearson Correlation Coefficient: %.10f\n", r)
}
