# Terraform

## What is this?
This is the backend configuration repository for the [cloud resume challenge 2.0](https://github.com/Thab310/cloud-resume-challenge-2.0). This repository houses a lambda function and the (iac) terraform script for the project. It runs it's own CI/CD pipeline through github actions. 

## Getting started
1. Clone the repo via SSH
```sh
#ssh
git clone git@github.com:Thab310/crc-2.0-backend.git
```    
2. CD into `terraform/env/dev` and create `terraform.tfvars`
```sh
cd terraform/env/dev && touch terraform.tfvars
```
3. Update `terraform.tfvars` if you're running this locally but if you're running this on github actions workflow update the environment [variables](https://docs.github.com/en/actions/learn-github-actions/variables) and [secrets](https://docs.github.com/en/actions/security-guides/using-secrets-in-github-actions) within github actions.

4. Run `terraform init` and `terraform apply` again only if you're running locally

5. Create github repo and give it a name like `crc-2.0-backend`

5. Upload to github to trigger github actions workflow
```sh
cd ../../..
git init
git add .
git commit -m "<commit message>"
git branch -M <branch>
git remote add origin git@github.com:<username>/crc-2.0-backend.git
git push -u origin main
```

>[!CAUTION]
When creating a lambda function using terraform we have to upload the lambda code or else the function will not get created.