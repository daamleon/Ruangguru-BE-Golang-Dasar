package main

import "fmt"

func GraduateStudent(score int, absent int) string {
        // Fungsi ini akan menentukan apakah seorang mahasiswa lulus atau tidak,
        // berdasarkan nilai dan jumlah absennya.

        // Jika nilai >= 70 dan absen < 5, maka mahasiswa dinyatakan lulus.
        // Jika tidak memenuhi kondisi di atas (nilai < 70 atau absen >= 5),
        // maka mahasiswa dinyatakan tidak lulus.

        var status string

        if score >= 70 && absent < 5 {
                status = "lulus"
        } else if score < 70 || absent >= 5 {
                status = "tidak lulus"
        }

        return status // Mengembalikan status kelulusan (lulus atau tidak lulus)
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(GraduateStudent(70, 4))
}
