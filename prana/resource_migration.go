package prana

import (
	"os"
	"path/filepath"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/jmoiron/sqlx"
	"github.com/phogolabs/prana/sqlmigr"
)

// NewMigrationResource creates a new migration resource
func NewMigrationResource() *schema.Resource {
	executor := func(d *schema.ResourceData, m interface{}) (*sqlmigr.Executor, error) {
		db := m.(*sqlx.DB)

		dir, err := filepath.Abs(d.Get("script_dir").(string))
		if err != nil {
			return nil, err
		}

		return &sqlmigr.Executor{
			Provider: &sqlmigr.Provider{
				FileSystem: os.DirFS(dir),
				DB:         db,
			},
			Runner: &sqlmigr.Runner{
				FileSystem: os.DirFS(dir),
				DB:         db,
			},
		}, nil
	}

	read := func(d *schema.ResourceData, m interface{}) error {
		executor, err := executor(d, m)
		if err != nil {
			return err
		}

		migrations, err := executor.Migrations()
		if err != nil {
			return err
		}

		for index, migration := range migrations {
			if migration.CreatedAt.IsZero() {
				if prev := index - 1; prev > 0 {
					d.SetId(migrations[prev].String())
				}
				break
			}
		}

		return nil
	}

	migrate := func(d *schema.ResourceData, m interface{}) error {
		executor, err := executor(d, m)
		if err != nil {
			return err
		}

		if _, err = executor.RunAll(); err != nil {
			return err
		}

		migrations, err := executor.Migrations()
		if err != nil {
			return err
		}

		if index := len(migrations) - 1; index >= 0 {
			d.SetId(migrations[index].String())
		}

		return err
	}

	revert := func(d *schema.ResourceData, m interface{}) error {
		executor, err := executor(d, m)
		if err != nil {
			return err
		}

		if _, err = executor.RevertAll(); err == nil {
			d.SetId("")
		}

		return err
	}

	return &schema.Resource{
		Create: migrate,
		Update: migrate,
		Read:   read,
		Delete: revert,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"script_dir": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
