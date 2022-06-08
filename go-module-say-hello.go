package module_say_hello

/*
	1. buat nama package modul sebagai nama dependency ketika diakase pd import
	2. buat go mod
	   dgn perintah: go mod init disertai url github project atau nama modules
	3. lakukan perintah simpan project/module di repository (github)
	   - git init
	   - git add .
	   - git commit -m "deskripsi"
	   - git add origin master disertai url github untuk menyimpan ke github
	   - git push origin master
	2. buat tag versi untuk merelease module berdasarkan versi dgn cara:
	   - git tag v1.0.0
	   - git push origin v1.0.0

	3. buat file dgn function main untuk mencetak data dari function modules	
*/

func SayHello(nama string) string {
	return "hello 1 " + nama
}