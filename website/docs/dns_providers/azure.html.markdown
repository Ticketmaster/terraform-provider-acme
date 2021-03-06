---
layout: "acme"
page_title: "ACME: Azure DNS Challenge Provider"
sidebar_current: "docs-acme-dns-providers-azure"
description: |-
  Provides a resource to manage certificates on an ACME CA.
---

# Azure DNS Challenge Provider

The `azure` DNS challenge provider can be used to perform DNS challenges for
the [`acme_certificate`][resource-acme-certificate] resource with [Microsoft
Azure][provider-service-page].

[resource-acme-certificate]: /docs/providers/acme/r/certificate.html
[provider-service-page]: https://azure.microsoft.com/en-ca/

For complete information on how to use this provider with the `acme_certifiate`
resource, see [here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenges]: /docs/providers/acme/r/certificate.html#using-dns-challenges

## Example

```hcl
resource "acme_certificate" "certificate" {
  ...

  dns_challenge {
    provider = "azure"
  }
}
```

## Argument Reference

The following arguments can be either passed as environment variables, or
directly through the `config` block in the
[`dns_challenge`][resource-acme-certificate-dns-challenge-arg] argument in the
[`acme_certificate`][resource-acme-certificate] resource. For more details, see
[here][resource-acme-certificate-dns-challenges].

[resource-acme-certificate-dns-challenge-arg]: /docs/providers/acme/r/certificate.html#dns_challenge

In addition, arguments can also be stored in a local file, with the path
supplied by supplying the argument with the `_FILE` suffix. See
[here][acme-certificate-file-arg-example] for more information.

[acme-certificate-file-arg-example]: /docs/providers/acme/r/certificate.html#using-variable-files-for-provider-arguments

* `AZURE_CLIENT_ID` - The Client ID of the Service Principal. Can also be
  supplied with `ARM_CLIENT_ID`.
* `AZURE_CLIENT_SECRET` - The Client Secret associated with the Service
  Principal. Can also be supplied with `ARM_CLIENT_SECRET`.
* `AZURE_SUBSCRIPTION_ID` - The ID of the Azure Subscription. Can also be
  supplied with `ARM_SUBSCRIPTION_ID`.
* `AZURE_TENANT_ID` - The Tenant ID to use. Can also be
  supplied with `ARM_TENANT_ID`.
* `AZURE_RESOURCE_GROUP` - The resource group to use to place the DNS records
  in. Can also be supplied with `ARM_RESOURCE_GROUP`.

The following additional optional variables are available:

* `AZURE_POLLING_INTERVAL` - The amount of time, in seconds, to wait between
  DNS propagation checks (default: `2`).
* `AZURE_PROPAGATION_TIMEOUT` - The amount of time, in seconds, to wait for DNS
  propagation (default: `120`).
* `AZURE_TTL` - The TTL to set on DNS challenge records, in seconds (default:
  `60`).
* `AZURE_METADATA_ENDPOINT` - The metadata endpoint to use.
