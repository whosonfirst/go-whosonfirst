GOMOD=$(shell test -f "go.work" && echo "readonly" || echo "vendor")
LDFLAGS=-s -w

vuln:
	govulncheck -show verbose ./...

bump-version:
	perl -i -p -e 's/github.com\/whosonfirst\/go-whosonfirst\/$(OLD)/github.com\/whosonfirst\/go-whosonfirst\/$(NEW)/g' go.mod
	perl -i -p -e 's/github.com\/whosonfirst\/go-whosonfirst\/$(OLD)/github.com\/whosonfirst\/go-whosonfirst\/$(NEW)/g' README.md
	find . -name '*.go' | xargs perl -i -p -e 's/github.com\/whosonfirst\/go-whosonfirst\/$(OLD)/github.com\/whosonfirst\/go-whosonfirst\/$(NEW)/g'
