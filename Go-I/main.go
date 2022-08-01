package main

func main() {
	/*
		Dasar-dasar Go
	*/
	// Cetak Hello, Go
	println("Hello, Go")

	// Cetak angka 7
	println(7)

	// Cetak penjumlahan dari 9 dan 3
	println(9 + 3)

	// Cetak "9 + 3" sebagai sebuah string
	println("9 + 3")

	// Cetak hasil dari 12 dibagi dengan 3
	println(12 / 3)

	// Cetak hasil dari 3 dikali 6
	println(3 * 6)

	// Cetak sisa hasil pembagian dari 8 dibagi dengan 3
	println(8 % 3)

	// Gabungkan string-string "Hello, " dan "world", lalu cetak
	println("Hello, " + "world")

	// Gabungkan string-string "38" dan "19", lalu cetak
	println("38" + "19")

	// Cetak penjumlahan 38 dan 19
	println(38 + 19)

	/*
		Dasar-dasar Variabel
	*/
	// Deklarasikan variable `message` dan tetapkan nilai "Hello, world"
	var message string = "Hello, world"

	// Cetak nilai dari dari variable `message`
	println(message)

	// Perbarui variable `message` dengan "Hello, Go"
	message = "Hello, Go"

	println(message)

	// Deklarasikan variable `message` dan tetapkan "Hello, world" pada-nya
	message2 := "Hello, world"

	// Deklarasikan variable `number` dan tetapkan `100` pada-nya
	number := 100

	// Cetak nilai dari `message` dan `number`
	println(message2, number)

	s1 := "Hi, "
	s2 := "world"

	// Ubah nilai yang diberikan pada variable saat pendefinisian
	s1 = "Hello, "

	// Cetak nilai-nilai dari variable s1 dan s2
	println(s1, s2)

	n := 100
	// Kurangkan 10 dari variable `n` lalu tetapkan hasilnya kembali pada `n`
	n -= 10

	println(n)

	m := 10
	// Tambahkan 1 pada variable `m` lalu tetapkan hasilnya kembali pada `m`
	m++

	println(m)

	/*
		Menggunakan Pernyataan IF
	*/
	x1 := 123
	y1 := 5 * 6

	// Ketika x lebih besar atau sama dengan 100, cetak "x lebih besar atau sama dengan 100"
	if x1 >= 100 {
		println("x lebih besar atau sama dengan 100")
	}

	// Ketika y kurang dari 40, cetak "y kurang dari 40"
	if y1 < 40 {
		println("y kurang dari 40")
	}

	x2 := 7 * 10
	y2 := 5 * 6

	// Buatlah sebuah pernyataan `if` untuk ketika nilai x sama dengan 70
	if x2 == 70 {
		println("x sama dengan 70")
	}

	// Buatlah sebuah pernyataan `if` untuk ketika nilai y tidak sama dengan 40
	if y2 != 40 {
		println("y tidak sama dengan 40")
	}

	money1 := 100
	price1 := 200

	// Tambahkan sebuah pernyataan `else`
	if money1 >= price1 {
		println("Anda bisa membeli produk ini")
	} else {
		println("Anda tidak memiliki cukup uang")
	}

	money2 := 200
	price2 := 200

	// Tambahkan pernyataan `else if` dan `else` berikut ini
	if money2 > price2 {
		println("Anda bisa membeli item ini")
	} else if money2 == price2 {
		println("Anda bisa membeli item ini, namun uang tidak akan tersisa")
	} else {
		println("Anda tidak memiliki cukup uang")
	}

	x3 := 20
	// Ketika x lebih besar dari atau sama dengan 10 dan kurang dari atau sama dengan 30,
	// cetak pesan "x paling sedikit 10 dan paling banyak 30"
	if x3 >= 10 && x3 <= 30 {
		println("x3 paling sedikit 10 dan paling banyak 30")
	}

	y3 := 60
	// Ketika y kurang dari 10 atau lebih besar dari 30,
	// cetak pesan "y kurang dari 10 atau lebih besar dari 30"
	if y3 < 10 || y3 > 30 {
		println("y3 kurang dari 10 atau lebih besar dari 30")
	}

	z := 55
	// Ketika z tidak sama dengan 77,
	// cetak pesan "z tidak sama dengan 77"
	if z != 77 {
		println("z tidak sama dengan 77")
	}

	n = 3
	switch n {
	// Tambahkan sebuah `case` untuk mencetak "Belum beruntung" ketika `n` bernilai 0
	case 0:
		println("Belum beruntung")

	// Tambahkan sebuah `case` untuk mencetak "Sedikit beruntung" ketika `n` bernilai 1 atau 2
	case 1, 2:
		println("Sedikit beruntung")

	// Tambahkan sebuah `case` untuk mencetak "Beruntung" ketika `n` bernilai 3 atau 4
	case 3, 4:
		println("Beruntung")

	// Tambahkan sebuah `case` untuk mencetak "Sangat beruntung" ketika `n` bernilai 5
	case 5:
		println("Sangat beruntung")

	}
}
