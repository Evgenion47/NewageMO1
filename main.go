package main

import (
	"fmt"
	"math"
	"os"
	"text/tabwriter"
)

func main() {
	fmt.Println("fibonacci")
	FibonacciMethod(1.1, 1.5)
	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("golden ratio")
	GoldenRatio(1.1, 1.5)
	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("half division")
	HalfDivision(1.1, 1.5)
	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("newton")
	Newton(1.1)
	fmt.Println("----------------------------------------------------------------------")

}
func calc(x float64) float64 {
	return math.Pow(x, 4) / math.Log(x)
}

func dxcalc(x float64) float64 {
	return (math.Pow(x, 3) * (-1 + 4*(math.Log(x)))) / (math.Pow(math.Log(x), 2))
}

func Newton(x0 float64) {
	EPS := math.Pow(10, -3)
	k := 1

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', tabwriter.AlignRight)
	fmt.Fprintln(w, "№       \tX       \tf(X)        \tf'(X)        \t|Xn-X(n-1)|      ")

	x1 := x0 - calc(x0)/dxcalc(x0)
	for math.Abs(x1-x0) >= EPS {
		_, err := fmt.Fprintf(w, "%d\t%f\t%f\t%f\t%f\n", k, x0, calc(x0), dxcalc(x0), math.Abs(x1-x0))
		if err != nil {
			print("oooups")
		}
		x0 = x1
		x1 = x0 - calc(x0)/dxcalc(x0)
		k++
	}
	_, err := fmt.Fprintf(w, "%d\t%f\t%f\t%f\t%f\n", k, x0, calc(x0), dxcalc(x0), math.Abs(x1-x0))
	if err != nil {
		print("oooups")
	}

	fmt.Fprintln(w)
	w.Flush()
}

func HalfDivision(a, b float64) {
	EPS := math.Pow(10, -3)
	DELTA := math.Pow(10, -4)
	var alpha, beta float64
	k := 1

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', tabwriter.AlignRight)
	fmt.Fprintln(w, "№       \tinterval       \tX*        \tf(X*)        ")

	for ; b-a >= EPS; k++ {
		mid := (a + b) / 2
		alpha = mid - DELTA
		beta = mid + DELTA

		_, err := fmt.Fprintf(w, "%d\t[%f,%f]\t%f\t%f\n", k, a, b, mid, calc(mid))
		if err != nil {
			print("oooups")
		}

		if calc(alpha) < calc(beta) {
			b = beta
		} else {
			a = alpha
		}
	}

	x := (a + b) / 2
	fx := calc(x)
	_, err := fmt.Fprintf(w, "%d\t[%f,%f]\t%f\t%f\n", k, a, b, x, fx)
	if err != nil {
		print("oooups")
	}
	fmt.Fprintln(w)
	w.Flush()
}

func GoldenRatio(a, b float64) {
	EPS := math.Pow(10, -3)
	k := 1
	lambda := (math.Sqrt(5) - 1) / 2
	alpha := a + (1-lambda)*(b-a)
	beta := a + lambda*(b-a)

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', tabwriter.AlignRight)
	fmt.Fprintln(w, "№       \tinterval       \tX*        \tf(X*)        ")

	for ; b-a >= EPS; k++ {
		mid := (a + b) / 2
		_, err := fmt.Fprintf(w, "%d\t[%f,%f]\t%f\t%f\n", k, a, b, mid, calc(mid))
		if err != nil {
			print("oooups")
		}
		if calc(alpha) < calc(beta) {
			b = beta
			beta = alpha
			alpha = a + (1-lambda)*(b-a)
		} else {
			a = alpha
			alpha = beta
			beta = a + lambda*(b-a)
		}
	}
	x := (a + b) / 2
	fx := calc(x)
	_, err := fmt.Fprintf(w, "%d\t[%f,%f]\t%f\t%f\n", k, a, b, x, fx)
	if err != nil {
		print("oooups")
	}
	fmt.Fprintln(w)
	w.Flush()
}

func FibonacciMethod(a, b float64) {
	EPS := math.Pow(10, -3)
	DELTA := math.Pow(10, -4)
	k := 1
	FN := precalcFib(92)

	n := 1
	for FN[n]*EPS <= (b - a) {
		n++
	}
	alpha := a + (b-a)*FN[n-2]/FN[n]
	beta := a + (b-a)*FN[n-1]/FN[n]

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', tabwriter.AlignRight)
	fmt.Fprintln(w, "№       \tinterval       \tX*        \tf(X*)        ")

	for ; k != n; k++ {
		mid := (a + b) / 2
		_, err := fmt.Fprintf(w, "%d\t[%f,%f]\t%f\t%f\n", k, a, b, mid, calc(mid))
		if err != nil {
			print("oooups")
		}
		if calc(alpha) >= calc(beta) {
			a = alpha
			alpha = beta
			beta = a + (b-a)*FN[n-k]/FN[n-k+1]
		} else {
			b = beta
			beta = alpha
			alpha = a + (b-a)*FN[n-k-1]/FN[n-k+1]
		}
	}
	alpha = a
	beta = alpha + DELTA
	if calc(alpha) > calc(beta) {
		a = alpha
	} else {
		b = beta
	}
	x := (a + b) / 2
	fx := calc(x)
	_, err := fmt.Fprintf(w, "%d\t[%f,%f]\t%f\t%f\n", k, a, b, x, fx)
	if err != nil {
		print("oooups")
	}
	fmt.Fprintln(w)
	w.Flush()
}

func precalcFib(count int) []float64 {
	fibNumbers := make([]float64, count, count)
	fibNumbers[1] = 1
	for i := 2; i < count; i++ {
		fibNumbers[i] = fibNumbers[i-1] + fibNumbers[i-2]
		if fibNumbers[i] < 0 {
			return fibNumbers
		}
	}
	return fibNumbers
}
