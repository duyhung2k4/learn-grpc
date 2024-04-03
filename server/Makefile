gen_code_grpc:
	protoc \
		--go_out=grpc --go_opt=paths=source_relative \
  	--go-grpc_out=grpc --go-grpc_opt=paths=source_relative \
  	proto/*.proto
export_bin_go:
	export PATH="$PATH:$(go env GOPATH)/bin"
gen_ssl:
	openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout keys/private.pem -out keys/public.pem -config keys/san.cfg -extensions 'v3_req'