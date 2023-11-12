// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Talos.Machine.Outputs
{

    /// <summary>
    /// A complete Machine Secrets Certificates configuration
    /// </summary>
    [OutputType]
    public sealed class CertificatesResult
    {
        public readonly Outputs.CertificateResult Etcd;
        public readonly Outputs.CertificateResult K8s;
        public readonly Outputs.CertificateResult K8s_aggregator;
        public readonly Outputs.KeyResult K8s_serviceaccount;
        public readonly Outputs.CertificateResult Os;

        [OutputConstructor]
        private CertificatesResult(
            Outputs.CertificateResult etcd,

            Outputs.CertificateResult k8s,

            Outputs.CertificateResult k8s_aggregator,

            Outputs.KeyResult k8s_serviceaccount,

            Outputs.CertificateResult os)
        {
            Etcd = etcd;
            K8s = k8s;
            K8s_aggregator = k8s_aggregator;
            K8s_serviceaccount = k8s_serviceaccount;
            Os = os;
        }
    }
}
