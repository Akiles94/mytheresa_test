package datasetimpl

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/Akiles94/mytheresa-test/domain/models"
	log "github.com/sirupsen/logrus"
)

type DatasetRepo struct {
	filePath string
	//dataset  *[]models.Product
}

func (d *DatasetRepo) Init(filepath string, env string) error {
	d.filePath = filepath
	var pwd string = "dev"
	//Getting current working dir
	if env == "dev" {
		pwd, _ = os.Getwd()
	} else {
		pwd = "/go/src"
	}
	//Reading file
	file, err := ioutil.ReadFile(pwd + filepath)
	if err != nil {
		log.Fatal("Error reading provided filepath: ", err)
		return err
	}
	//Product struct for binding
	data := []models.Product{}
	//Binding file data to struct
	err = json.Unmarshal([]byte(file), &data)
	if err != nil {
		log.Fatal("Error unmarshaling data from file: ", err)
		return err
	}

	log.Info("File readed succesfully!")

	return nil
}
