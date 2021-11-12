package main

type Keys interface {
	parse()
	verify()
	generate()
}
