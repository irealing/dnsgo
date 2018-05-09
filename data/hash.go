package data

func MurMurHash(data []byte, seed uint32) uint32 {
	const m uint32 = 0x5bd1e995
	dl := len(data)
	h := seed ^ uint32(dl)
	for i := 0; i+4 <= dl; i += 4 {
		arr := data[i : i+4]
		k := uint32(arr[3]) | uint32(arr[2])<<8 | uint32(arr[1])<<16 | uint32(arr[0])<<24
		k *= m
		k ^= k >> 24
		k *= m
		h *= m
		h ^= k
	}
	rm := (dl / 4) * 4
	ar := data[rm:]
	switch dl % 4 {
	case 3:
		h ^= uint32(ar[2] << 16)
		fallthrough
	case 2:
		h ^= uint32(ar[1] << 8)
		fallthrough
	case 1:
		h ^= uint32(ar[0])
		h *= m
	}
	h ^= h >> 13
	h *= m
	h ^= h >> 15
	return h
}
