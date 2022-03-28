package cmd

import (
	"github.com/Akiles94/mytheresa-test/domain/adapters"
	config "github.com/Akiles94/mytheresa-test/infrastructure/config"
	http "github.com/Akiles94/mytheresa-test/infrastructure/http/adapters"
	datasetimpl "github.com/Akiles94/mytheresa-test/infrastructure/repository/adapters/dataset_impl"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(serveCmd)
}

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run CD-Ecosys service",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Init server...")
		initServer()
	},
}

//Initialize the server with given configuration and implementations of infrastructure adapters
func initServer() {

	config := config.Init()

	//Init main repository
	repo := &datasetimpl.DatasetRepo{}
	repo.Init("/infrastructure/assets/products.json", config.Env)

	productsImpl := &adapters.ProductsImpl{
		Repo: repo,
	}

	api := &http.API{
		Port:         config.Port,
		ProductsImpl: productsImpl,
	}

	api.Start()

}
