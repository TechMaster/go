```go
greeting()

	indexLoop()
	fmt.Println("-----")
	rangeLoop()
	fmt.Println("-----")
	reverseLoop()
	fmt.Println("-----")
	demoArray2D()
	fmt.Println("-----")
	demoSlice()
	fmt.Println("-----")
	demoStruct()
	fmt.Println("-----")
	demoEnglishDictionary()
	fmt.Println("-----")
	demoMapInterface()
	fmt.Println("-----")
	demoSlice()

	fmt.Println("-----")
	a := removeItemSliceNotKeepOrder([]string{"a", "b", "c", "d", "e", "f"}, 2)
	fmt.Println(a)

	fmt.Println("-----")
	a = removeItemSliceNotKeepOrder2([]string{"a", "b", "c", "d", "e", "f"}, 2)
	fmt.Println(a)

	fmt.Println("-----")
	a = removeItemSliceKeepOrder([]string{"a", "b", "c", "d", "e", "f"}, 2)
	fmt.Println(a)

	fmt.Println("-----")
	a = removeItemSliceKeepOrder2([]string{"a", "b", "c", "d", "e", "f"}, 2)
	fmt.Println(a)

	fmt.Println("-----")
	b := []string{"a", "b", "c", "d", "e", "f"}
	reverseNotKeepInput(b)
	fmt.Println(b)

	fmt.Println("-----")
	b = []string{"a", "b", "c", "d", "e", "f"}
	c := reverseKeepInput(b)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Println("-----")
	b = []string{"a", "b", "b", "c", "d", "e", "f", "a"}
	c = removeDuplicate(b)
	fmt.Println(c)

	fmt.Println("-----")
	d := []interface{}{"a", true, 1, "abc", 'e', 3.14}
	e := reverseSliceAnyType(d)
	fmt.Println(d)
	fmt.Println(e)

	fmt.Println("-----")
	sortIntSlice()

	fmt.Println("-----")
	sortReverse()

	fmt.Println("-----")
	sortSliceWithFunc()

	fmt.Println("-----")
	search()
```