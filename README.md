This Repo contains Terraform code to create AWS lambda functions in Golang.

v1.0.0 - release-branch.v1
- lambda_func : Lambda function in Golang. It responses code 200 if request is successful.
                It can be configure to require IAM authentication in AWS.
                If IAM authentication is required, testing code sends credentials signed with AWS signer v4

v2.0.0 - release-branch.v2
- lambda_func_body_params : Lambda function in Golang. It sends query params in the BODY and gets a response message.
                It can be configure to require IAM authentication in AWS.
                If IAM authentication is required, testing code sends credentials signed with AWS signer v4

v3.0.0 - release-branch.v3
- lambda_func_query_params : Lambda function in Golang. It sends query params in the QUERY STRING and gets a response message.
                It can be configure to require IAM authentication in AWS.
                If IAM authentication is required, testing code sends credentials signed with AWS signer v4