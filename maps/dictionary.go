package main

//import "errors"

type Dictionary map[string]string
type DictionaryErr string

var (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add the word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot perform operations on word that does not exist")
)

func (e DictionaryErr) Error() string {
	return string(e)
}
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	// a map can return two values the second value can be a boolean which indicates if the key is returned or not
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word string, value string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		d[word] = value
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}
func (d Dictionary) Update(word, definiton string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definiton
	default:
		return err
	}
	return nil

}
func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		delete(d, word)
	default:
		return err
	}
	return nil
}

// func Search(dict map[string]string, keyword string) string {

// 	return dict[keyword]
// }
