package typeutil

import "go/types"

type Types struct {
	pkg  *types.Package
	info *types.Info
}

func newTypes(path, name string) Types {
	pkg, info := CreateInfoAndPkg(path, name)
	return Types{pkg, info}
}

func (t Types) Defs() map[string]types.Type {
	m := make(map[string]types.Type, len(t.info.Defs))
	for name, obj := range t.info.Defs {
		m[name.Name] = obj.Type()
	}

	return m
}

func (t Types) Types() []types.Type {
	m := make([]types.Type, len(t.info.Defs))

	i := 0
	for _, obj := range t.info.Defs {
		m[i] = obj.Type()
		i++
	}

	return m
}
