// This file contains constants, DO NOT MODIFY.

package core

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

var EpochLength uint64 = 1200
var NumberOfStates uint64 = 5
var ChainId = big.NewInt(0x109B4597)
var StateLength = EpochLength / NumberOfStates
var MaxRetries uint = 8
var NilHash = common.Hash{0x00}
var BlockCompletionTimeout = 30

//Following are the default config values for all the config parameters

var DefaultGasMultiplier = 1.0
var DefaultBufferPercent = 20
var DefaultGasPrice = 1
var DefaultWaitTime = 1
var DefaultGasLimit = 2
var DefaultGasLimitOverride = 0
var DefaultRPCTimeout = 10
var DefaultHTTPTimeout = 10
var DefaultLogLevel = ""

//Following are the default logFile parameters in config

var DefaultLogFileMaxSize = 200
var DefaultLogFileMaxBackups = 52
var DefaultLogFileMaxAge = 365

//DisputeGasMultiplier is a constant gasLimitMultiplier to increase gas Limit for function `disputeCollectionIdShouldBeAbsent` and `disputeCollectionIdShouldBePresent`
var DisputeGasMultiplier float32 = 5.5

// Following are the constants which will be used to derive different file paths

var DataFileDirectory = "data_files"
var CommitDataFile = "_commitData.json"
var ProposeDataFile = "_proposeData.json"
var DisputeDataFile = "_disputeData.json"
var AssetsDataFile = "assets.json"
var ConfigFile = "razor.yaml"
var LogFileDirectory = "logs"
var DefaultPathName = ".razor"

//BlockNumberInterval is the interval in seconds after which blockNumber needs to be calculated again
var BlockNumberInterval = 5

//SwitchClientDuration is the time after which alternate client from secondary RPC will be switched back to client from primary RPC
var SwitchClientDuration = 5 * EpochLength
