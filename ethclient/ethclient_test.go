// Copyright 2016 The go-onlinegamechain Authors
// This file is part of the go-onlinegamechain library.
//
// The go-onlinegamechain library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-onlinegamechain library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-onlinegamechain library. If not, see <http://www.gnu.org/licenses/>.

package ethclient

import "github.com/onlinegamechain/go-onlinegamechain"

// Verify that Client implements the onlinegamechain interfaces.
var (
	_ = onlinegamechain.ChainReader(&Client{})
	_ = onlinegamechain.TransactionReader(&Client{})
	_ = onlinegamechain.ChainStateReader(&Client{})
	_ = onlinegamechain.ChainSyncReader(&Client{})
	_ = onlinegamechain.ContractCaller(&Client{})
	_ = onlinegamechain.GasEstimator(&Client{})
	_ = onlinegamechain.GasPricer(&Client{})
	_ = onlinegamechain.LogFilterer(&Client{})
	_ = onlinegamechain.PendingStateReader(&Client{})
	// _ = onlinegamechain.PendingStateEventer(&Client{})
	_ = onlinegamechain.PendingContractCaller(&Client{})
)
