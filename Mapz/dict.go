package main

type Dict map[string]string

type DictErr string

func (e DictErr) Error() string {
	return string(e)
}

const (
	ErrDictBroken     = DictErr("all ur dict r belong to us")
	ErrDictTooFull    = DictErr("teh dict is about to bust lolz")
	ErrDictInadequate = DictErr("teh dict is missing teh key")
)

func (d Dict) Delete(k string) error {
	_, ok := d[k]
	if !ok {
		return ErrDictInadequate
	}
	delete(d, k)
	return nil
}

func (d Dict) Update(k, v string) error {
	_, ok := d[k]
	if !ok {
		return ErrDictInadequate
	}
	d[k] = v
	return nil
}

func (d Dict) Add(k, v string) error {
	_, ok := d[k]
	if ok {
		return ErrDictTooFull
	}
	d[k] = v
	return nil
}

func (d Dict) Search(word string) (error, string) {
	v, ok := d[word]
	if !ok {
		return ErrDictBroken, ""
	}
	return nil, v
}
