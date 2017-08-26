package gogs

import (
	"fmt"

	gogsclient "github.com/gogits/go-gogs-client"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceGogsUser() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGogsUserRead,
		Schema: map[string]*schema.Schema{
			"username": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"fullname": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"email": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"avatar_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceGogsUserRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*gogsclient.Client)

	username, usernameOk := d.GetOk("username")

	if usernameOk == false {
		return fmt.Errorf("username not set")
	}

	user, err := client.GetUserInfo(username.(string))
	if err != nil {
		return err
	}

	id := fmt.Sprintf("%d", user.ID)
	d.SetId(id)
	d.Set("username", user.UserName)
	d.Set("fullname", user.FullName)
	d.Set("email", user.Email)
	d.Set("avatar_url", user.AvatarUrl)
	return nil
}
