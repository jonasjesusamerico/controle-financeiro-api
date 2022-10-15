package enum

type FormaPagamento int

const (
	AVISTA FormaPagamento = 1
	PRAZO  FormaPagamento = 2
)

func (fp FormaPagamento) Description() string {
	switch fp {
	case AVISTA:
		return "À vista"
	case PRAZO:
		return "Débito"
	}
	return "Forma de pagamento desconhecido"

}
