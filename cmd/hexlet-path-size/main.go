package main

import (
	"context"
	"fmt"
	"hexlet-path-size/PathSize"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	fmt.Println("Hello from Hexlet!")
	cmd := &cli.Command{
		Name:  "hexlet-path-size",
		Usage: "print size of a file or directory",
		Action: func(c *cli.Context) error {
			if c.Args().Len() == 0 {
				return fmt.Errorf("path required")
			}

			path := c.Args().Get(0)
			size := PathSize.GetSize(path)
			fmt.Printf("%d\t%s\n", size, path)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
