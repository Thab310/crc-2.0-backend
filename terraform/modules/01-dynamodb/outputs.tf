output "db_arn" {
    value = aws_dynamodb_table.db.arn
}

output "db_name" {
    value = aws_dynamodb_table.db.name
}