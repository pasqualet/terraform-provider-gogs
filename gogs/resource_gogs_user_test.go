package gogs

import (
	"fmt"
	"testing"

	gogsclient "github.com/gogits/go-gogs-client"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

var testAccGogsUserConfig = fmt.Sprintf(`
resource "gogs_user" "testuser" {
    username = "testuser"
    fullname = "testuser"
    login_name = "testuser"
    email = "test@example.com"
    password = "pass"
    website = "https://github.com/pasqualet"
    location = "everywhere"
    active = true
    admin = true
    allow_git_hook = true
    allow_import_local = true
    max_repo_creation = 5
}
`)

func TestAccGogsUser_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccGogsUserDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccGogsUserConfig,
				Check: resource.ComposeTestCheckFunc(
					testCheckGogsUserExists("gogs_user.testuser", t),
				),
			},
		},
	})
}

func testCheckGogsUserExists(n string, t *testing.T) resource.TestCheckFunc {
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

		_, err := client.GetUserInfo(username)
		return err
	}
}

func testAccGogsUserDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*gogsclient.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "gogs_user" {
			continue
		}

		username := rs.Primary.Attributes["username"]
		_, err := client.GetUserInfo(username)

		if err == nil {
			return fmt.Errorf("User %s still exists", username)
		}
	}

	return nil
}
