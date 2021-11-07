package cmdReader

import (
	mongo "github.com/NickLovera/mongo-utils"
	"github.com/spf13/cobra"
	v "github.com/spf13/viper"
	"log"
)


var viper = v.New()
var rootCmd = &cobra.Command{
	Use: "Setup",
	Short: "Flooring Service",
	Long: "Service that handles flooring payments",
	Run: func(cmd *cobra.Command, args []string){
		dbConn, dbErr := mongo.ConnectAndMigrate(nil, "flooring")
		if dbErr != nil {
			log.Panicln("Cannot connect to mongo db: ", dbErr)
		}
		log.Println("Migration Done")
		_ = mongo.NewMongoClient(dbConn)
		log.Println("Connected")


	},

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Panicln(err)
	}
}