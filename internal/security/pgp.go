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
