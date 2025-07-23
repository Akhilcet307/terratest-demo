package test

import (
	"testing"
	"os"

	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/stretchr/testify/assert"
)

func TestS3BucketExists(t *testing.T) {
	region := "us-east-1"
	bucket := os.Getenv("BUCKET_NAME")
	assert.NotEmpty(t, bucket, "BUCKET_NAME must be set")

	exists := aws.S3BucketExists(t, region, bucket)
	assert.True(t, exists, "Expected S3 bucket to exist")
}