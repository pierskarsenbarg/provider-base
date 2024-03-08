package pkg

import (
	"github.com/google/uuid"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type Account struct {
}

func (acc *Account) Annotate(a infer.Annotator) {
	a.Describe(&acc, "Account to define")
}

type AccountArgs struct {
	Name string `pulumi:"name,optional"`
}

func (acc *AccountArgs) Annotate(a infer.Annotator) {
	a.Describe(&acc.Name, "Name of account")
}

type AccountState struct {
	Id          string `pulumi:"accountId"`
	Name        string `pulumi:"name"`
	Environment string `pulumi:"environment"`
}

func (acc *AccountState) Annotate(a infer.Annotator) {
	a.Describe(&acc.Id, "Id of account created")
	a.Describe(&acc.Name, "Name of account created")
	a.Describe(&acc.Environment, "Environment of account")
}

func (*Account) Create(ctx p.Context, name string, input AccountArgs, preview bool) (
	id string, output AccountState, err error) {
	if !preview {
		config := infer.GetConfig[Config](ctx)
		environment := config.Environment
		accountId := uuid.New().String()

		return name, AccountState{
			Id:          accountId,
			Name:        name,
			Environment: environment,
		}, nil
	}
	return "", AccountState{}, nil
}

func (*Account) Delete(ctx p.Context, id string, props AccountState) error {
	return nil
}

func (*Account) Read(ctx p.Context, id string, inputs AccountArgs, state AccountState) (
	string, AccountArgs, AccountState, error) {
	return inputs.Name, AccountArgs{
			Name: state.Name,
		}, AccountState{
			Id:          state.Id,
			Name:        state.Name,
			Environment: state.Environment,
		}, nil
}

func (*Account) Diff(ctx p.Context, id string, olds AccountState, news AccountArgs) (p.DiffResponse, error) {
	diff := map[string]p.PropertyDiff{}

	if olds.Name != news.Name {
		diff["name"] = p.PropertyDiff{Kind: p.Update}
	}

	return p.DiffResponse{
		DeleteBeforeReplace: true,
		HasChanges:          len(diff) > 0,
		DetailedDiff:        diff,
	}, nil
}

func (*Account) Update(ctx p.Context, id string, olds AccountState, news AccountArgs, preview bool) (AccountState, error) {
	var accountName string
	if !preview {
		if olds.Name != news.Name {
			accountName = news.Name
		} else {
			accountName = olds.Name
		}
	}
	return AccountState{
		Id:   olds.Id,
		Name: accountName,
	}, nil
}

type GetAccount struct{}

type GetAccountArgs struct {
	AccountName string `pulumi:"accountName"`
}

func (ga *GetAccount) Annotate(a infer.Annotator) {
	a.Describe(&ga, "GetOrganization gets the Organization information")
}

func (ga *GetAccountArgs) Annotate(a infer.Annotator) {
	a.Describe(&ga.AccountName, "Name of the Account")
}

func (GetAccount) Call(ctx p.Context, args GetAccountArgs) (AccountState, error) {
	config := infer.GetConfig[Config](ctx)
	return AccountState{
		Id:          uuid.New().String(),
		Name:        args.AccountName,
		Environment: config.Environment,
	}, nil
}
