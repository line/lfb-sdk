package keeper

import (
	sdk "github.com/line/lbm-sdk/types"
	"github.com/line/lbm-sdk/x/token"
)

var (
	balanceKeyPrefix       = []byte{0x00}
	classKeyPrefix         = []byte{0x01}
	grantKeyPrefix         = []byte{0x02}
	authorizationKeyPrefix = []byte{0x03}

	// statistics keys
	supplyKeyPrefix = []byte{0x04}
	mintKeyPrefix   = []byte{0x05}
	burnKeyPrefix   = []byte{0x06}
)

func classKey(id string) []byte {
	key := make([]byte, len(classKeyPrefix)+len(id))
	copy(key, classKeyPrefix)
	copy(key[len(classKeyPrefix):], id)
	return key
}

func balanceKey(classID string, address sdk.AccAddress) []byte {
	prefix := balanceKeyPrefixByContractID(classID)
	key := make([]byte, len(prefix)+len(address))

	copy(key, prefix)
	copy(key[len(prefix):], address)

	return key
}

func balanceKeyPrefixByContractID(classID string) []byte {
	key := make([]byte, len(balanceKeyPrefix)+1+len(classID))

	begin := 0
	copy(key, balanceKeyPrefix)

	begin += len(balanceKeyPrefix)
	key[begin] = byte(len(classID))

	begin++
	copy(key[begin:], classID)

	return key
}

func splitBalanceKey(key []byte) (classID string, address sdk.AccAddress) {
	begin := len(balanceKeyPrefix) + 1
	end := begin + int(key[begin-1])
	classID = string(key[begin:end])

	begin = end
	address = sdk.AccAddress(key[begin:])

	return
}

func statisticsKey(keyPrefix []byte, classID string) []byte {
	key := make([]byte, len(keyPrefix)+len(classID))
	copy(key, keyPrefix)
	copy(key[len(keyPrefix):], classID)
	return key
}

// func supplyKey(classID string) []byte {
// 	return statisticsKey(supplyKeyPrefix, classID)
// }

// func mintKey(classID string) []byte {
// 	return statisticsKey(mintKeyPrefix, classID)
// }

// func burnKey(classID string) []byte {
// 	return statisticsKey(burnKeyPrefix, classID)
// }

func splitStatisticsKey(key, keyPrefix []byte) (classID string) {
	return string(key[len(keyPrefix):])
}

// func splitSupplyKey(key []byte) (classID string) {
// 	return splitStatisticsKey(key, supplyKeyPrefix)
// }

// func splitMintKey(key []byte) (classID string) {
// 	return splitStatisticsKey(key, mintKeyPrefix)
// }

// func splitBurnKey(key []byte) (classID string) {
// 	return splitStatisticsKey(key, burnKeyPrefix)
// }

func grantKey(classID string, grantee sdk.AccAddress, permission token.Permission) []byte {
	prefix := grantKeyPrefixByGrantee(classID, grantee)
	key := make([]byte, len(prefix)+1)

	copy(key, prefix)
	key[len(prefix)] = byte(permission)

	return key
}

func grantKeyPrefixByGrantee(classID string, grantee sdk.AccAddress) []byte {
	prefix := grantKeyPrefixByContractID(classID)
	key := make([]byte, len(prefix)+1+len(grantee))

	begin := 0
	copy(key, prefix)

	begin += len(prefix)
	key[begin] = byte(len(grantee))

	begin++
	copy(key[begin:], grantee)

	return key
}

func grantKeyPrefixByContractID(classID string) []byte {
	key := make([]byte, len(grantKeyPrefix)+1+len(classID))

	begin := 0
	copy(key, grantKeyPrefix)

	begin += len(grantKeyPrefix)
	key[begin] = byte(len(classID))

	begin++
	copy(key[begin:], classID)

	return key
}

func splitGrantKey(key []byte) (classID string, grantee sdk.AccAddress, permission token.Permission) {
	begin := len(grantKeyPrefix) + 1
	end := begin + int(key[begin-1])
	classID = string(key[begin:end])

	begin = end + 1
	end = begin + int(key[begin-1])
	grantee = sdk.AccAddress(key[begin:end])

	begin = end
	permission = token.Permission(key[begin])

	return
}

func authorizationKey(classID string, operator, holder sdk.AccAddress) []byte {
	prefix := authorizationKeyPrefixByOperator(classID, operator)
	key := make([]byte, len(prefix)+len(holder))

	copy(key, prefix)
	copy(key[len(prefix):], holder)

	return key
}

func authorizationKeyPrefixByOperator(classID string, operator sdk.AccAddress) []byte {
	prefix := authorizationKeyPrefixByContractID(classID)
	key := make([]byte, len(prefix)+1+len(operator))

	begin := 0
	copy(key, prefix)

	begin += len(prefix)
	key[begin] = byte(len(operator))

	begin++
	copy(key[begin:], operator)

	return key
}

func authorizationKeyPrefixByContractID(classID string) []byte {
	key := make([]byte, len(authorizationKeyPrefix)+1+len(classID))

	begin := 0
	copy(key, authorizationKeyPrefix)

	begin += len(authorizationKeyPrefix)
	key[begin] = byte(len(classID))

	begin++
	copy(key[begin:], classID)

	return key
}

func splitAuthorizationKey(key []byte) (classID string, operator, holder sdk.AccAddress) {
	begin := len(authorizationKeyPrefix) + 1
	end := begin + int(key[begin-1])
	classID = string(key[begin:end])

	begin = end + 1
	end = begin + int(key[begin-1])
	operator = sdk.AccAddress(key[begin:end])

	begin = end
	holder = sdk.AccAddress(key[begin:])

	return
}
