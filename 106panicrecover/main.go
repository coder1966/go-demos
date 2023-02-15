package main

/*

 */

import (
	"fmt"
)

func main() {
	err := sub()

	fmt.Println("sub() 回来了", err)
	// time.Sleep(time.Millisecond * 200000)
}

func sub() (err error) {
	defer func() {
		if errPanic := recover(); errPanic != nil {
			fmt.Println("errPanic", errPanic)
			err = fmt.Errorf("error panic: %v", errPanic)
		} else {
			fmt.Println("else errPanic", errPanic)
		}
	}()
	sub01()
	// if err != nil {
	// 	return fmt.Errorf("error: %v", err)
	// }
	return err
}

func sub01() {
	sub02()
}

func sub02() {
	panic("人为的恐慌")
}
