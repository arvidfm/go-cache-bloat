package common

type Pair[A, B any] struct {
	Left  A
	Right B
}

func (p Pair[A, B]) Func1(a []Pair[A, B]) []Pair[A, B] {
	return appendData(a, p)
}

func (p Pair[A, B]) Func2(a []Pair[A, B]) []Pair[A, B] {
	return appendData(a, p)
}

func (p Pair[A, B]) Func3(a []Pair[A, B]) []Pair[A, B] {
	return appendData(a, p)
}

func (p Pair[A, B]) Func4(a []Pair[A, B]) []Pair[A, B] {
	return appendData(a, p)
}

func (p Pair[A, B]) Func5(a []Pair[A, B]) []Pair[A, B] {
	return appendData(a, p)
}

func (Pair[A, B]) Func6() MyInterface[A] { return nil }

func (Pair[A, B]) Func7() MyInterface[B] { return nil }

func (Pair[A, B]) Func8() MyInterface[Pair[A, B]] { return nil }

func (Pair[A, B]) Func9() MyInterface[Pair[A, int]] { return nil }

func (Pair[A, B]) Func10() MyInterface[Pair[B, int]] { return nil }

func (p Pair[A, B]) Func11() {
	// this results in a bunch of extra instantiations that significantly bloat the object file size
	switch any(p).(type) {
	case Pair[A, int], Pair[A, int64], Pair[A, int32], Pair[A, int16], Pair[A, int8]:
	case Pair[A, uint], Pair[A, uint8], Pair[A, uint16], Pair[A, uint32], Pair[A, uint64], Pair[A, uintptr]:
	case Pair[A, float32], Pair[A, float64]:
	}
}

type MyInterface[A any] interface {
	InterfaceFunc1()
	InterfaceFunc2()
	InterfaceFunc3()
	InterfaceFunc4()
	InterfaceFunc5()
	InterfaceFunc6()
	InterfaceFunc7()
	InterfaceFunc8()
	InterfaceFunc9()
	InterfaceFunc10()
}

// marking this as go:noinline cuts the file size by 30%
func appendData[A, B any](a []Pair[A, B], value Pair[A, B]) []Pair[A, B] {
	a = append(a, value)
	a = append(a, value)
	a = append(a, value)
	a = append(a, value)
	a = append(a, value)
	a = append(a, value)
	a = append(a, value)
	a = append(a, value)
	a = append(a, value)
	a = append(a, value)
	return a
}
