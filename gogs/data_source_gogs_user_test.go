package gogs

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccGogsUserDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccGogsUserDatasourceConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccGogsUserDatasource("data.gogs_user.test"),
				),
			},
		},
	})
}

func testAccGogsUserDatasource(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Can't find user data source: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Cannot find user %s", rs.Primary.Attributes["username"])
		}

		return nil
	}
}

var testAccGogsUserDatasourceConfig = fmt.Sprintf(`
data "gogs_user" "test" {
  username = "%s"
}
`, os.Getenv("GOGS_USERNAME"))
