package api

import (
	"listaativosinforme/aws"
	"strconv"
	"time"

	"context"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"

	"dev.azure.com/btgpactual/CLIENT-CLIENTTOOLS/_git/Panda.Golang.Framework.git/amazon"
	"dev.azure.com/btgpactual/CLIENT-CLIENTTOOLS/_git/Panda.Golang.Framework.git/database"
	"dev.azure.com/btgpactual/CLIENT-CLIENTTOOLS/_git/Panda.Golang.Framework.git/database/collections"
	"dev.azure.com/btgpactual/CLIENT-CLIENTTOOLS/_git/Panda.Golang.Framework.git/model"

	"archive/zip"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var URL_ARQUIVO = ""

func init() {
	// URL_ARQUIVO := env.GetUrlListaAtivosBacen()
	fmt.Println(URL_ARQUIVO)
}

func show(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "I´m alive 1"})
}

func healthcheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "available 1"})
}

func getBasic(c echo.Context, group, field string) error {
	item, err := aws.GetItem(group, field, c)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, item)

}
func AtualizaBacen(c echo.Context) error {
	total, erro := Processo(c.Request().Context())
	if erro != nil {
		fmt.Println("Erro Obtendo dados do bacen: ", erro.Error())
		return c.JSON(http.StatusOK, map[string]string{"Mensagem": "Erros ao executar o processo"})
	} else {
		return c.JSON(http.StatusOK, map[string]string{"total": strconv.FormatInt(total, 10)})
	}
}

func Processo(ctx context.Context) (int64, error) {

	client := amazon.GetDocumentDB()
	collectionLISTA_ATIVOS := client.Database(database.DOCUMENTDB).Collection(collections.LISTA_ATIVOS)

	listaAtivos, err := getBacen()
	if err != nil {
		fmt.Println("Erro Obtendo dados do bacen: ", err.Error())
		return 0, err
	}

	_, err = collectionLISTA_ATIVOS.DeleteMany(context.TODO(), bson.D{{}})

	if err != nil {
		log.Println("Erro inesperado ao tentar deletar a collecton Lista Ativos: ", err.Error())
		return 0, err
	}

	var inserts []interface{}
	for _, ativo := range listaAtivos {
		inserts = append(inserts, ativo)
	}
	_, err = collectionLISTA_ATIVOS.InsertMany(context.TODO(), inserts)

	if err != nil {
		fmt.Println("Erro no insert: ", err.Error())
	}
	total := int64(len(listaAtivos))
	return total, err
}

func getBacen() ([]model.ListaAtivos, error) {
	var lista_de_ativos []model.ListaAtivos
	URL_ARQUIVO := "http://sistemas.cvm.gov.br/download/listaativos/ListaAtivos.zip"
	resp, err := http.Get(URL_ARQUIVO)
	if err != nil {
		fmt.Println("Erro ao fazer download do ListaAtivos do bacen.")
		return lista_de_ativos, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erro ao tentar ler os dados do Bacen (zip).")
		return lista_de_ativos, err
	}

	zipReader, err := zip.NewReader(bytes.NewReader(body), int64(len(body)))
	if err != nil {
		fmt.Println("Erro ao tentar descompacar arquivo do Bacen.")
		return lista_de_ativos, err
	}

	// Read all the files from zip archive
	arquivoFull := ""
	for _, zipFile := range zipReader.File {
		fmt.Println("Reading file:", zipFile.Name)
		unzippedFileBytes, err := readZipFile(zipFile)
		if err != nil {
			log.Println("Erro lendo arquivo do bacen for if: ", err.Error())
			continue
		}

		// _ = unzippedFileBytes // this is unzipped file bytes
		arquivoFull = string(unzippedFileBytes)
	}

	linhas := strings.Split(strings.ReplaceAll(arquivoFull, "\r\n", "\n"), "\n")
	count := 0
	for _, row := range linhas[:len(linhas)-1] {
		if count > 0 {
			ativo := model.ListaAtivos{}
			coluna := strings.Split(row, ";")

			ativo.Cnpj_do_mercado = strings.TrimSpace(coluna[0])
			ativo.Denominacao_social_do_mercado = strings.TrimSpace(coluna[1])
			ativo.Codigo_do_ativo = strings.TrimSpace(coluna[2])
			ativo.Descricao_do_ativo = strings.TrimSpace(coluna[3])
			ativo.Data_inicio_vigencia = CorrigeData(strings.TrimSpace(coluna[4]))
			ativo.Data_fim_vigencia = CorrigeData(strings.TrimSpace(coluna[5]))
			ativo.Data_inicio_suspensao = CorrigeData(strings.TrimSpace(coluna[6]))
			ativo.Data_fim_suspensao = CorrigeData(strings.TrimSpace(coluna[7]))
			ativo.Codigo_tipo_ativo = strings.TrimSpace(coluna[8])
			ativo.Descricao_do_tipo_do_ativo = strings.TrimSpace(coluna[9])

			lista_de_ativos = append(lista_de_ativos, ativo)
		}
		count++
	}

	fmt.Println("Fim...", len(linhas), " Linhas percorridas!")
	return lista_de_ativos, nil

}

func readZipFile(zf *zip.File) ([]byte, error) {
	f, err := zf.Open()
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return ioutil.ReadAll(f)
}

func CorrigeData(date string) time.Time {
	if len(date) <= 0 {
		return time.Time{}
	}

	// date := "29/03/2015"
	in := "02/01/2006"
	out := "2006-01-02T00:00:00"
	dt, err := time.Parse(in, date)
	if err != nil {
		fmt.Println(err)
		// return nil, nil
	}
	fmt.Println(dt.Format(out))
	saida := dt.Format(out)

	dt1, _ := time.Parse(out, saida)
	return dt1
}
