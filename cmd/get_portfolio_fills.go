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

		if request.RefDatetime != "" || request.ResultLimit != utils.ZeroInt || request.ResultOffset != utils.ZeroInt {
			request.Pagination = &intx.PaginationParams{
				RefDatetime:  request.RefDatetime,
				ResultLimit:  request.ResultLimit,
				ResultOffset: request.ResultOffset,
			}
		}

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

	getPortfolioFillsCmd.Flags().StringP("portfolio-id", "p", "", "Portfolio ID (required). Uses environment variable if blank")
	getPortfolioFillsCmd.Flags().String("order-id", "", "Order ID")
	getPortfolioFillsCmd.Flags().String("client-order-id", "", "Client Order ID")
	getPortfolioFillsCmd.Flags().String("ref-datetime", "", "Reference datetime for the request")
	getPortfolioFillsCmd.Flags().Int("result-limit", utils.ZeroInt, "Result limit")
	getPortfolioFillsCmd.Flags().Int("result-offset", utils.ZeroInt, "Result offset")
	getPortfolioFillsCmd.Flags().String("time-from", "", "Time from which to get fills")
	getPortfolioFillsCmd.Flags().StringP(utils.FormatFlag, "z", "false", "Pass true for formatted JSON. Default is false")

	getPortfolioFillsCmd.MarkFlagRequired("portfolio-id")
}