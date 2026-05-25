package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Definition der CLI-Parameter
	bucketFlag := flag.String("bucket", "", "Name des AWS S3 Buckets (Erforderlich)")
	keyFlag := flag.String("key", "", "Ziel-Dateiname im S3 Bucket (Erforderlich)")
	fileFlag := flag.String("file", "", "Pfad zur lokalen Datei (Erforderlich)")

	flag.Parse()

	// Einfache Prüfung der CLI-Eingaben
	if *bucketFlag == "" || *keyFlag == "" || *fileFlag == "" {
		fmt.Println("Fehler: Fehlende Pflichtargumente.")
		flag.Usage()
		os.Exit(1)
	}

	// 1. Initialisierung der Datenschicht
	dataTier, err := NewDataTier()
	if err != nil {
		fmt.Printf("[Data Tier Fehler] %v\n", err)
		os.Exit(1)
	}

	// 2. Initialisierung der Logikschicht
	logicTier := NewLogicTier(dataTier)

	// 3. Ausführung des Prozesses
	fmt.Printf("Starte Validierung und Upload für '%s'...\n", *fileFlag)
	err = logicTier.ValidateAndUpload(*bucketFlag, *keyFlag, *fileFlag)
	if err != nil {
		fmt.Printf("[Prozess-Fehler] %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Erfolg: Datei wurde erfolgreich nach AWS S3 hochgeladen!")
}
