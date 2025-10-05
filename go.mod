module github.com/exercism/go-analyzer

go 1.25

require (
	github.com/logrusorgru/aurora v2.0.3+incompatible
	github.com/namsral/flag v1.7.4-pre
	github.com/pkg/errors v0.9.1
	github.com/pmezard/go-difflib v1.0.0
	github.com/shurcooL/vfsgen v0.0.0-20230704071429-0000e147ea92
	github.com/stretchr/testify v1.11.1
	github.com/tehsphinx/astpatt v0.3.2
	github.com/tehsphinx/astrav v0.5.1
	golang.org/x/lint v0.0.0-20241112194109-818c5a804067
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/shurcooL/httpfs v0.0.0-20230704072500-f1e31cf0ba5c // indirect
	golang.org/x/mod v0.28.0 // indirect
	golang.org/x/sync v0.17.0 // indirect
	golang.org/x/tools v0.37.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

// replace github.com/tehsphinx/astrav => ../astrav
// replace github.com/tehsphinx/astpatt => ../astpatt
