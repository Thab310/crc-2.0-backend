resource "aws_lambda_function" "lambda" {
  function_name = "${var.project_name}"
  handler       = var.handler 
  runtime       = var.runtime
  role          = var.lambda_role_arn
  filename      = var.filename
  architectures = var.architectures
  source_code_hash = var.source_code_hash
  memory_size   = 128

  environment {
    variables = {
      TABLE_NAME = var.dynamodb_name
      REGION = var.region
    }
  }

  //depends_on = [aws_iam_policy_attachment.lambda_dynamodb_policy]
}



resource "aws_lambda_function_url" "lambda" {
  function_name      = aws_lambda_function.lambda.function_name
  authorization_type = "NONE"

  cors {
    allow_credentials = true
    allow_origins     = var.allow_origin
    allow_methods     = var.allow_methods
    allow_headers     = ["date", "keep-alive"]
    expose_headers    = ["keep-alive", "date"]
    max_age           = 86400
  }
}