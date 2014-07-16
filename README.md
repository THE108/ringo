Ringo
=====

Simple fixed size ring container.

Copyleft Andrew Chernov <chernov.andrey1988@gmail.com>

Library licensed under GPLv3.

Install
-------

	go get github.com/the108/ringo

Examples
--------

Create ring, fill it and dump. You can use ShiftLeft to shift current pointer to next item.

```go
	// Create ring of size 10
	ring := ringo.New(5)
	for i := 0; i < 5; i++ {
		// Add some item
		ring.Insert(i)
	}

	// Dump in current position
	dump := ring.Dump()

	fmt.Printf("%+v\n", dump) // [0 1 2 3 4]

	ring.ShiftLeft()

	dump2 := ring.Dump()

	fmt.Printf("%+v\n", dump2) // [1 2 3 4 0]

	// Add some item again, current item will be overwritten
	prev := ring.Insert(100)
	// Previous item returned
	fmt.Printf("%+v\n", prev) // 1

	// Dump in current position
	dump3 := ring.Dump()

	fmt.Printf("%+v\n", dump3) // [2 3 4 0 100]

```