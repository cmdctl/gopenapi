package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"gopenapi/pkg/types"
	"io/ioutil"
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
					pkg := c.String("package")
					fmt.Println("Generating types from spec:", spec)
					fmt.Println("Output:", output)
					fmt.Println("Package:", pkg)
					_, b, err := types.Generate(types.Params{
						PackageName: pkg,
						OutputFile:  output,
						SpecFile:    spec,
					})
					if err != nil {
						return cli.Exit(err.Error(), 1)
					}
					err = ioutil.WriteFile(output, b, 0644)
					if err != nil {
						return cli.Exit(err.Error(), 1)
					}
					return nil
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "output",
						Aliases:     []string{"o"},
						Usage:       "gopenapi types './petstore.yaml -o ./path/to/output/file.go'",
						Required:    false,
						Value:       "types.go",
						DefaultText: "Set a path to output file",
					},
					&cli.StringFlag{
						Name:        "package",
						Aliases:     []string{"p"},
						Usage:       "gopenapi types './petstore.yaml --package=main'",
						Required:    true,
						Value:       "main",
						DefaultText: "Set a package name for the generated types",
					},
				},
			},
		},
	}
	return app
}
