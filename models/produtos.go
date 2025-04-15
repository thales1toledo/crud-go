package models

import (
	"GoLand/db"
	"log"
	"strconv"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosProdutos() []Produto {
	bd := db.ConectaBd()

	rows, err := bd.Query("SELECT * FROM produtos ORDER BY id ASC")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for rows.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = rows.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}

	defer rows.Close()
	defer bd.Close()

	return produtos
}

func CriaNovoProduto(nome, descricao string, preco float64, quantidade int) {
	bd := db.ConectaBd()

	insertDados, err := bd.Prepare("INSERT INTO produtos (nome, descricao, preco, quantidade) VALUES ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	_, err = insertDados.Exec(nome, descricao, preco, quantidade)
	if err != nil {
		panic(err.Error())
	}

	defer bd.Close()
}

func DeletaProduto(id string) {
	bd := db.ConectaBd()

	deleteProduto, err := bd.Prepare("DELETE FROM produtos WHERE id = $1")
	if err != nil {
		panic(err.Error())
	}

	_, err = deleteProduto.Exec(id)
	if err != nil {
		panic(err.Error())
	}

	defer bd.Close()
}

func FindOne(id string) Produto {
	bd := db.ConectaBd()
	defer bd.Close()

	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Println("ID inválido:", err)
		return Produto{}
	}

	row := bd.QueryRow("SELECT * FROM produtos WHERE id = $1", idInt)

	produto := Produto{}
	err = row.Scan(&produto.Id, &produto.Nome, &produto.Descricao, &produto.Preco, &produto.Quantidade)
	if err != nil {
		log.Println("Erro ao buscar produto:", err)
		return Produto{}
	}

	return produto
}

func AtualizaProduto(id int, nome, descricao string, preco float64, quantidade int) {
	bd := db.ConectaBd()

	updateProduto, err := bd.Prepare("UPDATE produtos SET nome = $1, descricao = $2, preco = $3, quantidade = $4 WHERE id = $5")
	if err != nil {
		panic(err.Error())
	}

	_, err = updateProduto.Exec(nome, descricao, preco, quantidade, id)
	if err != nil {
		panic(err.Error())
	}

	defer bd.Close()
}
