package main

import (
	"code"
	"context"
	"fmt"
	"log"
	"os"

	cli "github.com/urfave/cli/v3"
)

func main() {
	app := &cli.Command{
		Name:  "hexlet-path-size",
		Usage: "print size of a file or directory; supports -r (recursive), -H (human-readable), -a (include hidden)",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "recursive",
				Aliases: []string{"r"},
				Usage:   "recursive size of directories",
			},
			&cli.BoolFlag{
				Name:    "human",
				Aliases: []string{"H"},
				Usage:   "human-readable sizes (auto-select unit)",
			},
			&cli.BoolFlag{
				Name:    "all",
				Aliases: []string{"a"},
				Usage:   "include hidden files and directories",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error { // ← БЕЗ error!
			if cmd.Args().Len() == 0 {
				fmt.Fprintln(os.Stderr, "path required")

			}
			path := cmd.Args().Get(0)

			rawSize, err := code.GetPathSize(
				path,
				cmd.Bool("recursive"),
				cmd.Bool("human"),
				cmd.Bool("all"),
			)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %v\n", err)
				return err
			}

			formattedSize := code.FormatSize(rawSize, cmd.Bool("human"))
			fmt.Printf("%s\t%s\n", formattedSize, path)
			return nil
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
