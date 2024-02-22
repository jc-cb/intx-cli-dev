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

package utils

const (
	OrderIdFlag       = "order-id"
	ProductIdFlag     = "product-id"
	GenericIdFlag     = "id"
	TransactionIdFlag = "transaction-id"
	WalletIdFlag      = "wallet-id"
	InstrumentIdFlag  = "instrument-id"

	AssetIdFlag      = "asset-id"
	NetworkArnIdFlag = "network-arn-id"

	AddressFlag      = "address"
	SymbolFlag       = "symbol"
	NameFlag         = "name"
	AccountIdFlag    = "account-identifier"
	TypeFlag         = "type"
	SideFlag         = "side"
	SizeFlag         = "size"
	TifFlag          = "tif"
	BaseQuantityFlag = "base-quantity"
	QuoteValueFlag   = "quote-value"
	LimitPriceFlag   = "limit-price"
	StopPriceFlag    = "stop-price"
	StartTimeFlag    = "start-time"
	ExpiryTimeFlag   = "expiry-time"
	TimeInForceFlag  = "time-in-force"
	UserIdFlag       = "user-id"
	StpModeFlag      = "stp-mode"
	PostOnlyFlag     = "post-only"
	RefDatetimeFlag  = "ref-datetime"
	ResultLimitFlag  = "result-limit"
	ResultOffsetFlag = "result-offset"
	TimeFromFlag     = "time-from"

	FromFlag = "from"
	ToFlag   = "to"

	SourceWalletIdFlag       = "source-wallet-id"
	SourceSymbolFlag         = "source-symbol"
	DestinationWalletIdFlag  = "destination-wallet-id"
	DestinationSymbolFlag    = "destination-symbol"
	AmountFlag               = "amount"
	DestinationTypeFlag      = "destination-type"
	DepositTypeFlag          = "deposit-type"
	NonceFlag                = "nonce"
	CounterpartyIdFlag       = "counterparty-id"
	AddNetworkFeeToTotalFlag = "add-network-fee-to-total"

	AllocationIdFlag             = "allocation-id"
	SizeTypeFlag                 = "size-type"
	SourcePortfolioIdFlag        = "source-portfolio-id"
	RemainderDestPortfolioIdFlag = "remainder-destination-portfolio-id"
	OrderIdsFlag                 = "order-ids"
	AllocationLegsFlag           = "allocation-legs"

	PaymentMethodIdFlag   = "payment-method-id"
	BlockchainAddressFlag = "blockchain-address"
	AccountIdentifierFlag = "account-identifier"

	JsonIndent        = "  "
	SearchFlag        = "search"
	CursorFlag        = "cursor"
	LimitFlag         = "limit"
	SortDirectionFlag = "sort-direction"

	SymbolsFlag       = "symbols"
	CategoriesFlag    = "categories"
	StatusesFlag      = "statuses"
	StartFlag         = "start"
	EndFlag           = "end"
	ProductIdsFlag    = "product-ids"
	OrderStatusesFlag = "order-statuses"
	OrderTypeFlag     = "order-type"
	OrderSideFlag     = "order-side"
	TypesFlag         = "types"

	OrderStatusOpen = "OPEN"
	OrderTypeMarket = "MARKET"
	OrderTypeLimit  = "LIMIT"
	OrderTypeTwap   = "TWAP"
	OrderTypeVwap   = "VWAP"
	OrderSideBuy    = "BUY"
	OrderSideSell   = "SELL"

	InvoiceStatesFlag   = "states"
	InvoiceBillingYear  = "billing-year"
	InvoiceBillingMonth = "billing-month"

	TifFillOrKill         = "FILL_OR_KILL"
	TifGoodUntilDateTime  = "GOOD_UNTIL_DATE_TIME"
	TifGoodUntilCancelled = "GOOD_UNTIL_CANCELLED"
	TifImmediateOrCancel  = "IMMEDIATE_OR_CANCEL"

	FormatFlag = "format"
	ToggleFlag = "toggle"

	PortfolioIdFlag    = "portfolio-id"
	IdempotencyKeyFlag = "idempotency-key"
	ClientOrderIdFlag  = "client-order-id"

	LimitDefault         = "100"
	SortDirectionDefault = "DESC"

	ZeroInt = 0
)
