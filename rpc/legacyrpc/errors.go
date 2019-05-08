// Copyright (c) 2013-2015 The btcsuite developers
// Copyright (c) 2019 The Soteria DAG developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package legacyrpc

import (
	"errors"

	"github.com/soteria-dag/soterd/soterjson"
)

// TODO(jrick): There are several error paths which 'replace' various errors
// with a more appropiate error from the soterjson package.  Create a map of
// these replacements so they can be handled once after an RPC handler has
// returned and before the error is marshaled.

// Error types to simplify the reporting of specific categories of
// errors, and their *soterjson.RPCError creation.
type (
	// DeserializationError describes a failed deserializaion due to bad
	// user input.  It corresponds to soterjson.ErrRPCDeserialization.
	DeserializationError struct {
		error
	}

	// InvalidParameterError describes an invalid parameter passed by
	// the user.  It corresponds to soterjson.ErrRPCInvalidParameter.
	InvalidParameterError struct {
		error
	}

	// ParseError describes a failed parse due to bad user input.  It
	// corresponds to soterjson.ErrRPCParse.
	ParseError struct {
		error
	}
)

// Errors variables that are defined once here to avoid duplication below.
var (
	ErrNeedPositiveAmount = InvalidParameterError{
		errors.New("amount must be positive"),
	}

	ErrNeedPositiveMinconf = InvalidParameterError{
		errors.New("minconf must be positive"),
	}

	ErrAddressNotInWallet = soterjson.RPCError{
		Code:    soterjson.ErrRPCWallet,
		Message: "address not found in wallet",
	}

	ErrAccountNameNotFound = soterjson.RPCError{
		Code:    soterjson.ErrRPCWalletInvalidAccountName,
		Message: "account name not found",
	}

	ErrUnloadedWallet = soterjson.RPCError{
		Code:    soterjson.ErrRPCWallet,
		Message: "Request requires a wallet but wallet has not loaded yet",
	}

	ErrWalletUnlockNeeded = soterjson.RPCError{
		Code:    soterjson.ErrRPCWalletUnlockNeeded,
		Message: "Enter the wallet passphrase with walletpassphrase first",
	}

	ErrNotImportedAccount = soterjson.RPCError{
		Code:    soterjson.ErrRPCWallet,
		Message: "imported addresses must belong to the imported account",
	}

	ErrNoTransactionInfo = soterjson.RPCError{
		Code:    soterjson.ErrRPCNoTxInfo,
		Message: "No information for transaction",
	}

	ErrReservedAccountName = soterjson.RPCError{
		Code:    soterjson.ErrRPCInvalidParameter,
		Message: "Account name is reserved by RPC server",
	}
)
