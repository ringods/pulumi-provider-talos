package native

import (
	"context"

	gop "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/pkg/v3/resource/provider"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	rpc "github.com/pulumi/pulumi/sdk/v3/proto/go"
)

type Provider struct {
	prov gop.Provider
}

func NewProvider() *Provider {
	inst := Provider{}
	inst.prov = infer.Provider(infer.Options{
		Functions: []infer.InferredFunction{},
		ModuleMap: map[tokens.ModuleName]tokens.ModuleName{
			"native": "index",
		},
	})
	return &inst
}

func (p *Provider) GetInstance(ctx context.Context, name, version string, host *provider.HostClient) (rpc.ResourceProviderServer, error) {
	return gop.RawServer(name, version, p.prov)
}

func (p *Provider) GetSpec(ctx context.Context, name, version string) (schema.PackageSpec, error) {
	return gop.GetSchema(context.Background(), name, version, p.prov)
}
