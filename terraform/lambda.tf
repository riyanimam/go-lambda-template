resource "aws_lambda_function" "go_lambda" {
  function_name = "go_lambda_handler"
  filename      = "build/go_lambda.zip"
  handler       = "main"
  runtime       = "go1.x"
  role          = aws_iam_role.lambda_exec.arn

  source_code_hash = filebase64sha256("build/go_lambda.zip")

  environment {
    variables = {
      ENV = "production"
    }
  }

  tags = {
    Project     = "GoLambdaTemplate"
    Environment = "Production"
  }
}

resource "aws_cloudwatch_log_group" "go_lambda" {
  name              = "/aws/lambda/${aws_lambda_function.go_lambda.function_name}"
  retention_in_days = 14

  tags = {
    Project     = "GoLambdaTemplate"
    Environment = "Production"
    Owner       = "DevOps"
  }
}

resource "aws_iam_role" "lambda_exec" {
  name = "go_lambda_exec_role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [{
      Action = "sts:AssumeRole"
      Effect = "Allow"
      Principal = {
        Service = "lambda.amazonaws.com"
      }
    }]
  })

  tags = {
    Project     = "GoLambdaTemplate"
    Environment = "Production"
    Owner       = "DevOps"
  }
}

resource "aws_iam_role_policy_attachment" "lambda_logs" {
  role       = aws_iam_role.lambda_exec.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}
