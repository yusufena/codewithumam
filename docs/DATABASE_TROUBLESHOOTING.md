# Database Connection Troubleshooting

## Error

```
2026/02/08 17:46:07 Failed to initialize database:EOF
exit status 1
```

## Root Cause

EOF error saat `db.Ping()` menunjukkan masalah koneksi ke database. Kemungkinan penyebabnya:

1. **Supabase Pooler Issue** - Pooler connection mungkin tidak compatible dengan driver Go
2. **Network/Firewall** - Koneksi ke Supabase terblok
3. **Connection String** - Format atau parameter tidak sesuai

## Solutions Tried

✅ Added `sslmode=require`  
✅ URL encoded password (`!` → `%21`)  
❌ Still getting EOF error

## Recommended Solutions

### Option 1: Use Direct Database Connection (Bukan Pooler)

Supabase menyediakan 2 jenis connection:
- **Pooler** (port 6543) - Untuk connection pooling
- **Direct** (port 5432) - Direct ke PostgreSQL

Try using direct connection:

```env
# Ganti dari pooler (6543) ke direct (5432)
DB_CONN=postgresql://postgres.fdahxiqcopobqkoltzvl:Developer3128%21@aws-0-ap-southeast-1.pooler.supabase.com:5432/postgres?sslmode=require
```

### Option 2: Get Fresh Connection String dari Supabase

1. Login ke https://supabase.com
2. Go to your project
3. Settings → Database
4. Copy **Connection String** untuk Golang
5. Pastikan gunakan yang mode "Session" bukan "Transaction"

### Option 3: Use Local PostgreSQL

Untuk development, gunakan local PostgreSQL:

```bash
# Install PostgreSQL (jika belum)
brew install postgresql@14

# Start PostgreSQL
brew services start postgresql@14

# Create database
createdb kasir_db

# Update .env
DB_CONN=postgresql://localhost:5432/kasir_db?sslmode=disable
```

### Option 4: Use Docker PostgreSQL

```bash
# Run PostgreSQL container
docker run --name kasir-postgres \
  -e POSTGRES_PASSWORD=kasir123 \
  -e POSTGRES_DB=kasir_db \
  -p 5432:5432 \
  -d postgres:14-alpine

# Update .env
DB_CONN=postgresql://postgres:kasir123@localhost:5432/kasir_db?sslmode=disable
```

## Testing Connection

Use this test program:

```bash
go run test_db.go
```

## Next Steps

1. Try Option 1 (Direct connection) first
2. If still fails, check Supabase dashboard for connection string
3. If urgent, use local PostgreSQL or Docker
4. Run migration: `psql -d kasir_db -f database/migration_session3.sql`

## Contact Support

If none work, contact:
- Supabase Support - https://supabase.com/dashboard/support
- Check Supabase Status - https://status.supabase.com/
