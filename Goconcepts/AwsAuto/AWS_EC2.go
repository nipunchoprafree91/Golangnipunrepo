// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT-0
package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

const (
	aws_access_key_id     = "AKIAZSXYY55LWU7RIPMF"
	aws_secret_access_key = "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"
	region                = "us-east-1"
)

func GetEc2FromCreds() {
	awsec2creds := credentials.NewStaticCredentials(aws_access_key_id, aws_secret_access_key, "")

	sess := ec2.New(session.New(&aws.Config{Region: aws.String(region),
		Credentials: awsec2creds}))
	result, err := sess.DescribeInstances(nil)
	if err != nil {
		fmt.Errorf("Error getting the session for EC2 management.")
	}
	fmt.Println("Result is : %v", result)
	for _, r := range result.Reservations {
		fmt.Println("Reservation ID: " + *r.ReservationId)
		fmt.Println("Instance IDs:")
		for _, i := range r.Instances {
			fmt.Println("   " + *i.InstanceId)
		}
	}
}

func LoginWithCredsFile() (*ec2.DescribeInstancesOutput, error) {
	sess, err1 := session.NewSessionWithOptions(session.Options{
		Profile:           "default",
		SharedConfigFiles: []string{"c:/..Users/nipun.chopra/.aws/credentials", "c:/..Users/nipun.chopra/.aws/credentials"},
	})
	if err1 != nil {
		return nil, fmt.Errorf("Error getting the session for EC2 management %v", err1)
	}
	result := ec2.New(sess)
	runningInstances, err := result.DescribeInstances(nil)
	if err != nil {
		fmt.Errorf("Error getting the session for EC2 management %v", err)
	}

	return runningInstances, nil
}

func main() {
	result, err := LoginWithCredsFile()
	if err != nil {
		fmt.Errorf("Error getting the session for EC2 management.")
	}
	fmt.Printf("Result is  %v", result)
}
