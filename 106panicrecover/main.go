package main

/*

 */

import (
	"fmt"
	"time"
)

func main() {
	err := sub()

	fmt.Println("sub() 回来了", err)
	// time.Sleep(time.Millisecond * 10000)
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

	for {
		fmt.Println("执行: ")

		fmt.Println("返回 sub01() : ", sub01())
		sub01()
	}

	// if err != nil {
	// 	return fmt.Errorf("error: %v", err)
	// }
	return nil
}

func sub01() string {
	for {
		err := sub02()
		if err == nil {
			// Need get data to find backend errors.
			err = nil
			if err == nil {
				return "正常返回"
			}
			// Sleep, retry.
		}
		time.Sleep(time.Second * 2)
	}

}

func sub02() error {
	panic("人为的恐慌")
	return nil
}
