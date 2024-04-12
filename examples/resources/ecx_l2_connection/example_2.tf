data "equinix_ecx_l2_sellerprofile" "azure" {
  name = "Azure Express Route"
}

data "equinix_ecx_port" "sv-qinq-pri" {
  name = "CX-SV5-NL-Dot1q-BO-10G-PRI"
}

data "equinix_ecx_port" "sv-qinq-sec" {
  name = "CX-SV1-NL-Dot1q-BO-10G-SEC"
}

resource "equinix_ecx_l2_connection" "ports-2-azure" {
  name              = "tf-azure-pri"
  profile_uuid      = data.equinix_ecx_l2_sellerprofile.azure.id
  speed             = 50
  speed_unit        = "MB"
  notifications     = ["john@equinix.com", "marry@equinix.com"]
  port_uuid         = data.equinix_ecx_port.sv-qinq-pri.id
  vlan_stag         = 1482
  vlan_ctag         = 2512
  seller_metro_code = "SV"
  named_tag         = "PRIVATE"
  authorization_key = "c4dff8e8-b52f-4b34-b0d4-c4588f7338f3"
  secondary_connection {
    name      = "tf-azure-sec"
    port_uuid = data.equinix_ecx_port.sv-qinq-sec.id
    vlan_stag = 1904
    vlan_ctag = 1631
  }
}
