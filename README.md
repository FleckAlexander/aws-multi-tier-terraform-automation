# Multi-Tier Cloud File Uploader (Go & AWS)

Dieses Projekt demonstriert eine strukturierte, schichtweise verarbeitende Cloud-Anwendung (Multi-Tier-Ansatz). Ein performantes Kommandozeilen-Tool (CLI) in Go dient als Präsentationsschicht, welche lokale Dateien validiert (Logikschicht) und sicher in einen externen Cloud-Speicher (AWS S3 als Datenschicht) überträgt.

Das Projekt dient als praktischer Nachweis für meine Einarbeitung in die Backend-Entwicklung mit Golang sowie die Integration moderner, mehrschichtiger Cloud-Infrastrukturen.

## Architektur & Tiers (Schichten)
* **Presentation Tier:** CLI-Schnittstelle zur sicheren Entgegennahme von Benutzerparametern.
* **Logic Tier:** Go-Anwendungslogik zur Validierung von Dateiexistenz und -größe sowie Fehlerbehandlung.
* **Data Tier:** Persistente und skalierbare Datenspeicherung über das offizielle AWS SDK für Go im Cloud-Speicher.

## Features
* **Effizienter Upload:** Direkte Anbindung an den Cloud-Speicher über das aktuelle AWS SDK for Go v2.
* **Dateivalidierung:** Automatische Überprüfung von Dateieigenschaften vor dem Upload-Prozess.
* **Sicheres Rechtemanagement:** Trennung von Code und Konfiguration durch die Nutzung von Umgebungsvariablen (IAM-Zugriffsrechte).

## Technologien & Tools
* **Sprache:** Go (Golang)
* **Cloud-Infrastruktur:** Amazon Web Services (AWS S3)
* **SDKs:** AWS SDK for Go v2

## Voraussetzungen
Um das Tool lokal auszuführen, werden benötigt:
* Go (Version 1.22 oder neuer)
* Ein AWS-Konto mit einem aktiven S3-Bucket
* Eingerichtete AWS-Zugangsdaten (AWS_ACCESS_KEY_ID und AWS_SECRET_ACCESS_KEY)

## Installation & Start

1. Repository lokal klonen:
```bash
git clone https://github.com
cd DEIN-REPO-NAME
```

2. Abhängigkeiten installieren:
```bash
go mod tidy
```

3. Anwendung ausführen:
```bash
go run main.go -file="pfad/zur/datei.txt" -bucket="mein-cloud-bucket"
```

## Lernziele dieses Projekts
* Praktische Umsetzung einer sauberen Schichten-Einteilung (Multi-Tier) im Code.
* Vertiefung der Go-Syntax und des Fehlermanagements (Error Handling).
* Verständnis von Cloud-Infrastrukturen und sicheren IAM-Zugriffsrechten.

