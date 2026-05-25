package main

import (
	"errors"
	"fmt"
	"os"
)

// LogicTier verarbeitet die Geschäftslogik und Validierungen
type LogicTier struct {
	dataTier *DataTier
}

func NewLogicTier(data *DataTier) *LogicTier {
	return &LogicTier{dataTier: data}
}

// ValidateAndUpload prüft die Datei und startet den Upload-Prozess
func (l *LogicTier) ValidateAndUpload(bucket, key, filePath string) error {
	if bucket == "" || key == "" || filePath == "" {
		return errors.New("Fehlende Parameter: Bucket, Key und Dateipfad sind erforderlich")
	}

	// Dateivalidierung: Existenz und Größe prüfen
	fileInfo, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return fmt.Errorf("Validierungsfehler: Datei '%s' existiert nicht", filePath)
	}
	if err != nil {
		return fmt.Errorf("Fehler beim Lesen der Datei-Metadaten: %w", err)
	}

	if fileInfo.IsDir() {
		return errors.New("Validierungsfehler: Der Pfad verweist auf ein Verzeichnis, keine Datei")
	}

	if fileInfo.Size() == 0 {
		return errors.New("Validierungsfehler: Die Datei ist leer (0 Bytes)")
	}

	// Übergabe an die Datenschicht bei erfolgreicher Validierung
	return l.dataTier.UploadFile(bucket, key, filePath)
}
