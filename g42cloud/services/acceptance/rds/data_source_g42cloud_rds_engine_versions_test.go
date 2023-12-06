package rds

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/g42cloud-terraform/terraform-provider-g42cloud/g42cloud/services/acceptance"
)

func TestAccRdsEngineVersionsV3DataSource_basic(t *testing.T) {
	dataSourceName := "data.g42cloud_rds_engine_versions.test"

	dc := acceptance.InitDataSourceCheck(dataSourceName)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.TestAccPreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccRdsEngineVersionsV3DataSource_basic(),
				Check: resource.ComposeTestCheckFunc(
					dc.CheckResourceExists(),
					resource.TestCheckResourceAttr(dataSourceName, "type", "MySQL"),
					resource.TestMatchResourceAttr(dataSourceName, "versions.#", regexp.MustCompile("\\d+")),
				),
			},
		},
	})
}

func TestAccRdsEngineVersionsV3DataSource_PostgreSQL_basic(t *testing.T) {
	dataSourceName := "data.g42cloud_rds_engine_versions.test"

	dc := acceptance.InitDataSourceCheck(dataSourceName)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.TestAccPreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccRdsEngineVersionsV3DataSource_PostgreSQL_basic(),
				Check: resource.ComposeTestCheckFunc(
					dc.CheckResourceExists(),
					resource.TestCheckResourceAttr(dataSourceName, "type", "PostgreSQL"),
					resource.TestMatchResourceAttr(dataSourceName, "versions.#", regexp.MustCompile("\\d+")),
				),
			},
		},
	})
}

func TestAccRdsEngineVersionsV3DataSource_SQLServer_basic(t *testing.T) {
	dataSourceName := "data.g42cloud_rds_engine_versions.test"

	dc := acceptance.InitDataSourceCheck(dataSourceName)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.TestAccPreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccRdsEngineVersionsV3DataSource_SQLServer_basic(),
				Check: resource.ComposeTestCheckFunc(
					dc.CheckResourceExists(),
					resource.TestCheckResourceAttr(dataSourceName, "type", "SQLServer"),
					resource.TestMatchResourceAttr(dataSourceName, "versions.#", regexp.MustCompile("\\d+")),
				),
			},
		},
	})
}

func testAccRdsEngineVersionsV3DataSource_basic() string {
	return fmt.Sprintf(`
data "g42cloud_rds_engine_versions" "test" {
  type = "MySQL"
}
`)
}

func testAccRdsEngineVersionsV3DataSource_PostgreSQL_basic() string {
	return fmt.Sprintf(`
data "g42cloud_rds_engine_versions" "test" {
  type = "PostgreSQL"
}
`)
}

func testAccRdsEngineVersionsV3DataSource_SQLServer_basic() string {
	return fmt.Sprintf(`
data "g42cloud_rds_engine_versions" "test" {
  type = "SQLServer"
}
`)
}
