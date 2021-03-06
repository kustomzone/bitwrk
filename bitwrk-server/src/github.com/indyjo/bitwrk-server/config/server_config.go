//  BitWrk - A Bitcoin-friendly, anonymous marketplace for computing power
//  Copyright (C) 2013-2017  Jonas Eschenburg <jonas@bitwrk.net>
//
//  This program is free software: you can redistribute it and/or modify
//  it under the terms of the GNU General Public License as published by
//  the Free Software Foundation, either version 3 of the License, or
//  (at your option) any later version.
//
//  This program is distributed in the hope that it will be useful,
//  but WITHOUT ANY WARRANTY; without even the implied warranty of
//  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//  GNU General Public License for more details.
//
//  You should have received a copy of the GNU General Public License
//  along with this program.  If not, see <http://www.gnu.org/licenses/>.

// Package config contains settings that influence run-time behavior of the BitWrk server.
package config

const CfgBitcoinNetworkId byte = 0
const CfgRequireValidNonce = true
const CfgRequireValidSignature = true
const CfgRequireValidWorkerURL = true

// Account ID that is trusted when receiving a deposit
const CfgTrustedAccount = "1TrsjuCvBch1D9h6nRkadGKakv9KyaiP6"
