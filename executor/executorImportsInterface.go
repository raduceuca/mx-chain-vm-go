package elrondapi 

// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
// !!!!!!!!!!!!!!!!!!!!!! AUTO-GENERATED FILE !!!!!!!!!!!!!!!!!!!!!!
// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

type ImportsInterface interface {
	BigFloatNewFromParts(integralPart int32, fractionalPart int32, exponent int32) int32
	BigFloatNewFromFrac(numerator int64, denominator int64) int32
	BigFloatNewFromSci(significand int64, exponent int64) int32
	BigFloatAdd(destinationHandle int32, op1Handle int32, op2Handle int32)
	BigFloatSub(destinationHandle int32, op1Handle int32, op2Handle int32)
	BigFloatMul(destinationHandle int32, op1Handle int32, op2Handle int32)
	BigFloatDiv(destinationHandle int32, op1Handle int32, op2Handle int32)
	BigFloatNeg(destinationHandle int32, opHandle int32)
	BigFloatClone(destinationHandle int32, opHandle int32)
	BigFloatCmp(op1Handle int32, op2Handle int32) int32
	BigFloatAbs(destinationHandle int32, opHandle int32)
	BigFloatSign(opHandle int32) int32
	BigFloatSqrt(destinationHandle int32, opHandle int32)
	BigFloatPow(destinationHandle int32, opHandle int32, exponent int32)
	BigFloatFloor(destBigIntHandle int32, opHandle int32)
	BigFloatCeil(destBigIntHandle int32, opHandle int32)
	BigFloatTruncate(destBigIntHandle int32, opHandle int32)
	BigFloatSetInt64(destinationHandle int32, value int64)
	BigFloatIsInt(opHandle int32) int32
	BigFloatSetBigInt(destinationHandle int32, bigIntHandle int32)
	BigFloatGetConstPi(destinationHandle int32)
	BigFloatGetConstE(destinationHandle int32)
	BigIntGetUnsignedArgument(id int32, destinationHandle int32)
	BigIntGetSignedArgument(id int32, destinationHandle int32)
	BigIntStorageStoreUnsigned(keyOffset int32, keyLength int32, sourceHandle int32) int32
	BigIntStorageLoadUnsigned(keyOffset int32, keyLength int32, destinationHandle int32) int32
	BigIntGetCallValue(destinationHandle int32)
	BigIntGetESDTCallValue(destination int32)
	BigIntGetESDTCallValueByIndex(destinationHandle int32, index int32)
	BigIntGetExternalBalance(addressOffset int32, result int32)
	BigIntGetESDTExternalBalance(addressOffset int32, tokenIDOffset int32, tokenIDLen int32, nonce int64, resultHandle int32)
	BigIntNew(smallValue int64) int32
	BigIntUnsignedByteLength(referenceHandle int32) int32
	BigIntSignedByteLength(referenceHandle int32) int32
	BigIntGetUnsignedBytes(referenceHandle int32, byteOffset int32) int32
	BigIntGetSignedBytes(referenceHandle int32, byteOffset int32) int32
	BigIntSetUnsignedBytes(destinationHandle int32, byteOffset int32, byteLength int32)
	BigIntSetSignedBytes(destinationHandle int32, byteOffset int32, byteLength int32)
	BigIntIsInt64(destinationHandle int32) int32
	BigIntGetInt64(destinationHandle int32) int64
	BigIntSetInt64(destinationHandle int32, value int64)
	BigIntAdd(destinationHandle int32, op1Handle int32, op2Handle int32)
	BigIntSub(destinationHandle int32, op1Handle int32, op2Handle int32)
	BigIntMul(destinationHandle int32, op1Handle int32, op2Handle int32)
	BigIntTDiv(destinationHandle int32, op1Handle int32, op2Handle int32)
	BigIntTMod(destinationHandle int32, op1Handle int32, op2Handle int32)
	BigIntEDiv(destinationHandle int32, op1Handle int32, op2Handle int32)
	BigIntEMod(destinationHandle int32, op1Handle int32, op2Handle int32)
	BigIntSqrt(destinationHandle int32, opHandle int32)
	BigIntPow(destinationHandle int32, op1Handle int32, op2Handle int32)
	BigIntLog2(op1Handle int32) int32
	BigIntAbs(destinationHandle int32, opHandle int32)
	BigIntNeg(destinationHandle int32, opHandle int32)
	BigIntSign(opHandle int32) int32
	BigIntCmp(op1Handle int32, op2Handle int32) int32
	BigIntNot(destinationHandle int32, opHandle int32)
	BigIntAnd(destinationHandle int32, op1Handle int32, op2Handle int32)
	BigIntOr(destinationHandle int32, op1Handle int32, op2Handle int32)
	BigIntXor(destinationHandle int32, op1Handle int32, op2Handle int32)
	BigIntShr(destinationHandle int32, opHandle int32, bits int32)
	BigIntShl(destinationHandle int32, opHandle int32, bits int32)
	BigIntFinishUnsigned(referenceHandle int32)
	BigIntFinishSigned(referenceHandle int32)
	BigIntToString(bigIntHandle int32, destinationHandle int32)
	GetGasLeft() int64
	GetSCAddress(resultOffset int32)
	GetOwnerAddress(resultOffset int32)
	GetShardOfAddress(addressOffset int32) int32
	IsSmartContract(addressOffset int32) int32
	SignalError(messageOffset int32, messageLength int32)
	GetExternalBalance(addressOffset int32, resultOffset int32)
	BlockHash(nonce int64, resultOffset int32) int32
	GetESDTBalance(addressOffset int32, tokenIDOffset int32, tokenIDLen int32, nonce int64, resultOffset int32) int32
	GetESDTNFTNameLength(addressOffset int32, tokenIDOffset int32, tokenIDLen int32, nonce int64) int32
	GetESDTNFTAttributeLength(addressOffset int32, tokenIDOffset int32, tokenIDLen int32, nonce int64) int32
	GetESDTNFTURILength(addressOffset int32, tokenIDOffset int32, tokenIDLen int32, nonce int64) int32
	GetESDTTokenData(addressOffset int32, tokenIDOffset int32, tokenIDLen int32, nonce int64, valueHandle int32, propertiesOffset int32, hashOffset int32, nameOffset int32, attributesOffset int32, creatorOffset int32, royaltiesHandle int32, urisOffset int32) int32
	GetESDTLocalRoles(tokenIdHandle int32) int64
	ValidateTokenIdentifier(tokenIdHandle int32) int32
	TransferValue(destOffset int32, valueOffset int32, dataOffset int32, length int32) int32
	TransferValueExecute(destOffset int32, valueOffset int32, gasLimit int64, functionOffset int32, functionLength int32, numArguments int32, argumentsLengthOffset int32, dataOffset int32) int32
	TransferESDTExecute(destOffset int32, tokenIDOffset int32, tokenIDLen int32, valueOffset int32, gasLimit int64, functionOffset int32, functionLength int32, numArguments int32, argumentsLengthOffset int32, dataOffset int32) int32
	TransferESDTNFTExecute(destOffset int32, tokenIDOffset int32, tokenIDLen int32, valueOffset int32, nonce int64, gasLimit int64, functionOffset int32, functionLength int32, numArguments int32, argumentsLengthOffset int32, dataOffset int32) int32
	MultiTransferESDTNFTExecute(destOffset int32, numTokenTransfers int32, tokenTransfersArgsLengthOffset int32, tokenTransferDataOffset int32, gasLimit int64, functionOffset int32, functionLength int32, numArguments int32, argumentsLengthOffset int32, dataOffset int32) int32
	CreateAsyncCall(destOffset int32, valueOffset int32, dataOffset int32, dataLength int32, successOffset int32, successLength int32, errorOffset int32, errorLength int32, gas int64, extraGasForCallback int64) int32
	SetAsyncContextCallback(callback int32, callbackLength int32, data int32, dataLength int32, gas int64) int32
	UpgradeContract(destOffset int32, gasLimit int64, valueOffset int32, codeOffset int32, codeMetadataOffset int32, length int32, numArguments int32, argumentsLengthOffset int32, dataOffset int32)
	UpgradeFromSourceContract(destOffset int32, gasLimit int64, valueOffset int32, sourceContractAddressOffset int32, codeMetadataOffset int32, numArguments int32, argumentsLengthOffset int32, dataOffset int32)
	DeleteContract(destOffset int32, gasLimit int64, numArguments int32, argumentsLengthOffset int32, dataOffset int32)
	AsyncCall(destOffset int32, valueOffset int32, dataOffset int32, length int32)
	GetArgumentLength(id int32) int32
	GetArgument(id int32, argOffset int32) int32
	GetFunction(functionOffset int32) int32
	GetNumArguments() int32
	StorageStore(keyOffset int32, keyLength int32, dataOffset int32, dataLength int32) int32
	StorageLoadLength(keyOffset int32, keyLength int32) int32
	StorageLoadFromAddress(addressOffset int32, keyOffset int32, keyLength int32, dataOffset int32) int32
	StorageLoad(keyOffset int32, keyLength int32, dataOffset int32) int32
	SetStorageLock(keyOffset int32, keyLength int32, lockTimestamp int64) int32
	GetStorageLock(keyOffset int32, keyLength int32) int64
	IsStorageLocked(keyOffset int32, keyLength int32) int32
	ClearStorageLock(keyOffset int32, keyLength int32) int32
	GetCaller(resultOffset int32)
	CheckNoPayment()
	CallValue(resultOffset int32) int32
	GetESDTValue(resultOffset int32) int32
	GetESDTValueByIndex(resultOffset int32, index int32) int32
	GetESDTTokenName(resultOffset int32) int32
	GetESDTTokenNameByIndex(resultOffset int32, index int32) int32
	GetESDTTokenNonce() int64
	GetESDTTokenNonceByIndex(index int32) int64
	GetCurrentESDTNFTNonce(addressOffset int32, tokenIDOffset int32, tokenIDLen int32) int64
	GetESDTTokenType() int32
	GetESDTTokenTypeByIndex(index int32) int32
	GetNumESDTTransfers() int32
	GetCallValueTokenName(callValueOffset int32, tokenNameOffset int32) int32
	GetCallValueTokenNameByIndex(callValueOffset int32, tokenNameOffset int32, index int32) int32
	WriteLog(dataPointer int32, dataLength int32, topicPtr int32, numTopics int32)
	WriteEventLog(numTopics int32, topicLengthsOffset int32, topicOffset int32, dataOffset int32, dataLength int32)
	GetBlockTimestamp() int64
	GetBlockNonce() int64
	GetBlockRound() int64
	GetBlockEpoch() int64
	GetBlockRandomSeed(pointer int32)
	GetStateRootHash(pointer int32)
	GetPrevBlockTimestamp() int64
	GetPrevBlockNonce() int64
	GetPrevBlockRound() int64
	GetPrevBlockEpoch() int64
	GetPrevBlockRandomSeed(pointer int32)
	ReturnData(pointer int32, length int32)
	ExecuteOnSameContext(gasLimit int64, addressOffset int32, valueOffset int32, functionOffset int32, functionLength int32, numArguments int32, argumentsLengthOffset int32, dataOffset int32) int32
	ExecuteOnDestContext(gasLimit int64, addressOffset int32, valueOffset int32, functionOffset int32, functionLength int32, numArguments int32, argumentsLengthOffset int32, dataOffset int32) int32
	ExecuteReadOnly(gasLimit int64, addressOffset int32, functionOffset int32, functionLength int32, numArguments int32, argumentsLengthOffset int32, dataOffset int32) int32
	CreateContract(gasLimit int64, valueOffset int32, codeOffset int32, codeMetadataOffset int32, length int32, resultOffset int32, numArguments int32, argumentsLengthOffset int32, dataOffset int32) int32
	DeployFromSourceContract(gasLimit int64, valueOffset int32, sourceContractAddressOffset int32, codeMetadataOffset int32, resultAddressOffset int32, numArguments int32, argumentsLengthOffset int32, dataOffset int32) int32
	GetNumReturnData() int32
	GetReturnDataSize(resultID int32) int32
	GetReturnData(resultID int32, dataOffset int32) int32
	CleanReturnData()
	DeleteFromReturnData(resultID int32)
	GetOriginalTxHash(dataOffset int32)
	GetCurrentTxHash(dataOffset int32)
	GetPrevTxHash(dataOffset int32)
	ManagedSCAddress(destinationHandle int32)
	ManagedOwnerAddress(destinationHandle int32)
	ManagedCaller(destinationHandle int32)
	ManagedSignalError(errHandle int32)
	ManagedWriteLog(topicsHandle int32, dataHandle int32)
	ManagedGetOriginalTxHash(resultHandle int32)
	ManagedGetStateRootHash(resultHandle int32)
	ManagedGetBlockRandomSeed(resultHandle int32)
	ManagedGetPrevBlockRandomSeed(resultHandle int32)
	ManagedGetReturnData(resultID int32, resultHandle int32)
	ManagedGetMultiESDTCallValue(multiCallValueHandle int32)
	ManagedGetESDTBalance(addressHandle int32, tokenIDHandle int32, nonce int64, valueHandle int32)
	ManagedGetESDTTokenData(addressHandle int32, tokenIDHandle int32, nonce int64, valueHandle int32, propertiesHandle int32, hashHandle int32, nameHandle int32, attributesHandle int32, creatorHandle int32, royaltiesHandle int32, urisHandle int32)
	ManagedAsyncCall(destHandle int32, valueHandle int32, functionHandle int32, argumentsHandle int32)
	ManagedCreateAsyncCall(destHandle int32, valueHandle int32, functionHandle int32, argumentsHandle int32, successOffset int32, successLength int32, errorOffset int32, errorLength int32, gas int64, extraGasForCallback int64, callbackClosureHandle int32) int32
	ManagedGetCallbackClosure(callbackClosureHandle int32)
	ManagedUpgradeFromSourceContract(destHandle int32, gas int64, valueHandle int32, addressHandle int32, codeMetadataHandle int32, argumentsHandle int32, resultHandle int32)
	ManagedUpgradeContract(destHandle int32, gas int64, valueHandle int32, codeHandle int32, codeMetadataHandle int32, argumentsHandle int32, resultHandle int32)
	ManagedDeleteContract(destHandle int32, gasLimit int64, argumentsHandle int32)
	ManagedDeployFromSourceContract(gas int64, valueHandle int32, addressHandle int32, codeMetadataHandle int32, argumentsHandle int32, resultAddressHandle int32, resultHandle int32) int32
	ManagedCreateContract(gas int64, valueHandle int32, codeHandle int32, codeMetadataHandle int32, argumentsHandle int32, resultAddressHandle int32, resultHandle int32) int32
	ManagedExecuteReadOnly(gas int64, addressHandle int32, functionHandle int32, argumentsHandle int32, resultHandle int32) int32
	ManagedExecuteOnSameContext(gas int64, addressHandle int32, valueHandle int32, functionHandle int32, argumentsHandle int32, resultHandle int32) int32
	ManagedExecuteOnDestContext(gas int64, addressHandle int32, valueHandle int32, functionHandle int32, argumentsHandle int32, resultHandle int32) int32
	ManagedMultiTransferESDTNFTExecute(dstHandle int32, tokenTransfersHandle int32, gasLimit int64, functionHandle int32, argumentsHandle int32) int32
	ManagedTransferValueExecute(dstHandle int32, valueHandle int32, gasLimit int64, functionHandle int32, argumentsHandle int32) int32
	ManagedIsESDTFrozen(addressHandle int32, tokenIDHandle int32, nonce int64) int32
	ManagedIsESDTLimitedTransfer(tokenIDHandle int32) int32
	ManagedIsESDTPaused(tokenIDHandle int32) int32
	ManagedBufferToHex(sourceHandle int32, destHandle int32)
	MBufferNew() int32
	MBufferNewFromBytes(dataOffset int32, dataLength int32) int32
	MBufferGetLength(mBufferHandle int32) int32
	MBufferGetBytes(mBufferHandle int32, resultOffset int32) int32
	MBufferGetByteSlice(sourceHandle int32, startingPosition int32, sliceLength int32, resultOffset int32) int32
	MBufferCopyByteSlice(sourceHandle int32, startingPosition int32, sliceLength int32, destinationHandle int32) int32
	MBufferEq(mBufferHandle1 int32, mBufferHandle2 int32) int32
	MBufferSetBytes(mBufferHandle int32, dataOffset int32, dataLength int32) int32
	MBufferSetByteSlice(mBufferHandle int32, startingPosition int32, dataLength int32, dataOffset int32) int32
	MBufferAppend(accumulatorHandle int32, dataHandle int32) int32
	MBufferAppendBytes(accumulatorHandle int32, dataOffset int32, dataLength int32) int32
	MBufferToBigIntUnsigned(mBufferHandle int32, bigIntHandle int32) int32
	MBufferToBigIntSigned(mBufferHandle int32, bigIntHandle int32) int32
	MBufferFromBigIntUnsigned(mBufferHandle int32, bigIntHandle int32) int32
	MBufferFromBigIntSigned(mBufferHandle int32, bigIntHandle int32) int32
	MBufferToBigFloat(mBufferHandle int32, bigFloatHandle int32) int32
	MBufferFromBigFloat(mBufferHandle int32, bigFloatHandle int32) int32
	MBufferStorageStore(keyHandle int32, sourceHandle int32) int32
	MBufferStorageLoad(keyHandle int32, destinationHandle int32) int32
	MBufferStorageLoadFromAddress(addressHandle int32, keyHandle int32, destinationHandle int32)
	MBufferGetArgument(id int32, destinationHandle int32) int32
	MBufferFinish(sourceHandle int32) int32
	MBufferSetRandom(destinationHandle int32, length int32) int32
	SmallIntGetUnsignedArgument(id int32) int64
	SmallIntGetSignedArgument(id int32) int64
	SmallIntFinishUnsigned(value int64)
	SmallIntFinishSigned(value int64)
	SmallIntStorageStoreUnsigned(keyOffset int32, keyLength int32, value int64) int32
	SmallIntStorageStoreSigned(keyOffset int32, keyLength int32, value int64) int32
	SmallIntStorageLoadUnsigned(keyOffset int32, keyLength int32) int64
	SmallIntStorageLoadSigned(keyOffset int32, keyLength int32) int64
	Int64getArgument(id int32) int64
	Int64finish(value int64)
	Int64storageStore(keyOffset int32, keyLength int32, value int64) int32
	Int64storageLoad(keyOffset int32, keyLength int32) int64
	Sha256(dataOffset int32, length int32, resultOffset int32) int32
	ManagedSha256(inputHandle int32, outputHandle int32) int32
	Keccak256(dataOffset int32, length int32, resultOffset int32) int32
	ManagedKeccak256(inputHandle int32, outputHandle int32) int32
	Ripemd160(dataOffset int32, length int32, resultOffset int32) int32
	ManagedRipemd160(inputHandle int32, outputHandle int32) int32
	VerifyBLS(keyOffset int32, messageOffset int32, messageLength int32, sigOffset int32) int32
	ManagedVerifyBLS(keyHandle int32, messageHandle int32, sigHandle int32) int32
	VerifyEd25519(keyOffset int32, messageOffset int32, messageLength int32, sigOffset int32) int32
	ManagedVerifyEd25519(keyHandle int32, messageHandle int32, sigHandle int32) int32
	VerifyCustomSecp256k1(keyOffset int32, keyLength int32, messageOffset int32, messageLength int32, sigOffset int32, hashType int32) int32
	ManagedVerifyCustomSecp256k1(keyHandle int32, messageHandle int32, sigHandle int32, hashType int32) int32
	VerifySecp256k1(keyOffset int32, keyLength int32, messageOffset int32, messageLength int32, sigOffset int32) int32
	ManagedVerifySecp256k1(keyHandle int32, messageHandle int32, sigHandle int32) int32
	EncodeSecp256k1DerSignature(rOffset int32, rLength int32, sOffset int32, sLength int32, sigOffset int32) int32
	ManagedEncodeSecp256k1DerSignature(rHandle int32, sHandle int32, sigHandle int32) int32
	AddEC(xResultHandle int32, yResultHandle int32, ecHandle int32, fstPointXHandle int32, fstPointYHandle int32, sndPointXHandle int32, sndPointYHandle int32)
	DoubleEC(xResultHandle int32, yResultHandle int32, ecHandle int32, pointXHandle int32, pointYHandle int32)
	IsOnCurveEC(ecHandle int32, pointXHandle int32, pointYHandle int32) int32
	ScalarBaseMultEC(xResultHandle int32, yResultHandle int32, ecHandle int32, dataOffset int32, length int32) int32
	ManagedScalarBaseMultEC(xResultHandle int32, yResultHandle int32, ecHandle int32, dataHandle int32) int32
	ScalarMultEC(xResultHandle int32, yResultHandle int32, ecHandle int32, pointXHandle int32, pointYHandle int32, dataOffset int32, length int32) int32
	ManagedScalarMultEC(xResultHandle int32, yResultHandle int32, ecHandle int32, pointXHandle int32, pointYHandle int32, dataHandle int32) int32
	MarshalEC(xPairHandle int32, yPairHandle int32, ecHandle int32, resultOffset int32) int32
	ManagedMarshalEC(xPairHandle int32, yPairHandle int32, ecHandle int32, resultHandle int32) int32
	MarshalCompressedEC(xPairHandle int32, yPairHandle int32, ecHandle int32, resultOffset int32) int32
	ManagedMarshalCompressedEC(xPairHandle int32, yPairHandle int32, ecHandle int32, resultHandle int32) int32
	UnmarshalEC(xResultHandle int32, yResultHandle int32, ecHandle int32, dataOffset int32, length int32) int32
	ManagedUnmarshalEC(xResultHandle int32, yResultHandle int32, ecHandle int32, dataHandle int32) int32
	UnmarshalCompressedEC(xResultHandle int32, yResultHandle int32, ecHandle int32, dataOffset int32, length int32) int32
	ManagedUnmarshalCompressedEC(xResultHandle int32, yResultHandle int32, ecHandle int32, dataHandle int32) int32
	GenerateKeyEC(xPubKeyHandle int32, yPubKeyHandle int32, ecHandle int32, resultOffset int32) int32
	ManagedGenerateKeyEC(xPubKeyHandle int32, yPubKeyHandle int32, ecHandle int32, resultHandle int32) int32
	CreateEC(dataOffset int32, dataLength int32) int32
	ManagedCreateEC(dataHandle int32) int32
	GetCurveLengthEC(ecHandle int32) int32
	GetPrivKeyByteLengthEC(ecHandle int32) int32
	EllipticCurveGetValues(ecHandle int32, fieldOrderHandle int32, basePointOrderHandle int32, eqConstantHandle int32, xBasePointHandle int32, yBasePointHandle int32) int32
}
