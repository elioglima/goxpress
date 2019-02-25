package db_assistente

import (
	dbsql "goxpress/backend/database"
	"strconv"
)

type CTabela struct {
	dbsql.DBBase
	sSQL string
}

func New(MyConexao *dbsql.MyConexao) *CTabela {
	c := &CTabela{}
	c.Nome_Classe = "assistente"
	c.New(MyConexao)
	return c
}

func (c *CTabela) GetID(in_id_empresa, in_id int) error {
	c.sSQL = "select * from " + c.Nome_Tabela
	c.sSQL += " where id_empresa = " + strconv.Itoa(in_id_empresa)
	c.sSQL += " and id = " + strconv.Itoa(in_id)
	return c.Query(c.sSQL)
}

func (c *CTabela) GetAll(in_id_empresa int) error {
	c.sSQL = "select * from " + c.Nome_Tabela
	c.sSQL += " where id_empresa = " + strconv.Itoa(in_id_empresa)
	return c.Query(c.sSQL)
}
