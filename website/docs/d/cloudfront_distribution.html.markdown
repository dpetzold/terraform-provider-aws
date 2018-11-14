---
layout: "aws"
page_title: "AWS: cloudfront_distribution"
sidebar_current: "docs-aws-data-cloudfront-distribution"
description: |-
  Provides information about a CloudFront web distribution.
---

## Example Usage

The following example below creates a CloudFront distribution with an S3 origin.

```hcl
data "aws_cloudfront_distribution" "dist" {
    id = "E74FTE3EXAMPLE"
}
```

## Argument Reference

The following arguments are supported:

  * `id` (Required) - The distribution id.

## Attribute Reference

The following attributes are supported:

  * `id` - The identifier for the distribution. For example: `EDFDVBD632BHDS5`.

  * `arn` - The ARN (Amazon Resource Name) for the distribution. For example: arn:aws:cloudfront::123456789012:distribution/EDFDVBD632BHDS5, where 123456789012 is your AWS account ID.

  * `status` - The current status of the distribution. `Deployed` if the
    distribution's information is fully propagated throughout the Amazon
    CloudFront system.

  * `active_trusted_signers` - The key pair IDs that CloudFront is aware of for
    each trusted signer, if the distribution is set up to serve private content
    with signed URLs.

  * `domain_name` - The domain name corresponding to the distribution. For
    example: `d604721fxaaqy9.cloudfront.net`.

  * `last_modified_time` - The date and time the distribution was last modified.

  * `in_progress_validation_batches` - The number of invalidation batches
    currently in progress.

  * `etag` - The current version of the distribution's information. For example:
    `E2QWRUHAPOMQZL`.

  * `hosted_zone_id` - The CloudFront Route 53 zone ID that can be used to
     route an [Alias Resource Record Set][7] to. This attribute is simply an
     alias for the zone ID `Z2FDTNDATAQYW2`.
