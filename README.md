# Password Encryption and Decryption Utility

This Go (Golang) project is a simple utility for encrypting and decrypting passwords using the AES encryption algorithm. It provides a command-line interface to perform these operations and is designed for educational purposes and as a basic example of password encryption and decryption in Go.

## Features

- **Encryption**: Encrypt a password and save it to a file (`pwd.txt`).
- **Decryption**: Decrypt a previously encrypted password from `pwd.txt`.

## Getting Started

### Prerequisites

Before you begin, make sure you have Go installed on your system.

### Installation

1. Clone this repository:

   ```sh
   git clone https://github.com/zvdy/go-password-encryption.git
    ```
2. Navigate to the project directory:

   ```sh
   cd go-password-encryption
   ```
3. Build the project:

   ```sh
    go build
    ```

## Usage

### Encrypting a Password

To encrypt a password and save it to pwd.txt, run:

```sh
./main crypt <password>
```

### Decrypting a Password

To decrypt a password from pwd.txt, run:

```sh
./main decrypt <pwd.txt>
```

This will display the decrypted password on the console.

## Why?

This project was created to demonstrate basic password encryption and decryption using the AES encryption algorithm in Go. While it's not recommended for production use (as it uses a static key), it serves as an educational example to understand the fundamentals of encryption and decryption.

## Security Considerations


- __Static Key:__ In this example, a static encryption key is used for simplicity. In real-world applications, you should use a secure key management system to store and handle keys.
- __Secure Password Storage:__ In practice, passwords should not be stored in plaintext, even when encrypted. Instead, use secure password hashing algorithms like bcrypt for password storage.
- __Input Validation:__ This project does not perform extensive input validation and error handling. In a production system, you should handle errors and edge cases more robustly.

## License

This project is licensed under the MIT License. See [LICENSE](LICENSE) for more information.