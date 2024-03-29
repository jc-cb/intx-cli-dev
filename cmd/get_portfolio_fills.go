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

var getPortfolioFillsCmd = &cobra.Command{
	Use:   "get-portfolio-fills",
	Short: "Get portfolio fills.",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := utils.GetClientFromEnv()
		if err != nil {
			return fmt.Errorf("failed to initialize client: %w", err)
		}

		portfolioId, err := utils.GetPortfolioId(cmd, client)
		if err != nil {
			return err
		}

		resultLimit, _ := utils.GetFlagIntValue(cmd, utils.ResultLimitFlag)
		resultOffset, _ := utils.GetFlagIntValue(cmd, utils.ResultOffsetFlag)

		ctx, cancel := utils.GetContextWithTimeout()
		defer cancel()

		request := &intx.GetPortfolioFillsRequest{
			PortfolioId:   portfolioId,
			OrderId:       utils.GetFlagStringValue(cmd, utils.OrderIdFlag),
			ClientOrderId: utils.GetFlagStringValue(cmd, utils.ClientOrderIdFlag),
			RefDatetime:   utils.GetFlagStringValue(cmd, utils.RefDatetimeFlag),
			ResultLimit:   resultLimit,
			ResultOffset:  resultOffset,
			TimeFrom:      utils.GetFlagStringValue(cmd, utils.TimeFromFlag),
		}

		request.Pagination = utils.CreatePaginationParams(request.RefDatetime, request.ResultLimit, request.ResultOffset)

		response, err := client.GetPortfolioFills(ctx, request)
		if err != nil {
			return fmt.Errorf("cannot get portfolio fills: %w", err)
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
	rootCmd.AddCommand(getPortfolioFillsCmd)

	getPortfolioFillsCmd.Flags().StringP(utils.PortfolioIdFlag, "p", "", "Portfolio ID (required). Uses environment variable if blank")
	getPortfolioFillsCmd.Flags().String(utils.OrderIdFlag, "", "Order ID")
	getPortfolioFillsCmd.Flags().String(utils.ClientOrderIdFlag, "", "Client Order ID")
	getPortfolioFillsCmd.Flags().String(utils.RefDatetimeFlag, "", "Reference datetime for the request")
	getPortfolioFillsCmd.Flags().Int(utils.ResultLimitFlag, utils.ZeroInt, "Result limit")
	getPortfolioFillsCmd.Flags().Int(utils.ResultOffsetFlag, utils.ZeroInt, "Result offset")
	getPortfolioFillsCmd.Flags().String(utils.TimeFromFlag, "", "Time from which to get fills")
	getPortfolioFillsCmd.Flags().StringP(utils.FormatFlag, "z", "false", "Pass true for formatted JSON. Default is false")

	getPortfolioFillsCmd.MarkFlagRequired("portfolio-id")
}
