package main

import "time"

func sleep(sec int) {
	<-time.NewTimer(time.Second * time.Duration(sec)).C
}
func main() {
	sleep(5)
}
