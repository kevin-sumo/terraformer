// Copyright 2018 The Terraformer Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package cmd

import (
	datadog_terraforming "github.com/GoogleCloudPlatform/terraformer/providers/sumologic"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/spf13/cobra"
)

func newCmdDatadogImporter(options ImportOptions) *cobra.Command {
	var apiKey, appKey, apiURL, validate string
	cmd := &cobra.Command{
		Use:   "sumologic",
		Short: "Import current state to Terraform configuration from Sumo Logic",
		Long:  "Import current state to Terraform configuration from Sumo Logic",
		RunE: func(cmd *cobra.Command, args []string) error {
			provider := newSumoLogicProvider()
			err := Import(provider, options, []string{apiId, apiKey, deployment})
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.AddCommand(listCmd(newDataDogProvider()))
	baseProviderFlags(cmd.PersistentFlags(), &options, "monitors,users", "monitor=id1:id2:id4")
	cmd.PersistentFlags().StringVarP(&apiKey, "api-id", "", "", "YOUR_SUMOLOGIC_API_ID or env param SUMOLOGIC_API_ID")
	cmd.PersistentFlags().StringVarP(&appKey, "api-key", "", "", "YOUR_SUMOLOGIC_API_KEY or env param SUMOLOGIC_API_KEY")
	cmd.PersistentFlags().StringVarP(&apiURL, "deployment", "", "", "YOUR_SUMOLOGIC_DEPLOYMENT or env param SUMOLOGIC_DEPLOYMENT")
	return cmd
}

func newSumoLogicProvider() terraformutils.ProviderGenerator {
	return &sumologic_terraforming.SumoLogicProvider{}
}
