// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export namespace client {
    export interface GetConfigurationClientConfiguration {
        /**
         * The client CA certificate
         */
        caCertificate: string;
        /**
         * The client certificate
         */
        clientCertificate: string;
        /**
         * The client key
         */
        clientKey: string;
    }

}

export namespace cluster {
    export interface GetHealthClientConfiguration {
        /**
         * The client CA certificate
         */
        caCertificate: string;
        /**
         * The client certificate
         */
        clientCertificate: string;
        /**
         * The client key
         */
        clientKey: string;
    }

    export interface GetHealthTimeouts {
        /**
         * A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh is enabled.
         */
        read?: string;
    }

    export interface GetKubeconfigClientConfiguration {
        /**
         * The client CA certificate
         */
        caCertificate: string;
        /**
         * The client certificate
         */
        clientCertificate: string;
        /**
         * The client key
         */
        clientKey: string;
    }

    export interface GetKubeconfigKubernetesClientConfiguration {
        /**
         * The kubernetes CA certificate
         */
        caCertificate: string;
        /**
         * The kubernetes client certificate
         */
        clientCertificate: string;
        /**
         * The kubernetes client key
         */
        clientKey: string;
        /**
         * The kubernetes host
         */
        host: string;
    }

    export interface GetKubeconfigTimeouts {
        /**
         * A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh is enabled.
         */
        read?: string;
    }

    export interface KubeconfigClientConfiguration {
        /**
         * The client CA certificate
         */
        caCertificate: string;
        /**
         * The client certificate
         */
        clientCertificate: string;
        /**
         * The client key
         */
        clientKey: string;
    }

    export interface KubeconfigKubernetesClientConfiguration {
        /**
         * The kubernetes CA certificate
         */
        caCertificate: string;
        /**
         * The kubernetes client certificate
         */
        clientCertificate: string;
        /**
         * The kubernetes client key
         */
        clientKey: string;
        /**
         * The kubernetes host
         */
        host: string;
    }

    export interface KubeconfigTimeouts {
        /**
         * A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours).
         */
        create?: string;
        /**
         * A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours).
         */
        update?: string;
    }

}

export namespace imageFactory {
    export interface GetExtensionsVersionsExtensionsInfo {
        author: string;
        description: string;
        digest: string;
        name: string;
        ref: string;
    }

    export interface GetExtensionsVersionsFilters {
        /**
         * The name of the extension to filter by.
         */
        names?: string[];
    }

    export interface GetOverlaysVersionsFilters {
        /**
         * The name of the overlay to filter by.
         */
        name?: string;
    }

    export interface GetOverlaysVersionsOverlaysInfo {
        digest: string;
        image: string;
        name: string;
        ref: string;
    }

    export interface GetUrlsUrls {
        /**
         * The URL for the disk image.
         */
        diskImage: string;
        /**
         * The URL for the disk image with secure boot.
         */
        diskImageSecureboot: string;
        /**
         * The URL for the initramfs image.
         */
        initramfs: string;
        /**
         * The URL for the installer image.
         */
        installer: string;
        /**
         * The URL for the installer image with secure boot.
         */
        installerSecureboot: string;
        /**
         * The URL for the ISO image.
         */
        iso: string;
        /**
         * The URL for the ISO image with secure boot.
         */
        isoSecureboot: string;
        /**
         * The URL for the kernel image.
         */
        kernel: string;
        /**
         * The URL for the kernel command line.
         */
        kernelCommandLine: string;
        /**
         * The URL for the PXE image.
         */
        pxe: string;
        /**
         * The URL for the UKI image.
         */
        uki: string;
    }

    export interface GetVersionsFilters {
        /**
         * If set to true, only stable versions will be returned. If set to false, all versions will be returned.
         */
        stableVersionsOnly?: boolean;
    }

}

export namespace machine {
    export interface BootstrapTimeouts {
        /**
         * A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours).
         */
        create?: string;
    }

    /**
     * A Machine Secrets Certificate
     */
    export interface Certificate {
        /**
         * Certificate
         */
        cert: string;
        /**
         * Private Key
         */
        key: string;
    }

    /**
     * A complete Machine Secrets Certificates configuration
     */
    export interface Certificates {
        etcd: outputs.machine.Certificate;
        k8s: outputs.machine.Certificate;
        k8sAggregator: outputs.machine.Certificate;
        k8sServiceaccount: outputs.machine.Key;
        os: outputs.machine.Certificate;
    }

    /**
     * A Client Configuration
     */
    export interface ClientConfiguration {
        /**
         * The client CA certificate
         */
        caCertificate: string;
        /**
         * The client certificate
         */
        clientCertificate: string;
        /**
         * The client private key
         */
        clientKey: string;
    }

    /**
     * A Machine Secrets Cluster Info
     */
    export interface Cluster {
        /**
         * Certificate
         */
        id: string;
        /**
         * Private Key
         */
        secret: string;
    }

    export interface ConfigurationApplyOnDestroy {
        /**
         * Graceful indicates whether node should leave etcd before the upgrade, it also enforces etcd checks before leaving. Default true
         */
        graceful: boolean;
        /**
         * Reboot indicates whether node should reboot or halt after resetting. Default false
         */
        reboot: boolean;
        /**
         * Reset the machine to the initial state (STATE and EPHEMERAL will be wiped). Default false
         */
        reset: boolean;
    }

    export interface GetDisksClientConfiguration {
        /**
         * The client CA certificate
         */
        caCertificate: string;
        /**
         * The client certificate
         */
        clientCertificate: string;
        /**
         * The client key
         */
        clientKey: string;
    }

    export interface GetDisksDisk {
        /**
         * The bus path of the disk
         */
        busPath: string;
        /**
         * The modalias of the disk
         */
        modalias: string;
        /**
         * The model of the disk
         */
        model: string;
        /**
         * The name of the disk
         */
        name: string;
        /**
         * The serial number of the disk
         */
        serial: string;
        /**
         * The size of the disk
         */
        size: string;
        /**
         * The type of the disk
         */
        type: string;
        /**
         * The uuid of the disk
         */
        uuid: string;
        /**
         * The wwid of the disk
         */
        wwid: string;
    }

    export interface GetDisksFilters {
        /**
         * Filter disks by bus path
         */
        busPath?: string;
        /**
         * Filter disks by modalias
         */
        modalias?: string;
        /**
         * Filter disks by model
         */
        model?: string;
        /**
         * Filter disks by name
         */
        name?: string;
        /**
         * Filter disks by serial number
         */
        serial?: string;
        /**
         * Filter disks by size
         */
        size?: string;
        /**
         * Filter disks by type
         */
        type?: string;
        /**
         * Filter disks by uuid
         */
        uuid?: string;
        /**
         * Filter disks by wwid
         */
        wwid?: string;
    }

    export interface GetDisksTimeouts {
        /**
         * A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh is enabled.
         */
        read?: string;
    }

    /**
     * A Machine Secrets Private Key
     */
    export interface Key {
        /**
         * Private Key
         */
        key: string;
    }

    /**
     * A Machine Secrets Bootstrap data
     */
    export interface KubernetesSecrets {
        /**
         * The aescbc encryption secret for the talos kubernetes cluster
         */
        aescbcEncryptionSecret?: string;
        /**
         * The bootstrap token for the talos kubernetes cluster
         */
        bootstrapToken: string;
        /**
         * The secretbox encryption secret for the talos kubernetes cluster
         */
        secretboxEncryptionSecret: string;
    }

    /**
     * A complete Machine Secrets configuration
     */
    export interface MachineSecrets {
        certs: outputs.machine.Certificates;
        cluster: outputs.machine.Cluster;
        secrets: outputs.machine.KubernetesSecrets;
        trustdinfo: outputs.machine.TrustdInfo;
    }

    export interface Timeout {
        /**
         * A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours).
         */
        create?: string;
        /**
         * A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are saved into state before the destroy operation occurs.
         */
        delete?: string;
        /**
         * A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours).
         */
        update?: string;
    }

    /**
     * A Machine Secrets Trust daemon info
     */
    export interface TrustdInfo {
        /**
         * The trustd token for the talos kubernetes cluster
         */
        token: string;
    }

}
