package main

import (
	"fmt"
	"github.com/nullstyle/coinop/cmd"
	"github.com/spf13/viper"
	"os"
)

func main() {
	if err := cmd.Root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	viper.SetDefault("horizon-url", "https://horizon-testnet.stellar.org")
	viper.SetDefault("postgres-url", "postgres://localhost?sslmode=disable")
}
