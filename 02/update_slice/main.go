package main

import "fmt"

func main() {
	//arr_string := []string{"a", "b", "c"}
	var arr_string []string

	update_a_slice(&arr_string)
	fmt.Println(arr_string)

}

func update_a_slice(ref_arr *[]string) {
	if *ref_arr == nil {
		*ref_arr = []string{"a", "b", "c"}
		return
	}
	(*ref_arr)[0] = "x"
	*ref_arr = append(*ref_arr, "d")
}
