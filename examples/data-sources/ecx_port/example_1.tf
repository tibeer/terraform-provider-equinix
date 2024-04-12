data "equinix_ecx_port" "tf-pri-dot1q" {
  name = "sit-001-CX-NY5-NL-Dot1q-BO-10G-PRI-JP-157"
}

output "id" {
  value = data.equinix_ecx_port.tf-pri-dot1q.id
}
