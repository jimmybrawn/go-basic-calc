package main

import (
	"fmt"
	"math"
)

type number float64

func (n number) print(x number) {
	fmt.Println(x)
}

func (n number) multiplyBy(f number) {
	r := n * f
	n.print(r)
}

func (n number) divideBy(f number) {
	r := n / f
	n.print(r)
}

func (n number) powerBy(f number) {
	r := math.Pow(float64(n), float64(f))
	n.print(number(r))
}

func (n number) sqRoot() {
	r := n * n
	n.print(r)
}

func (n number) addBy(f number) {
	r := n + f
	n.print(r)
}

func (n number) subtrBy(f number) {
	r := n - f
	n.print(r)
}

func (n number) percOf(f number) {
	r := n / 100 * f
	n.print(r)
}

func (n number) afterPercOf(f number) {
	r := n - (n / 100 * f)
	n.print(r)
}
