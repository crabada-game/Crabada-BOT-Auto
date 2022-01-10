package service

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

func DecodeString(hexStr string) ([]byte, error) {
	cleaned := strings.Replace(hexStr, "0x", "", -1)
	return hex.DecodeString(cleaned)
}

func StringCompare(str1, str2 string) bool {
	str1 = strings.TrimSpace(str1)
	str2 = strings.TrimSpace(str2)
	return strings.ToLower(str1) == strings.ToLower(str2)
}

func StringContains(str1, str2 string) bool {
	return strings.Contains(strings.ToLower(str1), strings.ToLower(str2))
}

func Hex2Int64(hexStr string) int64 {
	// remove 0x suffix if found in the input string
	cleaned := strings.Replace(hexStr, "0x", "", -1)

	// base 16 for hexadecimal
	result, _ := strconv.ParseInt(cleaned, 16, 64)
	return result
}

func Hex2BigInt(hexStr string) *big.Int {
	cleaned := strings.Replace(hexStr, "0x", "", -1)
	result := new(big.Int)
	result.SetString(cleaned, 16)
	return result
}

func Remove0x(hexStr string) string {
	// remove 0x suffix if found in the input string
	return strings.Replace(hexStr, "0x", "", -1)
}

func HexString2Address(hexStr string) string {
	return "0x" + hexStr[24:]
}

const zeroHex = "0000000000000000000000000000000000000000000000000000000000000000"

func ToHexString64Chars(hexStr string) string {
	cleaned := strings.Replace(hexStr, "0x", "", -1)
	return zeroHex[:64-len(cleaned)] + cleaned
}

func String2BigInt(str string) *big.Int {
	result := new(big.Int)
	result, ok := result.SetString(str, 10)
	if !ok {
		fmt.Println("SetString: error")
		return nil
	}
	return result
}

func BigInt2HexString(n *big.Int) string {
	return fmt.Sprintf("%x", n)
}

func Int642HexString(n int64) string {
	return fmt.Sprintf("%x", n)
}
