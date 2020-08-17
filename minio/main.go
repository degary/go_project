package main

import (
	"github.com/minio/minio-go/v7"
	"log"
)

func main() {
	endpoint := "10.1.1.206:9001"
	accessKeyId := "minio"
	secretAccessKey := "minio123"
	useSSL := false

	minioClient, err := minio.New(endpoint, accessKeyId, secretAccessKey, useSSL)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%#v\n", minioClient)
}
