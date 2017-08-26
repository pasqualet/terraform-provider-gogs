package gogs

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"url": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("GOGS_URL", ""),
				Description: "The Gogs url.",
			},
			"token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("GOGS_TOKEN", ""),
				Description: "The Gogs API token.",
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"gogs_repository": dataSourceGogsRepository(),
			"gogs_user":       dataSourceGogsUser(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"gogs_repository": resourceGogsRepository(),
			"gogs_user":       resourceGogsUser(),
		},
		ConfigureFunc: configureProvider,
	}
}

func configureProvider(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		URL:   d.Get("url").(string),
		Token: d.Get("token").(string),
	}

	client, err := config.loadAndValidate()

	if err != nil {
		return nil, err
	}

	return client, nil
}
