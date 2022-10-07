
# E-Commerce ReST API with Echo Framework
Project ini merupakan REST API yang berkaitan dengan proses transaksi yang terdapat pada sebuah E-Commerce, diantaranya listing produk, sistem cart, dan sistem pembayaran.

Dibangun menggunakan bahasa Go dengan framework Echo, memanfaatkan GORM untuk proses transaksi database PostgreSQL, serta menerapkan konsep Clean Architecture.

Project ini dibuat sebagai bahan penilaian untuk assesment test program eFishery Academy Aqua Developer Batch 2.
## How To Use
Project ini dapat dijalankan menggunakan Docker melalui langkah berikut ini:
- Clone repository project
```
git clone https://github.com/asyrofi-dev/aqua-developer-assesment.git
```
- Jalankan perintah Docker compose
```
docker compose up
```
- Akses API endpoint via `localhost` port `1323`, sebagai contoh:
```
http://localhost:1323/api/v1/products
```
- Akses Swagger Documentation
```
http://localhost:1323/swagger/
```