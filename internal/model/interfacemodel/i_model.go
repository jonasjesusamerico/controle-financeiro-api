package interfacemodel

type IModel interface {
	GetId() uint64
	SetTenant(int64)
}
