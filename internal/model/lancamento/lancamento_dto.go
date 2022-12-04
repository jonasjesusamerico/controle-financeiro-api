package lancamento

import (
	"time"

	"github.com/joninhasamerico/controle-financeiro-api/internal/model/enum"
)

type LancamentoDto struct {
	ID             uint64
	Descricao      string
	ContaID        uint64
	SetorID        uint64
	FormaPagamento enum.FormaPagamento
	Valor          float64
	Data           time.Time
	Situacao       enum.SituacaoPagamento
}

func NewLancamentoDto() *LancamentoDto {
	return &LancamentoDto{}
}
