package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(promt string, r *bufio.Reader) (string, error) {
	fmt.Print(promt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)
	name, _ := getInput("create a new bill name:", reader)

	b := newBill(name)
	fmt.Println("create the bill - ", b.name)

	return b

}

func promtOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("choose option(a - add item, s - save bill, t - add tip):", reader)
	// fmt.Println(opt)
	switch opt {
	case "a":
		name, _ := getInput("Item name:", reader)
		price, _ := getInput("Item price:", reader)

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("the price must be a number")
			promtOptions(b)
		}
		b.addItem(name, p)

		fmt.Println("the item was added ")
		promtOptions(b)
	case "t":
		tip, _ := getInput("enter tip amount:", reader)
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("the tip must be a number")
		}
		b.updateTup(t)
		fmt.Println("tip was added")
		promtOptions(b)
	case "s":
		b.save()
		fmt.Println("you save the bill", b.name)
	default:
		fmt.Println("that was not a valid option.....")
		promtOptions(b)
	}

}

func main() {
	mybill := createBill()
	promtOptions(mybill)
}
