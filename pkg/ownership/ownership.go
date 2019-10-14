package ownership

type Item interface {
	GetKind() string
	GetName() string
}

type Store interface {
	Check(Item) bool
	Set(Item) error
	// ToDo: add remove method for when an item is deleted
}
