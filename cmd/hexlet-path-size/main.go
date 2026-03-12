package main

import (
    "context"
    "fmt"
    "log"
    "os"

    "github.com/urfave/cli/v3"
    "hexlet-path-size/code"
)

func main() {
    cmd := &cli.Command{
        Name:  "hexlet-path-size",
        Usage: "print size of a file or directory; supports -r (recursive), -H (human-readable), -a (include hidden)",
        Flags: []cli.Flag{
            &cli.BoolFlag{
                Name:    "recursive",
                Aliases: []string{"r"},
                Usage:   "recursive size of directories",
                Value:   false,
            },
            &cli.BoolFlag{
                Name:    "human",
                Aliases: []string{"H"},
                Usage:   "human-readable sizes (auto-select unit)",
                Value:   false,
            },
            &cli.BoolFlag{
                Name:    "all",
                Aliases: []string{"a"},
                Usage:   "include hidden files and directories",
                Value:   false,
            },
            &cli.BoolFlag{
                Name:    "help",
                Aliases: []string{"h"},
                Usage:   "show help",
            },
        },
        Action: func(ctx context.Context, c *cli.Command) error {
            if c.Bool("help") {
                return cli.ShowAppHelp(c)
            }

            if c.NArg() == 0 {
                return fmt.Errorf("path is required")
            }

            path := c.Args().First()
            
            if _, err := os.Stat(path); err != nil {
                return fmt.Errorf("path does not exist: %s", path)
            }

            size, err := code.GetPathSize(path, c.Bool("recursive"), c.Bool("all"))
            if err != nil {
                return fmt.Errorf("error calculating size: %w", err)
            }

            formattedSize := code.FormatSize(size, c.Bool("human"))
            fmt.Printf("%s\t%s\n", formattedSize, path)
            return nil
        },
    }

    if err := cmd.Run(context.Background(), os.Args); err != nil {
        log.Fatal(err)
    }
}