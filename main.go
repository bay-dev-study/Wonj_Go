package main

import "fmt"

func iterateNums(numbers ...int) {
	for index, number := range numbers {
		fmt.Println(index, number)
	}
	// Same
	for i := 0; i < len(numbers); i++ {
		fmt.Println(i, numbers[i])
	}
}

func canIDrink(age int) bool {
	// variable expression
	if koreanAge := age + 2; koreanAge < 20 {
		return false
	}
	return true
}

func switchExample(num int) bool {
	// Ordinary switch statement
	// switch numPlusTwo := num + 2; numPlusTwo {
	// case 10:
	// 	return false
	// case 20:
	// 	return true
	// default:
	// 	return false
	// }

	// if-like switch statement
	switch {
	case num < 20:
		return false
	case num > 70:
		return false
	default:
		return true
	}
}

func main() {
	iterateNums(2, 3, 4, 5, 6)
	fmt.Println(canIDrink(18))
	fmt.Println(switchExample(20))

	a := 900000000000
	b := 5
	fmt.Println(&a, &b)
}
