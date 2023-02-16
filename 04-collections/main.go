package main

import "fmt"

func main() {
	getArraysStructures()
	getSlicesStructures()
	getMapsStructures()
}

/*
Arrays are immutable list of anything
*/
func getArraysStructures() {
	fmt.Println("----------- ARRAYS -----------")

	// static arrays
	var staticEmptyNumberArray [5]int
	fmt.Println("empty array of", len(staticEmptyNumberArray), "elements", staticEmptyNumberArray)

	// assigning values to static array
	staticEmptyNumberArray[1] = 34
	staticEmptyNumberArray[0] = 3
	fmt.Println("array of", len(staticEmptyNumberArray), "elements", staticEmptyNumberArray)

	// creation and assignation of a array
	names := [3]string{"ale", "peter", "juan"}
	fmt.Println(len(names), "names creating and using", names)

	// creating array without say how long it is. But it have a size, that is fixed
	nonDeclaredSizeArray := [...]string{"1", "2", "3", "4"}
	fmt.Println("array of", len(nonDeclaredSizeArray), "elements, data:", nonDeclaredSizeArray)

}

/*
Slices are mutable list of anything
*/
func getSlicesStructures() {
	fmt.Println("----------- SLICES -----------")

	// declare initial slice
	numbers := []int{1, 2, 3}
	fmt.Println("data: ", numbers, ", size: ", len(numbers))

	// add data to slice
	numbers = append(numbers, 4)
	numbers = append(numbers, 5)
	fmt.Println("data: ", numbers, ", size: ", len(numbers))

	// slices has a pointer, length and capacity
	// length: it is the quantity of data
	// capacity: GO has reserved the largest array for this array; if I add more items, he will reserve more information.
	// pointer: If GO needs to add more items and the reserved capacity is insufficient, he will reserve additional larger memory space, and the pointer is a reference to it.
	fmt.Printf("Len: %v, Cap: %v, Point: %p \n", len(numbers), cap(numbers), numbers)

	// make function
	names := make([]string, 0, 3) // data type, length, capacity
	fmt.Println("data: ", names)
	fmt.Printf("Len: %v, Cap: %v, Point: %p \n", len(names), cap(names), names)

	names = append(names, "1")
	fmt.Println("data: ", names)
	fmt.Printf("Len: %v, Cap: %v, Point: %p \n", len(names), cap(names), names)
}

/*
Pairs key and value, it isn't ordered
*/
func getMapsStructures() {
	fmt.Println("----------- MAPS -----------")
	weekDays := make(map[string]int) // map[keyType]valueType

	weekDays["sunday"] = 0
	weekDays["monday"] = 1

	fmt.Println("weekDays: ", weekDays, ", size: ", len(weekDays))

	// map of arrays
	studentsBySubject := make(map[string][]string)
	studentsBySubject["math"] = []string{"Ale", "Juan"}
	studentsBySubject["english"] = []string{"Ale", "July"}
	fmt.Println("studentsBySubject: ", studentsBySubject, ", size: ", len(studentsBySubject))

	mathStudents := studentsBySubject["math"]
	fmt.Println("mathStudents: ", mathStudents, ", size: ", len(mathStudents))

	firstMathStudent := studentsBySubject["math"][0]
	fmt.Println("firstMathStudent:", firstMathStudent)

	// knowing if exist a value for a key
	users := make(map[string]string)
	users["ale"] = "ale@gmail.com"
	users["mirco"] = "mirco@gmail.com"

	_, okJuan := users["juan"]
	fmt.Printf("Juan has registered email %v\n", okJuan)

	names := [...]string{"ale", "juan", "mirco"}
	for i := 0; i < len(names); i++ {
		name := names[i]

		if email, ok := users[name]; ok {
			fmt.Println("Email for user", name, "is", email)
		} else {
			fmt.Println("Email for user", name, "doesn't exists")
		}
	}
}
