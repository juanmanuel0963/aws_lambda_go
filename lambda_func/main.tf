#############################################################################
# VARIABLES
#############################################################################

variable "region" {
  type    = string
}
variable "access_key" {
  type    = string
}
variable "secret_key" {
  type    = string
}

variable "lambda_function_name" {
  default = "lambda_function_name"
}

#############################################################################
# PROVIDERS
#############################################################################

terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.16"
    }
  }
}

provider "aws" {
  access_key = var.access_key
  secret_key = var.secret_key
}

provider "aws" {
  alias   = "east"
  region  = var.region
  access_key = var.access_key
  secret_key = var.secret_key 
}

#############################################################################
# DATA SOURCES
#############################################################################

data "aws_availability_zones" "azs" {}

#############################################################################
# RESOURCES
#############################################################################  

resource "aws_lambda_function" "test_lambda" {
  
  function_name = var.lambda_function_name
  provider =  aws.east
  tags = {
    Environment = "dev"
  }

}

#############################################################################
# OUTPUTS
#############################################################################
