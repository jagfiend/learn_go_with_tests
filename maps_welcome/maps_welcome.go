package maps_welcome

type Dictionary map[string]string

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

const (
	ErrWordAlreadyDefined = DictionaryErr("word has already been defined")
	ErrWordNotFound       = DictionaryErr("could not find the word you were looking for")
	ErrWordDoesntExist    = DictionaryErr("given word doesnt exist in the dictionary to update")
)

func (d Dictionary) Add(key string, defintion string) error {
	_, found := d[key]

	if found {
		return ErrWordAlreadyDefined
	}

	d[key] = defintion

	return nil
}

func (d Dictionary) Search(key string) (string, error) {
	definition, found := d[key]

	if !found {
		return "", ErrWordNotFound
	}

	return definition, nil
}

func (d Dictionary) Update(key string, definition string) error {
	_, found := d[key]

	if !found {
		return ErrWordDoesntExist
	}

	d[key] = definition

	return nil
}

func (d Dictionary) Delete(key string) {
	delete(d, key)
}
