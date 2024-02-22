/**
 * Copyright 2024-present Coinbase Global, Inc.
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

var getOrderDetailsCmd = &cobra.Command{
	Use:   "get-order-details",
	Short: "Get details for order.",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := utils.GetClientFromEnv()
		if err != nil {
			return fmt.Errorf("failed to initialize client: %w", err)
		}

		portfolioId, err := utils.GetPortfolioId(cmd, client)
		if err != nil {
			return err
		}

		ctx, cancel := utils.GetContextWithTimeout()
		defer cancel()

		request := &intx.GetOrderDetailsRequest{
			PortfolioId: portfolioId,
			OrderId:     utils.GetFlagStringValue(cmd, utils.OrderIdFlag),
		}

		response, err := client.GetOrderDetails(ctx, request)
		if err != nil {
			return fmt.Errorf("cannot get order details: %w", err)
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
	rootCmd.AddCommand(getOrderDetailsCmd)

	getOrderDetailsCmd.Flags().StringP(utils.OrderIdFlag, "i", "", "Order ID (Required)")
	getOrderDetailsCmd.Flags().StringP(utils.FormatFlag, "z", "false", "Pass true for formatted JSON. Default is false")

	getOrderDetailsCmd.MarkFlagRequired(utils.OrderIdFlag)
}
