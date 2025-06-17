package fake

type Deleter interface {
	Delete()
}

func (s StorageType) Delete() {

}
