package main

import "fmt"

func main() {
	var dia int
	var mes int

	fmt.Scan(&dia)
	fmt.Scan(&mes)

	switch mes {
	case 1:
		if dia >= 1 && dia <= 20 {
			fmt.Println("Capricornio")
		} else if dia >= 21 && dia <= 31 {
			fmt.Println("Acuario")
		}
	case 2:
		if dia >= 1 && dia <= 19 {
			fmt.Println("Acuario")
		} else if dia >= 20 && dia <= 29 {
			fmt.Println("Piscis")
		}
	case 3:
		if dia >= 1 && dia <= 20 {
			fmt.Println("Piscis")
		} else if dia >= 21 && dia <= 31 {
			fmt.Println("Aries")
		}
	case 4:
		if dia >= 1 && dia <= 20 {
			fmt.Println("Aries")
		} else if dia >= 21 && dia <= 30 {
			fmt.Println("Tauro")
		}
	case 5:
		if dia >= 1 && dia <= 20 {
			fmt.Println("Tauro")
		} else if dia >= 21 && dia <= 31 {
			fmt.Println("Geminis")
		}
	case 6:
		if dia >= 1 && dia <= 21 {
			fmt.Println("Geminis")
		} else if dia >= 22 && dia <= 30 {
			fmt.Println("Cancer")
		}
	case 7:
		if dia >= 1 && dia <= 22 {
			fmt.Println("Cancer")
		} else if dia >= 23 && dia <= 31 {
			fmt.Println("Leo")
		}
	case 8:
		if dia >= 1 && dia <= 23 {
			fmt.Println("Leo")
		} else if dia >= 24 && dia <= 31 {
			fmt.Println("Virgo")
		}
	case 9:
		if dia >= 1 && dia <= 22 {
			fmt.Println("Virgo")
		} else if dia >= 23 && dia <= 30 {
			fmt.Println("Libra")
		}
	case 10:
		if dia >= 1 && dia <= 22 {
			fmt.Println("Libra")
		} else if dia >= 23 && dia <= 31 {
			fmt.Println("Escorpio")
		}
	case 11:
		if dia >= 1 && dia <= 22 {
			fmt.Println("Escorpio")
		} else if dia >= 23 && dia <= 30 {
			fmt.Println("Sagitario")
		}
	case 12:
		if dia >= 1 && dia <= 21 {
			fmt.Println("Sagitario")
		} else if dia >= 22 && dia <= 31 {
			fmt.Println("Capricornio")
		}
	}
}
