package utils

func StringSlicePermutations(input []string, length int) [][]string {
	if length == -1 {
		length = len(input)
	}
	if length == 1 {
		return [][]string{
			input,
		}
	}

	outputs := [][]string{}
	for i := 0; i < length; i++ {
		tmpInput := make([]string, len(input))
		copy(tmpInput, input)
		subs := StringSlicePermutations(append(tmpInput[:i], tmpInput[i+1:]...), length-1)
		for _, j := range subs {
			output := []string{}
			output = append(output, input[i])
			output = append(output, j...)
			outputs = append(outputs, output)
		}
	}
	return outputs
}
