package model

type baseTable struct {
	rawPrefix string
	prefix    string
	name      string
}

func (bt baseTable) Name() string {
	return bt.name
}

func (bt baseTable) NameAlter() string {
	if bt.rawPrefix == "" || bt.rawPrefix == bt.Name() {
		return bt.Name()
	}

	return bt.Name() + " " + bt.rawPrefix
}

func (bt baseTable) withPrefix(p string) baseTable {
	if p == "" {
		return baseTable{
			name: bt.name,
		}
	}

	return baseTable{
		name:      bt.name,
		rawPrefix: p,
		prefix:    p + ".",
	}
}
