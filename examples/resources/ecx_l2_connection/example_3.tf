data "equinix_ecx_l2_sellerprofile" "gcp-1" {
  name = "Google Cloud Partner Interconnect Zone 1"
}

resource "equinix_ecx_l2_connection" "router-to-gcp" {
  name                = "tf-gcp-pri"
  profile_uuid        = data.equinix_ecx_l2_sellerprofile.gcp-1.id
  device_uuid         = equinix_network_device.myrouter.id
  device_interface_id = 5
  speed               = 100
  speed_unit          = "MB"
  notifications       = ["john@equinix.com", "marry@equinix.com"]
  seller_metro_code   = "SV"
  seller_region       = "us-west1"
  authorization_key   = "4d335adc-00fd-4a41-c9f3-782ca31ab3f7/us-west1/1"
}
