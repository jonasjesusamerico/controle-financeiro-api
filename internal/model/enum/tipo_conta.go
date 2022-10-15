package enum

type TipoConta int

const (
	CORRENTE TipoConta = 1
	DEBITO   TipoConta = 2
	CREDITO  TipoConta = 3
)

func (tp TipoConta) Description() string {
	switch tp {
	case CORRENTE:
		return "Corrente"
	case DEBITO:
		return "Débito"
	case CREDITO:
		return "Crédito"
	}
	return "Tipo de conta desconhecida"

}
