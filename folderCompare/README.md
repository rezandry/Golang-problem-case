# Problem 3 : Compare
## Problem 3.1 : Compare NEW DELETED
### __Cara Menjalankan__
1. Masuk ke folder task1
> cd task1/
2. Run compiled code 
> ./compare /sourcedir /targetdir
3. Hasil compare dari directory akan tampil beserta path dan status

### __Penjelasan__
1. Code akan dijalankan dengan membaca arguments dari running compiled code
2. Setelah mendapat directory yang akan dibandingkan, program akan membaca keseluruhan data dari 2 directory tersebut dan hasilnya akan disimpan pada array
3. Dari data di 2 directory tersebut, akan dicari panjangnya data untuk looping pembandingan dari list data source dengan list data taregt
4. Ketika proses pembandingan, file yang sama akan diset ke "" untuk membedakan file yang berbeda dari 2 directory tersebut dengan file yang benar benar hanya ada di directory source atau target
5. Maka setelah itu, list dari sourceList yang valuenya bukan "" akan dicetak dan statusnya NEW
6. Sedangkan list dari targetList yang valuenya bukan "" akan dicetak dan statusnya DELETED
7. Dalam compare 2 list ini, dibandingkan satu persatu dengan menggunakan package reflect dengan fungsi DeepEqual yang mana lebih cepat daripada pembandingan dengan menggunakan ==

## Problem 3.2 : Compare NEW DELETED EDITED
### __Cara Menjalankan__
1. Masuk ke folder task2
> cd task2/
2. Run compiled code 
> ./compare /sourcedir /targetdir
3. Hasil compare dari directory akan tampil beserta path dan status

### __Penjelasan__
1. Code akan dijalankan dengan membaca arguments dari running compiled code
2. Setelah mendapat directory yang akan dibandingkan, program akan membaca keseluruhan data dari 2 directory tersebut dan hasilnya akan disimpan pada array
3. Dari data di 2 directory tersebut, akan dicari panjangnya data untuk looping pembandingan dari list data source dengan list data taregt
4. Dalam perulangan dari setiap file, apabila namafilenya sama, maka akan dicek kembali dengan membandingkan dari 2 file tersebut dengan package reflect fungsi DeepEqual
5. Apabila 2 file tersebut sama, maka list dengan index tesebut akan diset "", selain itu, maka value pada index tersebut akan ditambah dengan string MODIFIED
6. Setelah itu dari 2 list, yakni sourceList dan targetList akan dicek apabila di sourceList tidak mengandung string MODIFIED maka akan dicetak dengan ditambah string NEW
7. Begitu juga dengan list targetList, apabila tidak mengandung string MODIFIED, maka akan dicetak dengan ditambah string DELETED
8. Selain keduanya, apabila mengandung string MODIFIED, maka akan dicetak
