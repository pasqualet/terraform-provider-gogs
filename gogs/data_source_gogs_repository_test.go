package gogs

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccGogsRepositoryDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccGogsRepositoryDatasourceConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccGogsRepositoryDatasource("data.gogs_repository.test"),
				),
			},
		},
	})
}

func testAccGogsRepositoryDatasource(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Can't find repository data source: %s", n)
		}

		if rs.Primary.ID == "" {
			username := rs.Primary.Attributes["username"]
			name := rs.Primary.Attributes["name"]
			return fmt.Errorf("Cannot find repository %s/%s", username, name)
		}

		return nil
	}
}

var testAccGogsRepositoryDatasourceConfig = fmt.Sprintf(`
data "gogs_repository" "test" {
  username = "%s"
  name = "%s"
}
`, os.Getenv("GOGS_USERNAME"), os.Getenv("GOGS_REPOSITORY"))
