# Todo CLI

Aplikasi todo list berbasis terminal yang ditulis dengan Go. Dibuat sebagai latihan setelah belajar Golang dasar.

## Fitur

- Melihat daftar tugas
- Menambah tugas beserta prioritas
- Menandai tugas selesai
- Menghapus tugas
- Data tersimpan otomatis ke file JSON, jadi tidak hilang saat program ditutup

## Prasyarat

- Go 1.21 atau lebih baru

## Cara Menjalankan

```bash
git clone <url-repository-ini>
cd todo-cli
go run main.go
```

## Contoh Tampilan

```
=== Todo List CLI ===
1. Lihat Tugas
2. Tambah Tugas
3. Tandai Selesai
4. Hapus Tugas
0. Keluar
Pilih menu: 1
[x] 1. Belajar Go (High)
[ ] 3. Latihan CLI (Basic)
```

`[x]` berarti tugas sudah selesai, `[ ]` berarti belum.

## Penyimpanan Data

Data disimpan di `todos.json` pada folder tempat program dijalankan. File ini dibuat otomatis saat ada perubahan pertama. ID tugas hanya naik dan tidak pernah dipakai ulang, bahkan setelah tugas dihapus.

## Konsep Go yang Dipakai

- Struct dan method (pointer receiver)
- Slice dan pointer
- Penanganan error
- `bufio` dan `os.Stdin` untuk input
- `switch`, `for`, dan `continue`
- `encoding/json` untuk menyimpan dan memuat data

## Struktur Proyek

```
todo-cli/
├── main.go
├── go.mod
├── .gitignore
└── README.md
```

## Rencana Pengembangan

- Memecah kode ke beberapa file
- Menambah unit test
- Mengedit nama atau prioritas tugas