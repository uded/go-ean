package main

import (
	"fmt"

	"github.com/victorsmirnov/go-ean/ean"
)

func main() {
	// UPC code has 12 digits, the last one is a checksum and it is ignored if given
	upcStr := "012345678905"
	if c, err := ean.ChecksumUpc(upcStr); err != nil {
		fmt.Printf("Checksum fail, error %+v\n", err)
	} else {
		fmt.Printf("UPC %s, checksum %d\n", upcStr, c)
		// Prints "UPC 012345678905, checksum 5"
	}

	// First 11 digits for the UPC
	upcStr = "01234567890"
	if c, err := ean.ChecksumUpc(upcStr); err != nil {
		fmt.Printf("Checksum fail, error %+v\n", err)
	} else {
		fmt.Printf("UPC %s, checksum %d\n", upcStr, c)
		// Prints "UPC 01234567890, checksum 5"
	}

	ean13Bad := "629104150021"
	valid := ean.ValidEan13(ean13Bad)
	fmt.Printf("Validate EAN13 %s, valid %t\n", ean13Bad, valid)
	// Prints "Validate EAN13 629104150021, valid false"

	// Calculate checksum and make sure resulting EAN code is valid
	eanPrefix := "629104150021"
	checkSum, _ := ean.ChecksumEan13(eanPrefix)
	ean13 := fmt.Sprintf("%s%d", eanPrefix, checkSum)
	valid = ean.ValidEan13(ean13)
	fmt.Printf("Validate EAN13 %s, valid %t\n", ean13, valid)
	// Prints "Validate EAN13 6291041500213, valid true"

	// Get the length from the string
	valid = ean.Valid("012345678905")
	fmt.Printf("General validation 629104150021, valid %t\n", valid)
	// Prints "General validation 629104150021, valid true"
}
