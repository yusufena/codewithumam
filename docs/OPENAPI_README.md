# Kasir API - OpenAPI Documentation

OpenAPI specification untuk Kasir API Bootcamp Session 3.

## File OpenAPI

📄 [openapi.yaml](file:///Users/maulanayusuf/golangProject/codewithUmam/openapi.yaml)

## Cara Menggunakan

### 1. View di Swagger Editor (Online)

1. Buka [Swagger Editor](https://editor.swagger.io/)
2. Copy isi file `openapi.yaml`
3. Paste ke Swagger Editor
4. Dokumentasi interaktif akan otomatis ter-generate

### 2. View di VS Code

Install extension **Swagger Viewer**:
1. Buka VS Code
2. Install extension: `Swagger Viewer` by Arjun G
3. Buka file `openapi.yaml`
4. Klik kanan → `Preview Swagger`

### 3. Serve dengan Swagger UI (Local)

```bash
# Install swagger-ui via npx
npx swagger-ui-watcher openapi.yaml
```

Atau menggunakan Docker:

```bash
docker run -p 8081:8080 -e SWAGGER_JSON=/openapi.yaml -v $(pwd):/usr/share/nginx/html swaggerapi/swagger-ui
```

Buka browser: http://localhost:8081

## Endpoints Summary

### Products (`/api/produk`)
- `GET /api/produk` - Get all products (with optional name filter)
- `POST /api/produk` - Create new product
- `GET /api/produk/{id}` - Get product by ID
- `PUT /api/produk/{id}` - Update product
- `DELETE /api/produk/{id}` - Delete product

### Categories (`/categories`)
- `GET /categories` - Get all categories
- `POST /categories` - Create new category
- `GET /categories/{id}` - Get category by ID
- `PUT /categories/{id}` - Update category
- `DELETE /categories/{id}` - Delete category

### Transactions (`/api/checkout`)
- `POST /api/checkout` - Create checkout transaction

### Reports (`/api/report`)
- `GET /api/report/hari-ini` - Get daily sales report
- `GET /api/report?start_date=xxx&end_date=xxx` - Get sales report by date range

### Health Check
- `GET /health` - API health check

## Testing dengan cURL

Semua contoh cURL sudah ada di [walkthrough.md](file:///Users/maulanayusuf/.gemini/antigravity/brain/7657fabb-df7d-4c19-a92a-31b3232454b2/walkthrough.md)

## Testing dengan Postman

1. Import `openapi.yaml` ke Postman
2. Postman otomatis akan generate collection dari OpenAPI spec
3. Tinggal test semua endpoint

Atau bisa juga:
1. File → Import → Pilih `openapi.yaml`
2. Collection akan ter-create otomatis dengan semua endpoints

## Schema Definitions

Semua schema sudah didefinisikan di `components/schemas`:
- `Product` - Model produk
- `ProductInput` - Input untuk create/update produk
- `Category` - Model kategori
- `CategoryInput` - Input untuk create/update kategori
- `Transaction` - Model transaksi
- `TransactionDetail` - Detail item transaksi
- `CheckoutRequest` - Request body untuk checkout
- `SalesReport` - Model laporan penjualan
- `BestSelling` - Model produk terlaris

## Next Steps

Untuk production, bisa menggunakan:
- **Swagger UI** - Interactive API documentation
- **ReDoc** - Beautiful API documentation
- **Stoplight** - API design & documentation platform
