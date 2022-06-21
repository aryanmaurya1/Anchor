package database

import (
	"database/sql"

	"github.com/aryanmaurya66/webhook/webhook-api/models"
	"github.com/sirupsen/logrus"
)

type DatabaseClient struct {
	logger *logrus.Logger
	db     *sql.DB
}

func (db *DatabaseClient) CreateBucket(bucketName, bucketId string) (string, error) {
	return "", nil
}

func (db *DatabaseClient) GetBucket(bucketId string) (models.Bucket, error) {
	var result models.Bucket
	return result, nil
}

func (db *DatabaseClient) GetBuckets() ([]models.Bucket, error) {}
func (db *DatabaseClient) DeleteBucket(bucketId string) error   {}

func (db *DatabaseClient) CreateRecord(bucketId, recordId, data string) (string, error) {}
func (db *DatabaseClient) GetRecords(bucketId string) ([]models.Record, error)          {}
func (db *DatabaseClient) DeleteRecord(recordId string) error                           {}
