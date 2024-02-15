package types

import (
	"encoding/base64"
	"math"
	"math/big"
	"unicode/utf8"

	"github.com/iden3/go-iden3-crypto/poseidon"
	"github.com/pkg/errors"

	sdkerrors "github.com/Finschia/finschia-sdk/types/errors"
)

const MaxHeaderLen = 992
const MaxExtIssLen = 32
const PackWidth = 248

// circom constants from main.circom / https://zkrepl.dev/?gist=30d21c7a7285b1b14f608325f172417b
// template RSAGroupSigVerify(n, k, levels) {
// component main { public [ modulus ] } = RSAVerify(121, 17)
// component main { public [ root, payload1 ] } = RSAGroupSigVerify(121, 17, 30)

const CircomBigintN = 121
const CircomBigintK = 17

func SplitToTwoFields(ephPkBytes []byte) []*big.Int {
	mid := len(ephPkBytes) - 16
	first := new(big.Int).SetBytes(ephPkBytes[0:mid])
	second := new(big.Int).SetBytes(ephPkBytes[mid:])

	return []*big.Int{first, second}
}

func ToField(val string) (*big.Int, bool) {
	return new(big.Int).SetString(val, 10)
}

// ChunkArray splits an array into chunks of a specified size.
func ChunkArray(array []byte, chunkSize int) [][]byte {
	// The length of the array divided by the chunk size is rounded up and used as the number of chunks.
	chunkCount := int(math.Ceil(float64(len(array)) / float64(chunkSize)))
	chunks := make([][]byte, chunkCount)

	shouldFirstPack := chunkSize
	if len(array)%chunkSize != 0 {
		shouldFirstPack = len(array) % chunkSize
	}
	for i := 0; i < chunkCount; i++ {
		if i == 0 {
			chunks[i] = array[0:shouldFirstPack]
		} else {
			chunks[i] = array[shouldFirstPack+((i-1)*chunkSize) : shouldFirstPack+(i*chunkSize)]
		}
	}
	return chunks
}

// BytesBEToBigInt converts a slice of bytes to a big.Int.
func BytesBEToBigInt(bytes []byte) *big.Int {
	if len(bytes) == 0 {
		return big.NewInt(0)
	}

	return new(big.Int).SetBytes(bytes)
}

// HashASCIIStrToField hashes an ASCII string to a field element.
func HashASCIIStrToField(val string, maxSize int) (*big.Int, error) {
	bytes := []byte(val)
	if len(bytes) > maxSize {
		return nil, errors.New("String is longer than allowed size")
	}

	// Padding with zeroes
	strPadded := make([]byte, maxSize)
	copy(strPadded, bytes)

	const chunkSize = PackWidth / 8
	chunks := ChunkArray(strPadded, chunkSize)
	packed := make([]*big.Int, 0, len(chunks))
	for _, chunk := range chunks {
		packed = append(packed, BytesBEToBigInt(chunk))
	}

	hash, err := PoseidonHash(packed)
	if err != nil {
		return nil, err
	}

	return hash, nil
}

// CircomBigIntToChunkedBytes converts a big integer to a slice of chunked *big.Int.
func CircomBigIntToChunkedBytes(num *big.Int) []*big.Int {
	bytesPerChunk, numChunks := CircomBigintN, CircomBigintK

	res := make([]*big.Int, 0, numChunks)
	msk := new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), uint(bytesPerChunk)), big.NewInt(1))

	for i := 0; i < numChunks; i++ {
		chunk := new(big.Int).And(new(big.Int).Rsh(num, uint(i*bytesPerChunk)), msk)
		res = append(res, chunk)
	}

	return res
}

// CircomBigIntToField hashes a circom style big integer to a field element.
func CircomBigIntToField(num *big.Int) (*big.Int, error) {
	packed := CircomBigIntToChunkedBytes(num)
	hash, err := PoseidonHash(packed)
	if err != nil {
		return nil, err
	}
	return hash, nil
}

// PoseidonHash hashes field elements of 32 or less.
// poseidon.Hash only supports less than 16.
func PoseidonHash(inpBI []*big.Int) (*big.Int, error) {
	arrayCnt := len(inpBI)
	switch true {
	case arrayCnt <= 16:
		return poseidon.Hash(inpBI)
	case arrayCnt > 16 && arrayCnt <= 32:
		hash1, err := poseidon.Hash(inpBI[0:16])
		if err != nil {
			return nil, err
		}
		hash2, err := poseidon.Hash(inpBI[16:])
		if err != nil {
			return nil, err
		}
		return poseidon.Hash([]*big.Int{hash1, hash2})
	default:
		return nil, errors.Errorf("Yet to implement: unable to hash a vector of length %d", arrayCnt)
	}
}

func (zk *ZKAuthInputs) CalculateAllInputsHash(ephPkBytes, modulus []byte, maxBlockHeight int64) (*big.Int, error) {
	if utf8.RuneCountInString(zk.HeaderBase64) > MaxHeaderLen {
		return nil, sdkerrors.Wrap(ErrInvalidZkAuthInputs, "zkAuth header should be less then MAX_HEADER_LEN")
	}

	addressSeedF, ok := ToField(zk.AddressSeed)
	if !ok {
		return nil, sdkerrors.Wrap(ErrInvalidZkAuthInputs, "invalid address_seed")
	}
	ephPkFs := SplitToTwoFields(ephPkBytes)
	maxBlockHeightF := new(big.Int).SetInt64(maxBlockHeight)

	issBytes, err := base64.StdEncoding.DecodeString(zk.IssF)
	if err != nil {
		return nil, err
	}
	issF, err := HashASCIIStrToField(string(issBytes), MaxExtIssLen)
	if err != nil {
		return nil, sdkerrors.Wrap(ErrInvalidZkAuthInputs, "invalid Iss base64")
	}
	headerBytes, err := base64.StdEncoding.DecodeString(zk.HeaderBase64)
	if err != nil {
		return nil, err
	}
	headerF, err := HashASCIIStrToField(string(headerBytes), MaxHeaderLen)
	if err != nil {
		return nil, sdkerrors.Wrap(ErrInvalidZkAuthInputs, "invalid jwt header")
	}
	modulusF, err := CircomBigIntToField(new(big.Int).SetBytes(modulus))
	if err != nil {
		return nil, sdkerrors.Wrap(ErrInvalidZkAuthInputs, "invalid modulus")
	}

	return poseidon.Hash([]*big.Int{
		ephPkFs[0],
		ephPkFs[1],
		addressSeedF,
		maxBlockHeightF,
		issF,
		headerF,
		modulusF,
	})
}
