package lib

import (
	"crypto/aes"
)

var CMD_VALIDATE_ENCRYPTED_MAIN_TOKEN = [...]byte {0x02}

func encryptMain(key []byte, seed []byte) []byte {
	return encryptWithKeyAndToken(key, buildMainToken(seed));
}

func buildMainToken(seed []byte) []byte {
	prepend := [...]byte {92, 101, 44, 182, 81, 212, 239, 235, 137, 90, 188, 111};
	return append(prepend[:], seed...);
}

func encryptWithKeyAndToken(key []byte, mainToken []byte) []byte {
	return decryptAes128Ecb(mainToken, key);
}

func decryptAes128Ecb(data, key []byte) []byte {
	cipher, _ := aes.NewCipher([]byte(key))
	decrypted := make([]byte, len(data))
	size := 16

	for bs, be := 0, size; bs < len(data); bs, be = bs+size, be+size {
		cipher.Decrypt(decrypted[bs:be], data[bs:be])
	}

	return decrypted
}

func Cryptor(key []byte, authSeed []byte) []byte {
	encryptedToken := encryptMain(key, authSeed);
	return append(CMD_VALIDATE_ENCRYPTED_MAIN_TOKEN[:], encryptedToken...);
}
