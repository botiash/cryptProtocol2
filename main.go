package main

import (
	"fmt"
	"crypto/rand"
	"crypto/sha256"
)

// Генерация случайного ключа
func generateKey() ([]byte, error) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}
	return key, nil
}

// Хэширование сообщения с использованием SHA-256
func hashMessage(message []byte) []byte {
	hash := sha256.Sum256(message)
	return hash[:]
}

func main() {
	// Генерация ключей для отправителя и получателя
	senderKey, err := generateKey()
	if err != nil {
		fmt.Println("Ошибка генерации ключа отправителя:", err)
		return
	}

	recipientKey, err := generateKey()
	if err != nil {
		fmt.Println("Ошибка генерации ключа получателя:", err)
		return
	}

	// Пример шифрования и расшифрования сообщения
	message := []byte("Секретное сообщение")

	// Шифрование
	encryptedMessage := xorEncrypt(message, senderKey)

	// Расшифрование
	decryptedMessage := xorDecrypt(encryptedMessage, recipientKey)

	// Проверка целостности сообщения с использованием хэша
	messageHash := hashMessage(message)
	decryptedHash := hashMessage(decryptedMessage)

	// Проверка соответствия хэшей
	if string(messageHash) == string(decryptedHash) {
		fmt.Println("Сообщение успешно передано и расшифровано.")
		fmt.Println("Исходное сообщение:", string(message))
		fmt.Println("Расшифрованное сообщение:", string(decryptedMessage))
	} else {
		fmt.Println("Ошибка передачи или расшифровки сообщения.")
	}
}

// Функция шифрования сообщения с использованием XOR
func xorEncrypt(message []byte, key []byte) []byte {
	encrypted := make([]byte, len(message))
	for i := 0; i < len(message); i++ {
		encrypted[i] = message[i] ^ key[i%len(key)]
	}
	return encrypted
}

// Функция расшифрования сообщения с использованием XOR
func xorDecrypt(message []byte, key []byte) []byte {
	return xorEncrypt(message, key)
}
