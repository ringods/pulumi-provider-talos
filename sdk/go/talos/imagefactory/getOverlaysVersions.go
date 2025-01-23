// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package imagefactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-talos/sdk/go/talos/internal"
)

// The image factory overlays versions data source provides a list of available overlays for a specific talos version from the image factory.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-talos/sdk/go/talos/imagefactory"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := imagefactory.GetOverlaysVersions(ctx, &imagefactory.GetOverlaysVersionsArgs{
//				TalosVersion: "v1.7.5",
//				Filters: imagefactory.GetOverlaysVersionsFilters{
//					Name: pulumi.StringRef("rock4cplus"),
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
func GetOverlaysVersions(ctx *pulumi.Context, args *GetOverlaysVersionsArgs, opts ...pulumi.InvokeOption) (*GetOverlaysVersionsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetOverlaysVersionsResult
	err := ctx.Invoke("talos:imageFactory/getOverlaysVersions:getOverlaysVersions", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOverlaysVersions.
type GetOverlaysVersionsArgs struct {
	// The filter to apply to the overlays list.
	Filters *GetOverlaysVersionsFilters `pulumi:"filters"`
	// The talos version to get overlays for.
	TalosVersion string `pulumi:"talosVersion"`
}

// A collection of values returned by getOverlaysVersions.
type GetOverlaysVersionsResult struct {
	// The filter to apply to the overlays list.
	Filters *GetOverlaysVersionsFilters `pulumi:"filters"`
	// The ID of this resource.
	Id string `pulumi:"id"`
	// The list of available extensions for the specified talos version.
	OverlaysInfos []GetOverlaysVersionsOverlaysInfo `pulumi:"overlaysInfos"`
	// The talos version to get overlays for.
	TalosVersion string `pulumi:"talosVersion"`
}

func GetOverlaysVersionsOutput(ctx *pulumi.Context, args GetOverlaysVersionsOutputArgs, opts ...pulumi.InvokeOption) GetOverlaysVersionsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetOverlaysVersionsResultOutput, error) {
			args := v.(GetOverlaysVersionsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("talos:imageFactory/getOverlaysVersions:getOverlaysVersions", args, GetOverlaysVersionsResultOutput{}, options).(GetOverlaysVersionsResultOutput), nil
		}).(GetOverlaysVersionsResultOutput)
}

// A collection of arguments for invoking getOverlaysVersions.
type GetOverlaysVersionsOutputArgs struct {
	// The filter to apply to the overlays list.
	Filters GetOverlaysVersionsFiltersPtrInput `pulumi:"filters"`
	// The talos version to get overlays for.
	TalosVersion pulumi.StringInput `pulumi:"talosVersion"`
}

func (GetOverlaysVersionsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOverlaysVersionsArgs)(nil)).Elem()
}

// A collection of values returned by getOverlaysVersions.
type GetOverlaysVersionsResultOutput struct{ *pulumi.OutputState }

func (GetOverlaysVersionsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOverlaysVersionsResult)(nil)).Elem()
}

func (o GetOverlaysVersionsResultOutput) ToGetOverlaysVersionsResultOutput() GetOverlaysVersionsResultOutput {
	return o
}

func (o GetOverlaysVersionsResultOutput) ToGetOverlaysVersionsResultOutputWithContext(ctx context.Context) GetOverlaysVersionsResultOutput {
	return o
}

// The filter to apply to the overlays list.
func (o GetOverlaysVersionsResultOutput) Filters() GetOverlaysVersionsFiltersPtrOutput {
	return o.ApplyT(func(v GetOverlaysVersionsResult) *GetOverlaysVersionsFilters { return v.Filters }).(GetOverlaysVersionsFiltersPtrOutput)
}

// The ID of this resource.
func (o GetOverlaysVersionsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetOverlaysVersionsResult) string { return v.Id }).(pulumi.StringOutput)
}

// The list of available extensions for the specified talos version.
func (o GetOverlaysVersionsResultOutput) OverlaysInfos() GetOverlaysVersionsOverlaysInfoArrayOutput {
	return o.ApplyT(func(v GetOverlaysVersionsResult) []GetOverlaysVersionsOverlaysInfo { return v.OverlaysInfos }).(GetOverlaysVersionsOverlaysInfoArrayOutput)
}

// The talos version to get overlays for.
func (o GetOverlaysVersionsResultOutput) TalosVersion() pulumi.StringOutput {
	return o.ApplyT(func(v GetOverlaysVersionsResult) string { return v.TalosVersion }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetOverlaysVersionsResultOutput{})
}
