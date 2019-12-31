module github.com/lukasjarosch/secret-santa

go 1.13

require (
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/go-playground/validator v9.31.0+incompatible
	github.com/golang/protobuf v1.3.2
	github.com/google/uuid v1.1.1
	github.com/grpc-ecosystem/grpc-gateway v1.12.1
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/lukasjarosch/genki v0.5.12
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.6.1 // indirect
	google.golang.org/genproto v0.0.0-20191230161307-f3c370f40bfb
	google.golang.org/grpc v1.26.0
)

replace github.com/lukasjarosch/genki => ../genki
