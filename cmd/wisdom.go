package main

import (
	"fmt"
	"os"

	"github.com/priyanka-jorasia/wisdom/lib/wisdom"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "wisdom"
	app.Usage = "dispense some programme quote"
	app.UsageText = "wisdom dispense"
	app.Version = wisdom.Version

	app.Commands = []cli.Command{
		cli.Command{
			Name:      "dispense",
			Usage:     "dispense a single programming fortune cookie",
			UsageText: "wisdom dispense",
			Action:    DispenseWisdom,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func DispenseWisdom(c *cli.Context) error {
	fmt.Println("Programming is fun")
	return nil
}
