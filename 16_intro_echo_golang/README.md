<h2 align="center">My Daily LogBook ✨</h2>

```
# STUDI INDEPENDEN ALTERRA ACADEMY #

18 September 2023
```

### Echo Go lang
Third party / library adalah kumpulan dari sebuah code yang memiliki fungsi - fungsi tertentu dan dimana fungsi itu dapat dipanggil oleh program lain. Echo sendiri adalah sebuah library dimana echo merupakan web framework golang yang memiliki performa tinggi, extensible dalam arti bisa dipasang dengan library lain seperti GORM, dan minimalist dalam arti cukup sederhana tidak memerlukan driver database atau ORM namun cukup powerful. Keuntungan dari echo adalah optimized router, middleware, data rendering, scaleable, dan data binding. Echo juga tidak memiliki struktur dalam pengembangannya.

Fitur yang ada di echo antara lain:

- **Fitur middleware yang lengkap.** Echo menyediakan berbagai macam fitur middleware, seperti logging, authentication, authorization, dan lainnya.
- **Dapat melakukan optimize router sehingga menjadi lebih cepat.** Echo menggunakan algoritma routing yang efisien untuk meningkatkan kinerja.
- **Dapat melakukan databinding.** Echo menyediakan berbagai macam fitur databinding untuk memudahkan pengembang dalam memproses data dari request.
- **Mempunyai scalable yang bagus untuk traffic besar.** Echo menggunakan arsitektur yang scalable untuk menangani traffic besar.

## Basic Routing & Controller
Pada saat ingin menggunakan Echo Golang, pertama yang harus dilakukan adalah melakukan setup awal pada direktori project kita dengan melakukan inisialisasi awal dan mengambil library dari echo golang

```go
go mod init <direktori>
go get github.com/labstack/echo/v4
```

Kemudian contoh kode program merupakan basic routing dan controller untuk mencetak "Latihan Echo"
```go
package main

import (
	"net/http"
	
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Latihan Echo")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
```

### Data Rendering, Data Retrieve, dan Binding Data

Data rendering merupakan bagaimana API dapat mengirim berbagai HTTP Response dalam JSON, XML, HTML, File, Attachment, Inline, Stream atau Blob. Data Retrieve adalah bagaimana kita mendapat data dengan spesifikasi tertentu, dalam hal ini kita bisa menggunakan URL Params, Query Params, atau Form Value. Binding data adalah untuk HTTP Request Payload.