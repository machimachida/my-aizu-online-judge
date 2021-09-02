package itp1

import (
	"fmt"
	"math"
)

func Distance() {
	var x1 float64
	var y1 float64
	var x2 float64
	var y2 float64
	_, err := fmt.Scan(&x1, &y1, &x2, &y2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%.8f\n", math.Sqrt(math.Pow(x1-x2, 2)+math.Pow(y1-y2, 2)))
}

func Triangle() {
	var a int
	var b int
	var c int
	_, err := fmt.Scan(&a, &b, &c)
	if err != nil {
		fmt.Println(err)
		return
	}
	af := float64(a)
	bf := float64(b)
	cf := float64(c) * math.Pi / 180
	S := af * bf * math.Sin(cf) / 2
	fmt.Printf("%.8f\n", S)
	fmt.Printf("%.8f\n", math.Sqrt(af*af+bf*bf-2*af*bf*math.Cos(cf))+af+bf)
	fmt.Printf("%.8f\n", 2*S/af)
}

func StandardDeviation() {
	for {
		var n int
		_, err := fmt.Scan(&n)
		if err != nil {
			fmt.Println(err)
			return
		}
		if n == 0 {
			return
		}
		if n < 1 || n > 1000 {
			fmt.Println("invalid n range")
			return
		}

		ss := make([]float64, n)
		var s int
		var m float64
		for i := 0; i < n; i++ {
			_, err := fmt.Scan(&s)
			if err != nil {
				fmt.Println(err)
				return
			}
			if s < 0 || s > 100 {
				fmt.Println("invalid s range")
				return
			}
			ss[i] = float64(s)
			m += ss[i]
		}
		nf := float64(n)
		m = m / nf
		var sd float64
		for i := 0; i < n; i++ {
			sd += math.Pow(ss[i]-m, 2) / nf
		}
		fmt.Printf("%.8f\n", math.Sqrt(sd))
	}
}

func Distance2() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println(err)
		return
	}
	if n == 0 {
		return
	}
	if n < 1 || n > 100 {
		fmt.Println("invalid n range")
		return
	}
	xs := make([]float64, n)
	ys := make([]float64, n)
	var tmp int
	for i := 0; i < n; i++ {
		_, err := fmt.Scan(&tmp)
		if err != nil {
			fmt.Println(err)
			return
		}
		if tmp < 0 || tmp > 1000 {
			fmt.Println("invalid xi range")
			return
		}
		xs[i] = float64(tmp)
	}
	for i := 0; i < n; i++ {
		_, err := fmt.Scan(&tmp)
		if err != nil {
			fmt.Println(err)
			return
		}
		if tmp < 0 || tmp > 1000 {
			fmt.Println("invalid yi range")
			return
		}
		ys[i] = float64(tmp)
	}

	ps := []float64{1, 2, 3}
	for _, p := range ps {
		var D float64
		for i := 0; i < n; i++ {
			D += math.Pow(math.Abs(xs[i]-ys[i]), p)
		}
		fmt.Printf("%.6f\n", math.Pow(D, 1/p))
	}
	max := math.Inf(-1)
	for i := 0; i < n; i++ {
		if a := math.Abs(xs[i] - ys[i]); a > max {
			max = a
		}
	}
	fmt.Printf("%.6f\n", max)
}
