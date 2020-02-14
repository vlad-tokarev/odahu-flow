//
//    Copyright 2019 EPAM Systems
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.
//

package main

import (
	"fmt"
	"github.com/odahu/odahu-flow/packages/operator/pkg/config"
	trainer_conf "github.com/odahu/odahu-flow/packages/operator/pkg/config/trainer"
	conn_http_storage "github.com/odahu/odahu-flow/packages/operator/pkg/repository/connection/http"
	train_http_storage "github.com/odahu/odahu-flow/packages/operator/pkg/repository/training/http"
	"github.com/odahu/odahu-flow/packages/operator/pkg/trainer"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
)

var log = logf.Log.WithName("trainer-main")

const (
	mtFileCLIParam               = "mt-file"
	mtIDCLIParam                 = "mt-id"
	outputConnectionNameCLIParam = "output-connection-name"
	apiURLCLIParam               = "api-url"
	outputTrainingDirCLIParam    = "output-dir"
)

var mainCmd = &cobra.Command{
	Use:              "trainer",
	Short:            "odahu-flow trainer cli",
	TraverseChildren: true,
}

var trainerSetupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Prepare environment for a trainer",
	Run: func(cmd *cobra.Command, args []string) {
		if err := newTrainerWithHTTPRepositories().Setup(); err != nil {
			log.Error(err, "Training setup failed")
			os.Exit(1)
		}
	},
}

var saveCmd = &cobra.Command{
	Use:   "result",
	Short: "Save a trainer result",
	Run: func(cmd *cobra.Command, args []string) {
		if err := newTrainerWithHTTPRepositories().SaveResult(); err != nil {
			log.Error(err, "Result saving failed")
			os.Exit(1)
		}
	},
}

func init() {
	currentDir, err := os.Getwd()
	if err != nil {
		// Impossible situation
		panic(err)
	}

	config.InitBasicParams(mainCmd)

	mainCmd.PersistentFlags().String(mtFileCLIParam, "mt.json", "File with model training content")
	config.PanicIfError(viper.BindPFlag(trainer_conf.MTFile, mainCmd.PersistentFlags().Lookup(mtFileCLIParam)))

	mainCmd.PersistentFlags().String(mtIDCLIParam, "", "ID of the model training")
	config.PanicIfError(viper.BindPFlag(trainer_conf.ModelTrainingID, mainCmd.PersistentFlags().Lookup(mtIDCLIParam)))

	mainCmd.PersistentFlags().String(apiURLCLIParam, "", "API URL")
	config.PanicIfError(viper.BindPFlag(trainer_conf.APIURL, mainCmd.PersistentFlags().Lookup(apiURLCLIParam)))

	mainCmd.PersistentFlags().String(
		outputTrainingDirCLIParam, currentDir,
		"The path to the dir when a user trainer will save their result",
	)
	config.PanicIfError(viper.BindPFlag(
		trainer_conf.OutputTrainingDir, mainCmd.PersistentFlags().Lookup(outputTrainingDirCLIParam),
	))

	mainCmd.PersistentFlags().String(outputConnectionNameCLIParam,
		"It is a connection ID, which specifies where a artifact trained artifact is stored.",
		"File with model training content",
	)
	config.PanicIfError(viper.BindPFlag(
		trainer_conf.OutputConnectionName,
		mainCmd.PersistentFlags().Lookup(outputConnectionNameCLIParam)),
	)

	mainCmd.AddCommand(trainerSetupCmd, saveCmd)
}

func newTrainerWithHTTPRepositories() *trainer.ModelTrainer {
	log.Info(fmt.Sprintf("OAuthOIDCTokenEndpoint: %s", viper.GetString(trainer_conf.OAuthOIDCTokenEndpoint)))
	trainRepo := train_http_storage.NewRepository(
		viper.GetString(trainer_conf.APIURL), viper.GetString(trainer_conf.APIToken),
		viper.GetString(trainer_conf.ClientID), viper.GetString(trainer_conf.ClientSecret),
		viper.GetString(trainer_conf.OAuthOIDCTokenEndpoint),
	)
	connRepo := conn_http_storage.NewRepository(
		viper.GetString(trainer_conf.APIURL), viper.GetString(trainer_conf.APIToken),
		viper.GetString(trainer_conf.ClientID), viper.GetString(trainer_conf.ClientSecret),
		viper.GetString(trainer_conf.OAuthOIDCTokenEndpoint),
	)

	return trainer.NewModelTrainer(trainRepo, connRepo, viper.GetString(trainer_conf.ModelTrainingID))
}

func main() {
	if err := mainCmd.Execute(); err != nil {
		log.Error(err, "trainer CLI command failed")

		os.Exit(1)
	}
}
