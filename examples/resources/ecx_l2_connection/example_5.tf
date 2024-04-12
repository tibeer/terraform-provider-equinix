data "equinix_ecx_port" "sv-qinq-pri" {
  name = "CX-SV5-NL-Dot1q-BO-10G-PRI"
}

resource "equinix_ecx_l2_connection" "port-to-token" {
  name                = "tf-port-token"
  zside_service_token = "e9c22453-d3a7-4d5d-9112-d50173531392"
  speed               = 200
  speed_unit          = "MB"
  notifications       = ["john@equinix.com", "marry@equinix.com"]
  seller_metro_code   = "FR"
  port_uuid           = data.equinix_ecx_port.sv-qinq-pri.id
  vlan_stag           = 1000
}
