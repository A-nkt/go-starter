# go-aws-lambda
AWS SAMを用いて、GoをLambdaにデプロイする。


## deployコマンド
### ``~/.aws/config``に下記の内容を記載
```
[profile PROFILE_NAME]
sso_start_url = https://wni.awsapps.com/start
sso_region = us-east-1
sso_account_id = <ACCOUNT_ID>
sso_role_name = <ROLE>
region = <REGION>
output = json
```
### 1. build
```
./build.sh
```
を実行し、

### 2. deploy
```
deploy.sh
```
を実行すればOK