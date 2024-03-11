package main

import "fmt"

func main() {
	{
		// a := 1
		if false {
			fmt.Println("true")
		} else {
			{
				b := 2
				if false {

				} else {

					{
						c := 3
						if false {

						} else {
							// println(a, b, c)
							println(b, c)
						}
					}
				}
			}
		}
	}
}
