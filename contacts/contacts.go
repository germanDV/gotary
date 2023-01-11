package contacts

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"gitlab.com/germandv/gotary/keyring"
	"gitlab.com/germandv/gotary/repository"
)

type Contact struct {
	Name   string          `json:"name"`
	Public *keyring.Public `json:"publicKey"`
}

// New creates a new contact and saves it.
func New(name string, publicKeyHex string) error {
	if name == "" {
		return errors.New("please indicate contact's name")
	}

	publicKey, err := keyring.PublicKeyFromHex(publicKeyHex)
	if err != nil {
		return err
	}

	contact := &Contact{
		Name:   normalizeName(name),
		Public: publicKey,
	}

	if !repository.ContactsFileExists() {
		contactJsonBytes, err := contactsToJsonBytes([]*Contact{contact})
		if err != nil {
			return err
		}
		return repository.WriteContacts(contactJsonBytes)
	}

	contacts, err := GetAll()
	if err != nil {
		return err
	}

	if isDuplicate(contacts, contact.Name) {
		return fmt.Errorf("contact with name %q already exists", contact.Name)
	}

	contacts = append(contacts, contact)
	contactsJsonBytes, err := contactsToJsonBytes(contacts)
	if err != nil {
		return err
	}
	return repository.WriteContacts(contactsJsonBytes)
}

// GetAll retrieves all contacts.
func GetAll() ([]*Contact, error) {
	bytes, err := repository.ReadContactFile()
	if err != nil {
		return nil, err
	}

	contactsMap := map[string]string{}
	json.Unmarshal(bytes, &contactsMap)

	contacts := []*Contact{}
	for name, hexKey := range contactsMap {
		key, err := keyring.PublicKeyFromHex(hexKey)
		if err != nil {
			return contacts, err
		}

		contacts = append(contacts, &Contact{
			Name:   name,
			Public: key,
		})
	}
	return contacts, nil
}

// Delete removes a contact.
func Delete(name string) error {
	contacts, err := GetAll()
	if err != nil {
		return err
	}

	normalizedName := normalizeName(name)

	filtered := []*Contact{}
	for _, c := range contacts {
		if c.Name != normalizedName {
			filtered = append(filtered, c)
		}
	}

	contactsJsonBytes, err := contactsToJsonBytes(filtered)
	if err != nil {
		return err
	}
	return repository.WriteContacts(contactsJsonBytes)
}

func contactsToJsonBytes(contacts []*Contact) ([]byte, error) {
	result := map[string]string{}
	for _, c := range contacts {
		result[c.Name] = c.Public.Hex
	}
	return json.Marshal(result)
}

func normalizeName(name string) string {
	return strings.ReplaceAll(name, " ", "_")
}

func isDuplicate(contacts []*Contact, name string) bool {
	found := false
	for _, c := range contacts {
		if c.Name == name {
			found = true
		}
	}
	return found
}
