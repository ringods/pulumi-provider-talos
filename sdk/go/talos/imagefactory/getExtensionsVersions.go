// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package imagefactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-talos/sdk/go/talos/internal"
)

// The image factory extensions versions data source provides a list of available extensions for a specific talos version from the image factory.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-talos/sdk/go/talos/imageFactory"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := imageFactory.GetExtensionsVersions(ctx, &imagefactory.GetExtensionsVersionsArgs{
//				TalosVersion: "v1.7.5",
//				Filters: imagefactory.GetExtensionsVersionsFilters{
//					Names: []string{
//						"amdgpu",
//						"tailscale",
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetExtensionsVersions(ctx *pulumi.Context, args *GetExtensionsVersionsArgs, opts ...pulumi.InvokeOption) (*GetExtensionsVersionsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetExtensionsVersionsResult
	err := ctx.Invoke("talos:imageFactory/getExtensionsVersions:getExtensionsVersions", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getExtensionsVersions.
type GetExtensionsVersionsArgs struct {
	// The filter to apply to the extensions list.
	Filters *GetExtensionsVersionsFilters `pulumi:"filters"`
	// The talos version to get extensions for.
	TalosVersion string `pulumi:"talosVersion"`
}

// A collection of values returned by getExtensionsVersions.
type GetExtensionsVersionsResult struct {
	// The list of available extensions for the specified talos version.
	ExtensionsInfos []GetExtensionsVersionsExtensionsInfo `pulumi:"extensionsInfos"`
	// The filter to apply to the extensions list.
	Filters *GetExtensionsVersionsFilters `pulumi:"filters"`
	// The ID of this resource.
	Id string `pulumi:"id"`
	// The talos version to get extensions for.
	TalosVersion string `pulumi:"talosVersion"`
}

func GetExtensionsVersionsOutput(ctx *pulumi.Context, args GetExtensionsVersionsOutputArgs, opts ...pulumi.InvokeOption) GetExtensionsVersionsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetExtensionsVersionsResultOutput, error) {
			args := v.(GetExtensionsVersionsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("talos:imageFactory/getExtensionsVersions:getExtensionsVersions", args, GetExtensionsVersionsResultOutput{}, options).(GetExtensionsVersionsResultOutput), nil
		}).(GetExtensionsVersionsResultOutput)
}

// A collection of arguments for invoking getExtensionsVersions.
type GetExtensionsVersionsOutputArgs struct {
	// The filter to apply to the extensions list.
	Filters GetExtensionsVersionsFiltersPtrInput `pulumi:"filters"`
	// The talos version to get extensions for.
	TalosVersion pulumi.StringInput `pulumi:"talosVersion"`
}

func (GetExtensionsVersionsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetExtensionsVersionsArgs)(nil)).Elem()
}

// A collection of values returned by getExtensionsVersions.
type GetExtensionsVersionsResultOutput struct{ *pulumi.OutputState }

func (GetExtensionsVersionsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetExtensionsVersionsResult)(nil)).Elem()
}

func (o GetExtensionsVersionsResultOutput) ToGetExtensionsVersionsResultOutput() GetExtensionsVersionsResultOutput {
	return o
}

func (o GetExtensionsVersionsResultOutput) ToGetExtensionsVersionsResultOutputWithContext(ctx context.Context) GetExtensionsVersionsResultOutput {
	return o
}

// The list of available extensions for the specified talos version.
func (o GetExtensionsVersionsResultOutput) ExtensionsInfos() GetExtensionsVersionsExtensionsInfoArrayOutput {
	return o.ApplyT(func(v GetExtensionsVersionsResult) []GetExtensionsVersionsExtensionsInfo { return v.ExtensionsInfos }).(GetExtensionsVersionsExtensionsInfoArrayOutput)
}

// The filter to apply to the extensions list.
func (o GetExtensionsVersionsResultOutput) Filters() GetExtensionsVersionsFiltersPtrOutput {
	return o.ApplyT(func(v GetExtensionsVersionsResult) *GetExtensionsVersionsFilters { return v.Filters }).(GetExtensionsVersionsFiltersPtrOutput)
}

// The ID of this resource.
func (o GetExtensionsVersionsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetExtensionsVersionsResult) string { return v.Id }).(pulumi.StringOutput)
}

// The talos version to get extensions for.
func (o GetExtensionsVersionsResultOutput) TalosVersion() pulumi.StringOutput {
	return o.ApplyT(func(v GetExtensionsVersionsResult) string { return v.TalosVersion }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetExtensionsVersionsResultOutput{})
}
