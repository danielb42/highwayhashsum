# highwayhashsum

![Tag](https://img.shields.io/github/v/tag/danielb42/highwayhashsum)
![Go Version](https://img.shields.io/github/go-mod/go-version/danielb42/highwayhashsum)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/danielb42/highwayhashsum)](https://pkg.go.dev/github.com/danielb42/highwayhashsum)
[![Go Report Card](https://goreportcard.com/badge/github.com/danielb42/highwayhashsum)](https://goreportcard.com/report/github.com/danielb42/highwayhashsum)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)

This is [minio/highwayhash](https://github.com/minio/highwayhash), a blazin' fast hash algorithm for large files. This implementation is built to behave like the more well-known hashing tools from coreutils (md5sum, sha512sum etc.).

Note that highwayhash uses an arbitrary "key" seed for hash generation. This implementation uses `12345678901234567890123456789012` and will thus generate different hashes to implementations with other seeding keys.