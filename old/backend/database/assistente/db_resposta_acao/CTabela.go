package db_resposta_acao

import (
	"errors"
	dbsql "goxpress/backend/database"
	"goxpress/libs"
	"strconv"
)

func New(MyConexao *dbsql.MyConexao) *CTabela {
	c := &CTabela{}
	c.Nome_Classe = "resposta_Acao"
	c.New(MyConexao)
	return c
}

type CTabela struct {
	dbsql.DBBase
	sSQL string
}

func (c *CTabela) ChkExistsDescricao(in_id_empresa, in_id_assistente int, in_descricao string) error {
	c.sSQL = "select id from " + c.Nome_Tabela
	c.sSQL += " where id_empresa = " + strconv.Itoa(in_id_empresa)
	c.sSQL += " and id_assistente = " + strconv.Itoa(in_id_assistente)
	c.sSQL += " and descricao = " + libs.Asp(in_descricao)
	c.sSQL += " limit 0,1"
	err := c.Query(c.sSQL)
	if err != nil {
		return err
	}

	if c.RowsCount > 0 {
		return errors.New("O registro ja existe na base de dados")
	}

	return nil
}

func (c *CTabela) ChkExistsId(in_id, in_id_empresa, in_id_assistente int) error {
	c.sSQL = "select id from " + c.Nome_Tabela
	c.sSQL += " where id_empresa = " + strconv.Itoa(in_id_empresa)
	c.sSQL += " and id_assistente = " + strconv.Itoa(in_id_assistente)
	c.sSQL += " and id = " + strconv.Itoa(in_id)
	c.sSQL += " limit 0,1"
	err := c.Query(c.sSQL)
	if err != nil {
		return err
	}

	if c.RowsCount > 0 {
		return errors.New("O registro ja existe na base de dados")
	}

	return nil
}

func (c *CTabela) Edd(dados CRow) error {
	err := c.ChkExistsId(dados.Id, dados.Id_empresa, dados.Id_assistente)
	if err != nil {
		c.sSQL = "update " + c.Nome_Tabela + " set "
		c.sSQL += " id_empresa = " + strconv.Itoa(dados.Id_empresa)
		c.sSQL += ", id_assistente = " + strconv.Itoa(dados.Id_assistente)
		c.sSQL += ", descricao = " + libs.Asp(dados.Descricao)
		c.sSQL += ", ativo = " + strconv.Itoa(1)
		c.sSQL += " where id = " + strconv.Itoa(dados.Id)
		return c.ExecSQL(c.sSQL)
	}

	return errors.New("Registro não localizado")
}

func (c *CTabela) Add(dados CRow) error {
	err := c.ChkExistsDescricao(dados.Id_empresa, dados.Id_assistente, dados.Descricao)
	if err != nil {
		return err
	}

	c.sSQL = "insert into " + c.Nome_Tabela + " "
	c.sSQL += " (id_empresa, id_assistente, descricao, ativo) "
	c.sSQL += " values ( "
	c.sSQL += strconv.Itoa(dados.Id_empresa)
	c.sSQL += "," + strconv.Itoa(dados.Id_assistente)
	c.sSQL += "," + libs.Asp(dados.Descricao)
	c.sSQL += "," + strconv.Itoa(1)
	c.sSQL += " ) "
	return c.ExecSQL(c.sSQL)
}
