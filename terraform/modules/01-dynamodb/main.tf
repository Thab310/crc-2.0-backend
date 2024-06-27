resource "aws_dynamodb_table" "db" {
  name         = "${var.project_name}"
  billing_mode = "PAY_PER_REQUEST"

  attribute {
    name = "ID"
    type = "S"
  }

  hash_key = "ID"
}

