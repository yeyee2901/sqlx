# Percobaan SQLX untuk Koneksi Database Menggunakan Bahasa Golang
## Technology Stacks
- [Gin](https://github.com/gin-gonic/gin) - for easy to configure web server framework
- [Gin-Swagger](https://github.com/swaggo/gin-swagger) - Swagger Docs integration for Gin web framework
- [Yaml v2](https://github.com/go-yaml/yaml/tree/v2.4.0) - Go package for processing yaml files
- [Sqlx](https://github.com/jmoiron/sqlx) - a package which provides a set of extensions on Go's builtin `database/sql` package

## Application Structure
```
app
├── config              // config
│   └── config.go       // logics to load the setting.yaml
├── controller          // routing, might change it to 'service', karena
│   └── controller.go   // package tidak menampung logic, hanya definisi routing & memanggil service lain
├── datasource          // koneksi ke datasource (database)
│   ├── datasource.go   // logics stored here
│   └── entity.go       // models generic untuk datasource, currently empty
├── entity              // generic models yang bisa digunakan dimana saja
│   └── entity.go
├── middlewares         // middlewares
│   └── middlewares.go  // untuk sementara hanya CORS supaya swagger jalan
└── user                // user management logic
    ├── entity.go       // models, isinya user model dan req & resp yg berhubungan dgn user
    └── user.go         // logics stored here
```

## How to try this project
1. Clone repository:
```bash
git clone https://github.com/yeyee2901/sqlx.git
```
2. Run go mod download & verify untuk melengkapi semua dependency
```bash
go mod download && go verify
```
3. Sesuaikan file `setting.yaml` terutama bagian mysql profile nya
```yaml
mysql:
  username: your_username
  password: your_password
  db: local_development
  host: 127.0.0.1
  port: 30000
  minpool: 1
  maxpool: 10
  parse_time: "true" # wajib true apabila ingin banding ke objek time.Time
```
4. Migrate database (script migrate ada di `./db/migration.sql`)
```sql
CREATE TABLE `users` (
    id INT auto_increment,
    name VARCHAR(255) NOT NULL,
    created_at DATETIME NOT NULL DEFAULT NOW(),

    PRIMARY KEY (`id`)
);
```
5. Jalankan bisa di compile dulu / langsung di go run
```bash
go run .    # Jalan langsung

go build -o compiled # di compile dulu
./compiled
```
6. Buka browser dan akses ke IP yang ada di `setting.yaml` untuk mengakses swagger supaya bisa coba-coba API nya
```
http://localhost:8767/swagger/index.html
```
7. Have fun trying!

You can see the source code, I put comments almost everywhere that might need explanation.
Welcome to the world of Go, sir.
