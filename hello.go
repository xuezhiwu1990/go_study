/**
 * go study
 * zhangyu 2018-06-02
 */
package main

import (
	"fmt"
	"time"
)

var message = make( chan  string )
func stest(){
	message <- "hello goroutine!"
}

func stest2(){
	time.Sleep( 2*time.Second )
	str := <-message
	str = str + "I 'm goroutine"
	message <- str
}

func main() {

	go stest()
	go stest2()

	time.Sleep( 3*time.Second )
	str := <-message
	fmt.Println( str )
	fmt.Println("Hello, Worldsss!")
}