package keyring

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type Contact struct {
	Name   string  `json:"name"`
	Public *Public `json:"publicKey"`
}

func NewContact(name string, publicKeyHex string) error {
	if name == "" {
		return errors.New("please indicate contact's name")
	}

	publicKey, err := PublicKeyFromHex(publicKeyHex)
	if err != nil {
		return err
	}

	contact := &Contact{
		Name:   normalizeName(name),
		Public: publicKey,
	}

	if !ContactsFileExists() {
		contactJsonBytes, err := contactsToJsonBytes([]*Contact{contact})
		if err != nil {
			return err
		}
		return WriteContacts(contactJsonBytes)
	}

	contacts, err := GetContacts()
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
	return WriteContacts(contactsJsonBytes)
}

func GetContacts() ([]*Contact, error) {
	bytes, err := ReadContactFile()
	if err != nil {
		return nil, err
	}

	contactsMap := map[string]string{}
	json.Unmarshal(bytes, &contactsMap)

	contacts := []*Contact{}
	for name, hexKey := range contactsMap {
		key, err := PublicKeyFromHex(hexKey)
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

func DeleteContact(name string) error {
	contacts, err := GetContacts()
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
	return WriteContacts(contactsJsonBytes)
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
