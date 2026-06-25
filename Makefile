GOMOD=$(shell test -f "go.work" && echo "readonly" || echo "vendor")
LDFLAGS=-s -w

TAGS=null

vuln:
	govulncheck -show verbose ./...

cli:
	@make cli-concordances
	@make cli-database
	@make cli-derivatives
	@make cli-edtf
	@make cli-fetch
	@make cli-findingaids
	@make cli-iterate
	@make cli-travel

cli-concordances:
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/wof-concordances-keys cmd/wof-concordances-keys/main.go

cli-edtf:
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/wof-edtf-find-invalid cmd/wof-edtf-find-invalid/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/wof-edtf-update-unknown-uncertain cmd/wof-edtf-update-unknown-uncertain/main.go

cli-fetch:
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/wof-fetch-records cmd/wof-fetch-records/main.go

cli-travel:
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/wof-travel-id cmd/wof-travel-id/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/wof-travel-belongsto cmd/wof-travel-belongsto/main.go

cli-derivatives:
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -tags $(TAGS) -o bin/wof-derivatives-server cmd/wof-derivatives-server/main.go

cli-iterate:
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -tags $(TAGS) -o bin/wof-iterate-count cmd/wof-iterate-count/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -tags $(TAGS) -o bin/wof-iterate-emit cmd/wof-iterate-emit/main.go

cli-database:
	@make cli-database-sql
	@make cli-database-opensearch

cli-database-sql:
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -tags $(TAGS) -o bin/wof-sql-create cmd/wof-sql-create/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -tags $(TAGS) -o bin/wof-sql-index cmd/wof-sql-index/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -tags $(TAGS) -o bin/wof-sql-prune cmd/wof-sql-prune/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -tags $(TAGS) -o bin/wof-sql-iterwriter cmd/wof-sql-iterwriter/main.go

cli-database-opensearch:
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/wof-opensearch-index cmd/wof-opensearch-index/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/wof-opensearch-query cmd/wof-opensearch-query/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/wof-opensearch-create-index cmd/wof-opensearch-create-index/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/wof-opensearch-delete-index cmd/wof-opensearch-delete-index/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/wof-opensearch-put-mapping cmd/wof-opensearch-put-mapping/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/wof-opensearch-get-mapping cmd/wof-opensearch-get-mapping/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/wof-opensearch-put-settings cmd/wof-opensearch-put-settings/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/wof-opensearch-list-indices cmd/wof-opensearch-list-indices/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/wof-opensearch-indices-stats cmd/wof-opensearch-indices-stats/main.go

cli-findingaids:
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/wof-findingaid-populate cmd/wof-findingaid-populate/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/wof-findingaid-sources cmd/wof-findingaid-sources/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/wof-findingaid-csv2sql cmd/wof-findingaid-csv2sql/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/wof-findingaid-csv2docstore cmd/wof-findingaid-csv2docstore/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/wof-findingaid-create-dynamodb-tables cmd/wof-findingaid-create-dynamodb-tables/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/wof-findingaid-create-dynamodb-import cmd/wof-findingaid-create-dynamodb-import/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/wof-findingaid-resolverd cmd/wof-findingaid-resolverd/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/wof-findingaid-resolve cmd/wof-findingaid-resolve/main.go


lambda:
	@make lambda-findingaids-resolverd

lambda-findingaids-resolverd:
	if test -f bootstrap; then rm -f bootstrap; fi
	if test -f resolverd.zip; then rm -f resolverd.zip; fi
	GOARCH=arm64 GOOS=linux go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -tags lambda.norpc -o bootstrap cmd/wof-findingaid-resolverd/main.go
	zip resolverd.zip bootstrap
	rm -f bootstrap


test-fetch:
	@make cli-fetch
	./bin/wof-fetch-records -verbose 1360695651


bump-version:
	perl -i -p -e 's/github.com\/whosonfirst\/go-whosonfirst\/$(OLD)/github.com\/whosonfirst\/go-whosonfirst\/$(NEW)/g' go.mod
	perl -i -p -e 's/github.com\/whosonfirst\/go-whosonfirst\/$(OLD)/github.com\/whosonfirst\/go-whosonfirst\/$(NEW)/g' README.md
	find . -name '*.go' | xargs perl -i -p -e 's/github.com\/whosonfirst\/go-whosonfirst\/$(OLD)/github.com\/whosonfirst\/go-whosonfirst\/$(NEW)/g'
