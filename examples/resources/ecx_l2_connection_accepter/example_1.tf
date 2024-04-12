resource "equinix_ecx_l2_connection_accepter" "accepter" {
  connection_id = equinix_ecx_l2_connection.awsConn.id
}
