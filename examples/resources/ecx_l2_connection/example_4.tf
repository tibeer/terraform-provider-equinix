data "equinix_ecx_l2_sellerprofile" "gcp" {
  name = "Google Cloud Partner Interconnect Zone 1"
}

resource "equinix_ecx_l2_connection" "token-to-gcp" {
  name                = "tf-gcp-pri"
  profile_uuid        = data.equinix_ecx_l2_sellerprofile.gcp-1.id
  service_token       = "e9c22453-d3a7-4d5d-9112-d50173531392"
  speed               = 100
  speed_unit          = "MB"
  notifications       = ["john@equinix.com", "marry@equinix.com"]
  seller_metro_code   = "SV"
  seller_region       = "us-west1"
  authorization_key   = "4d335adc-00fd-4a41-c9f3-782ca31ab3f7/us-west1/1"
}
