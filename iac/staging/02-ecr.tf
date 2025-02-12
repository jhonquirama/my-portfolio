resource "aws_ecr_repository" "ecr_repo" {
  name                 = var.app_name
  image_tag_mutability = "MUTABLE"
}


