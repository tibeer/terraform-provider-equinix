data "equinix_ecx_l2_sellerprofiles" "aws" {
  organization_global_name = "AWS"
  metro_codes              = ["SV", "DC"]
  speed_bands              = ["1GB", "500MB"]
}
