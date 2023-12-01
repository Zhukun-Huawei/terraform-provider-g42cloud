package acceptance

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/g42cloud-terraform/terraform-provider-g42cloud/g42cloud/services/acceptance"
	"github.com/g42cloud-terraform/terraform-provider-g42cloud/g42cloud/services/acceptance/common"
)

func TestAccEvsVolumesDataSource_basic(t *testing.T) {
	dataSourceName := "data.g42cloud_evs_volumes.test"
	dc := acceptance.InitDataSourceCheck(dataSourceName)
	rName := acceptance.RandomAccResourceName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.TestAccPreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccEvsVolumesDataSource_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					dc.CheckResourceExists(),
					resource.TestCheckResourceAttr(dataSourceName, "volumes.#", "5"),
				),
			},
		},
	})
}

func testAccEvsVolumesDataSource_base(rName string) string {
	return fmt.Sprintf(`
variable "volume_configuration" {
  type = list(object({
    suffix      = string
    size        = number
    device_type = string
    multiattach = bool
  }))
  default = [
    {suffix = "vbd_normal_volume", size = 100, device_type = "VBD", multiattach = false},
    {suffix = "vbd_share_volume", size = 100, device_type = "VBD", multiattach = true},
    {suffix = "scsi_normal_volume", size = 100, device_type = "SCSI", multiattach = false},
    {suffix = "scsi_share_volume", size = 100, device_type = "SCSI", multiattach = true},
  ]
}

%[1]s

resource "g42cloud_compute_instance" "test" {
  availability_zone = data.g42cloud_availability_zones.test.names[0]
  name              = "%[2]s"
  image_id          = data.g42cloud_images_image.test.id
  flavor_id         = "m6.large.8"

  system_disk_type = "SSD"
  system_disk_size = 50

  security_group_ids = [
    g42cloud_networking_secgroup.test.id
  ]

  network {
    uuid = g42cloud_vpc_subnet.test.id
  }
}

resource "g42cloud_evs_volume" "test" {
  count = length(var.volume_configuration)
  
  availability_zone = data.g42cloud_availability_zones.test.names[0]
  volume_type       = "SAS"
  name              = "%[2]s_${var.volume_configuration[count.index].suffix}"
  size              = var.volume_configuration[count.index].size
  device_type       = var.volume_configuration[count.index].device_type
  multiattach       = var.volume_configuration[count.index].multiattach

  tags = {
    index = tostring(count.index)
  }
}

resource "g42cloud_compute_volume_attach" "test" {
  count = length(g42cloud_evs_volume.test)

  instance_id = g42cloud_compute_instance.test.id
  volume_id   = g42cloud_evs_volume.test[count.index].id
}
`, common.TestBaseComputeResources(rName), rName)
}

func testAccEvsVolumesDataSource_basic(rName string) string {
	return fmt.Sprintf(`
%s

data "g42cloud_evs_volumes" "test" {
  depends_on = [g42cloud_compute_volume_attach.test]

  availability_zone = data.g42cloud_availability_zones.test.names[0]
  server_id         = g42cloud_compute_instance.test.id
  status            = "in-use"
}
`, testAccEvsVolumesDataSource_base(rName))
}
