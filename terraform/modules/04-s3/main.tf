resource "aws_s3_bucket" "s3" {
  bucket = "thabelo-${var.project_name}"

  tags = {
    Name = var.project_name
  }

  
}

# resource "aws_s3_bucket_acl" "acl" {
#   bucket = aws_s3_bucket.s3.id
#   acl    = "private"
# }

