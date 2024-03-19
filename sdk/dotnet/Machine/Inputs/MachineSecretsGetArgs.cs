// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Talos.Machine.Inputs
{

    /// <summary>
    /// A complete Machine Secrets configuration
    /// </summary>
    public sealed class MachineSecretsGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("certs", required: true)]
        public Input<Inputs.CertificatesGetArgs> Certs { get; set; } = null!;

        [Input("cluster", required: true)]
        public Input<Inputs.ClusterGetArgs> Cluster { get; set; } = null!;

        [Input("secrets", required: true)]
        public Input<Inputs.KubernetesSecretsGetArgs> Secrets { get; set; } = null!;

        [Input("trustdinfo", required: true)]
        public Input<Inputs.TrustdInfoGetArgs> Trustdinfo { get; set; } = null!;

        public MachineSecretsGetArgs()
        {
        }
        public static new MachineSecretsGetArgs Empty => new MachineSecretsGetArgs();
    }
}
