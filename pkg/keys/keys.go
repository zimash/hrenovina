package main

type Keys interface {
	parse()
	verify()
	generate()
}

type MockKey struct {
	key string
}

func (mk *MockKey) pasrse() {
}

func (mk *MockKey) verify() {
}

func (mk *MockKey) generate() {
}
