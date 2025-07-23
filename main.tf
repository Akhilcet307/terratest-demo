provider "aws" {
  region = var.aws_region
}

resource "aws_s3_bucket" "example" {
  bucket = var.bucket_name
  force_destroy = true
}

output "bucket_name" {
  value = aws_s3_bucket.example.bucket
}