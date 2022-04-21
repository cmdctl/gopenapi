package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

func InitApp() *cli.App {
	app := &cli.App{
		Name:  "gopenapi",
		Usage: "tool to generate go source code from OpenAPI 3 spec",
		Commands: []*cli.Command{
			{
				Name:      "types",
				Aliases:   []string{"t"},
				Usage:     "Generate go model structs from OpenAPI 3 spec",
				UsageText: "gopenapi types [options] <spec.yaml>",
				Action: func(c *cli.Context) error {
					spec := c.Args().First()
					if spec == "" {
						return cli.Exit("spec file is required", 1)
					}
					output := c.String("output")
					fmt.Println("Generating types from spec:", spec)
					fmt.Println("Output:", output)
					return nil
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "output",
						Aliases:     []string{"o"},
						Usage:       "openapi types './petstore.yaml -o ./path/to/output/file.go'",
						Required:    false,
						Value:       "generated.go",
						DefaultText: "Set a path to output file",
					},
				},
			},
		},
	}
	return app
}
