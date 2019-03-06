package main

// Specify Filter function here
func Filter(arr []int, predicate func(elem, _ int)bool) []int{
	var result []int
	for _, value := range arr{
		if predicate(value, value) {
			result = append(result, value)
		}
	}
	return result
}



func main() {

}
