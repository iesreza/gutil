package byte

func Contains(needle []byte, haystack []byte) bool {
	indicator := 0
	for _, haystackByte := range haystack {
		if haystackByte == needle[indicator] {
			indicator++
			if indicator == len(needle) {
				return true
			}
		} else {
			indicator = 0
		}
	}

	return false
}

/*func Like(needle string,haystack []byte)  bool{

}*/
