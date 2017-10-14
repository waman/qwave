package bb84

func AppendMatchingBit(key, bits, matches []bool, max int) []bool {
	for i, match := range matches {
		if match {
			key = append(key, bits[i])
			if len(key) == max { break}
		}
	}
	return key
}