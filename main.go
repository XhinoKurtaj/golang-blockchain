package main

import (
	"os"

	// "github.com/XhinoKurtaj/golang-blockchain/cli"
	"github.com/XhinoKurtaj/golang-blockchain/wallet"
)

func main() {
	defer os.Exit(0)
	// cli := cli.CommandLine{}
	// cli.Run()

	w := wallet.MakeWallet()
	w.Address()

}
