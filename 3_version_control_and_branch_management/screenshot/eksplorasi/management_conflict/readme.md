_Case untuk menghandle conflict_

- Terdapat 1 file dengan nama readme yang dimana didalamnya sudah terdapat tulisan, dan casenya saya sudah mempunyai 2 branch, yaitu `main` dan `new_lesson`

_BRANCH MAIN_

- Kemudian berpindah ke branch `main`, lalu saya melakukan perubahan di beberapa line disana, contoh: `- `git config --global`: Untuk mengkonfigurasi pengaturan git seperti username, email dan lainnya`.
- Lalu, dari perubahan tersebut saya lakukan commit dan push ke branch `main`.

_BRANCH NEW LESSON_

- Kemudian berpindah ke branch `new_lesson`. Saya lakukan perubahan di line yang sama dengan di branch `main`, hanya saya konten nya berbeda, contoh: `- `git config`: Untuk mengkonfigurasi pengaturan git seperti username, email dan lainnya`.
- Lalu, saya melakukan commit dan push ke branch `new_lesson`. Disini 2 branch sudah melakukan commit di file yang sama dan di line yang sama.
- Masih di branch `new_lesson`, saya ingin melakukan merge dengan branch `main` hanya saja conflict sudah muncul.

_Bagaimana cara menghandle nya?_

- Dikarenakan di 2 branch ini sama-sama melakukan perubahan di file dan line yang sama. Jadi kita sebagai developer perlu me-review dari 2 file ini, yang mana kah akan di use kode nya.

- Sebenarnya kita bisa langsung menggunakan perubahan di ke-2 branch ini, namun perlu disesuaikan juga dengan kondisinya. Karena di branch `new_lesson`, perubahannya hanya sekedar duplikasi sedangkan di branch `main` tidak, maka untuk fix conflict ini saya menggunakan perubahan dari branch `main`, lalu klik tombol complete merge.
