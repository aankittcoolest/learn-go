// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX - License - Identifier: Apache - 2.0
package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/joho/godotenv"
)

// S3DeleteObjectAPI defines the interface for the DeleteObject function.
// We use this interface to test the function using a mocked service.
type S3DeleteObjectAPI interface {
	DeleteObject(ctx context.Context,
		params *s3.DeleteObjectInput,
		optFns ...func(*s3.Options)) (*s3.DeleteObjectOutput, error)
}

// DeleteItem deletes an object from an Amazon Simple Storage Service (Amazon S3) bucket
// Inputs:
//     c is the context of the method call, which includes the AWS Region
//     api is the interface that defines the method call
//     input defines the input arguments to the service call.
// Output:
//     If success, a DeleteObjectOutput object containing the result of the service call and nil
//     Otherwise, an error from the call to DeleteObject
func DeleteItem(c context.Context, api S3DeleteObjectAPI, input *s3.DeleteObjectInput) (*s3.DeleteObjectOutput, error) {
	return api.DeleteObject(c, input)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	bucket := flag.String("b", "", "The bucket from which the object is deleted")
	objectName := flag.String("o", "", "The object to delete")
	flag.Parse()

	if *bucket == "" || *objectName == "" {
		fmt.Println("You must supply the bucket (-b BUCKET), and object to delete (-o OBJECT")
		return
	}

	var er aws.EndpointResolverWithOptions = aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL:  os.Getenv("ENDPOINT_URL"),
			HostnameImmutable: true,
		}, nil
	})

	var loadConfig = config.WithEndpointResolverWithOptions(er)
	cfg, err := config.LoadDefaultConfig(context.TODO(),loadConfig)
	if err != nil {
		panic("configuration error, " + err.Error())
	}

	client := s3.NewFromConfig(cfg)

	input := &s3.DeleteObjectInput{
		Bucket: bucket,
		Key:    objectName,
	}

	_, err = DeleteItem(context.TODO(), client, input)
	if err != nil {
		fmt.Println("Got an error deleting item:")
		fmt.Println(err)
		return
	}

	fmt.Println("Deleted " + *objectName + " from " + *bucket)
}