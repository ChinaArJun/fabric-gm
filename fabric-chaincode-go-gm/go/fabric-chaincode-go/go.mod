module github.com/hyperledger/fabric/scripts/fabric-samples/chaincode/abstore/go/fabric-chaincode-go

go 1.14

require (
	github.com/golang/protobuf v1.4.2
	github.com/hyperledger/fabric-chaincode-go v0.0.0-00010101000000-000000000000
	github.com/hyperledger/fabric-protos-go v0.0.0-20200728190333-526bfc137380
	github.com/tjfoc/gmsm v1.2.0
	github.com/tjfoc/gmtls v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.30.0
)

replace (
	github.com/hyperledger/fabric-chaincode-go => ../fabric-chaincode-go
	github.com/tjfoc/gmsm => ../tjfoc/gmsm
	github.com/tjfoc/gmtls => ../tjfoc/gmtls
)
