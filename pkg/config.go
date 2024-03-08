package pkg

import (
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type Config struct {
	AccessToken string `pulumi:"accessToken" provider:"secret"`
	Environment string `pulumi:"environment,optional"`
}

var _ = (infer.Annotated)((*Config)(nil))

func (c *Config) Annotate(a infer.Annotator) {
	a.Describe(&c.AccessToken, "Your access token")
	a.Describe(&c.Environment, "Environment")
}

var _ = (infer.CustomConfigure)((*Config)(nil))

func (c *Config) Configure(ctx p.Context) error {

	if len(c.Environment) == 0 {
		c.Environment = "dev"
	}

	return nil
}
