# terraform-provider-prana

A Terraform Provider for Prana

# Getting Started

```hcl
provider "prana" {
  sql_driver     = "cloudsqlpostgres"
  sql_connection = local.prana_sql_connection
}

resource "prana_migration" "publication_api" {
  script_dir = "${path.module}/../../database/migration"
}
```
