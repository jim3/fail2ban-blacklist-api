# Fail2Ban Blacklist API

A lightweight Go HTTP server that exposes fail2ban's banned IP addresses via a JSON API endpoint.

## Overview

A security monitoring API in Go to automate the export of real-time threat intelligence (banned IPs) on a VPS. This allows external systems to query and retrieve the blacklist dynamically.

## Architecture

- **Web Server**: Caddy (configured as reverse proxy)
- **Application**: Go HTTP server
- **Security**: fail2ban (SSH brute-force protection)
- **API Response**: JSON format

## Features

- Real-time retrieval of banned IPs from fail2ban
- RESTful JSON API endpoint
- Secure reverse proxy configuration via Caddy
- Efficient command execution to fetch live ban data

## API Endpoint

```bash
GET /blacklist
```

**Response:**
```json
{
  "blacklist": [
    "92.118.39.95",
    "64.225.77.207",
    "209.38.34.67",
    ...
  ]
}
```

## Usage

Query the blacklist endpoint:
```bash
curl https://yourdomain.com/blacklist
```

## Technical Implementation

The server executes `fail2ban-client get sshd banip` to retrieve currently banned IPs and marshals them into JSON format. The application runs on `localhost:8080` behind Caddy, which handles TLS termination and public-facing requests.

## Requirements

- Go 1.25+
- fail2ban (configured for SSH protection)
- Caddy web server
- Appropriate sudo permissions for fail2ban-client

## Security Note

The server requires sudo access to run fail2ban-client commands. Ensure proper sudoers configuration is in place for the application user.
