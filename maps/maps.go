package maps

type Dictionary map[string]string

const (
	ErrNotFound       = DictionaryErr("word not found")
	ErrKeyExists      = DictionaryErr("key already exists")
	ErrKeyNotFound    = DictionaryErr("key not found")
	ErrDelKeyNotFound = DictionaryErr("key not found, no deletion")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(key string) (string, error) {
	word, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}
	return word, nil
}

func (d Dictionary) Add(key, value string) error {
	_, ok := d[key]
	if ok {
		return ErrKeyExists
	}
	d[key] = value
	return nil
}

func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)
	switch err {
	case ErrNotFound:
		return ErrKeyNotFound
	default:
		d[key] = value
		return nil
	}
}

func (d Dictionary) Delete(key string) error {
	_, err := d.Search(key)
	switch err {
	case ErrNotFound:
		return ErrDelKeyNotFound
	default:
		delete(d, key)
		return nil
	}
}
