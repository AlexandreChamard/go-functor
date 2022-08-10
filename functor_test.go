package functor

import (
	"fmt"
	"testing"
)

func TestFunctor(t *testing.T) {
	functors := []Functor{
		MakeFunctor1_0(
			func(i int) {
				t.Log(i)
			},
			42,
		),
		MakeFunctor2_0(
			func(i int, s string) {
				t.Log(i, s)
			},
			1, "foobar",
		),
		MakeFunctor4_0(
			func(i int, s string, b bool, pf *float64) {
				t.Log(i, s, b, pf)
			},
			1, "foobar", true, func(f float64) *float64 { return &f }(1.23),
		),
	}

	for _, f := range functors {
		f()
	}

	f := MakeFunctorR4_2(func(a, b int, c, d float64) float64 {
		return float64(a+b) * (c - d)
	}, 4, 3)
	f1 := MakeFunctorR2_1(f, 42.42)
	if f1(3) != (4+3)*(42.42-3) {
		t.Errorf("expected %v got %v", (4+3)*(42.42-3), f1(3))
	}

	if f1(-2) != (4+3)*(42.42+2) {
		t.Errorf("expected %v got %v", (4+3)*(42.42+2), f1(-2))
	}
}

func BenchmarkFunctor(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(1)
	b.StartTimer()

	functors := []Functor{}
	for i := 0; i < 1000; i++ {

		functors = append(functors,
			MakeFunctor1_0(
				func(i int) {
					fmt.Println(i)
				},
				42,
			),
			MakeFunctor2_0(
				func(i int, s string) {
					fmt.Println(i, s)
				},
				1, "foobar",
			),
			MakeFunctor4_0(
				func(i int, s string, b bool, pf *float64) {
					fmt.Println(i, s, b, pf)
				},
				1, "foobar", true, func(f float64) *float64 { return &f }(1.23),
			),
		)
	}
	for _, f := range functors {
		f()
	}

	b.StopTimer()
}
