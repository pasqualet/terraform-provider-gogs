package gogs

import (
	"fmt"

	gogsclient "github.com/gogits/go-gogs-client"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceGogsUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceGogsUserCreate,
		Update: resourceGogsUserUpdate,
		Read:   resourceGogsUserRead,
		Exists: resourceGogsUserExists,
		Delete: resourceGogsUserDelete,
		Schema: map[string]*schema.Schema{
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"login_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},
			"fullname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},
			"email": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},
			"password": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},
			"send_notify": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: false,
				Default:  true,
			},
			"website": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: false,
			},
			"location": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: false,
			},
			"active": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: false,
				Default:  true,
			},
			"admin": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: false,
				Default:  false,
			},
			"allow_git_hook": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: false,
				Default:  false,
			},
			"allow_import_local": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: false,
				Default:  false,
			},
			"max_repo_creation": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: false,
				Default:  -1,
			},
		},
	}
}

func readGogsUser(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceGogsUserRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceGogsUserCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*gogsclient.Client)

	createUserOption := gogsclient.CreateUserOption{
		LoginName:  d.Get("login_name").(string),
		Username:   d.Get("username").(string),
		FullName:   d.Get("fullname").(string),
		Email:      d.Get("email").(string),
		Password:   d.Get("password").(string),
		SendNotify: d.Get("send_notify").(bool),
	}

	user, err := client.AdminCreateUser(createUserOption)
	if err != nil {
		return err
	}

	id := fmt.Sprintf("%d", user.ID)
	d.SetId(id)

	return resourceGogsUserUpdate(d, meta)
}

func resourceGogsUserUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*gogsclient.Client)

	username := d.Get("username").(string)

	editUserOption := gogsclient.EditUserOption{
		LoginName:        d.Get("login_name").(string),
		FullName:         d.Get("fullname").(string),
		Email:            d.Get("email").(string),
		Password:         d.Get("password").(string),
		Website:          d.Get("website").(string),
		Location:         d.Get("location").(string),
		Active:           Bool(d.Get("active").(bool)),
		Admin:            Bool(d.Get("admin").(bool)),
		AllowGitHook:     Bool(d.Get("allow_git_hook").(bool)),
		AllowImportLocal: Bool(d.Get("allow_import_local").(bool)),
		MaxRepoCreation:  Int(d.Get("max_repo_creation").(int)),
	}

	err := client.AdminEditUser(username, editUserOption)
	if err != nil {
		return err
	}

	return readGogsUser(d, meta)
}

func resourceGogsUserDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*gogsclient.Client)

	username := d.Get("username").(string)

	err := client.AdminDeleteUser(username)
	return err
}

func resourceGogsUserExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	client := meta.(*gogsclient.Client)

	username := d.Get("username").(string)

	_, err := client.GetUserInfo(username)
	return err == nil, err
}
