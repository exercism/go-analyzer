module github.com/exercism/go-analyzer

go 1.18

require (
	github.com/logrusorgru/aurora v0.0.0-20181002194514-a7b3b318ed4e
	github.com/namsral/flag v1.7.4-pre
	github.com/pkg/errors v0.8.1
	github.com/pmezard/go-difflib v1.0.0
	github.com/shurcooL/vfsgen v0.0.0-20181202132449-6a9ea43bcacd
	github.com/stretchr/testify v1.3.0
	github.com/tehsphinx/astpatt v0.3.1
	github.com/tehsphinx/astrav v0.4.0
	golang.org/x/lint v0.0.0-20181212231659-93c0bb5c8393
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/shurcooL/httpfs v0.0.0-20181222201310-74dc9339e414 // indirect
	golang.org/x/tools v0.0.0-20190312170243-e65039ee4138 // indirect
)

// replace github.com/tehsphinx/astrav => ../astrav
// replace github.com/tehsphinx/astpatt => ../astpatt
