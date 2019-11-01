output "bootstrap_instances" {
  value = google_compute_instance.bootstrap.*
}

output "ip_addresses" {
  value = google_compute_instance.bootstrap[0].network_interface.0.network_ip
}

