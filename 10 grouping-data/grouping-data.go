// AKA Data Structures or Data Aggregation
package main

func main() {
	// Arrays (generally, you use slices instead of Arrays)
	// var x [5]int
	// fmt.Println(x)
	// x[3] = 42 // zero-based index
	// fmt.Println(x)
	// fmt.Println(len(x))

	// Slices
	//x := type{values} <- composite literal
	//x := []int{1, 2, 3, 4, 5, 6, 7} // SLICE = group values of the same type
	//fmt.Println(x)

	// Slice - for range
	// x := []int{1, 2, 3, 4, 5, 6, 7}
	// fmt.Println(len(x))
	// // Accessing by index position
	// fmt.Println(x)
	// fmt.Println(x[0])
	// fmt.Println(x[1])
	// fmt.Println(x[2])
	// fmt.Println(x[3])
	// // Accessing by looping
	// for i, v := range x {
	// 	fmt.Println(i, v)

	// Slicing a slice
	// x := []int{1, 2, 3, 4, 5, 6, 7}
	// fmt.Println(x)
	// fmt.Println(x[1])
	// fmt.Println(x[1:])  // slice of slice
	// fmt.Println(x[:3])  // slice of slice
	// fmt.Println(x[1:5]) // 1 to where position 5 starts

	// // for loop without range
	// for i := 0; i < len(x); i++ {
	// 	fmt.Println(x[i])
	// }

	// Slice - append to a slice
	// x := []int{1, 2, 3, 4, 5, 6, 7}
	// fmt.Println(x)
	// x = append(x, 8, 9, 10)
	// fmt.Println(x)
	// y := []int{11, 12, 13, 14, 15}
	// x = append(x, y...) // Dump values of y into x
	// fmt.Println(x)
	// fmt.Println(y)

	// Deleting a slice
	// x := []int{1, 2, 3, 4, 5, 6, 7}
	// x = append(x, 8, 9, 10)
	// y := []int{11, 12, 13, 14, 15}
	// x = append(x, y...) // Dump values of y into x
	// fmt.Println(x)
	// x = append(x[:1], x[6:]...) // Remove numbers between 2nd and 6th index
	// fmt.Println(x)

	// Slice - make
	// // x := []int{1, 2, 3, 4, 5, 6, 7} // composite literal (typical way to create a slice)
	// x := make([]int, 10, 12) // make is a built in function to make a slice a specific size and capacity
	// fmt.Println(x)
	// fmt.Println(len(x))
	// fmt.Println(cap(x))
	// x[0] = 42
	// x[1] = 999
	// fmt.Println(x)
	// // x[10] = 1000 outside the range of the slice
	// x = append(x, 1000) // extend slice to 11
	// fmt.Println(len(x))
	// x = append(x, 1000) // extend slice to 12
	// fmt.Println(len(x))
	// // note: once cap is met, the slice will re-create itself in a new, larger cap slice

	// Slice - multi-dimensional slice
	// jb := []string{"James", "Bond", "Chocolate", "Martini"}
	// fmt.Println(jb)
	// mp := []string{"Miss", "Moneypenny", "Strawberry", "Hazelnut"}
	// fmt.Println(mp)
	// xp := [][]string{jb, mp}
	// fmt.Println(xp)

	// Map - intro (similar to Python dicts)
	// Key-value store
	// m := map[string]int{ // make a map with key and values
	// 	"James":  32,
	// 	"Golden": 20,
	// } // composite literal is the entire type which is map[string]int{}
	// fmt.Println(m)
	// fmt.Println(m["James"]) // 32
	// fmt.Println(m["Someone Else"]) // returns 0
	// v, ok := m["Someone Else"]     // comma ok idium
	// fmt.Println(v)                 // 0
	// fmt.Println(ok)                // false
	// if v, ok := m["James"]; ok {
	// 	fmt.Println("THIS IS THE IF PRINT", v)
	// }

	// Map - add element & range
	// m := map[string]int{ // make a map with key and values
	// 	"James":  32,
	// 	"Golden": 20,
	// } // composite literal is the entire type which is map[string]int{}
	// m["todd"] = 33 // add a new element
	// for k, v := range m {
	// 	fmt.Println(k, v) // print all the keys and values using range
	// }

	// xi := []int{1, 2, 3, 4, 5, 6} // iterating over slice with range
	// for i, v := range xi {
	// 	fmt.Println(i, v)
	// }

	// Map - delete
	// m := map[string]int{
	// 	"Geo":            123,
	// 	"Something Else": 345,
	// }
	// fmt.Println(m)
	// delete(m, "Name")
	// fmt.Println(m)
	// delete(m, "Something Else") // If you delete something that doesn't exist, nothing happens
	// fmt.Println(m)
	// if v, ok := m["Something Else"]; ok {
	// 	fmt.Println("value:", v)
	// 	delete(m, "Something Else") // Delete something else
	// 	fmt.Println(m)
	// }

}
