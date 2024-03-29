/**
 * Copyright 2023-present Coinbase Global, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cmd

import (
	"fmt"
	"github.com/coinbase-samples/intx-sdk-go"
	"github.com/jc-cb/intx-cli-dev/utils"
	"github.com/spf13/cobra"
)

var createCryptoAddressCmd = &cobra.Command{
	Use:   "create-crypto-address",
	Short: "Create crypto address.",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := utils.GetClientFromEnv()
		if err != nil {
			return fmt.Errorf("failed to initialize client: %w", err)
		}

		portfolioId, err := utils.GetPortfolioId(cmd, client)
		if err != nil {
			return err
		}

		assetId, err := cmd.Flags().GetString(utils.AssetIdFlag)
		if err != nil {
			return fmt.Errorf("cannot cancel order: %w", err)
		}

		networkArnId, err := cmd.Flags().GetString(utils.NetworkArnIdFlag)
		if err != nil {
			return fmt.Errorf("cannot cancel order: %w", err)
		}

		ctx, cancel := utils.GetContextWithTimeout()
		defer cancel()

		request := &intx.CreateCryptoAddressRequest{
			PortfolioId:  portfolioId,
			AssetId:      assetId,
			NetworkArnId: networkArnId,
		}

		response, err := client.CreateCryptoAddress(ctx, request)
		if err != nil {
			return fmt.Errorf("cannot create address: %w", err)
		}

		jsonResponse, err := utils.FormatResponseAsJson(cmd, response)
		if err != nil {
			return err
		}

		fmt.Println(jsonResponse)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(createCryptoAddressCmd)

	createCryptoAddressCmd.Flags().StringP(utils.AssetIdFlag, "a", "", "ID of the asset (Required)")
	createCryptoAddressCmd.Flags().StringP(utils.NetworkArnIdFlag, "n", "", "Network Arn ID of the asset (Required)")
	createCryptoAddressCmd.Flags().StringP(utils.PortfolioIdFlag, "", "", "Portfolio ID. Uses environment variable if blank")
	createCryptoAddressCmd.Flags().StringP(utils.FormatFlag, "z", "false", "Pass true for formatted JSON. Default is false")

	createCryptoAddressCmd.MarkFlagRequired(utils.AssetIdFlag)
	createCryptoAddressCmd.MarkFlagRequired(utils.NetworkArnIdFlag)
}
