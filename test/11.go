package main

import (
	"fmt"
	apk2 "gitee.com/mryy1996/androidbinary/apk"
)

func main() {

	apk, _ := apk2.OpenFile("apk/testdata/touyingtong3.0.1_2265.com.apk")

	fmt.Println(apk.Label(nil))

	//apk2.OpenFile()

}
