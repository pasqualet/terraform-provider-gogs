---
layout: "gogs"
page_title: "Gogs: gogs_repository"
sidebar_current: "docs-gogs-resource-repository-x"
description: |-
  Creates and manages repositories.
---

# gogs\_repository

This resource allows you to create and manage repositories.


## Example Usage

```hcl
resource "gogs_user" "sample_user" {
    username = "sample"
    fullname = "Sample User"
    login_name = "sample"
    email = "sample@sample.com"
    password = "password"
    active = true
}

resource "gogs_repository" "sample" {
    username = "${gogs_user.sample_user.username}"
    name = "sample-repository"
}

```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the repository.

* `description` - (Required) A short description of the repository.

* `private` - (Optional) Either true to create a private repository, or false to create a public one. Boolean value. (Default: `false`).

* `auto_init` - (Optional) Pass true to create an initial commit with README, .gitignore and LICENSE. Boolean value. (Default: `false`).

* `gitignores` - (Optional) Desired language .gitignore templates to apply. Use the name of the templates. For example, "Go" or "Go,SublimeText".

* `license` - (Optional) Desired LICENSE template to apply. Use the name of the template. For example, "Apache v2 License" or "MIT License".

* `readme` - (Optional) Desired README template to apply. Use the name of the template. (Default: `Default`).

## Attributes Reference

The following additional attributes are exported:

* `id` - Integer that uniquely identifies the repository.
