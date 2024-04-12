data "equinix_fabric_service_profiles" "service_profiles_data_name" {
  filter {
    property = "/name"
    operator = "="
    values   = ["<list_of_profiles_to_return>"]
  }
}
