// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package machine

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
	"github.com/pulumiverse/pulumi-talos/sdk/go/talos/internal"
)

// Generate machine secrets for Talos cluster.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-talos/sdk/go/talos/machine"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := machine.NewSecrets(ctx, "machineSecrets", nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// terraform machine secrets can be imported from an existing secrets file
//
// ```sh
//
//	$ pulumi import talos:machine/secrets:Secrets this <path-to-secrets.yaml>
//
// ```
type SecretsType struct {
	pulumi.CustomResourceState

	// The generated client configuration data
	ClientConfiguration SecretsClientConfigurationOutput `pulumi:"clientConfiguration"`
	// The secrets for the talos cluster
	MachineSecrets MachineSecretsOutput `pulumi:"machineSecrets"`
	// The version of talos features to use in generated machine configuration
	TalosVersion pulumi.StringOutput `pulumi:"talosVersion"`
}

// NewSecretsType registers a new resource with the given unique name, arguments, and options.
func NewSecretsType(ctx *pulumi.Context,
	name string, args *SecretsTypeArgs, opts ...pulumi.ResourceOption) (*SecretsType, error) {
	if args == nil {
		args = &SecretsTypeArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SecretsType
	err := ctx.RegisterResource("talos:machine/secrets:Secrets", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretsType gets an existing SecretsType resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretsType(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretsTypeState, opts ...pulumi.ResourceOption) (*SecretsType, error) {
	var resource SecretsType
	err := ctx.ReadResource("talos:machine/secrets:Secrets", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretsType resources.
type secretsTypeState struct {
	// The generated client configuration data
	ClientConfiguration *SecretsClientConfiguration `pulumi:"clientConfiguration"`
	// The secrets for the talos cluster
	MachineSecrets *MachineSecrets `pulumi:"machineSecrets"`
	// The version of talos features to use in generated machine configuration
	TalosVersion *string `pulumi:"talosVersion"`
}

type SecretsTypeState struct {
	// The generated client configuration data
	ClientConfiguration SecretsClientConfigurationPtrInput
	// The secrets for the talos cluster
	MachineSecrets MachineSecretsPtrInput
	// The version of talos features to use in generated machine configuration
	TalosVersion pulumi.StringPtrInput
}

func (SecretsTypeState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretsTypeState)(nil)).Elem()
}

type secretsTypeArgs struct {
	// The version of talos features to use in generated machine configuration
	TalosVersion *string `pulumi:"talosVersion"`
}

// The set of arguments for constructing a SecretsType resource.
type SecretsTypeArgs struct {
	// The version of talos features to use in generated machine configuration
	TalosVersion pulumi.StringPtrInput
}

func (SecretsTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretsTypeArgs)(nil)).Elem()
}

type SecretsTypeInput interface {
	pulumi.Input

	ToSecretsTypeOutput() SecretsTypeOutput
	ToSecretsTypeOutputWithContext(ctx context.Context) SecretsTypeOutput
}

func (*SecretsType) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretsType)(nil)).Elem()
}

func (i *SecretsType) ToSecretsTypeOutput() SecretsTypeOutput {
	return i.ToSecretsTypeOutputWithContext(context.Background())
}

func (i *SecretsType) ToSecretsTypeOutputWithContext(ctx context.Context) SecretsTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretsTypeOutput)
}

func (i *SecretsType) ToOutput(ctx context.Context) pulumix.Output[*SecretsType] {
	return pulumix.Output[*SecretsType]{
		OutputState: i.ToSecretsTypeOutputWithContext(ctx).OutputState,
	}
}

// SecretsTypeArrayInput is an input type that accepts SecretsTypeArray and SecretsTypeArrayOutput values.
// You can construct a concrete instance of `SecretsTypeArrayInput` via:
//
//	SecretsTypeArray{ SecretsTypeArgs{...} }
type SecretsTypeArrayInput interface {
	pulumi.Input

	ToSecretsTypeArrayOutput() SecretsTypeArrayOutput
	ToSecretsTypeArrayOutputWithContext(context.Context) SecretsTypeArrayOutput
}

type SecretsTypeArray []SecretsTypeInput

func (SecretsTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretsType)(nil)).Elem()
}

func (i SecretsTypeArray) ToSecretsTypeArrayOutput() SecretsTypeArrayOutput {
	return i.ToSecretsTypeArrayOutputWithContext(context.Background())
}

func (i SecretsTypeArray) ToSecretsTypeArrayOutputWithContext(ctx context.Context) SecretsTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretsTypeArrayOutput)
}

func (i SecretsTypeArray) ToOutput(ctx context.Context) pulumix.Output[[]*SecretsType] {
	return pulumix.Output[[]*SecretsType]{
		OutputState: i.ToSecretsTypeArrayOutputWithContext(ctx).OutputState,
	}
}

// SecretsTypeMapInput is an input type that accepts SecretsTypeMap and SecretsTypeMapOutput values.
// You can construct a concrete instance of `SecretsTypeMapInput` via:
//
//	SecretsTypeMap{ "key": SecretsTypeArgs{...} }
type SecretsTypeMapInput interface {
	pulumi.Input

	ToSecretsTypeMapOutput() SecretsTypeMapOutput
	ToSecretsTypeMapOutputWithContext(context.Context) SecretsTypeMapOutput
}

type SecretsTypeMap map[string]SecretsTypeInput

func (SecretsTypeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretsType)(nil)).Elem()
}

func (i SecretsTypeMap) ToSecretsTypeMapOutput() SecretsTypeMapOutput {
	return i.ToSecretsTypeMapOutputWithContext(context.Background())
}

func (i SecretsTypeMap) ToSecretsTypeMapOutputWithContext(ctx context.Context) SecretsTypeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretsTypeMapOutput)
}

func (i SecretsTypeMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*SecretsType] {
	return pulumix.Output[map[string]*SecretsType]{
		OutputState: i.ToSecretsTypeMapOutputWithContext(ctx).OutputState,
	}
}

type SecretsTypeOutput struct{ *pulumi.OutputState }

func (SecretsTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretsType)(nil)).Elem()
}

func (o SecretsTypeOutput) ToSecretsTypeOutput() SecretsTypeOutput {
	return o
}

func (o SecretsTypeOutput) ToSecretsTypeOutputWithContext(ctx context.Context) SecretsTypeOutput {
	return o
}

func (o SecretsTypeOutput) ToOutput(ctx context.Context) pulumix.Output[*SecretsType] {
	return pulumix.Output[*SecretsType]{
		OutputState: o.OutputState,
	}
}

// The generated client configuration data
func (o SecretsTypeOutput) ClientConfiguration() SecretsClientConfigurationOutput {
	return o.ApplyT(func(v *SecretsType) SecretsClientConfigurationOutput { return v.ClientConfiguration }).(SecretsClientConfigurationOutput)
}

// The secrets for the talos cluster
func (o SecretsTypeOutput) MachineSecrets() MachineSecretsOutput {
	return o.ApplyT(func(v *SecretsType) MachineSecretsOutput { return v.MachineSecrets }).(MachineSecretsOutput)
}

// The version of talos features to use in generated machine configuration
func (o SecretsTypeOutput) TalosVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretsType) pulumi.StringOutput { return v.TalosVersion }).(pulumi.StringOutput)
}

type SecretsTypeArrayOutput struct{ *pulumi.OutputState }

func (SecretsTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretsType)(nil)).Elem()
}

func (o SecretsTypeArrayOutput) ToSecretsTypeArrayOutput() SecretsTypeArrayOutput {
	return o
}

func (o SecretsTypeArrayOutput) ToSecretsTypeArrayOutputWithContext(ctx context.Context) SecretsTypeArrayOutput {
	return o
}

func (o SecretsTypeArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*SecretsType] {
	return pulumix.Output[[]*SecretsType]{
		OutputState: o.OutputState,
	}
}

func (o SecretsTypeArrayOutput) Index(i pulumi.IntInput) SecretsTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SecretsType {
		return vs[0].([]*SecretsType)[vs[1].(int)]
	}).(SecretsTypeOutput)
}

type SecretsTypeMapOutput struct{ *pulumi.OutputState }

func (SecretsTypeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretsType)(nil)).Elem()
}

func (o SecretsTypeMapOutput) ToSecretsTypeMapOutput() SecretsTypeMapOutput {
	return o
}

func (o SecretsTypeMapOutput) ToSecretsTypeMapOutputWithContext(ctx context.Context) SecretsTypeMapOutput {
	return o
}

func (o SecretsTypeMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*SecretsType] {
	return pulumix.Output[map[string]*SecretsType]{
		OutputState: o.OutputState,
	}
}

func (o SecretsTypeMapOutput) MapIndex(k pulumi.StringInput) SecretsTypeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SecretsType {
		return vs[0].(map[string]*SecretsType)[vs[1].(string)]
	}).(SecretsTypeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecretsTypeInput)(nil)).Elem(), &SecretsType{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretsTypeArrayInput)(nil)).Elem(), SecretsTypeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretsTypeMapInput)(nil)).Elem(), SecretsTypeMap{})
	pulumi.RegisterOutputType(SecretsTypeOutput{})
	pulumi.RegisterOutputType(SecretsTypeArrayOutput{})
	pulumi.RegisterOutputType(SecretsTypeMapOutput{})
}
