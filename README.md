[![Build Status](https://travis-ci.com/victorsmirnov/go-ean.svg?branch=master)](https://travis-ci.com/victorsmirnov/go-ean)

go-ean
======

go-ean is a simple utility library for calculating EAN checksum and validating EAN-8, EAN-13 and UPC numbers.
This project is a fork of the project https://github.com/nicholassm/go-ean adopted for the new [Go 1.11 Modules].

## Installation

Import the package in the application code. For example

    import "github.com/victorsmirnov/go-ean/ean"

When you build the project using `go buid` it will install the latest package version automatically
and update `go.mod` and `go.sum` files.

Example output from the `go buid` command:

    $ go build main/*
    go: finding github.com/victorsmirnov/go-ean/ean latest

If you do not use 1.11 Modules you can still install the package using `go get` command

    go get github.com/victorsmirnov/go-ean/ean

## Usage

To calculate a checksum use the `ChecksumEan8`, `ChecksumEan13` or `ChecksumUpc` functions:

    upcStr := "012345678905"
    if c, err := ean.ChecksumUpc(upcStr); err != nil {
        fmt.Printf("Checksum fail, error %+v\n", err)
    } else {
        fmt.Printf("UPC %s, checksum %d\n", upcStr, c)
        // Prints "UPC 012345678905, checksum 5"
    }

To check the validity of a string as EAN-8, EAN-13 or UPC use `Valid`:

    valid = ean.Valid("012345678905")
    fmt.Printf("General validation 629104150021, valid %t\n", valid)
    // Prints "General validation 629104150021, valid true"

See [examples](../blob/master/example/main.go) for complete working application.

## Changes from original library

In the original library `nicholassm/go-ean` the `checksum()` function expected the `ean` string argument
to include the last check digit. In my version the string may not have the last digit.

[Go 1.11 Modules]: https://github.com/golang/go/wiki/Modules

## License

(The MIT License)

Copyright (c) 2013 Nicholas Schultz-Møller
Copyright (c) 2019 Victor Smirnov

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
'Software'), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED 'AS IS', WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
