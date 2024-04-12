data "equinix_fabric_routing_protocol" "routing_protocol_data_name" {
  connection_uuid = "<uuid_of_connection_routing_protocol_is_applied_to>"
  uuid = "<uuid_of_routing_protocol>"
}
