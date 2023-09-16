- **Berapa banyak kekurangan dalam penulisan kode tersebut?** 

- **Bagian mana saja terjadi kekurangan tersebut?**

- **Tuliskan alasan dari setiap kekurangan tersebut. Alasan bisa diberikan dalam bentuk komentar pada kode yang disediakan berikut.**

#### Jawaban

1. Tipe data dari username dan password salah. Username dan password seharusnya bertipe string, bukan int.

```go
package main

type user struct {
	id       int
	username string
	password string
}

type userservice struct {
	t []user
}

func (u userservice) getallusers() []user {
	return u.t
}

func (u userservice) getuserbyid(id int) user {
	for _, r := range u.t {
		if id == r.id {
			return r
		}
	}

	return user{}
}
```

Selain itu, ada beberapa hal yang dapat di update dalam penulisan kodenya, yaitu:

- Gunakan konstanta untuk nilai ID kosong. Hal ini akan membuat kode lebih mudah dibaca dan dimaintain.
- Gunakan struktur data yang lebih sesuai untuk menyimpan data pengguna. Misalnya, jika pengguna memiliki banyak data, Anda dapat menggunakan map atau slice of maps.
- Gunakan error handling untuk menangani kesalahan yang mungkin terjadi. Misalnya, jika tidak ada pengguna dengan ID tertentu, Anda dapat mengembalikan error.


2. Kekurangan dalam penulisan kode tersebut terjadi pada bagian type user struct. Pada bagian ini tipe data dari username dan password adalah int, seharusnya string.

```go
// Kode yang sudah diperbaiki

type user struct {
	id       int
	username string
	password string
}
```

dan hasil akhirnya seperti ini

```go
const (
	EmptyUserID = 0
)

type user struct {
	id       int
	username string
	password string
}

type userservice struct {
	users map[int]user
}

func newUserService() *userservice {
	return &userservice{users: make(map[int]user)}
}

func (u *userservice) AddUser(user user) error {
	if user.id == EmptyUserID {
		return fmt.Errorf("user ID cannot be empty")
	}

	u.users[user.id] = user
	return nil
}

func (u *userservice) GetUserByID(id int) (user, error) {
	user, ok := u.users[id]
	if !ok {
		return user, fmt.Errorf("user with ID %d does not exist", id)
	}

	return user, nil
}

func main() {
	userService := newUserService()

	userOne := user{id: 1, username: "rey", password: "josaselole"}
	userTwo := user{id: 2, username: "udin", password: "momoyeye"}

	userService.AddUser(userOne)
	userService.AddUser(userTwo)

	user, err := userService.GetUserByID(1)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(user)
}
```

3. Berikut alasan dan perbaikannya

**Kekurangan 1**: Tipe data username dan password salah

Alasan:

Username dan password adalah data teks, sehingga tipe datanya harus string.
Jika tipe datanya int, maka username dan password dapat berisi angka, bukan hanya teks.
Komentar:

```go
// Username dan password adalah data teks,
// sehingga tipe datanya harus string.
type user struct {
	id       int
	username string // <-- Diubah menjadi string
	password string // <-- Diubah menjadi string
}
```

**Kekurangan 2**: Tidak ada error handling untuk menangani kesalahan

Alasan:

Jika tidak ada pengguna dengan ID tertentu, maka fungsi GetUserByID() akan mengembalikan nilai kosong. Nilai kosong ini dapat menyebabkan masalah jika tidak ditangani dengan benar.
Komentar:

```go
// Jika tidak ada pengguna dengan ID tertentu,
// maka fungsi `GetUserByID()` akan mengembalikan nilai kosong.
// Oleh karena itu, perlu dilakukan error handling.
func (u *userservice) GetUserByID(id int) (user, error) {
	user, ok := u.users[id]
	if !ok {
		return user, fmt.Errorf("user with ID %d does not exist", id)
	}

	return user, nil
}
```

**Kekurangan 3**: Tidak menggunakan struktur data yang lebih sesuai

Alasan:

Jika pengguna memiliki banyak data, maka menggunakan slice of users akan memakan banyak memori.Struktur data map lebih sesuai untuk menyimpan data pengguna dengan kunci unik, yaitu ID.

Komentar:

```go
// Jika pengguna memiliki banyak data,
// maka menggunakan slice of users akan memakan banyak memori.
// Oleh karena itu, lebih baik menggunakan struktur data map.
const (
	EmptyUserID = 0
)

type user struct {
	id       int
	username string
	password string
}

type userservice struct {
	users map[int]user
}

```
Dengan menambahkan komentar pada kode, akan membuat kode lebih mudah dibaca dan dipahami.