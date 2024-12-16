# infrastructure/terraform/modules/github-iam-role/main.tf



data "aws_iam_policy_document" "github_actions_assume_role" {
  statement {
    actions = ["sts:AssumeRoleWithWebIdentity"]
    effect  = "Allow"

    principals {
      type        = "Federated"
      identifiers = [var.oidc_provider_arn]
    }

    condition {
      test     = "StringLike"
      variable = "token.actions.githubusercontent.com:sub"
      values   = ["repo:${var.github_org}/${var.github_repo}:*"]
    }
  }
}

resource "aws_iam_role" "github_actions" {
  name               = var.role_name
  assume_role_policy = data.aws_iam_policy_document.github_actions_assume_role.json
}

# 環境ごとの権限を設定するポリシーをアタッチ
# 例: ステージング環境用のポリシー
resource "aws_iam_role_policy_attachment" "dev_policy" {
  count      = var.environment == "dev" ? 1 : 0
  policy_arn = "arn:aws:iam::aws:policy/AdministratorAccess"  # 例として Lambda の完全アクセス権を付与
  role       = aws_iam_role.github_actions.name
}

# 例: 本番環境用のポリシー
resource "aws_iam_role_policy_attachment" "production_policy" {
  count      = var.environment == "production" ? 1 : 0
  policy_arn = "arn:aws:iam::aws:policy/AWSLambdaExecute"  # 例として Lambda の実行権限のみを付与
  role       = aws_iam_role.github_actions.name
}

