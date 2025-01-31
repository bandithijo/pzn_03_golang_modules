package main

/*
Pengenalan

Sebelum Ada Go Modules
- Saat kita membuat aplikasi, biasanya kita akan menggunakan library atau dependency dari project lain
- Sebelum ada Go Modules, management untuk dependency sangat sulit dilakukan di Go
- Biasanya kita akan meng-copy semua source code library atau dependency lain ke project kita, hal ini membuat project kita menjadi bengkak karena penuh dengan library orang lain
- Atau biasanya library orang lain akan kita save di GOPATH directory, namun hal ini akan sulit jika ternyata beberapa aplikasi butuh library yang sama dengan versi yang berbeda

Go Modules
- Go Modules mulai dikenalkan di Go versi 1.11 dan 1.12
- Dengan Go Modules, kita bisa membuat project dengan mudah dan dependency management yang sangat mudah
*/
