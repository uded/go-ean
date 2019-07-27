package ean

import (
	"fmt"
	"strconv"
)

func Valid(ean string) bool {
	return ValidEan8(ean) || ValidEan13(ean) || ValidGTIN14(ean) || ValidUpc(ean)
}

func ValidEan8(ean string) bool {
	return validCode(ean, 8)
}

func ValidEan13(ean string) bool {
	return validCode(ean, 13)
}

func ValidGTIN14(gtin string) bool {
	return validCode(gtin, 14)
}

func ValidUpc(upc string) bool {
	return validCode(upc, 12)
}

func validCode(ean string, size int) bool {
	if len(ean) != size {
		return false
	}
	checksum, err := checksum(ean, size)

	return err == nil && strconv.Itoa(checksum) == ean[size-1:size]
}

// Calculate checksum for the EAN code.
// The ean string might be of length 7 or 8. If it has length 8 then the last checksum digit is ignored.
func ChecksumEan8(ean string) (int, error) {
	return checksum(ean, 8)
}

func ChecksumEan13(ean string) (int, error) {
	return checksum(ean, 13)
}

func ChecksumGTIN14(gtin string) (int, error) {
	return checksum(gtin, 14)
}

func ChecksumUpc(upc string) (int, error) {
	return checksum(upc, 12)
}

func checksum(ean string, size int) (int, error) {
	if len(ean) != size && len(ean) != size-1 {
		return -1, fmt.Errorf("incorrect ean %v to compute a checksum", ean)
	}

	code := ean[:size-1]
	multiplyWhenEven := size%2 == 0
	sum := 0

	for i, v := range code {
		value, err := strconv.Atoi(string(v))

		if err != nil {
			return -1, fmt.Errorf("contains non-digit: %q", v)
		}

		if (i%2 == 0) == multiplyWhenEven {
			sum += 3 * value
		} else {
			sum += value
		}
	}

	return (10 - sum%10) % 10, nil
}
