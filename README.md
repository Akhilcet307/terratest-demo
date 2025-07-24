# Terraform S3 + Terratest Example

This project shows how to:
- Deploy an S3 bucket using Terraform
- Verify bucket creation with Terratest in Go

## Usage
###
It uses AWS Secret Key and AWS Secret Password set inside git settings
### Deploy S3 Bucket
```bash
terraform init
terraform apply -auto-approve
```

### Run Test
```bash
export BUCKET_NAME=$(terraform output -raw bucket_name)
go test -v ./test
```

### Destroy Resources
```bash
terraform destroy -auto-approve
```