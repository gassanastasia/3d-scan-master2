# 3D Scan Platform

## WebSocket Protocol

### User Authentification
```json
{
	"type": "auth",
	"payload": {
		"token": "JWT_TOKEN"
	}
}
```
### Printer Registration
```json
{
	"type": "register_printer",
	"payload": {
		"printer_id": "printer_001",
		"token": "DEVICE_TOKEN"
	}
}
```
### Printer Registration
```json
{
	"type": "telemetry",
	"payload": {
		"temp": 215.3,
		"progress": 42
	}
}
```