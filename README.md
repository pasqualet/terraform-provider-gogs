Terraform Provider
==================

Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) 0.10.x
-	[Go](https://golang.org/doc/install) 1.8 (to build the provider plugin)

Building The Provider
---------------------

Clone repository to: `$GOPATH/src/github.com/pasqualet/terraform-provider-gogs`

```sh
$ mkdir -p $GOPATH/src/github.com/pasqualet; cd $GOPATH/src/github.com/pasqualet
$ git clone git@github.com:pasqualet/terraform-provider-gogs
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/pasqualet/terraform-provider-gogs
$ make build
```

Using the provider
----------------------

```
provider "gogs" {
    token = "3a626845fc4922fff0e84c6516b6adc8b14ff7b2"
    url = "http://localhost:3000"
}

resource "gogs_user" "sample_user" {
    username = "sample"
    fullname = "Sample User"
    login_name = "sample"
    email = "sample@sample.com"
    password = "password"
}

resource "gogs_repository" "sample" {
    username = "${gogs_user.sample_user.username}"
    name = "sample-repository"
}
```

Developing the Provider
---------------------------

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.8+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make build
...
$ $GOPATH/bin/terraform-provider-gogs
...
```

In order to test the provider, you can simply run `make test`.

```sh
$ make test
```

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Before run the acceptance tests set the following environment variables:

```sh
export GOGS_URL=https://gogs.example.com # The target Gogs instance.
export GOGS_TOKEN=3a626845fc4922fff0e84c6516b6adc8b14ff7b2 # The Gogs personal access token.
export GOGS_USERNAME=user1 # The user already existing in Gogs.
export GOGS_REPOSITORY=repository # The repository alreaddy existing in Gogs.
```

```sh
$ make testacc
```
