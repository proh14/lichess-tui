package security

import (
	"github.com/ProtonMail/gopenpgp/v3/crypto"
)

func EncryptToken(token string, password string) string {
	pgp := crypto.PGP()

	encHandle, _ := pgp.Encryption().Password([]byte(password)).New()
	pgpMessage, _ := encHandle.Encrypt([]byte(token))
	armored, _ := pgpMessage.ArmorBytes()

	return string(armored)
}

func DecryptToken(encryptedToken string, password string) (string, error) {
	pgp := crypto.PGP()

	armored := []byte(encryptedToken)
	passphrase := []byte(password)

	decHandle, err := pgp.Decryption().Password(passphrase).New()

	if err != nil {
		return "", err
	}
	decrypted, err := decHandle.Decrypt(armored, crypto.Armor)

	if err != nil {
		return "", err
	}

	myMessage := decrypted.Bytes()

	return string(myMessage), nil
}
