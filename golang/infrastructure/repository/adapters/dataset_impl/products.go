package datasetimpl

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/Akiles94/mytheresa-test/domain/dto"
	"github.com/Akiles94/mytheresa-test/domain/models"
	"github.com/Akiles94/mytheresa-test/infrastructure/httputils"
	log "github.com/sirupsen/logrus"
)

func (d *DatasetRepo) GetProducts(params dto.QueryParams) (*[]models.Product, error) {
	response := []models.Product{}
	var pwd string = ""
	//Getting current working dir
	if d.env == "dev" {
		pwd, _ = os.Getwd()
	} else {
		pwd = "/go/src"
	}
	//Reading file
	file, err := ioutil.ReadFile(pwd + d.filePath)
	if err != nil {
		log.Error("Error reading provided filepath: ", err)
		return nil, err
	}
	//Product struct for binding
	data := []models.Product{}
	//Binding file data to struct
	err = json.Unmarshal([]byte(file), &data)
	if err != nil {
		log.Error("Error unmarshaling data from file: ", err)
		return nil, err
	}

	response = httputils.GetProductsFiltered(data, params)

	return &response, nil
}
