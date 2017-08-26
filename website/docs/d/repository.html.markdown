---
layout: "gogs"
page_title: "Gogs: gogs_repository"
sidebar_current: "docs-gogs-datasource-repository-x"
description: |-
    Provides details about a specific repository
---

# gogs\_repository

`gogs_repository` provides details about a specific repository.

## Example Usage


```hcl
data "gogs_repository" "sample" {
    username = "sample-user"
    name = "sample-repository"
}
```

## Argument Reference

The given filters must match exactly the repository whose data will
be exported as attributes.

* `username` - (Required) The owner of the repository.

* `name` - (Required) The repository name.

## Attributes Reference

The following attribute are additionally exported:

* `id` - Integer that uniquely identifies the repository.

* `fullname` - Full name of repository.

* `email` - Unique email address of the repository.

* `avatar_url` - Avatar URL of the repository.
