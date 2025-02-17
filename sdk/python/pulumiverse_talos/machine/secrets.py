# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['SecretsArgs', 'Secrets']

@pulumi.input_type
class SecretsArgs:
    def __init__(__self__, *,
                 talos_version: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Secrets resource.
        :param pulumi.Input[str] talos_version: The version of talos features to use in generated machine configuration
        """
        if talos_version is not None:
            pulumi.set(__self__, "talos_version", talos_version)

    @property
    @pulumi.getter(name="talosVersion")
    def talos_version(self) -> Optional[pulumi.Input[str]]:
        """
        The version of talos features to use in generated machine configuration
        """
        return pulumi.get(self, "talos_version")

    @talos_version.setter
    def talos_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "talos_version", value)


@pulumi.input_type
class _SecretsState:
    def __init__(__self__, *,
                 client_configuration: Optional[pulumi.Input['ClientConfigurationArgs']] = None,
                 machine_secrets: Optional[pulumi.Input['MachineSecretsArgs']] = None,
                 talos_version: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Secrets resources.
        :param pulumi.Input['ClientConfigurationArgs'] client_configuration: The generated client configuration data
        :param pulumi.Input['MachineSecretsArgs'] machine_secrets: The secrets for the talos cluster
        :param pulumi.Input[str] talos_version: The version of talos features to use in generated machine configuration
        """
        if client_configuration is not None:
            pulumi.set(__self__, "client_configuration", client_configuration)
        if machine_secrets is not None:
            pulumi.set(__self__, "machine_secrets", machine_secrets)
        if talos_version is not None:
            pulumi.set(__self__, "talos_version", talos_version)

    @property
    @pulumi.getter(name="clientConfiguration")
    def client_configuration(self) -> Optional[pulumi.Input['ClientConfigurationArgs']]:
        """
        The generated client configuration data
        """
        return pulumi.get(self, "client_configuration")

    @client_configuration.setter
    def client_configuration(self, value: Optional[pulumi.Input['ClientConfigurationArgs']]):
        pulumi.set(self, "client_configuration", value)

    @property
    @pulumi.getter(name="machineSecrets")
    def machine_secrets(self) -> Optional[pulumi.Input['MachineSecretsArgs']]:
        """
        The secrets for the talos cluster
        """
        return pulumi.get(self, "machine_secrets")

    @machine_secrets.setter
    def machine_secrets(self, value: Optional[pulumi.Input['MachineSecretsArgs']]):
        pulumi.set(self, "machine_secrets", value)

    @property
    @pulumi.getter(name="talosVersion")
    def talos_version(self) -> Optional[pulumi.Input[str]]:
        """
        The version of talos features to use in generated machine configuration
        """
        return pulumi.get(self, "talos_version")

    @talos_version.setter
    def talos_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "talos_version", value)


class Secrets(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 talos_version: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Generate machine secrets for Talos cluster.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_talos as talos

        machine_secrets = talos.machine.Secrets("machine_secrets")
        ```

        ## Import

        terraform

        machine secrets can be imported from an existing secrets file

        ```sh
        $ pulumi import talos:machine/secrets:Secrets this <path-to-secrets.yaml>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] talos_version: The version of talos features to use in generated machine configuration
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[SecretsArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Generate machine secrets for Talos cluster.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_talos as talos

        machine_secrets = talos.machine.Secrets("machine_secrets")
        ```

        ## Import

        terraform

        machine secrets can be imported from an existing secrets file

        ```sh
        $ pulumi import talos:machine/secrets:Secrets this <path-to-secrets.yaml>
        ```

        :param str resource_name: The name of the resource.
        :param SecretsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SecretsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 talos_version: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SecretsArgs.__new__(SecretsArgs)

            __props__.__dict__["talos_version"] = talos_version
            __props__.__dict__["client_configuration"] = None
            __props__.__dict__["machine_secrets"] = None
        super(Secrets, __self__).__init__(
            'talos:machine/secrets:Secrets',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            client_configuration: Optional[pulumi.Input[Union['ClientConfigurationArgs', 'ClientConfigurationArgsDict']]] = None,
            machine_secrets: Optional[pulumi.Input[Union['MachineSecretsArgs', 'MachineSecretsArgsDict']]] = None,
            talos_version: Optional[pulumi.Input[str]] = None) -> 'Secrets':
        """
        Get an existing Secrets resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union['ClientConfigurationArgs', 'ClientConfigurationArgsDict']] client_configuration: The generated client configuration data
        :param pulumi.Input[Union['MachineSecretsArgs', 'MachineSecretsArgsDict']] machine_secrets: The secrets for the talos cluster
        :param pulumi.Input[str] talos_version: The version of talos features to use in generated machine configuration
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SecretsState.__new__(_SecretsState)

        __props__.__dict__["client_configuration"] = client_configuration
        __props__.__dict__["machine_secrets"] = machine_secrets
        __props__.__dict__["talos_version"] = talos_version
        return Secrets(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clientConfiguration")
    def client_configuration(self) -> pulumi.Output['outputs.ClientConfiguration']:
        """
        The generated client configuration data
        """
        return pulumi.get(self, "client_configuration")

    @property
    @pulumi.getter(name="machineSecrets")
    def machine_secrets(self) -> pulumi.Output['outputs.MachineSecretsResult']:
        """
        The secrets for the talos cluster
        """
        return pulumi.get(self, "machine_secrets")

    @property
    @pulumi.getter(name="talosVersion")
    def talos_version(self) -> pulumi.Output[str]:
        """
        The version of talos features to use in generated machine configuration
        """
        return pulumi.get(self, "talos_version")

