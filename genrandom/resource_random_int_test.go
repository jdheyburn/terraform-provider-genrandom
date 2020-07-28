package genrandom

import (
	"fmt"
	"regexp"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/jdheyburn/terraform-provider-genrandom/model"
)

var testProviders = map[string]terraform.ResourceProvider{
	"genrandom": Provider(),
}

func TestRandomInt_Deterministic(t *testing.T) {
	resourceName := "genrandom_int.rando"
	var rInt model.RandomInt
	min, max := 1, 1
	resource.UnitTest(t, resource.TestCase{
		Providers: testProviders,
		Steps: []resource.TestStep{
			{
				Config: formatResource(min, max),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRandomIntExists(resourceName, &rInt),
					resource.TestCheckResourceAttr(resourceName, "value", "1"),
					testAccCheckRandomValueInRange(min, max, &rInt),
				),
			},
		},
	})
}

func TestRandomInt_expectError(t *testing.T) {
	min, max := 10, -10
	resource.UnitTest(t, resource.TestCase{
		Providers: testProviders,
		Steps: []resource.TestStep{
			{
				Config:      formatResource(min, max),
				ExpectError: regexp.MustCompile("max cannot be lower than min"),
			},
		},
	})
}

func TestRandomInt_withinRange(t *testing.T) {
	resourceName := "genrandom_int.rando"
	var rInt model.RandomInt
	min, max := 5000, 10000
	resource.UnitTest(t, resource.TestCase{
		Providers: testProviders,
		Steps: []resource.TestStep{
			{
				Config: formatResource(min, max),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRandomIntExists(resourceName, &rInt),
					testAccCheckRandomValueInRange(min, max, &rInt),
				),
			},
		},
	})
}

func testAccCheckRandomIntExists(n string, rInt *model.RandomInt) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No RandomInt ID is set")
		}

		// In place of a proper API endpoint to query...
		min, _ := strconv.Atoi(rs.Primary.Attributes["min"])
		max, _ := strconv.Atoi(rs.Primary.Attributes["max"])
		value, _ := strconv.Atoi(rs.Primary.Attributes["value"])

		*rInt = *&model.RandomInt{
			Min:   min,
			Max:   max,
			Value: value,
		}

		return nil
	}
}

func testAccCheckRandomValueInRange(min, max int, rInt *model.RandomInt) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if rInt.Value < min {
			return fmt.Errorf("bad value, expected < %d, got: %#v", min, rInt.Value)
		}
		if rInt.Value > max {
			return fmt.Errorf("bad value, expected > %d, got: %#v", max, rInt.Value)
		}
		return nil
	}
}

func formatResource(min, max int) string {
	return fmt.Sprintf(`
resource "genrandom_int" "rando" {
  min = %d
  max = %d
}
`, min, max)
}
