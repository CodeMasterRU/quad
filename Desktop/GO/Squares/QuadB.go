package challenge

import "github.com/01-edu/z01"

func QuadB(x, y int) {

	if x == 1 && y == 1 {
		z01.PrintRune('/')
	}

	if x > 1 && y == 1 {
		z01.PrintRune('/')
		for index1 := 2; index1 < x; index1++ {
			z01.PrintRune('*')
		}
		z01.PrintRune(92)
	}
	if x == 1 && y > 1 {
		z01.PrintRune('/')
		for index2 := 2; index2 < y; index2++ {
			z01.PrintRune('*')
			z01.PrintRune('\n')

		}
		z01.PrintRune(92)
	}
	if x > 1 && y > 1 {
		z01.PrintRune('/')
		for index3 := 2; index3 < x; index3++ {
			z01.PrintRune('*')
		}
		z01.PrintRune(92)
		z01.PrintRune('\n')

		for index4 := 2; index4 < y; index4++ {
			z01.PrintRune('*')
			for index5 := 2; index5 < x; index5++ {
				z01.PrintRune(' ')
			}
			z01.PrintRune('*')
			z01.PrintRune('\n')
		}
		z01.PrintRune(92)
		for index6 := 2; index6 < x; index6++ {
			z01.PrintRune('*')
		}
		z01.PrintRune('/')

	}

}
