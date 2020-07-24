package main

import (
	"fmt"
)

func main() {
	
	var wc int
	fmt.Println("Enter the week count?")
	fmt.Scanf("%v",&wc)

	switch wc{
	case 1:
		fmt.Println("Sunday")
		break
	case 2:
		fmt.Println("Monday")
		break
	case 3:
		fmt.Println("Tuesday")
		break
	case 4:
		fmt.Println("Wednesday")
		break
	case 5:
		fmt.Println("Thursday")
		break
	case 6:
		fmt.Println("Friday")
		break
	case 7:
		fmt.Println("Saturday")
		break
	default:
		fmt.Println("Select valid option")
		break
	}

}