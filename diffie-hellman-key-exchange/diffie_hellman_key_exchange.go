package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Hardcoded values
	g := big.NewInt(5)                     // Generator
	n := big.NewInt(232395928811)          // Modulus (12-digit prime)
	clientPrivateKey := big.NewInt(123456) // Client private key
	serverPrivateKey := big.NewInt(789012) // Server private key

	fmt.Println("Generator :", g)
	fmt.Println("Prime N: ", n)

	clientPublicKey := generatePublicKey(g, n, clientPrivateKey)
	serverPublicKey := generatePublicKey(g, n, serverPrivateKey)

	fmt.Println("Client's Public Key (A):", clientPublicKey)
	fmt.Println("Server's Public Key (B):", serverPublicKey)

	// Step 4: Compute shared secret

	clientSharedSecret := compuateSharedSecret(serverPublicKey, clientPrivateKey, n)
	serverSharedSecret := compuateSharedSecret(clientPublicKey, serverPrivateKey, n)

	fmt.Println("Client's Shared Secret (S):", clientSharedSecret)
	fmt.Println("Server's Shared Secret (S):", serverSharedSecret)

	// Ensure both shared secrets are equal
	if clientSharedSecret.Cmp(serverSharedSecret) != 0 {
		fmt.Println("Shared secrets do not match! Something went wrong.")
		return
	}

	// Step 5: Securely exchange data

	originalValue := int64(22723) // Value to encrypt

	encryptedValue := encrypt(originalValue, clientSharedSecret, n)

	fmt.Println("Original Value:", originalValue)
	fmt.Println("Encrypted Value:", encryptedValue)

	// Decryption (server side)

	decryptedValue, err := decrypt(encryptedValue, serverSharedSecret, n)

	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Println("Decrypted Value:", decryptedValue)
	}

}

func generatePublicKey(generator, prime, privateKey *big.Int) *big.Int {
	return new(big.Int).Exp(generator, privateKey, prime) // A = G^a mod N
}

func compuateSharedSecret(yourPublicyKey, myPrivateKey, prime *big.Int) *big.Int {
	return new(big.Int).Exp(yourPublicyKey, myPrivateKey, prime)
}

func encrypt(original int64, clientSharedSecret, prime *big.Int) *big.Int {

	originalValue := big.NewInt(original)

	encryptedValue := new(big.Int).Mod(
		new(big.Int).Mul(originalValue, clientSharedSecret),
		prime,
	)

	return encryptedValue

}

func decrypt(encrypted, sharedSecret, prime *big.Int) (*big.Int, error) {

	// Modular inverse of shared secret (S^-1 mod N)
	sharedSecretInverse := new(big.Int).ModInverse(sharedSecret, prime)
	if sharedSecretInverse == nil {

		return &big.Int{}, fmt.Errorf("Modular inverse not found. Decryption failed.")
	}

	// Decrypted_Value = Encrypted_Value * S^-1 mod N
	decryptedValue := new(big.Int).Mod(
		new(big.Int).Mul(encrypted, sharedSecretInverse),
		prime,
	)

	return decryptedValue, nil

}

