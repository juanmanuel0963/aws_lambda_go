This Repo contains code to create AWS lambda functions in Golang.

v1.0.0 - release-branch.v1
- lambda_func.
    Simple lambda function. 
    It responses code 200 if request is successful.
    Testing code sends credentials signed with AWS signer v4 to authenticate if IAM is required on the server side.

v2.0.0 - release-branch.v2
- lambda_func_body_params.
    It sends query params in the BODY and gets a response message.
    Testing code sends credentials signed with AWS signer v4 to authenticate if IAM is required on the server side.

v3.0.0 - release-branch.v3
- lambda_func_query_params.
    It sends query params in the QUERY STRING and gets a response message.
    Testing code sends credentials signed with AWS signer v4 to authenticate if IAM is required on the server side.