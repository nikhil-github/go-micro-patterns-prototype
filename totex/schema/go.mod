module github.com/yourusername/schema

go 1.24.3

require (
	github.com/bufbuild/connect-go v1.10.0
	google.golang.org/protobuf v1.36.1
)

require github.com/google/go-cmp v0.6.0 // indirect

replace github.com/yourusername/foundation => ../foundation
