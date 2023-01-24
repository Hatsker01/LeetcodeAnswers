package main

func reverseString(s []byte) {
	var k []byte
	for i := 0; i < len(s); i++ {
		k = append(k, s[i])
	}

	for i := 0; i < len(s); i++ {

		s[i] = k[len(s)-i-1]
	}

}
