package database

import "github.com/aryanmaurya66/webhook/webhook-api/models"

type DatabaseRepo interface {
	CreateBucket(bucketName, bucketId string) (string, error)
	GetBucket(bucketId string) (models.Bucket, error)
	GetBuckets() ([]models.Bucket, error)
	DeleteBucket(bucketId string) error

	CreateRecord(bucketId, recordId, data string) (string, error)
	GetRecords(bucketId string) ([]models.Record, error)
	DeleteRecord(recordId string) error
}
