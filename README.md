# Cryptography Implementations in Go

This repository contains implementations of various cryptographic algorithms and schemes in the Go programming language.

## Projects

### BigNumber

* **Purpose:** An arbitrary-precision arithmetic library. This is often a fundamental building block for many cryptographic algorithms.
* **Usage:**
    ```bash
    cd bignumber
    go run main.go
    ```

### Curve

* **Purpose:** Demonstrates Elliptic Curve Integrated Encryption Scheme (ECIES) using the `ecies/go/v2` library. ECIES is a hybrid encryption scheme providing data confidentiality.
* **Usage:**
    ```bash
    cd curve
    go run main.go
    ```

### ElGamal

* **Purpose:** An implementation of the ElGamal encryption scheme. ElGamal is an asymmetric key encryption algorithm for public-key cryptography. This project supports encryption, decryption, signing, and verification.
* **Usage:**
    ```bash
    cd elgamal
    go run main.go --help # To see available flags and options
    # Example: Encrypt a message
    go run main.go --encrypt --p 23 --alfa 5 --B 10 --m 9
    ```

### SHA-1

* **Purpose:** An implementation of the SHA-1 (Secure Hash Algorithm 1) hashing function. SHA-1 produces a 160-bit (20-byte) hash value.
* **Usage:**
    ```bash
    cd sha1
    go run main.go # Hashes a predefined string
    go run main.go test.txt # Hashes the content of test.txt (create this file or use another)
    ```

## How to Build

To build any of the projects, navigate to the project's directory and use the `go build` command:

```bash
cd <project-directory>
go build
```

For example, to build the `bignumber` project:

```bash
cd bignumber
go build
```

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue.
