package gogs

import (
	"os"
	"testing"

	gogsclient "github.com/gogits/go-gogs-client"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider
var testAccGogsClient *gogsclient.Client

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"gogs": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

func testAccPreCheck(t *testing.T) {
	v := os.Getenv("GOGS_URL")
	if v == "" {
		t.Fatal("GOGS_URL must be set for acceptance tests")
	}

	v = os.Getenv("GOGS_TOKEN")
	if v == "" {
		t.Fatal("GOGS_TOKEN must be set for acceptance tests")
	}

	if testAccGogsClient == nil {
		config := Config{
			URL:   os.Getenv("GOGS_URL"),
			Token: os.Getenv("GOGS_TOKEN"),
		}

		if client, err := config.loadAndValidate(); err != nil {
			t.Fatalf("could not load Gogs Client: %s", err)
		} else {
			testAccGogsClient = client
		}
	}
}
