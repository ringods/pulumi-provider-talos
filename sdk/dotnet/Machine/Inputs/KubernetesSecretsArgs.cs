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
    /// A Machine Secrets Bootstrap data
    /// </summary>
    public sealed class KubernetesSecretsInputArgs : global::Pulumi.ResourceArgs
    {
        [Input("aescbcEncryptionSecret")]
        private Input<string>? _aescbcEncryptionSecret;

        /// <summary>
        /// The aescbc encryption secret for the talos kubernetes cluster
        /// </summary>
        public Input<string>? AescbcEncryptionSecret
        {
            get => _aescbcEncryptionSecret;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _aescbcEncryptionSecret = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("bootstrapToken", required: true)]
        private Input<string>? _bootstrapToken;

        /// <summary>
        /// The bootstrap token for the talos kubernetes cluster
        /// </summary>
        public Input<string>? BootstrapToken
        {
            get => _bootstrapToken;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _bootstrapToken = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("secretboxEncryptionSecret", required: true)]
        private Input<string>? _secretboxEncryptionSecret;

        /// <summary>
        /// The secretbox encryption secret for the talos kubernetes cluster
        /// </summary>
        public Input<string>? SecretboxEncryptionSecret
        {
            get => _secretboxEncryptionSecret;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _secretboxEncryptionSecret = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public KubernetesSecretsInputArgs()
        {
        }
        public static new KubernetesSecretsInputArgs Empty => new KubernetesSecretsInputArgs();
    }
}
