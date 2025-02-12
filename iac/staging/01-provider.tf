terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 4.50"

    }
    archive = {
      source = "hashicorp/archive"
    }
    null = {
      source = "hashicorp/null"
    }
  }
  backend "s3" {
    bucket = "blt-stage-terraform"
    key    = "services/hexagonal-scaffolding/terraform.tfstate"
    region = "us-east-1"
  }
}

provider "aws" {
  region = var.region
}