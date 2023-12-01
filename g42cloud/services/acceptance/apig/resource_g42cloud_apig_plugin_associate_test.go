package apig

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/chnsz/golangsdk/openstack/apigw/dedicated/v2/plugins"

	"github.com/g42cloud-terraform/terraform-provider-g42cloud/g42cloud/services/acceptance"
	"github.com/g42cloud-terraform/terraform-provider-g42cloud/g42cloud/services/acceptance/common"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
)

func getPluginAssociateFunc(cfg *config.Config, state *terraform.ResourceState) (interface{}, error) {
	client, err := cfg.ApigV2Client(acceptance.G42_REGION_NAME)
	if err != nil {
		return nil, fmt.Errorf("error creating APIG v2 client: %s", err)
	}

	listOpts := plugins.ListBindOpts{
		InstanceId: state.Primary.Attributes["instance_id"],
		PluginId:   state.Primary.Attributes["plugin_id"],
		EnvId:      state.Primary.Attributes["env_id"],
	}
	return plugins.ListBind(client, listOpts)
}

func TestAccPluginAssociate_basic(t *testing.T) {
	var (
		bindList []plugins.BindApiInfo

		rName1 = "g42cloud_apig_plugin_associate.cors_bind"
		rName2 = "g42cloud_apig_plugin_associate.http_resp_bind"
		rName3 = "g42cloud_apig_plugin_associate.rate_limit_bind"
		rName4 = "g42cloud_apig_plugin_associate.kafka_log_bind"
		rName5 = "g42cloud_apig_plugin_associate.breaker_bind"
		name   = acceptance.RandomAccResourceName()

		rc1 = acceptance.InitResourceCheck(rName1, &bindList, getPluginAssociateFunc)
		rc2 = acceptance.InitResourceCheck(rName2, &bindList, getPluginAssociateFunc)
		rc3 = acceptance.InitResourceCheck(rName3, &bindList, getPluginAssociateFunc)
		rc4 = acceptance.InitResourceCheck(rName4, &bindList, getPluginAssociateFunc)
		rc5 = acceptance.InitResourceCheck(rName5, &bindList, getPluginAssociateFunc)
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acceptance.TestAccPreCheck(t)
		},
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      rc1.CheckResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccPluginAssociate_basic_step1(name),
				Check: resource.ComposeTestCheckFunc(
					rc1.CheckResourceExists(),
					resource.TestCheckResourceAttrPair(rName1, "instance_id",
						"g42cloud_apig_instance.test", "id"),
					resource.TestCheckResourceAttrPair(rName1, "plugin_id",
						"g42cloud_apig_plugin.cors", "id"),
					resource.TestCheckResourceAttrPair(rName1, "env_id",
						"g42cloud_apig_environment.test.0", "id"),
					resource.TestCheckResourceAttrPair(rName1, "api_ids.0",
						"g42cloud_apig_api.test", "id"),
					rc2.CheckResourceExists(),
					resource.TestCheckResourceAttrPair(rName2, "instance_id",
						"g42cloud_apig_instance.test", "id"),
					resource.TestCheckResourceAttrPair(rName2, "plugin_id",
						"g42cloud_apig_plugin.http_resp", "id"),
					resource.TestCheckResourceAttrPair(rName2, "env_id",
						"g42cloud_apig_environment.test.0", "id"),
					resource.TestCheckResourceAttrPair(rName2, "api_ids.0",
						"g42cloud_apig_api.test", "id"),
					rc3.CheckResourceExists(),
					resource.TestCheckResourceAttrPair(rName3, "instance_id",
						"g42cloud_apig_instance.test", "id"),
					resource.TestCheckResourceAttrPair(rName3, "plugin_id",
						"g42cloud_apig_plugin.rate_limit", "id"),
					resource.TestCheckResourceAttrPair(rName3, "env_id",
						"g42cloud_apig_environment.test.0", "id"),
					resource.TestCheckResourceAttrPair(rName3, "api_ids.0",
						"g42cloud_apig_api.test", "id"),
					rc4.CheckResourceExists(),
					resource.TestCheckResourceAttrPair(rName4, "instance_id",
						"g42cloud_apig_instance.test", "id"),
					resource.TestCheckResourceAttrPair(rName4, "plugin_id",
						"g42cloud_apig_plugin.kafka_log", "id"),
					resource.TestCheckResourceAttrPair(rName4, "env_id",
						"g42cloud_apig_environment.test.0", "id"),
					resource.TestCheckResourceAttrPair(rName4, "api_ids.0",
						"g42cloud_apig_api.test", "id"),
					rc5.CheckResourceExists(),
					resource.TestCheckResourceAttrPair(rName5, "instance_id",
						"g42cloud_apig_instance.test", "id"),
					resource.TestCheckResourceAttrPair(rName5, "plugin_id",
						"g42cloud_apig_plugin.breaker", "id"),
					resource.TestCheckResourceAttrPair(rName5, "env_id",
						"g42cloud_apig_environment.test.0", "id"),
					resource.TestCheckResourceAttrPair(rName5, "api_ids.0",
						"g42cloud_apig_api.test", "id"),
				),
			},
			{
				Config: testAccPluginAssociate_basic_step2(name),
				Check: resource.ComposeTestCheckFunc(
					rc1.CheckResourceExists(),
					resource.TestCheckResourceAttrPair(rName1, "instance_id",
						"g42cloud_apig_instance.test", "id"),
					resource.TestCheckResourceAttrPair(rName1, "plugin_id",
						"g42cloud_apig_plugin.cors", "id"),
					resource.TestCheckResourceAttrPair(rName1, "env_id",
						"g42cloud_apig_environment.test.1", "id"),
					resource.TestCheckResourceAttrPair(rName1, "api_ids.0",
						"g42cloud_apig_api.test", "id"),
					rc2.CheckResourceExists(),
					resource.TestCheckResourceAttrPair(rName2, "instance_id",
						"g42cloud_apig_instance.test", "id"),
					resource.TestCheckResourceAttrPair(rName2, "plugin_id",
						"g42cloud_apig_plugin.http_resp", "id"),
					resource.TestCheckResourceAttrPair(rName2, "env_id",
						"g42cloud_apig_environment.test.1", "id"),
					resource.TestCheckResourceAttrPair(rName2, "api_ids.0",
						"g42cloud_apig_api.test", "id"),
					rc3.CheckResourceExists(),
					resource.TestCheckResourceAttrPair(rName3, "instance_id",
						"g42cloud_apig_instance.test", "id"),
					resource.TestCheckResourceAttrPair(rName3, "plugin_id",
						"g42cloud_apig_plugin.rate_limit", "id"),
					resource.TestCheckResourceAttrPair(rName3, "env_id",
						"g42cloud_apig_environment.test.1", "id"),
					resource.TestCheckResourceAttrPair(rName3, "api_ids.0",
						"g42cloud_apig_api.test", "id"),
					rc4.CheckResourceExists(),
					resource.TestCheckResourceAttrPair(rName4, "instance_id",
						"g42cloud_apig_instance.test", "id"),
					resource.TestCheckResourceAttrPair(rName4, "plugin_id",
						"g42cloud_apig_plugin.kafka_log", "id"),
					resource.TestCheckResourceAttrPair(rName4, "env_id",
						"g42cloud_apig_environment.test.1", "id"),
					resource.TestCheckResourceAttrPair(rName4, "api_ids.0",
						"g42cloud_apig_api.test", "id"),
					rc5.CheckResourceExists(),
					resource.TestCheckResourceAttrPair(rName5, "instance_id",
						"g42cloud_apig_instance.test", "id"),
					resource.TestCheckResourceAttrPair(rName5, "plugin_id",
						"g42cloud_apig_plugin.breaker", "id"),
					resource.TestCheckResourceAttrPair(rName5, "env_id",
						"g42cloud_apig_environment.test.1", "id"),
					resource.TestCheckResourceAttrPair(rName5, "api_ids.0",
						"g42cloud_apig_api.test", "id"),
				),
			},
			{
				ResourceName:      rName1,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      rName2,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      rName3,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      rName4,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      rName5,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccPluginAssociate_base(name string) string {
	return fmt.Sprintf(`
%[1]s

data "g42cloud_dms_kafka_flavors" "test" {
  type = "cluster"
}

locals {
  query_results     = data.g42cloud_dms_kafka_flavors.test
  flavor            = data.g42cloud_dms_kafka_flavors.test.flavors[0]
  connect_addresses = split(",", g42cloud_dms_kafka_instance.test.connect_address)
  plugin_ids = [
    g42cloud_apig_plugin.cors.id,
    g42cloud_apig_plugin.http_resp.id,
    g42cloud_apig_plugin.rate_limit.id,
    g42cloud_apig_plugin.kafka_log.id,
    g42cloud_apig_plugin.breaker.id
  ]
}

resource "g42cloud_dms_kafka_instance" "test" {
  name              = "%[2]s"
  vpc_id            = g42cloud_vpc.test.id
  network_id        = g42cloud_vpc_subnet.test.id
  security_group_id = g42cloud_networking_secgroup.test.id

  flavor_id          = local.flavor.id
  storage_spec_code  = local.flavor.ios[0].storage_spec_code
  availability_zones = ["ae-ad-1a"]
  engine_version     = element(local.query_results.versions, length(local.query_results.versions)-1)
  storage_space      = local.flavor.properties[0].min_broker * local.flavor.properties[0].min_storage_per_node
  broker_num         = 3

  access_user      = "user"
  password         = "Kafkatest@123"
  manager_user     = "kafka-user"
  manager_password = "Kafkatest@123"

  lifecycle {
    ignore_changes = [
      availability_zones, manager_password, password,
    ]
  }
}

resource "g42cloud_dms_kafka_topic" "test" {
  instance_id = g42cloud_dms_kafka_instance.test.id
  name        = "%[2]s"
  partitions  = 20
}

resource "g42cloud_apig_instance" "test" {
  name                  = "%[2]s"
  edition               = "BASIC"
  vpc_id                = g42cloud_vpc.test.id
  subnet_id             = g42cloud_vpc_subnet.test.id
  security_group_id     = g42cloud_networking_secgroup.test.id
  enterprise_project_id = "0"

  availability_zones = try(slice(data.g42cloud_availability_zones.test.names, 0, 1), null)
}

resource "g42cloud_compute_instance" "test" {
  name               = "%[2]s"
  image_id           = data.g42cloud_images_image.test.id
  flavor_id          = "m6.large.8"
  security_group_ids = [g42cloud_networking_secgroup.test.id]
  availability_zone  = data.g42cloud_availability_zones.test.names[0]
  system_disk_type   = "SSD"

  network {
    uuid = g42cloud_vpc_subnet.test.id
  }
}

resource "g42cloud_apig_group" "test" {
  name        = "%[2]s"
  instance_id = g42cloud_apig_instance.test.id
}

resource "g42cloud_apig_vpc_channel" "test" {
  name        = "%[2]s"
  instance_id = g42cloud_apig_instance.test.id
  port        = 80
  algorithm   = "WRR"
  protocol    = "HTTP"
  path        = "/"
  http_code   = "201"

  members {
    id = g42cloud_compute_instance.test.id
  }
}

resource "g42cloud_apig_api" "test" {
  instance_id             = g42cloud_apig_instance.test.id
  group_id                = g42cloud_apig_group.test.id
  name                    = "%[2]s"
  type                    = "Public"
  request_protocol        = "HTTP"
  request_method          = "GET"
  request_path            = "/user_info/{user_age}"
  security_authentication = "APP"
  matching                = "Exact"

  request_params {
    name     = "user_age"
    type     = "NUMBER"
    location = "PATH"
    required = true
    maximum  = 200
    minimum  = 0
  }
  
  backend_params {
    type     = "REQUEST"
    name     = "userAge"
    location = "PATH"
    value    = "user_age"
  }

  web {
    path             = "/getUserAge/{userAge}"
    vpc_channel_id   = g42cloud_apig_vpc_channel.test.id
    request_method   = "GET"
    request_protocol = "HTTP"
    timeout          = 30000
  }
}

resource "g42cloud_apig_environment" "test" {
  count = 2

  name        = "%[2]s_${count.index}"
  instance_id = g42cloud_apig_instance.test.id
}

resource "g42cloud_apig_api_publishment" "test" {
  count = 2

  instance_id = g42cloud_apig_instance.test.id
  api_id      = g42cloud_apig_api.test.id
  env_id      = g42cloud_apig_environment.test[count.index].id
}

resource "g42cloud_apig_plugin" "cors" {
  instance_id = g42cloud_apig_instance.test.id
  name        = "%[2]s_cors"
  type        = "cors"
  content     = jsonencode(
    {
      allow_origin      = "*"
      allow_methods     = "GET,PUT,DELETE,HEAD,PATCH"
      allow_headers     = "Content-Type,Accept,Cache-Control"
      expose_headers    = "X-Request-Id,X-Apig-Latency"
      max_age           = 12700
      allow_credentials = true
    }
  )
}

resource "g42cloud_apig_plugin" "http_resp" {
  instance_id = g42cloud_apig_instance.test.id
  name        = "%[2]s_http_resp"
  type        = "set_resp_headers"
  content     = jsonencode(
    {
      response_headers = [{
        name   = "X-Custom-Pwd"
        value  = "**********"
        action = "override"
      }]
    }
  )
}

resource "g42cloud_apig_plugin" "rate_limit" {
  instance_id = g42cloud_apig_instance.test.id
  name        = "%[2]s_rate_limit"
  type        = "rate_limit"
  content     = jsonencode(
    {
      "scope": "basic",
      "default_time_unit": "minute",
      "default_interval": 1,
      "api_limit": 25,
      "app_limit": 15,
      "user_limit": 15,
      "ip_limit": 10,
      "algorithm": "counter",
      "specials": [],
      "parameters": [],
      "rules": []
    }
  )
}

resource "g42cloud_apig_plugin" "kafka_log" {
  instance_id = g42cloud_apig_instance.test.id
  name        = "%[2]s_kafka_log"
  type        = "kafka_log"
  content     = jsonencode(
    {
      "broker_list": [for v in local.connect_addresses: format("%%s:%%d", v, g42cloud_dms_kafka_instance.test.port)],
      "topic": "${g42cloud_dms_kafka_topic.test.name}",
      "key": "",
      "max_retry_count": 0,
      "retry_backoff": 1,
      "sasl_config": {
        "security_protocol": "PLAINTEXT",
        "sasl_mechanisms": "PLAIN",
        "sasl_username": "",
        "sasl_password": "",
        "ssl_ca_content": ""
      },
      "meta_config": {
        "system": {
          "start_time": false,
          "request_id": false,
          "client_ip": false,
          "api_id": false,
          "user_name": false,
          "app_id": false,
          "access_model1": false,
          "request_time": true,
          "http_status": true,
          "server_protocol": false,
          "scheme": true,
          "request_method": true,
          "host": false,
          "api_uri_mode": false,
          "uri": false,
          "request_size": false,
          "response_size": false,
          "upstream_uri": false,
          "upstream_addr": false,
          "upstream_status": true,
          "upstream_connect_time": false,
          "upstream_header_time": false,
          "upstream_response_time": false,
          "all_upstream_response_time": false,
          "region_id": true,
          "auth_type": false,
          "http_x_forwarded_for": false,
          "http_user_agent": false,
          "error_type": false,
          "access_model2": false,
          "inner_time": false,
          "proxy_protocol_vni": false,
          "proxy_protocol_vpce_id": false,
          "proxy_protocol_addr": false,
          "body_bytes_sent": false,
          "api_name": true,
          "app_name": true,
          "provider_app_id": false,
          "provider_app_name": false,
          "custom_data_log01": false,
          "custom_data_log02": false,
          "custom_data_log03": false,
          "custom_data_log04": false,
          "custom_data_log05": false,
          "custom_data_log06": false,
          "custom_data_log07": false,
          "custom_data_log08": false,
          "custom_data_log09": false,
          "custom_data_log10": false,
          "response_source": false
        },
        "call_data": {
          "log_request_header": false,
          "log_request_query_string": false,
          "log_request_body": false,
          "log_response_header": false,
          "log_response_body": false,
          "request_header_filter": "",
          "request_query_string_filter": "",
          "response_header_filter": "",
          "custom_authorizer": {
            "frontend": [],
            "backend": []
          }
        }
      }
    }
  )
}

resource "g42cloud_apig_plugin" "breaker" {
  instance_id = g42cloud_apig_instance.test.id
  name        = "%[2]s_breaker"
  type        = "breaker"
  content     = jsonencode(
    {
      "breaker_condition": {
        "breaker_type": "timeout",
        "breaker_mode": "percentage",
        "unhealthy_condition": "",
        "unhealthy_threshold": 30,
        "min_call_threshold": 20,
        "unhealthy_percentage": 51,
        "time_window": 15,
        "open_breaker_time": 15
      },
      "downgrade_default": null,
      "downgrade_parameters": null,
      "downgrade_rules": null,
      "scope": "share"
    }
  )
}
`, common.TestBaseComputeResources(name), name)
}

func testAccPluginAssociate_basic_step1(name string) string {
	return fmt.Sprintf(`
%[1]s

resource "g42cloud_apig_plugin_associate" "cors_bind" {
  instance_id = g42cloud_apig_instance.test.id
  plugin_id   = local.plugin_ids[0]
  env_id      = g42cloud_apig_environment.test[0].id
  api_ids     = [g42cloud_apig_api.test.id]
}

resource "g42cloud_apig_plugin_associate" "http_resp_bind" {
  instance_id = g42cloud_apig_instance.test.id
  plugin_id   = local.plugin_ids[1]
  env_id      = g42cloud_apig_environment.test[0].id
  api_ids     = [g42cloud_apig_api.test.id]
}

resource "g42cloud_apig_plugin_associate" "rate_limit_bind" {
  instance_id = g42cloud_apig_instance.test.id
  plugin_id   = local.plugin_ids[2]
  env_id      = g42cloud_apig_environment.test[0].id
  api_ids     = [g42cloud_apig_api.test.id]
}

resource "g42cloud_apig_plugin_associate" "kafka_log_bind" {
  instance_id = g42cloud_apig_instance.test.id
  plugin_id   = local.plugin_ids[3]
  env_id      = g42cloud_apig_environment.test[0].id
  api_ids     = [g42cloud_apig_api.test.id]
}

resource "g42cloud_apig_plugin_associate" "breaker_bind" {
  instance_id = g42cloud_apig_instance.test.id
  plugin_id   = local.plugin_ids[4]
  env_id      = g42cloud_apig_environment.test[0].id
  api_ids     = [g42cloud_apig_api.test.id]
}
`, testAccPluginAssociate_base(name), name)
}

func testAccPluginAssociate_basic_step2(name string) string {
	return fmt.Sprintf(`
%[1]s

resource "g42cloud_apig_plugin_associate" "cors_bind" {
  instance_id = g42cloud_apig_instance.test.id
  plugin_id   = local.plugin_ids[0]
  env_id      = g42cloud_apig_environment.test[1].id
  api_ids     = [g42cloud_apig_api.test.id]
}

resource "g42cloud_apig_plugin_associate" "http_resp_bind" {
  instance_id = g42cloud_apig_instance.test.id
  plugin_id   = local.plugin_ids[1]
  env_id      = g42cloud_apig_environment.test[1].id
  api_ids     = [g42cloud_apig_api.test.id]
}

resource "g42cloud_apig_plugin_associate" "rate_limit_bind" {
  instance_id = g42cloud_apig_instance.test.id
  plugin_id   = local.plugin_ids[2]
  env_id      = g42cloud_apig_environment.test[1].id
  api_ids     = [g42cloud_apig_api.test.id]
}

resource "g42cloud_apig_plugin_associate" "kafka_log_bind" {
  instance_id = g42cloud_apig_instance.test.id
  plugin_id   = local.plugin_ids[3]
  env_id      = g42cloud_apig_environment.test[1].id
  api_ids     = [g42cloud_apig_api.test.id]
}

resource "g42cloud_apig_plugin_associate" "breaker_bind" {
  instance_id = g42cloud_apig_instance.test.id
  plugin_id   = local.plugin_ids[4]
  env_id      = g42cloud_apig_environment.test[1].id
  api_ids     = [g42cloud_apig_api.test.id]
}
`, testAccPluginAssociate_base(name), name)
}
