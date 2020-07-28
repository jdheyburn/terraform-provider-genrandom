package genrandom

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceRandomInt() *schema.Resource {
	return &schema.Resource{
		Create: resourceRandomIntCreate,
		Read:   resourceRandomIntRead,
		Delete: resourceRandomIntDelete,

		Schema: map[string]*schema.Schema{
			"min": {
				Type:        schema.TypeInt,
				Description: "The lower bound for the random number",
				Required:    true,
				ForceNew:    true,
			},
			"max": {
				Type:        schema.TypeInt,
				Description: "The upper bound for the random number",
				Required:    true,
				ForceNew:    true,
			},
			"value": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func resourceRandomIntRead(d *schema.ResourceData, meta interface{}) error {
	min := d.Get("min").(int)
	max := d.Get("max").(int)

	id := generateID(min, max)
	if id != d.Id() {
		d.SetId("")
		return nil
	}

	return nil
}

func resourceRandomIntCreate(d *schema.ResourceData, meta interface{}) error {

	min := d.Get("min").(int)
	max := d.Get("max").(int)

	if max < min {
		return fmt.Errorf("max cannot be lower than min")
	}

	rand.Seed(time.Now().UnixNano())

	random := rand.Intn(max-min+1) + min

	id := generateID(min, max)

	d.SetId(id)
	d.Set("value", random)

	return nil
}

func resourceRandomIntDelete(d *schema.ResourceData, _ interface{}) error {
	d.SetId("")

	return nil
}

func generateID(min, max int) string {
	return fmt.Sprintf("%d:%d", min, max)
}
