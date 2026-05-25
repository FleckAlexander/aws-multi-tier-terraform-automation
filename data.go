package main

import (
	"context"
	"fmt"
	"os"

	"://github.com"
	"://github.com"
)

// DataTier kapselt den AWS S3 Client
type DataTier struct {
	s3Client *s3.Client
}

// NewDataTier initialisiert den S3-Client über Umgebungsvariablen
func NewDataTier() (*DataTier, error) {
	// Lädt AWS_ACCESS_KEY_ID und AWS_SECRET_ACCESS_KEY automatisch aus der Umgebung
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, fmt.Errorf("AWS Konfiguration konnte nicht geladen werden: %w", err)
	}

	return &DataTier{
		s3Client: s3.NewFromConfig(cfg),
	}, nil
}

// UploadFile überträgt die Datei in den angegebenen S3 Bucket
func (d *DataTier) UploadFile(bucketName, objectKey, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("Datei konnte nicht geöffnet werden: %w", err)
	}
	defer file.Close()

	_, err = d.s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: &bucketName,
		Key:    &objectKey,
		Body:   file,
	})
	if err != nil {
		return fmt.Errorf("S3 Upload fehlgeschlagen: %w", err)
	}

	return nil
}

