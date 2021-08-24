package piscine

func IndexWithNotes(s string, toFind string) int {
	string := []rune(s)
	// this converts the string into the runes
	substring := []rune(toFind)
	// this converts the toFind substring into runes
	counter := 0
	if len(toFind) > len(s) {
		return -1
		// if the toFind substring is longer than the string, then that the toFind substring will
		// never be found and so return -1
	} else {
		for i := 0; i <= len(s)-len(toFind); i++ {
			//otherwise, starting from index 0, go through every letter the string (i)
			// until the last index that can fit the length of the toFind substring.
			// for example:
			// if s was 'encyclopedia' (which has 12 letters )and toFind was 'clo'(3)
			// i<= len(s)-len(toFind) means :
			// i<= 12-3
			// i<=9
			// so index will look for 'clo' until index 9 as that is the last starting index
			// that can fit the toFind substring.
			for j := 0; j < len(toFind); j++ {
				// this is for the toFind substring.
				// This is saying that, starting from index 0, until the last index of the
				// substring (which is length of the word -1, hence j<len(toFind))
				if string[j+i] == substring[j] {
					counter += 1
					// if the letter in the string at position [j+i] --> so if
					// the letter at [0+4]== substring[0]
					// and letter at [1+4]== substring[1]
					//(remember that the this is a loop inside the loop hence why i
					// stays the same as it is the starting index and we are trying to see if
					// following letters at the following sequence will also match the
					// the rest of the sequence of the substring)
					// matches the corresponding letters of the substring (toFind)
					// (remember this is a loop)
					// then add 1 to the counter
				} else {
					counter = 0
					// otherwise return the counter back back to 0
					// so if letter at [2+4] != substring[2] then the
					// counter will return to 0 regardless of the previously added
				}
				if counter == len(toFind) {
					return i
					// if the counter is equivalent to the length of the substring
					// meaning if the counter has found all the letters in the
					// substring, then return i.
				}
			}
		}
		return -1
	}
}
