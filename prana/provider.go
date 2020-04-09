package prana

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jmoiron/sqlx"
)

// NewProvider creates a new provider
func NewProvider() *schema.Provider {
	configure := func(d *schema.ResourceData) (interface{}, error) {
		var (
			driver = d.Get("sql_driver").(string)
			conn   = d.Get("sql_connection").(string)
		)

		return sqlx.Open(driver, conn)
	}

	return &schema.Provider{
		ConfigureFunc: configure,
		Schema: map[string]*schema.Schema{
			"sql_driver": {
				Type:        schema.TypeString,
				Description: "SQL Database Driver",
				Required:    true,
			},
			"sql_connection": {
				Type:        schema.TypeString,
				Description: "SQL Database Connection",
				Required:    true,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"prana_migration": NewMigrationResource(),
		},
	}
}
