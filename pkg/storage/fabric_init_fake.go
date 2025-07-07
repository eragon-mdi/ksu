package storage

import (
	"github.com/eragon-mdi/ksu/pkg/config"
	"github.com/eragon-mdi/ksu/pkg/fake"
)

/*
func init() {
	register("internal", &InitFuncs{
		Connect: connectFakeAdapter,
		Migrate: migrateStub,
	})
}

func connectFakeAdapter(cfg config.Config) (storageImplement, error) {
	return fake.New()
}

func migrateStub(cfg config.Config) error {
	return nil
}
*/

type fakeStub struct {
	*fake.StorageType
}

func init() {
	//var fake fakeStub
	register("internal", &fakeStub{})
}

func (f *fakeStub) Connect(cfg config.Config) (err error) {
	f.StorageType, err = fake.New()
	return
}

func (f *fakeStub) Migrate(cfg config.Config) error {
	return nil
}
