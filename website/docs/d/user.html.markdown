---
layout: "gogs"
page_title: "Gogs: gogs_user"
sidebar_current: "docs-gogs-datasource-user-x"
description: |-
    Provides details about a specific user
---

# gogs\_user

`gogs_user` provides details about a specific user.

## Example Usage


```hcl
data "gogs_user" "sample" {
    username = "sample"
}
```

## Argument Reference

The given filters must match exactly the userwhose data will
be exported as attributes.

* `username` - (Required) Unique user name.

## Attributes Reference

The following attribute are additionally exported:

* `id` - Integer that uniquely identifies the user.

* `fullname` - Full name of user.

* `email` - Unique email address of the user.

* `avatar_url` - Avatar URL of the user.
