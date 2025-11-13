# spam-complaint
# 📨 Automatic Spam Complaint System

## Overview
This project aims to automate the process of identifying and reporting spam emails.  
It parses email headers, extracts IP addresses, identifies responsible network operators,  
and generates professional complaint emails to help reduce spam at its source.

## 🧠 How It Works
1. The user pastes raw email headers.
2. The system extracts routing IP addresses.
3. WHOIS lookups are performed for each IP.
4. Abuse contacts are identified.
5. A complaint email template is generated.

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

## 🧑‍💻 How to Run
```bash
go run main.go