package example

import (
	"context"

	"github.com/getopends/libprovider"
)

var _ libprovider.Provider = &ExampleProvider{}

const (
	envHost = "HOST"
	envPort = "PORT"
)

func NewExampleProvider(context.Context) libprovider.Provider {
	return &ExampleProvider{}
}

type ExampleProvider struct {
	opts             *libprovider.Options
	apiHost, apiPort string
}

func (p *ExampleProvider) Info() libprovider.Info {
	return libprovider.Info{
		Name:    "Example",
		Slug:    "example-moz",
		Author:  "Edson Michaque",
		Version: "0.1.0",
		Secrets: []string{
			"USERNAME",
			"PASSWORD",
		},
		Env: []string{
			"HOST",
			"PORT",
		},
	}
}

func (p *ExampleProvider) Configure(ctx context.Context, opts *libprovider.Options) error {
	apiHost, err := p.opts.Env.Get(envHost)
	if err != nil {
		return err
	}

	apiPort, err := p.opts.Env.Get(envPort)
	if err != nil {
		return err
	}

	p.apiHost = apiHost
	p.apiPort = apiPort

	return nil
}

func (p *ExampleProvider) CreateTransaction(ctx context.Context, t *libprovider.Transaction) (*libprovider.TransactionResult, error) {
	return nil, libprovider.ErrNotSupported
}
