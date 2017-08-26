package gogs

import (
	"fmt"

	gogsclient "github.com/gogits/go-gogs-client"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceGogsRepository() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGogsRepositoryRead,

		Schema: map[string]*schema.Schema{
			"username": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"full_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"private": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"fork": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"parent_username": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"empty": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"mirror": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"size": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"html_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ssh_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"clone_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"website": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"stars": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"forks": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"watchers": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"open_issue_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"default_branch": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"permission_admin": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"permission_push": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"permission_pull": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func dataSourceGogsRepositoryRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*gogsclient.Client)

	username, usernameOk := d.GetOk("username")
	name, nameOk := d.GetOk("name")

	if usernameOk == false {
		return fmt.Errorf("username not set")
	}

	if nameOk == false {
		return fmt.Errorf("name not set")
	}

	repository, err := client.GetRepo(username.(string), name.(string))
	if err != nil {
		return err
	}

	id := fmt.Sprintf("%d", repository.ID)
	d.SetId(id)
	d.Set("name", repository.Name)
	d.Set("description", repository.Description)
	d.Set("full_name", repository.FullName)
	d.Set("description", repository.Description)
	d.Set("private", repository.Private)
	d.Set("fork", repository.Fork)
	if repository.Parent != nil {
		d.Set("parent_username", repository.Parent)
		d.Set("parent_name", repository.Parent)
	}
	d.Set("empty", repository.Empty)
	d.Set("mirror", repository.Mirror)
	d.Set("size", repository.Size)
	d.Set("html_url", repository.HTMLURL)
	d.Set("ssh_url", repository.SSHURL)
	d.Set("clone_url", repository.CloneURL)
	d.Set("website", repository.Website)
	d.Set("stars", repository.Stars)
	d.Set("forks", repository.Forks)
	d.Set("watchers", repository.Watchers)
	d.Set("open_issue_count", repository.OpenIssues)
	d.Set("default_branch", repository.DefaultBranch)
	d.Set("created", repository.Created)
	d.Set("updated", repository.Updated)
	d.Set("permission_admin", repository.Permissions.Admin)
	d.Set("permission_push", repository.Permissions.Push)
	d.Set("permission_pull", repository.Permissions.Pull)
	return nil
}
