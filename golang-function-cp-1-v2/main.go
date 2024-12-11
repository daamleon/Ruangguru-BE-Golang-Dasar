package main

import (
	"fmt"
)

// DateFormat mengubah tanggal numerik menjadi format yang mudah dibaca
func DateFormat(day, month, year int) string {
	months := []string{
		"January", "February", "March", "April", "May", "June",
		"July", "August", "September", "October", "November", "December",
	}
	
	// Format tanggal dengan menambahkan nol di depan jika perlu
	return fmt.Sprintf("%02d-%s-%d", day, months[month-1], year)
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(DateFormat(1, 1, 2012))
}
