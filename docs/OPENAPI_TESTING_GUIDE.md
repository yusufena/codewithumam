# OpenAPI Testing Guide

## 🚀 Cara Paling Mudah: Swagger Editor Online

### Step 1: Copy OpenAPI file
```bash
cat openapi.yaml | pbcopy
```

### Step 2: Buka Swagger Editor
Buka browser ke: **https://editor.swagger.io/**

### Step 3: Paste OpenAPI spec
1. Di Swagger Editor, delete semua content yang ada
2. Paste `openapi.yaml` yang sudah di-copy
3. Documentation akan otomatis ter-generate!

### Step 4: Testing API
1. Di sisi kanan, klik endpoint yang ingin di-test
2. Klik **"Try it out"**
3. Isi parameter (jika ada)
4. Pastikan server sudah running: `go run main.go`
5. Klik **"Execute"**

> ⚠️ **PENTING**: Swagger Editor default menggunakan `http://localhost:8080`. Pastikan server sudah running!

---

## 🔧 Cara 2: Import ke Postman

### Step 1: Buka Postman
Download dari: https://www.postman.com/downloads/

### Step 2: Import OpenAPI file
1. Click **File** → **Import**
2. Pilih **`openapi.yaml`** dari project folder
3. Postman akan auto-generate collection dengan semua endpoints!

### Step 3: Testing
1. Set environment variable `baseUrl` = `http://localhost:8080`
2. Pilih endpoint dari collection
3. Click **Send**

---

## 💻 Cara 3: VS Code Extension

### Step 1: Install Extension
1. Buka VS Code
2. Install extension: **"Swagger Viewer"** by Arjun G
3. Install extension: **"OpenAPI (Swagger) Editor"** by 42Crunch

### Step 2: Preview
1. Buka file `openapi.yaml`
2. Right-click → **"Preview Swagger"**
3. Documentation akan muncul di preview panel

---

## 🐳 Cara 4: Local Swagger UI dengan Docker

```bash
# Run Swagger UI container
docker run -p 8081:8080 \
  -e SWAGGER_JSON=/openapi.yaml \
  -v $(pwd)/openapi.yaml:/openapi.yaml \
  swaggerapi/swagger-ui

# Buka browser: http://localhost:8081
```

---

## 🧪 Testing Endpoints

### Health Check
```bash
curl http://localhost:8080/health
```

### Get All Products
```bash
curl http://localhost:8080/api/produk
```

### Search Products
```bash
curl "http://localhost:8080/api/produk?name=indom"
```

### Create Transaction
```bash
curl -X POST http://localhost:8080/api/checkout \
  -H "Content-Type: application/json" \
  -d '{
    "items": [
      {"product_id": 1, "quantity": 2},
      {"product_id": 2, "quantity": 1}
    ]
  }'
```

### Daily Sales Report
```bash
curl http://localhost:8080/api/report/hari-ini
```

### Date Range Report
```bash
curl "http://localhost:8080/api/report?start_date=2026-02-01&end_date=2026-02-28"
```

---

## 📝 Quick Reference

| Tool | URL | Purpose |
|------|-----|---------|
| Swagger Editor | https://editor.swagger.io/ | Online testing & docs |
| Postman | https://postman.com | API testing & collection |
| ReDoc | https://redocly.github.io/redoc/ | Beautiful docs |
| Server | http://localhost:8080 | Your API server |

---

## ✅ Checklist

- [ ] Server running: `go run main.go`
- [ ] Database migrated: `go run migrate.go`
- [ ] OpenAPI file copied
- [ ] Swagger Editor opened
- [ ] Testing endpoints

**Status API Server**: Check dengan `curl http://localhost:8080/health`

Expected response:
```json
{"status": "OK", "message": "API Running"}
```

Selamat testing! 🎉
