package provider

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestToken(t *testing.T){
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccTokenConfiguration(),
			},
		},
	})
}

func testAccTokenConfiguration() string {
	return fmt.Sprintf(`
resource "astra_token" "example" {
  roles = ["puppies"]
}
`)
}