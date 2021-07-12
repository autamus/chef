package cli

import (
	"fmt"
	"github.com/DataDrake/cli-ng/v2/cmd"
	"github.com/autamus/chef/config"
	"github.com/autamus/chef/container"
)

// GenerateArgs holds arguments for Generate
type GenerateArgs struct {
	Chefyaml []string `zero:"true" desc:"The chef config file to read packages from."`
}

type GenerateFlags struct {
	SkipValidation bool `zero:"true" short:"s" long:"skip-validation" desc:"Skip validating the names and tags (requests internet connection).`
}

// Generate loads a chef configuration file and generates a Dockerfile
var Generate = cmd.Sub{
	Name:  "generate",
	Alias: "g",
	Short: "Generate a Dockerfile from a chef config file.",
	Flags: &GenerateFlags{},
	Args:  &GenerateArgs{},
	Run:   RunGenerate,
}

func init() {
	cmd.Register(&Generate)
}

// RunGenerate loads the config file and generates a Dockerfile
func RunGenerate(r *cmd.Root, c *cmd.Sub) {
	args := c.Args.(*GenerateArgs)
	flags := c.Flags.(*GenerateFlags)

	// If a config isn't provided, chef.yaml is the default
	if len(args.Chefyaml) == 0 {
		args.Chefyaml = append(args.Chefyaml, "chef.yaml")
	}
	fmt.Println(flags.SkipValidation)
	// Create and load a new config
	conf := config.Load(args.Chefyaml[0])

	// Generate the Dockerfile template
	dockerfile := container.Dockerfile(conf.Packages, !flags.SkipValidation)
	fmt.Println(dockerfile)
}
