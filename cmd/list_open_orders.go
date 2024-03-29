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

var listOpenOrdersCmd = &cobra.Command{
	Use:   "list-open-orders",
	Short: "List open orders.",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := utils.GetClientFromEnv()
		if err != nil {
			return fmt.Errorf("cannot get client from environment: %w", err)
		}

		portfolioId, err := utils.GetPortfolioId(cmd, client)
		if err != nil {
			return fmt.Errorf("failed to get portfolio ID: %w", err)
		}

		resultLimit, _ := utils.GetFlagIntValue(cmd, utils.ResultLimitFlag)
		resultOffset, _ := utils.GetFlagIntValue(cmd, utils.ResultOffsetFlag)

		ctx, cancel := utils.GetContextWithTimeout()
		defer cancel()

		request := &intx.ListOpenOrdersRequest{
			PortfolioId:   portfolioId,
			InstrumentId:  utils.GetFlagStringValue(cmd, utils.InstrumentIdFlag),
			ClientOrderId: utils.GetFlagStringValue(cmd, utils.ClientOrderIdFlag),
			EventType:     utils.GetFlagStringValue(cmd, utils.EventTypeFlag),
			RefDatetime:   utils.GetFlagStringValue(cmd, utils.RefDatetimeFlag),
			ResultLimit:   resultLimit,
			ResultOffset:  resultOffset,
		}

		request.Pagination = utils.CreatePaginationParams(request.RefDatetime, request.ResultLimit, request.ResultOffset)

		response, err := client.ListOpenOrders(ctx, request)
		if err != nil {
			return fmt.Errorf("cannot list open orders: %w", err)
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
	rootCmd.AddCommand(listOpenOrdersCmd)

	listOpenOrdersCmd.Flags().StringP(utils.PortfolioIdFlag, "", "", "Portfolio ID(s). Uses environment variable if blank")
	listOpenOrdersCmd.Flags().StringP(utils.InstrumentIdFlag, "i", "", "Filter open orders by instrument ID")
	listOpenOrdersCmd.Flags().StringP(utils.ClientOrderIdFlag, "c", "", "Filter open orders by client order ID")
	listOpenOrdersCmd.Flags().StringP(utils.EventTypeFlag, "e", "", "Filter open orders by event type")
	listOpenOrdersCmd.Flags().StringP(utils.RefDatetimeFlag, "", "", "Filter open orders by datetime")
	listOpenOrdersCmd.Flags().Int(utils.ResultLimitFlag, 0, "Limit the number of open orders returned")
	listOpenOrdersCmd.Flags().Int(utils.ResultOffsetFlag, 0, "Offset for the list of open orders returned")
	listOpenOrdersCmd.Flags().StringP(utils.TimeFromFlag, "", "", "Filter open orders from this time")
	listOpenOrdersCmd.Flags().StringP(utils.FormatFlag, "z", "false", "Pass true for formatted JSON. Default is false")
}
