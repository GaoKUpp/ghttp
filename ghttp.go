package main

import (
	"fmt"
	"ghttp/command"
	"ghttp/util"
)

func main() {

	userParam := command.Parse()

	if command.Debug {
		fmt.Println(userParam.ToPrettyJsonString())
		return
	}

	result := command.Request(userParam)

	util.Echo(result)

}
