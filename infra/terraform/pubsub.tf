resource "google_pubsub_topic" "terraform-topic"{
  name = "shs-web-messenger"
}

resource "google_pubsub_topic" "dead_letter" {
  name = "shs-web-dead-letter"
}

resource "google_pubsub_subscription" "terraform-subscription" {
  name = "shs-web-messenger"
  topic = google_pubsub_topic.terraform-topic.name

  message_retention_duration = "604800s"
  retain_acked_messages      = true
  ack_deadline_seconds = 10

  enable_message_ordering = false

  dead_letter_policy {
    dead_letter_topic = google_pubsub_topic.dead_letter.id
    max_delivery_attempts = 10
  }
}
