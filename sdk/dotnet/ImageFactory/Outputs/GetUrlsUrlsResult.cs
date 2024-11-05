// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Talos.ImageFactory.Outputs
{

    [OutputType]
    public sealed class GetUrlsUrlsResult
    {
        /// <summary>
        /// The URL for the disk image.
        /// </summary>
        public readonly string DiskImage;
        /// <summary>
        /// The URL for the disk image with secure boot.
        /// </summary>
        public readonly string DiskImageSecureboot;
        /// <summary>
        /// The URL for the initramfs image.
        /// </summary>
        public readonly string Initramfs;
        /// <summary>
        /// The URL for the installer image.
        /// </summary>
        public readonly string Installer;
        /// <summary>
        /// The URL for the installer image with secure boot.
        /// </summary>
        public readonly string InstallerSecureboot;
        /// <summary>
        /// The URL for the ISO image.
        /// </summary>
        public readonly string Iso;
        /// <summary>
        /// The URL for the ISO image with secure boot.
        /// </summary>
        public readonly string IsoSecureboot;
        /// <summary>
        /// The URL for the kernel image.
        /// </summary>
        public readonly string Kernel;
        /// <summary>
        /// The URL for the kernel command line.
        /// </summary>
        public readonly string KernelCommandLine;
        /// <summary>
        /// The URL for the PXE image.
        /// </summary>
        public readonly string Pxe;
        /// <summary>
        /// The URL for the UKI image.
        /// </summary>
        public readonly string Uki;

        [OutputConstructor]
        private GetUrlsUrlsResult(
            string diskImage,

            string diskImageSecureboot,

            string initramfs,

            string installer,

            string installerSecureboot,

            string iso,

            string isoSecureboot,

            string kernel,

            string kernelCommandLine,

            string pxe,

            string uki)
        {
            DiskImage = diskImage;
            DiskImageSecureboot = diskImageSecureboot;
            Initramfs = initramfs;
            Installer = installer;
            InstallerSecureboot = installerSecureboot;
            Iso = iso;
            IsoSecureboot = isoSecureboot;
            Kernel = kernel;
            KernelCommandLine = kernelCommandLine;
            Pxe = pxe;
            Uki = uki;
        }
    }
}
