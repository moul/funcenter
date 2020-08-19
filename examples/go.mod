module moul.io/funcenter/examples

go 1.12

require (
	github.com/kr/pretty v0.1.0 // indirect
	go.uber.org/zap v1.15.0
	moul.io/funcenter v1.0.0
	moul.io/funcenter/logging/zap v0.0.0
)

replace moul.io/funcenter => ../

replace moul.io/funcenter/logging/zap => ../logging/zap/
