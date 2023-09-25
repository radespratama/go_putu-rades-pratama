<h2 align="center">My Daily LogBook ✨</h2>

```
# STUDI INDEPENDEN ALTERRA ACADEMY #

22 September 2023
```

### Middleware
Middlware adalah entitas yang terpasang pada proses server request/response, middleware memperbolehkan kita untuk implementasi berbagai fungsionalitas diantara http request yang datang ke server dan response keluar. Ada banyak third party middleware yang ada seperti negroni, echo, interpose, Alice, atau bisa juga membuat sendiri.

### Log, Auth Middleware
Log middleware berfungsi untuk mencatat apa saja http request yang dilakukan dan apa respon nya, biasanya digunakan untuk logging, footprint, dan analisa data. Auth Middleware berfungsi untuk memberikan user identification yaitu siapa saja yang bisa meminta request ke server dengan tujuan untuk menjaga data dan informasi sehingga tidak sembarang orang yang bisa melakukan request ke server.

### JWT Middleware
JSON Web Token (JWT) merupakan bagian dari auth middleware dimana JWT merupakan teknik untuk autentifikasi request dengan memberikan informasi berupa token pada request header. JWT Memiliki 3 bagian yaitu header, payload, dan verify signature.

```
Contoh token JWT:
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c

HEADER nya itu ini : eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9

PAYLOAD nya ini: eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ

SIGNATURE nya ini: SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c

```