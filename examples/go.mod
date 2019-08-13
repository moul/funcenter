module moul.io/funcenter/examples

go 1.12

require (
	go.uber.org/zap v1.10.0
	moul.io/funcenter v0.0.0
	moul.io/funcenter/logging/zap v0.0.0
)

replace moul.io/funcenter => ../

replace moul.io/funcenter/logging/zap => ../logging/zap/
