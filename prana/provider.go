package prana

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jmoiron/sqlx"
	"github.com/phogolabs/prana"
)

// NewProvider creates a new provider
func NewProvider() *schema.Provider {
	configure := func(d *schema.ResourceData) (interface{}, error) {
		driver, conn, err := prana.ParseURL(d.Get("database_url").(string))
		if err != nil {
			return nil, err
		}

		return sqlx.Open(driver, conn)
	}

	return &schema.Provider{
		ConfigureFunc: configure,
		Schema: map[string]*schema.Schema{
			"database_url": {
				Type:        schema.TypeString,
				DefaultFunc: schema.EnvDefaultFunc("PRANA_DB_URL", nil),
				Description: "Database URL",
				Required:    true,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"prana_migration": NewMigrationResource(),
		},
	}
}
