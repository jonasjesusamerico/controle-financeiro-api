package model

type IModel interface {
	GetId() uint64

	Validate() error
}
