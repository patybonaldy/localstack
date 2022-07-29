package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/aws/aws-sdk-go/service/kms/kmsiface"
	"os"
)

func main() {
	// comment this out if you want to work on a real AWS account
	err := os.Setenv("LOCALSTACK_ENDPOINT", "http://localhost:4566")
	if err != nil {
		panic(err)
	}

	// when running locally, we're going to have the LOCALSTACK_ENDPOINT environment variable set
	sess, err := createSession("us-east-1")
	if err != nil {
		panic(err)
	}

	// Create KMS service client
	svc := kms.New(sess)

	key := "test-key"
	value := "test-value"
	result, err := makeKey(svc, &key, &value)
	if err != nil {
		fmt.Println("Got error creating key:")
		fmt.Println(err)
		return
	}
	fmt.Println(result)

	text := "probando texto 123"
	keyId := result.KeyMetadata.Arn
	// Encrypt the data
	encryptResult, err := svc.Encrypt(&kms.EncryptInput{
		KeyId:     keyId,
		Plaintext: []byte(text),
	})

	if err != nil {
		fmt.Println("Got error encrypting data: ", err)
		os.Exit(1)
	}

	fmt.Println("Blob (base-64 byte array):")
	fmt.Println(encryptResult.CiphertextBlob)

	decryptResult, err := svc.Decrypt(&kms.DecryptInput{
		CiphertextBlob: encryptResult.CiphertextBlob,
		KeyId:          keyId,
	})
	if err != nil {
		fmt.Println("Got error encrypting data: ", err)
		os.Exit(1)
	}

	fmt.Println(string(decryptResult.Plaintext))
}

func makeKey(svc kmsiface.KMSAPI, key, value *string) (*kms.CreateKeyOutput, error) {
	result, err := svc.CreateKey(&kms.CreateKeyInput{
		Tags: []*kms.Tag{
			{
				TagKey:   key,
				TagValue: value,
			},
		},
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}

func createSession(region string) (*session.Session, error) {
	awsConfig := &aws.Config{
		Region: aws.String(region),
	}

	if localStackEndpoint := os.Getenv("LOCALSTACK_ENDPOINT"); localStackEndpoint != "" {
		os.Setenv("AWS_PROFILE", "localstack")
		awsConfig.S3ForcePathStyle = aws.Bool(true)
		awsConfig.Endpoint = aws.String(localStackEndpoint)
		awsConfig.Credentials = credentials.NewSharedCredentials("", "localstack")
	}

	return session.NewSessionWithOptions(session.Options{
		Config:  *awsConfig,
		Profile: "localstack",
	})
}
