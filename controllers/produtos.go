package controllers

import (
	"GoLand/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosProdutos := models.BuscaTodosProdutos()
	err := temp.ExecuteTemplate(w, "Index", todosProdutos)
	if err != nil {
		return
	}
}

func New(w http.ResponseWriter, r *http.Request) {
	err := temp.ExecuteTemplate(w, "New", nil)
	if err != nil {
		return
	}
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoFormatted, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão do preço: ", err)
		}

		quantidadeFormatted, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversão da quantidade: ", err)
		}

		models.CriaNovoProduto(nome, descricao, precoFormatted, quantidadeFormatted)
	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Delete(w http.ResponseWriter, r *http.Request) {

	idProduto := r.URL.Query().Get("id")
	models.DeletaProduto(idProduto)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Editar(w http.ResponseWriter, r *http.Request) {

	idProduto := r.URL.Query().Get("id")

	produto := models.FindOne(idProduto)

	err := temp.ExecuteTemplate(w, "Editar", produto)
	if err != nil {
		return
	}
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		idProduto := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		idInt, err := strconv.Atoi(idProduto)
		if err != nil {
			log.Println("Erro na conversao do id: ", err)
		}

		precoFormatted, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão do preço: ", err)
		}

		quantidadeFormatted, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversão da quantidade: ", err)
		}

		models.AtualizaProduto(idInt, nome, descricao, precoFormatted, quantidadeFormatted)

		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
}
