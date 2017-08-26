---
layout: "gogs"
page_title: "Gogs: gogs_user"
sidebar_current: "docs-gogs-resource-user-x"
description: |-
  Creates and manages Gogs users
---

# gogs\_user

This resource allows you to create and manage Gogs users.


## Example Usage

```hcl
resource "gogs_user" "sample_user" {
    username = "sample"
    fullname = "Sample User"
    login_name = "sample"
    email = "sample@sample.com"
    password = "password"
}
```

## Argument Reference

The following arguments are supported:

* `username` - (Required) Unique user name.

* `email` - (Required) Unique email address of user.

* `fullname` - (Required) Full name of user.

* `login_name` - (Required) Authentication source login name.

* `password` - (Required) Default password for user.

* `send_notify` - (Optional) Send a notification email for this creation, require mailer service enabled. Boolean value. (Default: `true`).

* `website` - (Optional) Personal website of user.

* `location` - (Optional) The user location.

* `active` - (Optional) Active or deactive the user. Boolean value. (Default: `true`).

* `admin` - (Optional) Promote or depromote the user to be site admin. Boolean value. (Default: `false`).

* `allow_git_hook` - (Optional) Allow or disallow the user for using Git hooks. Boolen value. (Default: `false`).

* `allow_import_local` - (Optional) Allow or disallow the user for importing local repositories. Boolean value. (Default: `false`).

* `max_repo_creation` - (Optional) The max number of repositories that can be created. Integer value. (Default: `-1`).

## Attributes Reference

The following additional attributes are exported:

* `id` - Integer that uniquely identifies the user.
