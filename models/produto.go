package models

type Produto struct {
	Id            int     `json:"id"`
	Nome          string  `json:"nome"`
	Quantidade    int32   `json:"quantidade"`
	ValorUnitario float64 `json:"valorUnitario"`
}

func (p *Produto) GetQuantidade() int32 {
	return p.Quantidade
}

func (p *Produto) GetId() int {
	return p.Id
}

func (p *Produto) SetQuantidade(quantidade int32) {
	p.Quantidade = quantidade
}

var Produtos []Produto
