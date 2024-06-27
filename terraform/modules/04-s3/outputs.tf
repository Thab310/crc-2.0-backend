output "s3_bucket" {
    value = aws_s3_bucket.s3.bucket
}

output "s3_bucket_arn" {
    value = aws_s3_bucket.s3.arn
}

output "s3_regional_domain" {
  value = aws_s3_bucket.s3.bucket_regional_domain_name
}