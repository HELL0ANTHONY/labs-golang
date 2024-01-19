package repository

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	"github.com/HELL0ANTHONY/labs-golang/lambda/Shape/pkg/models"
)

type DynamoDBClient interface {
	GetItem(
		context.Context,
		*dynamodb.GetItemInput,
		...func(*dynamodb.Options),
	) (*dynamodb.GetItemOutput, error)
	PutItem(
		context.Context,
		*dynamodb.PutItemInput,
		...func(*dynamodb.Options),
	) (*dynamodb.PutItemOutput, error)
}

type DynamoDBService interface {
	FindClientBy(string) (models.Record, error)
}

type DynamoDBRepository struct {
	client DynamoDBClient
}

func NewDynamoDBRepository() DynamoDBService {
	c, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}
	return DynamoDBRepository{client: dynamodb.NewFromConfig(c)}
}

func (r DynamoDBRepository) FindClientBy(id string) (models.Record, error) {
	tn := "TABLE_NAME"
	table := os.Getenv(tn)
	log.Println("TABLE_NAME", table)
	if table == "" {
		return models.Record{}, errors.New(tn + " is missing")
	}

	input := &dynamodb.GetItemInput{
		TableName: aws.String(table),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{
				Value: id,
			},
		},
	}

	out, err := r.client.GetItem(context.TODO(), input)
	if err != nil {
		return models.Record{}, err
	}

	if out.Item == nil {
		return models.Record{}, errors.New("client not found")
	}

	var o models.Record
	err = attributevalue.UnmarshalMap(out.Item, &o)

	return o, err
}
