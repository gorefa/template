package main


import (
	"fmt"

	"gorefa/template/config"
	"github.com/spf13/pflag"
)

var (
	cfg  = pflag.StringP("config", "c", "", "apiserver config file path.")
)

func main() {
	pflag.Parse()

	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	fmt.Println("Init success!")

}