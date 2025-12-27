package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	fmt.Printf("Memproses string: %q\n", s)
	// karakterTerlihat menyimpan posisi terakhir setiap rune yang sudah dilewati.
	karakterTerlihat := make(map[rune]int)
	kiri := 0  // ujung kiri jendela aktif
	kanan := 0 // ujung kanan jendela aktif
	panjangMaksimal := 0
	for i := 0; i < len(s); i++ {
		if posisi, ada := karakterTerlihat[rune(s[i])]; ada {
			kiri = max(kiri, posisi+1)
		}
		kanan++
		karakterTerlihat[rune(s[i])] = i
		panjangMaksimal = max(kanan-kiri, panjangMaksimal)
	}
	return panjangMaksimal
}
