package main

import (
	"encoding/hex"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"os"
	"strconv"

	twos "github.com/ElrondNetwork/big-int-util/twos-complement"
)

const numberOfDataSets = 5000

var positiveEncodedBigFloatPrefix = [...]byte{1, 10, 0, 0, 0, 53}
var negativeEncodedBigFloatPrefix = [...]byte{1, 11, 0, 0, 0, 53}
var positiveEncodedBigFloatForPowPrefix = [...]byte{1, 10, 0, 0, 53, 0, 0, 0}
var negativeEncodedBigFloatForPowPrefix = [...]byte{1, 11, 0, 0, 53, 0, 0, 0}

const copyBigFloatGasCost = 18000

var maxValForExponent = big.NewInt(0).SetInt64(math.MaxInt8)
var minValForExponent = big.NewInt(0).SetInt64(math.MinInt8)

// contract endpoint names
const (
	bfToManagedBufferName = "BigFloatToManagedBufferTest"
	bfFromPartsName       = "BigFloatNewFromPartsTest"
	bfFromFracName        = "BigFloatNewFromFracTest"
	bfFromSciName         = "BigFloatNewFromSciTest"
	bfAddName             = "BigFloatAddTest"
	bfSubName             = "BigFloatSubTest"
	bfMulName             = "BigFloatMulTest"
	bfDivName             = "BigFloatDivTest"
	bfTruncName           = "BigFloatTruncateTest"
	bfAbsName             = "BigFloatAbsTest"
	bfNegName             = "BigFloatNegTest"
	bfCmpName             = "BigFloatCmpTest"
	bfSignName            = "BigFloatSignTest"
	bfCloneName           = "BigFloatCloneTest"
	bfSqrtName            = "BigFloatSqrtTest"
	bfPowName             = "BigFloatPowTest"
	bfFloorName           = "BigFloatFloorTest"
	bfCeilName            = "BigFloatCeilTest"
	bfIsIntName           = "BigFloatIsIntTest"
	bfSetInt64Name        = "BigFloatSetInt64Test"
	bfSetIntName          = "BigFloatSetIntTest"
	bfGetPiName           = "BigFloatGetConstPiTest"
	bfGetEName            = "BigFloatGetConstETest"
)

const (
	bfNewFromPartsGasCost = 3000
	bfAddGasCost          = 10000
	bfSubGasCost          = 10000
	bfMulGasCost          = 10000
	bfDivGasCost          = 10000
	bfTruncateGasCost     = 7000
	bfNegGasCost          = 7000
	bfCloneGasCost        = 7000
	bfCmpGasCost          = 4500
	bfAbsGasCost          = 7000
	bfSqrtGasCost         = 10000
	bfPowGasCost          = 10000
	bfFloorGasCost        = 7000
	bfCeilGasCost         = 7000
	bfIsIntGasCost        = 4000
	bfSetBigIntGasCost    = 3000
	bfSetInt64GasCost     = 1000
	bfGetConstGasCost     = 1000
)

//POW AND SETBIGINT

func main() {
	generateBigFloatFromPartsData()
	generateBigFloatsFromFracData()
	generateBigFloatsFromSciData()
	generateDataForEndpoint(1, bfToManagedBufferName, 55000)
	generateDataForEndpoint(2, bfAddName, 55000)
	generateDataForEndpoint(2, bfSubName, 55000)
	generateDataForEndpoint(2, bfMulName, 55000)
	generateDataForEndpoint(2, bfDivName, 55000)
	generateDataForEndpoint(1, bfTruncName, 30000)
	generateDataForEndpoint(1, bfAbsName, 30000)
	generateDataForEndpoint(1, bfNegName, 30000)
	generateDataForEndpoint(2, bfCmpName, 55000)
	generateDataForEndpoint(1, bfSignName, 30000)
	generateDataForEndpoint(1, bfCloneName, 30000)
	generateDataForEndpoint(1, bfSqrtName, 30000)
	generateDataForBigFloatPow()
	generateDataForEndpoint(1, bfFloorName, 30000)
	generateDataForEndpoint(1, bfCeilName, 30000)
	generateDataForEndpoint(1, bfIsIntName, 30000)
	generateDataForEndpoint(1, bfSetInt64Name, 30000)

}

func generateBigFloatFromPartsData() {
	file, _ := os.Create(bfFromPartsName + ".data")
	for i := 0; i < numberOfDataSets; i++ {
		file.WriteString("BigFloatNewFromPartsTest@")

		// integralPart
		integralPart := rand.Intn(math.MaxInt32 - 1)
		bigIntegralPart := big.NewInt(0).SetInt64(int64(integralPart))
		if rand.Intn(2) == 1 {
			bigIntegralPart.Neg(bigIntegralPart)
		}
		hexEncodedIntegralPart := hex.EncodeToString(twos.ToBytes(bigIntegralPart))
		file.WriteString(hexEncodedIntegralPart + "@")
		// fractionalPart
		fractionalPart := rand.Intn(math.MaxInt32 - 1)
		bigFractionalPart := big.NewInt(0).SetInt64(int64(fractionalPart))
		hexEncodedFractionalPart := hex.EncodeToString(twos.ToBytes(bigFractionalPart))
		file.WriteString(hexEncodedFractionalPart + "@")
		// exponent
		exponent := rand.Intn(400)
		bigExponent := big.NewInt(0).SetInt64(int64(exponent))
		validExponent := big.NewInt(0).Neg(bigExponent)
		hexEncodedExponent := hex.EncodeToString(validExponent.Bytes())
		file.WriteString(hexEncodedExponent + ":30000" + "\n")

	}
	defer file.Close()
}

func generateBigFloatsFromFracData() {
	file, _ := os.Create(bfFromFracName + ".data")

	for i := 0; i < numberOfDataSets; i++ {
		file.WriteString("BigFloatNewFromFracTest@")

		// numerator
		numeratorPart := rand.Intn(math.MaxInt64 - 1)
		bigNumeratorPart := big.NewInt(0).SetInt64(int64(numeratorPart))
		if rand.Intn(2) == 1 {
			bigNumeratorPart.Neg(bigNumeratorPart)
		}
		hexEncodedNumerator := hex.EncodeToString(bigNumeratorPart.Bytes())
		file.WriteString(hexEncodedNumerator + "@")
		// denominator
		denominatorPart := rand.Intn(math.MaxInt64 - 1)
		bigDenominatorPart := big.NewInt(0).SetInt64(int64(denominatorPart))
		hexEncodedDenominator := hex.EncodeToString(bigDenominatorPart.Bytes())
		file.WriteString(hexEncodedDenominator + ":30000" + "\n")

	}
	defer file.Close()
}

func generateBigFloatsFromSciData() {
	file, _ := os.Create(bfFromSciName + ".data")

	for i := 0; i < numberOfDataSets; i++ {
		file.WriteString("BigFloatNewFromSciTest@")

		// significand
		significandPart := rand.Intn(math.MaxInt64 - 1)
		bigSignificandPart := big.NewInt(0).SetInt64(int64(significandPart))
		if rand.Intn(2) == 1 {
			bigSignificandPart.Neg(bigSignificandPart)
		}
		hexEncodedSignificand := hex.EncodeToString(bigSignificandPart.Bytes())
		file.WriteString(hexEncodedSignificand + "@")
		// exponent
		exponentPart := rand.Intn(400)
		bigExponentPart := big.NewInt(0).SetInt64(int64(exponentPart))
		hexEncodedExponent := hex.EncodeToString(bigExponentPart.Bytes())
		file.WriteString(hexEncodedExponent + ":30000" + "\n")
	}

	defer file.Close()
}

func generateDataForBigFloatPow() {
	file, _ := os.Create(bfPowName + ".data")
	for i := 0; i < numberOfDataSets/10; i++ {
		file.WriteString("BigFloatPowTest@")
		hexEncodedFloat := generateHexEncodedBigFloatForPow()
		file.WriteString(hexEncodedFloat + "@")

		bytes, _ := hex.DecodeString(hexEncodedFloat)
		floatVal := big.NewFloat(0)
		_ = floatVal.GobDecode(bytes)
		intOp := big.NewInt(0)
		floatVal.Int(intOp)

		//exponent
		exponentBytes := make([]byte, 1)
		rand.Read(exponentBytes)
		bigExponent := big.NewInt(0).SetBytes(exponentBytes)
		if rand.Intn(2) == 1 {
			bigExponent.Neg(bigExponent)
		}
		exponentBytes, _ = twos.ToBytesOfLength(bigExponent, 4)

		//gas cost
		lengthOfResult := big.NewInt(0).Div(big.NewInt(0).Mul(bigExponent, big.NewInt(int64(intOp.BitLen()))), big.NewInt(8))
		gasForPow := big.NewInt(0).Mul(lengthOfResult, big.NewInt(1000))
		//fmt.Println(gasForPow)
		gasToUse := uint64(math.MaxUint64)
		if gasForPow.Cmp(big.NewInt(0).SetUint64(math.MaxUint64)) < 0 {
			gasToUse = gasForPow.Uint64() + 10000
		}
		hexEncodedExponent := hex.EncodeToString(exponentBytes)
		file.WriteString(hexEncodedExponent)
		file.WriteString(":" + strconv.Itoa(int(gasToUse)) + "\n")
	}
	defer file.Close()
}

func generateDataForEndpoint(numberOfBigFloats int, endpointName string, gasLimit int) {
	fileName := fmt.Sprintf("%s.data", endpointName)
	file, _ := os.Create(fileName)

	for i := 0; i < numberOfDataSets; i++ {
		file.WriteString(endpointName)
		for j := 0; j < numberOfBigFloats; j++ {
			bigFloatValue := generateHexEncodedBigFloat()
			file.WriteString("@" + bigFloatValue)
		}
		file.WriteString(":" + strconv.Itoa(gasLimit) + "\n")
	}
	defer file.Close()
}

func generateHexEncodedBigFloat() string {
	encodedBigFloat := make([]byte, 0)
	if rand.Intn(2) == 1 {
		encodedBigFloat = append(encodedBigFloat, positiveEncodedBigFloatPrefix[:]...)
	} else {
		encodedBigFloat = append(encodedBigFloat, negativeEncodedBigFloatPrefix[:]...)
	}
	randomExponentAndMantissa := make([]byte, 12)
	rand.Read(randomExponentAndMantissa)
	encodedBigFloat = append(encodedBigFloat, randomExponentAndMantissa...)
	hexEncodedBigFloat := hex.EncodeToString(encodedBigFloat)
	return hexEncodedBigFloat
}

func generateHexEncodedBigFloatForPow() string {
	encodedBigFloat := make([]byte, 0)
	if rand.Intn(2) == 1 {
		encodedBigFloat = append(encodedBigFloat, positiveEncodedBigFloatForPowPrefix[:]...)
	} else {
		encodedBigFloat = append(encodedBigFloat, negativeEncodedBigFloatForPowPrefix[:]...)
	}
	randomExponentAndMantissa := make([]byte, 9)
	rand.Read(randomExponentAndMantissa)
	encodedBigFloat = append(encodedBigFloat, randomExponentAndMantissa...)
	hexEncodedBigFloat := hex.EncodeToString(encodedBigFloat)
	return hexEncodedBigFloat
}
