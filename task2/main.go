package main

// Specify MapTo function here
func MapTo(arr []int, converter func(elem, _ int)string) []string{
	var result []string
	for _, value := range arr{
		{
			result = append(result, converter(value, value))
		}
	}
	return result
}

func Convert(arr []int) []string {
	// converting nimber code should be here
	cnv := func(elem, _ int) string {
		elements := make(map[int]string)
		elements[1] = "one"
		elements[2] = "two"
		elements[3] = "three"
		elements[4] = "four"
		elements[5] = "five"
		elements[6] = "six"
		elements[7] = "seven"
		elements[8] = "eight"
		elements[9] = "nine"
		s := elements[elem]
		if s == "" {
			s = "unknown"
		}
		return s
	}
	result := MapTo(arr, cnv)
	return result
}
func main() {
}
