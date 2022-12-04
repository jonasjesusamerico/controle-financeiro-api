package lancamento

import (
	"time"

	"github.com/joninhasamerico/controle-financeiro-api/internal/model"
	"github.com/joninhasamerico/controle-financeiro-api/internal/model/conta"
	"github.com/joninhasamerico/controle-financeiro-api/internal/model/enum"
	"github.com/joninhasamerico/controle-financeiro-api/internal/model/setor"
	"gorm.io/gorm"
)

type Lancamento struct {
	ID             uint64
	Descricao      string
	ContaID        uint64
	Conta          conta.Conta
	SetorID        uint64
	Setor          setor.Setor
	FormaPagamento enum.FormaPagamento
	Valor          float64
	Data           time.Time
	Situacao       enum.SituacaoPagamento
	model.Tenant
}

func NewLancamento() *Lancamento {
	return &Lancamento{}
}

func NewSliceLancamento() []Lancamento {
	return []Lancamento{}
}

func (l Lancamento) GetId() uint64 {
	return l.ID
}

func (l *Lancamento) SetTenant(tenantId int64) {
	l.TenantID = tenantId
}

func (u *Lancamento) BeforeCreate(tx *gorm.DB) (err error) {
	if u.TenantID == 0 {
		tx.Model(u).Update("tenant_id", u.ID)
	}
	return
}

func LancamentoConverter(dto LancamentoDto) *Lancamento {
	return &Lancamento{
		ID:             dto.ID,
		Descricao:      dto.Descricao,
		ContaID:        dto.ContaID,
		SetorID:        dto.SetorID,
		FormaPagamento: dto.FormaPagamento,
		Valor:          dto.Valor,
		Data:           dto.Data,
		Situacao:       dto.Situacao,
	}
}
