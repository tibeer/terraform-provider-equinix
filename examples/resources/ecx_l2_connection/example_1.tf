data "equinix_ecx_l2_sellerprofile" "aws" {
  name = "AWS Direct Connect"
}

data "equinix_ecx_port" "sv-qinq-pri" {
  name = "CX-SV5-NL-Dot1q-BO-10G-PRI"
}

resource "equinix_ecx_l2_connection" "port-2-aws" {
  name              = "tf-aws"
  profile_uuid      = data.equinix_ecx_l2_sellerprofile.aws.id
  speed             = 200
  speed_unit        = "MB"
  notifications     = ["marry@equinix.com", "john@equinix.com"]
  port_uuid         = data.equinix_ecx_port.sv-qinq-pri.id
  vlan_stag         = 777
  vlan_ctag         = 1000
  seller_region     = "us-west-1"
  seller_metro_code = "SV"
  authorization_key = "345742915919"
}
