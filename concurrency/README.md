# Problem 4 : Concurrency
### __Cara Menjalankan__
1. Run compiled code 
> ./concurrency -concurrent_limit=2 -output=/home/rezandry/data/museum 
2. Maka data museum akan masuk pada directory output dengan worker goroutine dibatasi sesuai dengan concurrent_limit

### __Penjelasan__
1. Pertama perlu didefine model yang sesuai untuk menerima data dari API, yakni ResponseReg, ResponseMuseum, Region dan Museum, selain itu diperlukan varible global dirPath dan  Limit untuk menerima inputan berupa flag dari console
2. Selanjutnya, dilakukan parse dari flag dan define url API, setelah itu panggil fungsi saveCSV untuk memulai proses
3. Pada saveCSV akan dilakukan pengecekan apakah directory ada, apabila tidak ada, maka akan dibuat directory tersebut
4. Setelah itu, akan request body dari url yang sudah ada, prosesnya adalah mengambil data Provinsi, selanjutnya data Kota, selanjutnya baru data museum
5. Pada getbody menggunakan http request selain itu pada body dilakukan TrimPrefix untuk menghilangkan special character yang ada pada body responseAPI
6. Selanjutnya dilakukan parsing data json dengan menggunakan Unmarshal dan data hasil parsing disimpan pada struct ResponseReg
7. Setelah data provinsi sudah didapat, selanjutnya akan dilakukan pengambilan data Kota dengan memanggil fungsi getKota
8. Pada fungsi ini diimplementasikan concurrency
9. Pada fungsi ini, didefine channel queue dengan type data []Region
10. Lalu dilakukan looping untuk worker sesuai dengan Limit yang sudah ditentukan selanjutnya akan dilakukan pemanggilan go function untuk menjalankan concurrency
11. Dalam function tersebut, dilakukan pemanggilan fungsi untuk queue yang ada, selama queue terisi, maka akan dijalankan
12. Selanjutnya pada posisi setelah itu, akan dilakukan pencarian data kota, dimana data kota itu nanti yang akan dimasukan pada queue
13. go routine tersebut akan menjalankan fungsi getMuseum dimana di fungsi tersebut akan membuat file sesuai dengan nama kota/kabupaten
14. Setelah itu akan ada pengambilan body dan diparsing menjadi array dan dari setiap data tersebut dimasukkan pada struct Museum
15. Dari setiap data tersebut ditambahkan pada array record untuk keperluan memasukkan data ke dalam file csv
16. Setelah itu data akan masuk pada file terkait
