// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	big "math/big"
	bindings "razor/pkg/bindings"

	bind "github.com/ethereum/go-ethereum/accounts/abi/bind"

	ethclient "github.com/ethereum/go-ethereum/ethclient"

	mock "github.com/stretchr/testify/mock"
)

// BlockManagerUtils is an autogenerated mock type for the BlockManagerUtils type
type BlockManagerUtils struct {
	mock.Mock
}

// GetBlock provides a mock function with given fields: _a0, _a1, _a2
func (_m *BlockManagerUtils) GetBlock(_a0 *ethclient.Client, _a1 *bind.CallOpts, _a2 uint32) (bindings.StructsBlock, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 bindings.StructsBlock
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.CallOpts, uint32) bindings.StructsBlock); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(bindings.StructsBlock)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.CallOpts, uint32) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNumProposedBlocks provides a mock function with given fields: _a0, _a1, _a2
func (_m *BlockManagerUtils) GetNumProposedBlocks(_a0 *ethclient.Client, _a1 *bind.CallOpts, _a2 uint32) (uint8, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 uint8
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.CallOpts, uint32) uint8); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(uint8)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.CallOpts, uint32) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProposedBlock provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *BlockManagerUtils) GetProposedBlock(_a0 *ethclient.Client, _a1 *bind.CallOpts, _a2 uint32, _a3 uint32) (bindings.StructsBlock, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 bindings.StructsBlock
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.CallOpts, uint32, uint32) bindings.StructsBlock); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		r0 = ret.Get(0).(bindings.StructsBlock)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.CallOpts, uint32, uint32) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MaxAltBlocks provides a mock function with given fields: _a0, _a1
func (_m *BlockManagerUtils) MaxAltBlocks(_a0 *ethclient.Client, _a1 *bind.CallOpts) (uint8, error) {
	ret := _m.Called(_a0, _a1)

	var r0 uint8
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.CallOpts) uint8); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(uint8)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.CallOpts) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MinStake provides a mock function with given fields: _a0, _a1
func (_m *BlockManagerUtils) MinStake(_a0 *ethclient.Client, _a1 *bind.CallOpts) (*big.Int, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.CallOpts) *big.Int); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.CallOpts) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SortedProposedBlockIds provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *BlockManagerUtils) SortedProposedBlockIds(_a0 *ethclient.Client, _a1 *bind.CallOpts, _a2 uint32, _a3 *big.Int) (uint32, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 uint32
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.CallOpts, uint32, *big.Int) uint32); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.CallOpts, uint32, *big.Int) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}