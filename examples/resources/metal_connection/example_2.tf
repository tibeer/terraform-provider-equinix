resource "equinix_metal_vlan" "example" {
    project_id      = local.my_project_id
    metro           = "FR"
}

resource "equinix_metal_connection" "example" {
    name               = "tf-port-to-metal"
    project_id         = local.project_id
    type               = "shared"
    redundancy         = "primary"
    metro              = "FR"
    speed              = "200Mbps"
    service_token_type = "z_side"
    contact_email      = "username@example.com"
    vlans              = [
      equinix_metal_vlan.example.vxlan
    ]
}

data "equinix_ecx_port" "example" {
  name = "CX-FR5-NL-Dot1q-BO-1G-PRI"
}

resource "equinix_ecx_l2_connection" "example" {
  name                = "tf-port-to-metal"
  zside_service_token = equinix_metal_connection.example.service_tokens.0.id
  speed               = "200"
  speed_unit          = "MB"
  notifications       = ["example@equinix.com"]
  port_uuid           = data.equinix_ecx_port.example.id
  vlan_stag           = 1020
}
