package pkg

import (
	"context"

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

func (*Account) Create(ctx context.Context, req infer.CreateRequest[AccountArgs]) (infer.CreateResponse[AccountState], error) {
	if req.DryRun {
		return infer.CreateResponse[AccountState]{}, nil
	}
	config := infer.GetConfig[Config](ctx)
	return infer.CreateResponse[AccountState]{
		ID: req.Name,
		Output: AccountState{
			Id:          uuid.New().String(),
			Name:        req.Name,
			Environment: config.Environment,
		},
	}, nil
}

func (*Account) Delete(ctx context.Context, req infer.DeleteRequest[AccountState]) (infer.DeleteResponse, error) {
	return infer.DeleteResponse{}, nil
}

func (*Account) Read(ctx context.Context, req infer.ReadRequest[AccountArgs, AccountState]) (infer.ReadResponse[AccountArgs, AccountState], error) {
	return infer.ReadResponse[AccountArgs, AccountState]{
		ID:     req.Inputs.Name,
		Inputs: AccountArgs{Name: req.State.Name},
		State:  req.State,
	}, nil
}

func (*Account) Diff(ctx context.Context, req infer.DiffRequest[AccountArgs, AccountState]) (p.DiffResponse, error) {
	diff := map[string]p.PropertyDiff{}

	if req.State.Name != req.Inputs.Name {
		diff["name"] = p.PropertyDiff{Kind: p.Update}
	}

	return p.DiffResponse{
		DeleteBeforeReplace: true,
		HasChanges:          len(diff) > 0,
		DetailedDiff:        diff,
	}, nil
}

func (*Account) Update(ctx context.Context, req infer.UpdateRequest[AccountArgs, AccountState]) (infer.UpdateResponse[AccountState], error) {
	if req.DryRun {
		return infer.UpdateResponse[AccountState]{}, nil
	}
	name := req.State.Name
	if req.State.Name != req.Inputs.Name {
		name = req.Inputs.Name
	}
	return infer.UpdateResponse[AccountState]{
		Output: AccountState{
			Id:   req.State.Id,
			Name: name,
		},
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

func (*GetAccount) Invoke(ctx context.Context, req infer.FunctionRequest[GetAccountArgs]) (infer.FunctionResponse[AccountState], error) {
	config := infer.GetConfig[Config](ctx)
	return infer.FunctionResponse[AccountState]{
		Output: AccountState{
			Id:          uuid.New().String(),
			Name:        req.Input.AccountName,
			Environment: config.Environment,
		},
	}, nil
}
