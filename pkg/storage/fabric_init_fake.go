package storage

import (
	"github.com/eragon-mdi/ksu/pkg/config"
	"github.com/eragon-mdi/ksu/pkg/fake"
)

type fakeStub struct {
	*fake.StorageType
}

func init() {
	register("internal", &fakeStub{})
}

func (f *fakeStub) Connect(cfg config.Config) (err error) {
	f.StorageType, err = fake.New()
	return
}

func (f *fakeStub) Migrate(cfg config.Config) error {
	return nil
}

func (f *fakeStub) Shutdown() error {
	return nil
}
