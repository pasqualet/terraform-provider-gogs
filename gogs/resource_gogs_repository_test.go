package gogs

import (
	"fmt"
	"testing"

	gogsclient "github.com/gogits/go-gogs-client"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

var testAccGogsRepositoryConfig = fmt.Sprintf(`
resource "gogs_repository" "test" {
    username = "test"
	name = "resourcetest"
	description = "Just a test"
}
`)

func TestAccGogsRepository_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccGogsRepositoryDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccGogsRepositoryConfig,
				Check: resource.ComposeTestCheckFunc(
					testCheckGogsRepositoryExists("gogs_repository.test", t),
				),
			},
		},
	})
}

func testCheckGogsRepositoryExists(n string, t *testing.T) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*gogsclient.Client)

		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		username := rs.Primary.Attributes["username"]
		if username == "" {
			return fmt.Errorf("No Username is set")
		}

		name := rs.Primary.Attributes["name"]
		if name == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.GetRepo(username, name)
		return err
	}
}

func testAccGogsRepositoryDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*gogsclient.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "gogs_user" {
			continue
		}

		username := rs.Primary.Attributes["username"]
		name := rs.Primary.Attributes["name"]
		_, err := client.GetRepo(username, name)

		if err == nil {
			return fmt.Errorf("Repository %s/%s still exists", username, name)
		}
	}

	return nil
}
