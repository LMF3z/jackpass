package database

import (
	"encoding/json"
	"github/lmf3z/jack-pass/pkg/encrypt"
	"os"
)

var DbVaultFileName = "sparrow-vault.db"

type DBVault struct {
	Entries map[string]string `json:"entries"`
}

func NewDbVault() *DBVault {
	return &DBVault{
		Entries: make(map[string]string),
	}
}

func (v *DBVault) ToBytes() ([]byte, error) {
	return json.Marshal(v)
}

func (v *DBVault) SaveVault(password string) error {
	plaintext, err := v.ToBytes()
	if err != nil {
		return err
	}

	encryptedData, err := encrypt.EncryptFile(plaintext, password)
	if err != nil {
		return err
	}

	return os.WriteFile(DbVaultFileName, encryptedData, 0600)
}
