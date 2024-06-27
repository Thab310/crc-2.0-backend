########################################
#lambda Role
########################################
resource "aws_iam_role" "lambda_role" {
  name = "${var.project_name}-role"

  # Terraform's "jsonencode" function converts a
  # Terraform expression result to valid JSON syntax.
  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Sid    = ""
        Principal = {
          Service = "lambda.amazonaws.com"
        }
      },
    ]
  })
}

//Grant access to Cloudwatch logs
resource "aws_iam_policy_attachment" "lambda_cloudwatch_policy" {
  name       = "${var.project_name}-lambda_cloudwatch_policy"
  roles      = [aws_iam_role.lambda_role.name]
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSlambdaBasicExecutionRole"
}

//Grant access to Dynamodb
data "aws_iam_policy_document" "dynamodb_policy" {
  statement {
    sid = "AllowDynamodbServicePrincipal"

    effect = "Allow"

    actions = [
      "dynamodb:GetItem",
      "dynamodb:PutItem"
    ]

    resources = [
      "${var.dynamodb_arn}"
    ]
  }
}

resource "aws_iam_policy" "dynamodb_policy" {
  name   = "${var.project_name}-lambda-dynamodb-policy"
  policy = data.aws_iam_policy_document.dynamodb_policy.json
}
resource "aws_iam_policy_attachment" "lambda_dynamodb_policy" {
  name       = "lambda_dynamodb_policy"
  roles      = [aws_iam_role.lambda_role.name]
  policy_arn = aws_iam_policy.dynamodb_policy.arn
}
