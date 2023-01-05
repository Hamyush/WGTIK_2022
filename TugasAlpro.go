package main

import "fmt"

const NMAX = 1000

type mahasiswa struct {
	nama           string
	kuis, uts, uas float64
	nilai_akhir    float64
}
type kelas [NMAX]mahasiswa

func main() {
	var data kelas
	var N, i int
	var nama string

	fmt.Scan(&data, &N)
	fmt.Scan(&nama)

	i = search_data(data, N, nama)
	hitung_nilai_akhir(&data, N)

	if i != -1 {
		fmt.Print("Nilai akhir dari", data[i].nama, "adalah", data[i].nilai_akhir)
	}
	sort_data(data, N)
	print_data(data, N)
}

func input_data(T *kelas, N int) { //proc
	/*
	   IS Data siswa telah siap pada piranti masukan
	   Proses, melakukan proses masukan dari user hingga nama yang diinputkan adalah "NONE" atau array telah penuh
	   FS Array T berisi sejumlah N siswa
	*/
	var nama string

	N = 0
	fmt.Scan(&nama)
	for nama != "NONE" && N < NMAX {
		T[N].nama = nama
		fmt.Scan(&T[N].kuis, &T[N].uts, &T[N].uas)
		N++
		fmt.Scan(&nama)
	}
}

func hitung_nilai_akhir(T *kelas, N int) { //proc
	/*
	   IS terdefinisi array t yang berisi N data mahasiswa
	   proses, nilai akhir = 20% kuis + 40% uts + 40% uas
	   FS field NilaiAkhir pada array t berisi nilai akhir yang dihitung dari nilai kuis, UTS dan UAS
	*/
	var i int

	i = 0
	for i < N {
		T[i].nilai_akhir = 0.2*T[i].kuis + 0.2*T[i].uts + 0.2*T[i].uas
		i++
	}
}

func search_data(T kelas, N int, nama string) int { //function
	/*
	   mengembalikan index data mahasiswa yang dicari berdasarkan inputan naman, atau -1 apalbila tidak ditemukan
	*/
	var i, ketemu int
	ketemu = -1
	i = 0

	for ketemu == -1 && i < N {
		if T[i].nama == nama {
			ketemu = i
		}
	}
	return ketemu
}

func sort_data(T kelas, N int) { //proc
	/*
	   IS terdefinis array t yang berisi N data mahasiswa
	   FS data array T terurut mengecil berdasarkan nilai akhir dengan algoritma insertion sort atau selection sort
	*/
	var pass, i int
	var temp mahasiswa

	pass = 1
	for pass > N-1 {
		i = pass
		temp = T[pass]
		for i > 0 && temp.nilai_akhir > T[i-1].nilai_akhir {
			T[i] = T[i-1]
			i++
		}
		T[i] = temp
	}
}

func print_data(T kelas, N int) { // proc
	/*
	   IS terdefinisi array t yang berisi N data mahasiswa
	   FS tercetak top-3 data siswa dengan nilai akhir tertinggi
	*/

	i := 0
	for i < 3 {
		fmt.Println(T[i].nama, "rangking", i+1, "dengan nilai akhir", T[i].nilai_akhir)
		i++
	}
}
