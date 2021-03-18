package temp

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func DeliveryFee() {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	if err := stdin.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	params := strings.Split(stdin.Text(), " ")
	if len(params) != 2 {
		fmt.Println("'A B'形式で入力してください")
		return
	}
	var price, deliveryFee int
	var err error
	price, err = strconv.Atoi(params[0])
	if err != nil || price < 1 || price > int(math.Pow10(6)) {
		fmt.Println("商品代金Aは1<=A<=10^6の整数(int)を入力してください。")
	}
	deliveryFee, err = strconv.Atoi(params[1])
	if err != nil || deliveryFee < 100 || deliveryFee > 5000 || deliveryFee%10 != 0 {
		fmt.Println("送料Bは100<=B<=5000の10の倍数を入力してください。")
		return
	}
	if price >= 3000 {
		if deliveryFee <= 700 {
			deliveryFee = 0
		} else {
			deliveryFee -= 700
		}
	} else {
		if deliveryFee <= 700 {
			deliveryFee /= 2
		} else {
			deliveryFee -= 350
		}
	}
	fmt.Println(price + deliveryFee)
}

// func getStdin() []string {
// 	stdin := bufio.NewScanner(os.Stdin)
// 	lines := []string{}
// 	// for stdin.Scan() {
// 	// 	if err := stdin.Err(); err != nil {
// 	// 		fmt.Fprintln(os.Stderr, err)
// 	// 	}
// 	// 	if stdin.Text() == ":q" {
// 	// 		break
// 	// 	}
// 	// 	lines = append(lines, stdin.Text())
// 	// }
// 	stdin.Scan()
// 	lines = append(lines, stdin.Text())
// 	return lines
// }
