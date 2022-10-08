package database

import (
	"github.com/IBM/cloudant-go-sdk/cloudantv1"
)

// OpenDBConnection func for opening database connection.
func OpenDBConnection() *cloudantv1.CloudantV1 {
	return CloudantConnection()
}
