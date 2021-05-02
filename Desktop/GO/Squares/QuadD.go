package challenge

import "github.com/01-edu/z01"

func QuadD(x, y int) {

	if x == 1 && y == 1 {
		z01.PrintRune('A')
	}

	if x > 1 && y == 1 {
		z01.PrintRune('A')
		for index1 := 2; index1 < x; index1++ {
			z01.PrintRune('B')
		}
		z01.PrintRune('C')
	}
	if x == 1 && y > 1 {
		z01.PrintRune('A')
		for index2 := 2; index2 < y; index2++ {
			z01.PrintRune('B')
			z01.PrintRune('\n')

		}
		z01.PrintRune('A')
	}
	if x > 1 && y > 1 {
		z01.PrintRune('A')
		for index3 := 2; index3 < x; index3++ {
			z01.PrintRune('B')
		}
		z01.PrintRune('C')
		z01.PrintRune('\n')

		for index4 := 2; index4 < y; index4++ {
			z01.PrintRune('B')
			for index5 := 2; index5 < x; index5++ {
				z01.PrintRune(' ')
			}
			z01.PrintRune('B')
			z01.PrintRune('\n')
		}

		z01.PrintRune('A')
		for index6 := 2; index6 < x; index6++ {
			z01.PrintRune('B')
		}
		z01.PrintRune('C')
	}

}
