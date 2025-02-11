---
subcategory: "Cloud Container Engine (CCE)"
---

# g42cloud_cce_node

Add a node to a CCE cluster.

## Basic Usage

```hcl
data "g42cloud_availability_zones" "myaz" {}

resource "g42cloud_compute_keypair" "mykp" {
  name       = "mykp"
  public_key = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDAjpC1hwiOCCmKEWxJ4qzTTsJbKzndLo1BCz5PcwtUnflmU+gHJtWMZKpuEGVi29h0A/+ydKek1O18k10Ff+4tyFjiHDQAT9+OfgWf7+b1yK+qDip3X1C0UPMbwHlTfSGWLGZquwhvEFx9k3h/M+VtMvwR1lJ9LUyTAImnNjWG7TAIPmui30HvM2UiFEmqkr4ijq45MyX2+fLIePLRIFuu1p4whjHAQYufqyno3BS48icQb4p6iVEZPo4AE2o9oIyQvj2mx4dk5Y8CgSETOZTYDOR3rU2fZTRDRgPJDH9FWvQjF5tA0p3d9CoWWd2s6GKKbfoUIi8R/Db1BSPJwkqB jrp-hp-pc"
}

resource "g42cloud_cce_cluster" "mycluster" {
  name                   = "mycluster"
  cluster_type           = "VirtualMachine"
  flavor_id              = "cce.s1.small"
  vpc_id                 = g42cloud_vpc.myvpc.id
  subnet_id              = g42cloud_vpc_subnet.mysubnet.id
  container_network_type = "overlay_l2"
}

resource "g42cloud_cce_node" "node" {
  cluster_id        = g42cloud_cce_cluster.mycluster.id
  name              = "node"
  flavor_id         = "s3.large.2"
  availability_zone = data.g42cloud_availability_zones.myaz.names[0]
  key_pair          = g42cloud_compute_keypair.mykp.name

  root_volume {
    size       = 40
    volumetype = "SATA"
  }
  data_volumes {
    size       = 100
    volumetype = "SATA"
  }
}
```

## Node with Eip

```hcl
resource "g42cloud_cce_node" "mynode" {
  cluster_id        = g42cloud_cce_cluster.mycluster.id
  name              = "mynode"
  flavor_id         = "s3.large.2"
  availability_zone = data.g42cloud_availability_zones.myaz.names[0]
  key_pair          = g42cloud_compute_keypair.mykp.name

  root_volume {
    size       = 40
    volumetype = "SATA"
  }
  data_volumes {
    size       = 100
    volumetype = "SATA"
  }

  // Assign EIP
  iptype                = "5_bgp"
  bandwidth_charge_mode = "traffic"
  sharetype             = "PER"
  bandwidth_size        = 100
}
```

## Node with Existing Eip

```hcl
resource "g42cloud_vpc_eip" "myeip" {
  publicip {
    type = "5_bgp"
  }
  bandwidth {
    name        = "test"
    size        = 8
    share_type  = "PER"
    charge_mode = "traffic"
  }
}

resource "g42cloud_cce_node" "mynode" {
  cluster_id        = g42cloud_cce_cluster.mycluster.id
  name              = "mynode"
  flavor_id         = "s3.large.2"
  availability_zone = data.g42cloud_availability_zones.myaz.names[0]
  key_pair          = g42cloud_compute_keypair.mykp.name

  root_volume {
    size       = 40
    volumetype = "SATA"
  }
  data_volumes {
    size       = 100
    volumetype = "SATA"
  }

  // Assign existing EIP
  eip_id = g42cloud_vpc_eip.myeip.id
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) Specifies the region in which to create the CCE node resource.
  If omitted, the provider-level region will be used. Changing this creates a new CCE node resource.

* `cluster_id` - (Required, String, ForceNew) Specifies the ID of the cluster.
  Changing this parameter will create a new resource.

* `name` - (Optional, String) Specifies the the node name.

* `flavor_id` - (Required, String, ForceNew) Specifies the flavor ID. Changing this parameter will create a new
  resource.

* `availability_zone` - (Required, String, ForceNew) Specifies the name of the available partition (AZ). Changing this
  parameter will create a new resource.

* `os` - (Optional, String, ForceNew) Specifies the operating system of the node.
  Changing this parameter will create a new resource.
  + For VM nodes, clusters of v1.13 and later support *EulerOS 2.5* and *CentOS 7.6*.
  + For BMS nodes purchased in the yearly/monthly billing mode, only *EulerOS 2.3* is supported.

* `key_pair` - (Optional, String, ForceNew) Specifies the key pair name when logging in to select the key pair mode.
  This parameter and `password` are alternative. Changing this parameter will create a new resource.

* `password` - (Optional, String, ForceNew) Specifies the root password when logging in to select the password mode.
  This parameter can be plain or salted and is alternative to `key_pair`.
  Changing this parameter will create a new resource.

* `root_volume` - (Required, List, ForceNew) Specifies the configuration of the system disk.
  Changing this parameter will create a new resource.

  + `size` - (Required, Int, ForceNew) Specifies the disk size in GB.
    Changing this parameter will create a new resource.
  + `volumetype` - (Required, String, ForceNew) Specifies the disk type.
    Changing this parameter will create a new resource.
  + `extend_params` - (Optional, Map, ForceNew) Specifies the disk expansion parameters.
    Changing this parameter will create a new resource.
  + `kms_key_id` - (Optional, String, ForceNew) Specifies the KMS key ID. This is used to encrypt the volume.
  Changing this parameter will create a new resource.

* `data_volumes` - (Required, List, ForceNew) Specifies the configurations of the data disk.
  Changing this parameter will create a new resource.

  + `size` - (Required, Int, ForceNew) Specifies the disk size in GB.
    Changing this parameter will create a new resource.
  + `volumetype` - (Required, String, ForceNew) Specifies the disk type.
    Changing this parameter will create a new resource.
  + `extend_params` - (Optional, Map, ForceNew) Specifies the disk expansion parameters.
    Changing this parameter will create a new resource.
  + `kms_key_id` - (Optional, String, ForceNew) Specifies the ID of a KMS key. This is used to encrypt the volume.
    Changing this parameter will create a new resource.

* `subnet_id` - (Optional, String, ForceNew) Specifies the ID of the subnet to which the NIC belongs.
  Changing this parameter will create a new resource.

* `fixed_ip` - (Optional, String, ForceNew) Specifies the fixed IP of the NIC.
  Changing this parameter will create a new resource.

* `eip_id` - (Optional, String, ForceNew) Specifies the ID of the EIP.
  Changing this parameter will create a new resource.

-> **NOTE:** If the eip_id parameter is configured, you do not need to configure the bandwidth parameters:
`iptype`, `bandwidth_charge_mode`, `bandwidth_size` and `share_type`.

* `iptype` - (Optional, String, ForceNew) Specifies the elastic IP type.
  Changing this parameter will create a new resource.

* `bandwidth_charge_mode` - (Optional, String, ForceNew) Specifies the bandwidth billing type.
  Changing this parameter will create a new resource.

* `sharetype` - (Optional, String, ForceNew) Specifies the bandwidth sharing type.
  Changing this parameter will create a new resource.

* `bandwidth_size` - (Optional, Int, ForceNew) Specifies the bandwidth size.
  Changing this parameter will create a new resource.

* `max_pods` - (Optional, Int, ForceNew) Specifies the maximum number of instances a node is allowed to create.
  Changing this parameter will create a new resource.

* `ecs_group_id` - (Optional, String, ForceNew) Specifies the ECS group ID. If specified, the node will be created under
  the cloud server group. Changing this parameter will create a new resource.

* `preinstall` - (Optional, String, ForceNew) Specifies the script to be executed before installation.
  The input value can be a Base64 encoded string or not. Changing this parameter will create a new resource.

* `postinstall` - (Optional, String, ForceNew) Specifies the script to be executed after installation.
  The input value can be a Base64 encoded string or not. Changing this parameter will create a new resource.

* `runtime` - (Optional, String, ForceNew) Specifies the runtime of the CCE node. Valid values are *docker* and
  *containerd*. Changing this creates a new resource.

* `extend_param` - (Optional, Map, ForceNew) Specifies the extended parameter.
  Changing this parameter will create a new resource.
  The available keys are as follows:
  + **agency_name**: The agency name to provide temporary credentials for CCE node to access other cloud services.
  + **alpha.cce/NodeImageID**: The custom image ID used to create the BMS nodes.
  + **dockerBaseSize**: The available disk space of a single docker container on the node in device mapper mode.
  + **DockerLVMConfigOverride**: Specifies the data disk configurations of Docker.

  The following is an example default configuration:

```hcl
extend_param = {
  DockerLVMConfigOverride = "dockerThinpool=vgpaas/90%VG;kubernetesLV=vgpaas/10%VG;diskType=evs;lvType=linear"
}
```

* `labels` - (Optional, Map, ForceNew) Specifies the tags of a Kubernetes node, key/value pair format.
  Changing this parameter will create a new resource.

* `tags` - (Optional, Map) Specifies the tags of a VM node, key/value pair format.

* `taints` - (Optional, List, ForceNew) Specifies the taints configuration of the nodes to set anti-affinity.
  Changing this parameter will create a new resource. Each taint contains the following parameters:

  + `key` - (Required, String, ForceNew) A key must contain 1 to 63 characters starting with a letter or digit.
    Only letters, digits, hyphens (-), underscores (_), and periods (.) are allowed. A DNS subdomain name can be used
    as the prefix of a key. Changing this parameter will create a new resource.
  + `value` - (Required, String, ForceNew) A value must start with a letter or digit and can contain a maximum of 63
    characters, including letters, digits, hyphens (-), underscores (_), and periods (.). Changing this parameter will
    create a new resource.
  + `effect` - (Required, String, ForceNew) Available options are NoSchedule, PreferNoSchedule, and NoExecute.
    Changing this parameter will create a new resource.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The resource ID in UUID format.
* `server_id` - ID of the ECS instance associated with the node.
* `private_ip` - Private IP of the CCE node.
* `public_ip` - Public IP of the CCE node.

## Timeouts

This resource provides the following timeouts configuration options:

* `create` - Default is 20 minute.
* `delete` - Default is 20 minute.

## Import

CCE node can be imported using the cluster ID and node ID separated by a slash, e.g.:

```
$ terraform import g42cloud_cce_node.my_node 5c20fdad-7288-11eb-b817-0255ac10158b/e9287dff-7288-11eb-b817-0255ac10158b
```

Note that the imported state may not be identical to your resource definition, due to some attributes missing from the
API response, security or some other reason. The missing attributes include:
`password`, `fixed_ip`, `eip_id`, `preinstall`, `postinstall`, `iptype`, `bandwidth_charge_mode`, `bandwidth_size`,
`share_type`, `max_pods`, `extend_param`, `labels`, `taints` and arguments for pre-paid. It is generally recommended
running `terraform plan` after importing a node. You can then decide if changes should be applied to the node, or the
resource definition should be updated to align with the node. Also you can ignore changes as below.

```
resource "g42cloud_cce_node" "my_node" {
    ...

  lifecycle {
    ignore_changes = [
      extend_param, labels,
    ]
  }
}
```
