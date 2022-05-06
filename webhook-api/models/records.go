package models

type Record struct {
	RecordId  string                 `json:"id"`
	BucketId  string                 `json:"bucket_id"`
	Body      map[string]interface{} `json:"body"`
	Headers   map[string]string      `json:"headers"`
	CreatedAt string                 `json:"createdAt"`
}
