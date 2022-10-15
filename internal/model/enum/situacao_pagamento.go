package enum

type SituacaoPagamento int

const (
	QUITADO   SituacaoPagamento = 1
	EM_ABERTO SituacaoPagamento = 2
)

func (tp SituacaoPagamento) Description() string {
	switch tp {
	case QUITADO:
		return "Quitado"
	case EM_ABERTO:
		return "Em aberto"
	}
	return "Situação do pagamento desconhecido"

}
