package generics

// Foo ...
type Foo struct {
	Bar  *Bar
	Name string
}

// NewFoo ...
func NewFoo(name string, bar *Bar) Foo {
	return Foo{
		Bar:  bar,
		Name: name,
	}
}

// Bar ...
type Bar struct {
	Name string
}

// NewBar ...
func NewBar(name string) Bar {
	return Bar{
		Name: name,
	}
}

// Baz ...
type Baz struct {
	Name string
	Foos []Foo
}

// NewBaz ...
func NewBaz(name string, foos ...Foo) Baz {
	return Baz{
		Name: name,
		Foos: foos,
	}
}

// NewFooBarPair ...
func NewFooBarPair(name string) (Foo, Bar) {
	bar := NewBar(name)
	foo := NewFoo(name, &bar)

	return foo, bar
}

// NewFooBarBazGroup ...
func NewFooBarBazGroup(name string) (Foo, Bar, Baz) {
	bar := NewBar(name)
	foo := NewFoo(name, &bar)
	baz := NewBaz(name, foo)

	return foo, bar, baz
}
