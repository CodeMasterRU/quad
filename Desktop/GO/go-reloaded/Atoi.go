package student

func Atoi(s string) int {
	lenght := Strlen(s)
	result := 0
	chaine := s

	if s[0] == '-' || s[0] == '+' {
		chaine = s[1:]
	} else {
		chaine = s
	}
	for index, char := range chaine {
		if char >= 48 && char <= 57 {
			chiffre := int(char) - 48
			if chiffre != 0 {
				result += chiffre * RecursivePower(10, lenght-index-1)
			}

		} else {
			return 0
		}
	}
	if s[0] == '-' {
		result = -result
	}
	return result

}
func Strlen(s string) int {
	lenght := 0
	for range s {
		lenght++
	}
	return lenght
}
func RecursivePower(nb int, power int) int {
	res := 0
	if power == 0 {
		res = 1
	} else {
		res = RecursivePower(nb, power-1) * nb
	}
	return res

}
