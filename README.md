# Golang with LocalStack

## Installation and execution
This application will run with LocalStack by default.
If you want to disable it, comment lines 14-17 in main.go

start docker compose by running:
```shell
docker network create localstacknetwork
docker-compose up -d
```

then run the application:
```shell
go mod tidy
go run main.go
```

```shell
aws configure --profile localstack

AWS Access Key ID [None]: test
AWS Secret Access Key [None]: test
Default region name [None]: us-east-1
Default output format [None]: 
```


The bootstrap.sh file creates all the resource we want on localstack.
The interaction with LocalStack is seamless, only need to override the endpoint-url.
for example, when using AWSCLI:

create secret
```shell
aws --endpoint http://localhost:4566 secretsmanager create-secret --name test-encrypt --description "Some secret" --secret-string '{"user":"test","password":"test","decryptionkey":"123"}' --tags '[{"Key":"user", "Value":"adm"},{"Key":"password","Value":"adm"},{"Key":"decryptionkey","Value":"123"}]' --region us-east-1
```

result
```shell
{
"ARN": "arn:aws:secretsmanager:us-east-1:000000000000:secret:bd-test2-95bd9e",
"Name": "bd-test2",
"VersionId": "be4136c6-36d2-42cc-a24d-345ec167e7a1"
}
```



create secret manager
```shell
aws --endpoint-url=http://${LOCALSTACK_HOST}:4566 s3 ls s3://
aws --endpoint http://localhost:4566 secretsmanager \
 create-secret --name test-encrypt \
 --description "Segredo para efetuar um teste com LocalStack" \
 --secret-string "Bruce Wayne. Você descobriu o segredo do Batman" \
 --region us-east-1
```

result
```shell
{
    "ARN": "arn:aws:secretsmanager:us-east-1:000000000000:secret:test-encrypt-eebb29",
    "Name": "test-encrypt",
    "VersionId": "1be2edaa-79d7-4e35-98e4-ce133abd75b5"
}


aws --endpoint http://localhost:4566 secretsmanager \
 get-secret-value \
 --secret-id test-encrypt  --region us-east-1
```