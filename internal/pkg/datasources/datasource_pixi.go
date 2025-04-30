package datasources

var _ Loader = &PixiLoader{}

type PixiLoader struct {
	Options Options
}

func NewPixiLoader(datasourceOption map[string]string, options Options, secrets Secrets) (*PixiLoader, error) {
	return &PixiLoader{}, nil
}

func (l *PixiLoader) Sync(fromURI string, toPath string) error {
	return nil
}
