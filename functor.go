package functor

type Functor func()
type FunctorR[R any] func() R

type Functor1[T1 any] func(T1)
type FunctorR1[R, T1 any] func(T1) R

type Functor2[T1, T2 any] func(T1, T2)
type FunctorR2[R, T1, T2 any] func(T1, T2) R

type Functor3[T1, T2, T3 any] func(T1, T2, T3)
type FunctorR3[R, T1, T2, T3 any] func(T1, T2, T3) R

// func(T1) -> func()
func MakeFunctor1_0[T1 any](f func(T1), a1 T1) Functor {
	return func() { f(a1) }
}

// func(T1) R -> func() R
func MakeFunctorR1_0[R, T1 any](f func(T1) R, a1 T1) FunctorR[R] {
	return func() R { return f(a1) }
}

// func(T1, T2) -> func()
func MakeFunctor2_0[T1, T2 any](f func(T1, T2), a1 T1, a2 T2) Functor {
	return func() { f(a1, a2) }
}

// func(T1, T2) -> func(T1)
func MakeFunctor2_1[T1, T2 any](f func(T1, T2), a1 T1) Functor1[T2] {
	return func(a2 T2) { f(a1, a2) }
}

// func(T1, T2) R -> func() R
func MakeFunctorR2_0[R, T1, T2 any](f func(T1, T2) R, a1 T1, a2 T2) FunctorR[R] {
	return func() R { return f(a1, a2) }
}

// func(T1, T2) R -> func(T1) R
func MakeFunctorR2_1[R, T1, T2 any](f func(T1, T2) R, a1 T1) FunctorR1[R, T2] {
	return func(a2 T2) R { return f(a1, a2) }
}

// func(T1, T2, T3) -> func()
func MakeFunctor3_0[T1, T2, T3 any](f func(T1, T2, T3), a1 T1, a2 T2, a3 T3) Functor {
	return func() { f(a1, a2, a3) }
}

// func(T1, T2, T3) -> func(T1)
func MakeFunctor3_1[T1, T2, T3 any](f func(T1, T2, T3), a1 T1, a2 T2) Functor1[T3] {
	return func(a3 T3) { f(a1, a2, a3) }
}

// func(T1, T2, T3) -> func(T1, T2, T3)
func MakeFunctor3_2[T1, T2, T3 any](f func(T1, T2, T3), a1 T1) Functor2[T2, T3] {
	return func(a2 T2, a3 T3) { f(a1, a2, a3) }
}

// func(T1, T2, T3) R -> func() R
func MakeFunctorR3_0[R, T1, T2, T3 any](f func(T1, T2, T3) R, a1 T1, a2 T2, a3 T3) FunctorR[R] {
	return func() R { return f(a1, a2, a3) }
}

// func(T1, T2, T3) R -> func(T1) R
func MakeFunctorR3_1[R, T1, T2, T3 any](f func(T1, T2, T3) R, a1 T1, a2 T2) FunctorR1[R, T3] {
	return func(a3 T3) R { return f(a1, a2, a3) }
}

// func(T1, T2, T3) R -> func(T1, T2) R
func MakeFunctorR3_2[R, T1, T2, T3 any](f func(T1, T2, T3) R, a1 T1) FunctorR2[R, T2, T3] {
	return func(a2 T2, a3 T3) R { return f(a1, a2, a3) }
}

// func(T1, T2, T3, T4) -> func()
func MakeFunctor4_0[T1, T2, T3, T4 any](f func(T1, T2, T3, T4), a1 T1, a2 T2, a3 T3, a4 T4) Functor {
	return func() { f(a1, a2, a3, a4) }
}

// func(T1, T2, T3, T4) -> func(T1)
func MakeFunctor4_1[T1, T2, T3, T4 any](f func(T1, T2, T3, T4), a1 T1, a2 T2, a3 T3) Functor1[T4] {
	return func(a4 T4) { f(a1, a2, a3, a4) }
}

// func(T1, T2, T3, T4) R -> func(T1, T2, T3, T4) R
func MakeFunctor4_2[T1, T2, T3, T4 any](f func(T1, T2, T3, T4), a1 T1, a2 T2) Functor2[T3, T4] {
	return func(a3 T3, a4 T4) { f(a1, a2, a3, a4) }
}

// func(T1, T2, T3, T4) R -> func(T1, T2, T3, T4) R
func MakeFunctor4_3[T1, T2, T3, T4 any](f func(T1, T2, T3, T4), a1 T1) Functor3[T2, T3, T4] {
	return func(a2 T2, a3 T3, a4 T4) { f(a1, a2, a3, a4) }
}

// func(T1, T2, T3, T4) R -> func(T1, T2, T3, T4) R
func MakeFunctorR4_0[R, T1, T2, T3, T4 any](f func(T1, T2, T3, T4) R, a1 T1, a2 T2, a3 T3, a4 T4) FunctorR[R] {
	return func() R { return f(a1, a2, a3, a4) }
}

// func(T1, T2, T3, T4) R -> func(T1, T2, T3, T4) R
func MakeFunctorR4_1[R, T1, T2, T3, T4 any](f func(T1, T2, T3, T4) R, a1 T1, a2 T2, a3 T3) FunctorR1[R, T4] {
	return func(a4 T4) R { return f(a1, a2, a3, a4) }
}

// func(T1, T2, T3, T4) R -> func(T1, T2, T3, T4) R
func MakeFunctorR4_2[R, T1, T2, T3, T4 any](f func(T1, T2, T3, T4) R, a1 T1, a2 T2) FunctorR2[R, T3, T4] {
	return func(a3 T3, a4 T4) R { return f(a1, a2, a3, a4) }
}

// func(T1, T2, T3, T4) R -> func(T1, T2, T3, T4) R
func MakeFunctorR4_3[R, T1, T2, T3, T4 any](f func(T1, T2, T3, T4) R, a1 T1) FunctorR3[R, T2, T3, T4] {
	return func(a2 T2, a3 T3, a4 T4) R { return f(a1, a2, a3, a4) }
}
