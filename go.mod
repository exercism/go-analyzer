module github.com/exercism/go-analyzer

go 1.12

require (
	github.com/namsral/flag v1.7.4-pre
	github.com/shurcooL/httpfs v0.0.0-20181222201310-74dc9339e414 // indirect
	github.com/shurcooL/vfsgen v0.0.0-20181202132449-6a9ea43bcacd
	github.com/stretchr/testify v1.3.0
	github.com/tehsphinx/astpatt v0.1.1
	github.com/tehsphinx/astrav v0.0.0-20190316220234-6d5960953913
	golang.org/x/tools v0.0.0-20190312170243-e65039ee4138 // indirect
)

// replace github.com/tehsphinx/astrav => ../astrav
// replace github.com/tehsphinx/astpatt => ../astpatt
