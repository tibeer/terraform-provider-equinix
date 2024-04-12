data "equinix_ecx_l2_sellerprofile" "aws" {
  name = "AWS Direct Connect"
}

output "id" {
  value = data.equinix_ecx_l2_sellerprofile.aws.id
}
