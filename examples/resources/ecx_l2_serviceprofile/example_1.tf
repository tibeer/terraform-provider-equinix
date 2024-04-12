resource "equinix_ecx_l2_serviceprofile" "private-profile" {
  name                               = "private-profile"
  description                        = "my private profile"
  connection_name_label              = "Connection"
  bandwidth_threshold_notifications  = ["John.Doe@example.com", "Marry.Doe@example.com"]
  profile_statuschange_notifications = ["John.Doe@example.com", "Marry.Doe@example.com"]
  vc_statuschange_notifications      = ["John.Doe@example.com", "Marry.Doe@example.com"]
  private                            = true
  private_user_emails                = ["John.Doe@example.com", "Marry.Doe@example.com"]
  features {
    allow_remote_connections = true
    test_profile = false
  }
  port {
    uuid       = "a867f685-422f-22f7-6de0-320a5c00abdd"
    metro_code = "NY"
  }
  port {
    uuid       = "a867f685-4231-2317-6de0-320a5c00abdd"
    metro_code = "NY"
  }
  speed_band {
    speed      = 1000
    speed_unit = "MB"
  }
  speed_band {
    speed      = 500
    speed_unit = "MB"
  }
  speed_band {
    speed      = 100
    speed_unit = "MB"
  }
}
