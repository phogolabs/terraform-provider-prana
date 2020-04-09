package prana

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jmoiron/sqlx"
)

// NewProvider creates a new provider
func NewProvider() *schema.Provider {
	configure := func(d *schema.ResourceData) (interface{}, error) {
		var (
			driver = d.Get("database_driver").(string)
			conn   = d.Get("database_connection").(string)
		)

		return sqlx.Open(driver, conn)
	}

	return &schema.Provider{
		ConfigureFunc: configure,
		Schema: map[string]*schema.Schema{
			"database_driver": {
				Type:        schema.TypeString,
				Description: "Database Driver",
				Required:    true,
			},
			"database_connection": {
				Type:        schema.TypeString,
				Description: "Database Connection",
				Required:    true,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"prana_migration": NewMigrationResource(),
		},
	}
}
