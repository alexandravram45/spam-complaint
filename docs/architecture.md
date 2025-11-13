# 🧱 Architecture – Automatic Spam Complaint System

## 1. Overview

The **Automatic Spam Complaint System** is a command-line tool written in **Go** that helps users analyze spam email headers, identify the origin of spam messages, and generate automated complaint emails to the responsible network operators.

The project integrates:
- Email header parsing  
- Network intelligence gathering (WHOIS lookups)  
- Abuse contact discovery  
- Complaint email generation  

This document describes the high-level architecture, main components, and data flow within the system.

---

## 2. System Architecture

### 🗂️ High-Level Structure

spam-complaint/
├── cmd/ # Application entry point  
│ └── main.go  
├── internal/ # Internal reusable packages  
│ ├── parser/ # Email header parser  
│ ├── whois/ # WHOIS lookup logic  
│ ├── abusefinder/ # Finds abuse contacts from WHOIS data  
│ └── emailgen/ # Generates complaint emails from templates  
├── docs/ # Documentation (architecture, decisions, screenshots)  
├── tests/ # Unit tests for each module  
├── go.mod # Go module definition  
├── README.md # Project overview  
└── .gitignore  

## 3. Component Responsibilities

### 🧩 `cmd/main.go`
- Entry point of the application.  
- Handles user input (pasting raw email headers).  
- Orchestrates the workflow across internal packages.  
- Displays parsed results and generated complaint text.

---

### 🧩 `internal/parser`
- Extracts routing information (IP addresses) from email headers.  
- Filters lines starting with `Received:`.  
- Uses regular expressions to identify valid IPv4/IPv6 addresses.  
- Provides a clean list of unique IPs for further analysis.

---

### 🧩 `internal/whois`
- Performs WHOIS lookups for each IP obtained by the parser.  
- Extracts key data: network name, country, organization, and abuse contacts (if present).  
- Uses either the `net` package or external WHOIS APIs (like `whois.arin.net`).

---

### 🧩 `internal/abusefinder`
- If WHOIS data lacks abuse contacts, queries public sources such as:
  - `abuse.net`
  - `abuseipdb.com`
- Normalizes contact information (email addresses, URLs).  
- Returns structured results for each IP hop.

---

### 🧩 `internal/emailgen`
- Generates professional complaint emails using templates.  
- Includes:
  - Evidence section (headers, timestamps)
  - Clear description of the violation
  - Contact info of the reporter
- Supports customizable templates (plain text or Markdown).
