package contacts

import (
	"strings"
	"testing"
)

// MemoryRepo implements repository.Repo interface.
type MemoryRepo struct {
	contactsBuf strings.Builder
	mnemnoicBuf strings.Builder
}

func (mr *MemoryRepo) WriteContacts(contactJsonBytes []byte) error {
	mr.contactsBuf.Reset() // to simulate truncating a file during writes.
	_, err := mr.contactsBuf.Write(contactJsonBytes)
	return err
}

func (mr *MemoryRepo) ReadContactFile() ([]byte, error) {
	return []byte(mr.contactsBuf.String()), nil
}

func (mr *MemoryRepo) ContactsFileExists() bool {
	return mr.contactsBuf.Len() > 0
}

func (mr *MemoryRepo) WriteMnemonic(mnemonic string) error {
	mr.mnemnoicBuf.Reset() // to simulate truncating a file during writes.
	_, err := mr.mnemnoicBuf.WriteString(mnemonic)
	return err
}

func (mr *MemoryRepo) GetSavedMnemonic() (string, bool) {
	if mr.mnemnoicBuf.Len() == 0 {
		return "", false
	}
	return mr.mnemnoicBuf.String(), true
}

func (mr *MemoryRepo) DeleteMnemonic() error {
	mr.mnemnoicBuf.Reset()
	return nil
}

func (mr *MemoryRepo) WriteSignature(path string, signature string) error {
	return nil
}

func TestContacts(t *testing.T) {
	t.Run("DoesNotCreateContactWithoutName", func(t *testing.T) {
		memoryRepo := &MemoryRepo{}
		err := New(memoryRepo, "", "")
		if err == nil {
			t.Error("Want error when trying to create contact without a name, got nil.")
		}
		if memoryRepo.ContactsFileExists() {
			t.Error("Want no writes to the repository.")
		}
	})

	t.Run("DoesNotCreateContactWithBadPublicKey", func(t *testing.T) {
		memoryRepo := &MemoryRepo{}
		err := New(memoryRepo, "Rob", "invalid-public-key")
		if err == nil {
			t.Error("Want error when trying to create contact without a bad public key, got nil.")
		}
		if memoryRepo.ContactsFileExists() {
			t.Error("Want no writes to the repository.")
		}
	})

	t.Run("CreatesContact", func(t *testing.T) {
		memoryRepo := &MemoryRepo{}
		err := New(memoryRepo, "Rob", "dd399d223656d1d3db00bdeab0c9b587dffa6a91d4f5af4e385c2e7097ff1e0f")
		if err != nil {
			t.Errorf("Want no error, got: %s", err)
		}
		if !memoryRepo.ContactsFileExists() {
			t.Error("Want contact to be written to the repository.")
		}
	})

	t.Run("DoesNotCreateContactWithDuplicateName", func(t *testing.T) {
		memoryRepo := &MemoryRepo{}
		err := New(memoryRepo, "Ken", "b0f3bc6b5fa843a456cc96d3443ca49b1916c20e034874a7cfb4674f271c1870")
		if err != nil {
			t.Errorf("Want no error, got: %s", err)
		}
		err = New(memoryRepo, "Ken", "b0f3bc6b5fa843a456cc96d3443ca49b1916c20e034874a7cfb4674f271c1870")
		if err == nil {
			t.Error("Want error when trying to create contact without a bad public key, got nil.")
		}
	})

	t.Run("GetContacts", func(t *testing.T) {
		memoryRepo := &MemoryRepo{}
		New(memoryRepo, "Rob", "dd399d223656d1d3db00bdeab0c9b587dffa6a91d4f5af4e385c2e7097ff1e0f")
		New(memoryRepo, "Ken", "b0f3bc6b5fa843a456cc96d3443ca49b1916c20e034874a7cfb4674f271c1870")

		contacts, err := GetAll(memoryRepo)
		if err != nil {
			t.Errorf("Want no error, got: %s", err)
		}

		if len(contacts) != 2 {
			t.Errorf("Want two contacts to have been stored, got: %d", len(contacts))
		}

		seenRob := false
		seenKen := false
		for _, c := range contacts {
			if c.Name == "Rob" {
				seenRob = true
			} else if c.Name == "Ken" {
				seenKen = true
			}
		}
		if !seenRob || !seenKen {
			t.Error("Want names of stored contacts to match.")
		}
	})

	t.Run("DeleteContact", func(t *testing.T) {
		memoryRepo := &MemoryRepo{}
		New(memoryRepo, "Rob", "dd399d223656d1d3db00bdeab0c9b587dffa6a91d4f5af4e385c2e7097ff1e0f")
		New(memoryRepo, "Ken", "b0f3bc6b5fa843a456cc96d3443ca49b1916c20e034874a7cfb4674f271c1870")
		Delete(memoryRepo, "Rob")

		contacts, err := GetAll(memoryRepo)
		if err != nil {
			t.Errorf("Want no error, got: %s", err)
		}
		if len(contacts) != 1 {
			t.Errorf("Want one contact after adding 2 and removing 1, got: %d", len(contacts))
		}
		if contacts[0].Name != "Ken" {
			t.Errorf("Want only contact to be Ken, got: %v", contacts[0])
		}
	})
}
