package genrandom

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

// func testTemplateDirWriteFiles(files map[string]testTemplate) (in, out string, err error) {
// 	in, err = ioutil.TempDir(os.TempDir(), "terraform_template_dir")
// 	if err != nil {
// 		return
// 	}

// 	for name, file := range files {
// 		path := filepath.Join(in, name)

// 		err = os.MkdirAll(filepath.Dir(path), 0777)
// 		if err != nil {
// 			return
// 		}

// 		err = ioutil.WriteFile(path, []byte(file.template), 0777)
// 		if err != nil {
// 			return
// 		}
// 	}

// 	out = fmt.Sprintf("%s.out", in)
// 	return
// }

var testProviders = map[string]terraform.ResourceProvider{
	"genrandom": Provider(),
}

func TestRandomIntDeterministic(t *testing.T) {
	resourceName := "genrandom_int.rando"
	// Run test case.
	resource.UnitTest(t, resource.TestCase{
		Providers: testProviders,
		Steps: []resource.TestStep{
			{
				Config: testRandomIntBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "value", "1"),
				),
			},
			// {
			// 	Config: testRandomIntBasicUpdated(),
			// 	Check: r.ComposeTestCheckFunc(
			// 		r.TestCheckResourceAttr(resourceName, "value", "2"),
			// 		// func(*terraform.State) error {
			// 		// 	if *before.BaselineId != *after.BaselineId {
			// 		// 		t.Fatal("Baseline IDs changed unexpectedly")
			// 		// 	}
			// 		// 	return nil
			// 		// },
			// 	),
			// },
		},
		// CheckDestroy: func(*terraform.State) error {
		// 	if _, err := os.Stat(out); os.IsNotExist(err) {
		// 		return nil
		// 	}
		// 	return errors.New("random_int did not get destroyed")
		// },
	})
}

// func TestRandomIntNonDeterministic(t *testing.T) {
// 	resourceName := "random_int.rando"
// 	// Run test case.
// 	resource.UnitTest(t, resource.TestCase{
// 		Providers: testProviders,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testRandomIntBasicUpdated(),
// 				Check:  resource.ComposeTestCheckFunc(
// 					resource.TestCheckResourceAttr(resourceName, "value", "1"),
// 				),
// 			},
// 		},
// 		// CheckDestroy: func(*terraform.State) error {
// 		// 	if _, err := os.Stat(out); os.IsNotExist(err) {
// 		// 		return nil
// 		// 	}
// 		// 	return errors.New("random_int did not get destroyed")
// 		// },
// 	})
// }

// func testRandomIntNonDeterministicCheck(s *terraform.State) error {

// }

func testRandomIntBasic() string {
	return fmt.Sprintf(`
resource "genrandom_int" "rando" {
  min = 1
  max = 1
}
`)
}

func testRandomIntBasicUpdated() string {
	return fmt.Sprintf(`
resource "genrandom_int" "rando" {
  min = 1
  max = 10
}
`)
}
