# highwayhashsum

This is [minio/highwayhash](https://github.com/minio/highwayhash), a fast hash algorithm for large files, built to behave just like other more well-known hashing binaries from coreutils (e.g. md5sum, sha512sum etc.).

Note that highwayhash uses an arbitrary "key" seed for hash generation. This implementation uses `12345678901234567890123456789012` and will thus generate different hashes to implementations with other seeding keys.