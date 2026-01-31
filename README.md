# spam-complaint
# 📨 Automatic Spam Complaint System

## Overview
A Go-based REST API that detects spam messages and automatically logs spam complaints.  
This project demonstrates Clean Architecture, dependency injection, configuration via YAML, persistence, and RESTful API design.

## 🧠 Features
- Detect spam messages using configurable keywords from `config/config.yaml`
- Persist spam complaints in JSON (`data/complaints.json`)
- Expose REST API to send messages and retrieve complaints
- Modular architecture:
  - Orchestrator
  - Detector
  - Complaint service and repository
  - Ingestor
  - Configuration
  - Logger
- Extensible and testable design

## 🚀 Future Work
- Add abuse.net integration
- Add customizable templates
- Add tracking for sent complaints
- Optional: Direct email sending via SMTP

---

## 🛠️ Tech Stack
- **Language:** Go (Golang)
- **Libraries:** net, net/mail, net/smtp, regexp, and third-party WHOIS library
- **Version Control:** Git + GitHub

---

## Getting Started

### Prerequisites

- Go 1.20+ installed
- PowerShell (Windows) or terminal (Linux / Git Bash)
- Git

---

### 1️⃣ Clone repository

```bash
git clone https://github.com/USERNAME/spam-complaint.git
cd spam-complaint
```
### Create data file
```bash
mkdir data
echo [] > data/complaints.json
```

### Run the API
```bash
go run cmd/main.go
```

### You should see
```bash
Automatic Spam Complaint System
API running on :8080
```


### API Endpoints
POST Complaints
Send a message to check for spam.
Powershell
```bash
Invoke-RestMethod `
  -Uri http://localhost:8080/complaints `
  -Method POST `
  -Headers @{ "Content-Type" = "application/json" } `
  -Body '{"message":"Win money now!!!"}'
```
Linux/ Git Bash/ Curl
```bash
curl -X POST http://localhost:8080/complaints \
  -H "Content-Type: application/json" \
  -d '{"message":"Win money now!!!"}'
```

GET Complaints
Retrieve all recorded spam complaints.

Powershell
```bash
Invoke-RestMethod `
  -Uri http://localhost:8080/complaints `
  -Method GET
```
Linux/ Git Bash/ Curl
```bash
curl http://localhost:8080/complaints
```


### Configuration
configs/config.yaml contains spam keywords:
```yaml
spam:
  keywords:
    - win
    - money
    - free
    - click
    - prize
```
