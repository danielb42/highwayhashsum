# highwayhashsum

![Build](https://github.com/danielb42/highwayhashsum/workflows/Build/badge.svg)
![Tag](https://img.shields.io/github/v/tag/danielb42/highwayhashsum)
![Go Version](https://img.shields.io/github/go-mod/go-version/danielb42/highwayhashsum)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/danielb42/highwayhashsum)](https://pkg.go.dev/github.com/danielb42/highwayhashsum)
[![Go Report Card](https://goreportcard.com/badge/github.com/danielb42/highwayhashsum)](https://goreportcard.com/report/github.com/danielb42/highwayhashsum)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)

This is [minio/highwayhash](https://github.com/minio/highwayhash), a **blazin'** fast hash algorithm for large files. This implementation is built to behave like the more well-known hashing tools from coreutils (md5sum, sha512sum etc.).

```bash
$ ./highwayhashsum fileA largeFileB giantFileC
$ a46507b2b3b6237a8f1d15dfd91dd6a8edb6ca3b7af978d1ecbb63863005656b  fileA
$ f5ae398a066fe999805704b249c2ef7b79cd0a92f5ae608703ae92ab15634109  largeFileB
$ d799b2e6ff862794433e8da31d28ec150098abec6e6ffb90b5f71610697042e3  giantFileC
$
```

## Install

### Either download a precompiled binary ...

Available for Linux, Windows and MacOS: [Latest Release](https://github.com/danielb42/highwayhashsum/releases/latest)

### ... or use go get

`go get github.com/danielb42/highwayhashsum`

Note that highwayhash uses an arbitrary "key" seed for hash generation. This implementation uses `12345678901234567890123456789012` and will thus generate different hashes to implementations with other seeding keys.
