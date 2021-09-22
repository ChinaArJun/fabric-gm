module github.com/hyperledger/fabric/scripts/fabric-samples/chaincode/abstore/go

go 1.14

require (
	github.com/hyperledger/fabric-chaincode-go v0.0.0-20200728190242-9b3ae92d8664 // indirect
	github.com/hyperledger/fabric-contract-api-go v1.1.0
	github.com/tjfoc/gmtls v0.0.0-00010101000000-000000000000 // indirect
)

replace (
	github.com/hyperledger/fabric-chaincode-go => ./fabric-chaincode-go
	github.com/tjfoc/gmsm => ./tjfoc/gmsm
	github.com/tjfoc/gmtls => ./tjfoc/gmtls
)
