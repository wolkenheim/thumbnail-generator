package main

import (
	"wolkenheim.cloud/thumbnail-generator/app"
	"fmt"
)

func main(){
	fmt.Println("Hello")
}

func init()  {
	app.InitLogger()
	app.InitConfig()
}