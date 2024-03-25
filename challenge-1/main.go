package main

import (
	"fmt"
	"os"
)

type Biodata struct {
	Nama       string
	Alamat    string
	Pekerjaan string
	Alasan     string
}

var biodataFriends = []Biodata{
 	{
 		Nama: "Herman",
 		Alamat: "Banyuwangi",
 		Pekerjaan: "Pegawai",
 		Alasan: "Ingin meningkatkan skillset dan peluang karir di bidang teknologi",
 	},
 	{
 		Nama: "Bima",
 		Alamat: "Banyuwangi",
 		Pekerjaan: "Mahasiswa",
 		Alasan: "Memenuhi persyaratan mata kuliah",
 	},
 	{
 		Nama: "Ali",
 		Alamat: "Banyuwangi",
 		Pekerjaan: "Frontend Developer",
 		Alasan: "Ingin Mendalami konsep-konsep dasar Golang dan penerapannya dalam pengembangan aplikasi",
	},
}

func main() {
	args := os.Args
	number := args[1]
	ShowData(number)
}

func ShowData(number string) {
	var num int
	_, err := fmt.Sscanf(number, "%d", &num)

	if err != nil {
		fmt.Println("Harus integer")
		return
	}

	if num < 1 || num > len(biodataFriends) {
		fmt.Println("Tidak ada data")
		return
	}

	biodata := biodataFriends[num-1]
	fmt.Println("Nama:", biodata.Nama)
	fmt.Println("Alamat:", biodata.Alamat)
	fmt.Println("Pekerjaan:", biodata.Pekerjaan)
	fmt.Println("Alasan memilih kelas Golang:", biodata.Alasan)
}