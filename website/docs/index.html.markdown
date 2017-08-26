---
layout: "gogs"
page_title: "Provider: Gogs"
sidebar_current: "docs-gogs-index"
description: |-
  The Gogs provider is used to interact with Gogs instance.
---

# Gogs Provider

The Gogs provider is used to interact with Gogs instance.

It needs to be configured with the proper credentials before it can be used.

Use the navigation to the left to read about the available resources.

## Example Usage

```hcl
# Configure the Gogs Provider
provider "gogs" {
    url = "${var.gogs_url}"
    token = "${var.gogs_token}"
}

# Add user
resource "gogs_user" "sample_user" {
    username = "sample"
    fullname = "Sample User"
    login_name = "sample"
    email = "sample@sample.com"
    password = "password"
    active = true
}
```

## Argument Reference

The following arguments are supported in the `provider` block:

* `token` - (Optional) This is the Gogs personal access token. It must be provided, but
  it can also be sourced from the `GOGS_TOKEN` environment variable.

* `url` - (Optional) This is the target Gogs instance.
  It is optional to provide this value and it can also be sourced from the `GOGS_BASE_URL` environment variable.
