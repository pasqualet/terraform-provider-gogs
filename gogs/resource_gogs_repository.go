package gogs

import (
	"fmt"

	gogsclient "github.com/gogits/go-gogs-client"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceGogsRepository() *schema.Resource {
	return &schema.Resource{
		Create: resourceGogsRepositoryCreate,
		Read:   resourceGogsRepositoryRead,
		Exists: resourceGogsRepositoryExists,
		Delete: resourceGogsRepositoryDelete,
		Schema: map[string]*schema.Schema{
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"private": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
			"auto_init": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
			"gitignores": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"license": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"readme": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func readGogsRepository(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceGogsRepositoryRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceGogsRepositoryCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*gogsclient.Client)

	username := d.Get("username").(string)
	createRepoOption := gogsclient.CreateRepoOption{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Private:     d.Get("private").(bool),
		AutoInit:    d.Get("auto_init").(bool),
		Gitignores:  d.Get("gitignores").(string),
		License:     d.Get("license").(string),
		Readme:      d.Get("readme").(string),
	}

	repository, err := client.AdminCreateRepo(username, createRepoOption)
	if err != nil {
		return err
	}

	id := fmt.Sprintf("%d", repository.ID)
	d.SetId(id)

	return nil
}

func resourceGogsRepositoryDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*gogsclient.Client)

	username := d.Get("username").(string)
	name := d.Get("name").(string)

	err := client.DeleteRepo(username, name)
	return err
}

func resourceGogsRepositoryExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	client := meta.(*gogsclient.Client)

	username := d.Get("username").(string)
	name := d.Get("name").(string)

	_, err := client.GetRepo(username, name)
	return err == nil, nil
}
