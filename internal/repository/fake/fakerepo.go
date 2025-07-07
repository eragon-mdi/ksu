package fakerepo

func New(s FakeStorage) *fakeRepository {
	return &fakeRepository{
		storage: s,
	}
}

type fakeRepository struct {
	storage FakeStorage
}
