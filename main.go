package main

import (
	"fmt"
	"os"

	"github.com/pierskarsenbarg/provider-base/pkg"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi-go-provider/middleware/schema"
	gen "github.com/pulumi/pulumi/pkg/v3/codegen/go"
	"github.com/pulumi/pulumi/sdk/go/common/tokens"
)

func main() {
	err := p.RunProvider("base", "0.1.0", provider())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		os.Exit(1)
	}
}

func provider() p.Provider {
	return infer.Provider(infer.Options{
		Metadata: schema.Metadata{
			DisplayName: "gitpod",
			Description: "Gitpod provider",
			LanguageMap: map[string]any{
				"go": gen.GoPackageInfo{
					Generics:       gen.GenericsSettingGenericsOnly,
					ImportBasePath: "github.com/pierskarsenbarg/provider-base/sdk/go/base",
				},
			},
		},
		Resources: []infer.InferredResource{
			infer.Resource[*pkg.Account, pkg.AccountArgs, pkg.AccountState](),
		},
		ModuleMap: map[tokens.ModuleName]tokens.ModuleName{
			"pkg": "index", // required because the folder with everything in is "pkg"
		},
		Functions: []infer.InferredFunction{
			infer.Function[*pkg.GetAccount, pkg.GetAccountArgs, pkg.AccountState](),
		},
		Config: infer.Config[*pkg.Config](),
	})
}
