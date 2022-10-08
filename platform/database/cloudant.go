package database

import (
	"github.com/IBM/cloudant-go-sdk/cloudantv1"
	"github.com/cesarvspr/api-go/app/models"
)

var client *cloudantv1.CloudantV1
var dbName = "books"
var ContentType = "application/json"

// CloudantConnection func for connection to PostgreSQL database.
func CloudantConnection() *cloudantv1.CloudantV1 {
	var err error
	// Check if database connection already exists.
	if client == nil {
		client, err = cloudantv1.NewCloudantV1UsingExternalConfig(&cloudantv1.CloudantV1Options{})
		if err != nil {
			panic(err)
		}
	}
	return client
}

func CreateBook(book *models.Book) (*cloudantv1.DocumentResult, error) {
	CloudantBook := cloudantv1.Document{}
	CloudantBook.SetProperties(book.Value())
	// Save the document in the database with "PostDocument" function
	document := CloudantConnection().NewPostDocumentOptions(dbName).SetDocument(&CloudantBook)
	rew, _, err := CloudantConnection().PostDocument(document)
	return rew, err
}

func GetAllBooks() ([]map[string]interface{}, error) {
	// Get all documents in the database with "GetAllDocuments" function
	res, _, err := CloudantConnection().PostAllDocs(&cloudantv1.PostAllDocsOptions{
		Db:          &dbName,
		IncludeDocs: &[]bool{true}[0],
	})

	count := int(*res.TotalRows)
	n := 0 // number of documents

	books := make([]map[string]interface{}, 0)
	for n < count {
		rawBookJson := res.Rows[n].Doc.GetProperties()

		books = append(books, rawBookJson)
		n++
	}

	return books, err
}
