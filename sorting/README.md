# Problem 1 : Sorting
### __Cara Menjalankan__
1. Run compiled code 
> ./main
2. Input array
> 2,3,4,7,1,2
3. Pilih metode sorting (1 untuk Accending, 2 untuk Deccending)
> 1
4. Visualisasi step insertion sort akan tertampil sesuai dengan metode sorting

### __Penjelasan__
1. Struktur code mengunakan package terpisah untuk memudahkan pencarian fungsi dan mensederhanakan fungsi main
2. Code menggunakan input dari console yang mana inputan dipisahkan dengan comma
3. Hasil inputan user diparsing berdasarkan comma dan dilakukan preprocessing untuk inputan yang menggunakan [] seperti test case
..3.1. Didalam fungsi parseInput, dilakukan convert type string ke integer dengan menggunakan package strconv
..3.2. Hasil dari convert akan dimasukkan ke array of integer untuk menjadi return value
4. Setelah data sudah menjadi array of integer, dilakukan proses insertionSort
..4.1. Pada fungsi insertionSort digunakan fungsi recursive yang mana akan memanggil dirinya sendiri dan pada setiap step, akan dilakukan proses visualisasi dengan memanggil fungsi visualize
5. Pada fungsi visualize, diperlukan nilai terbesar dari array untuk menentukan tinggi dari bar, sehingga diperlukan fungsi untuk mencari nilai terbesar pada fungsi maxValue
..5.1. Setelah mendapat nilai terbesar, maka dilakukan looping untuk mencetak bar sesuai dengan nilai dari array index
