data "equinix_fabric_ports" "ports_data_name" {
  filters {
    name = "<name_of_port||port_prefix>"
  }
}
