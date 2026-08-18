package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/url"
	"os"
	"path/filepath"
	"sort"

	_ "github.com/whosonfirst/go-whosonfirst/v4/iterate/git"

	"github.com/sfomuseum/go-flags/flagset"
	"github.com/sfomuseum/go-flags/multi"
	"github.com/whosonfirst/go-whosonfirst-github/organizations"
	"github.com/whosonfirst/go-whosonfirst/v4/iterate"
	"github.com/whosonfirst/go-whosonfirst/v4/parquet"
)

func main() {

	var iterator_uri string
	var writer_root string
	var git_org string
	var git_repos multi.MultiString
	var git_prefixes multi.MultiString
	var git_exclude multi.MultiString
	var verbose bool

	fs := flagset.NewFlagSet("parquet")
	fs.StringVar(&iterator_uri, "iterator-uri", "git:///tmp", "One of two registered whosonfirst/go-whosonfirst/v4/iterate.Iterator URIs: repo:// or git://")
	fs.StringVar(&writer_root, "writer-root", "/usr/local/data/whosonfirst-parquet", "The path on the local filesystem where Parquet files should be written to.")
	fs.StringVar(&git_org, "git-org", "whosonfirst-data", "If -iterator-uri is git://, the organization to use when cloning repos.")
	fs.Var(&git_prefixes, "git-prefix", "If -iterator-uri is git://, zero or more repository prefixe to filter by (if -repo flag is empty). Default is [whosonfirst-data-admin-].")
	fs.Var(&git_exclude, "git-exclude", "If -iterator-uri is git://, zero or more repository prefixe to exclude by (if -repo flag is empty). Default is [whosonfirst-data-admin-alt].")
	fs.Var(&git_repos, "git-repo", "If -iterator-uri is git://, zero or more repos to clone and create Parquet exports for. If empty then repos will be limited to those starting in whosonfirst-data-admin.")
	fs.BoolVar(&verbose, "verbose", false, "Enable verbose (debug) logging.")

	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, "Write data from one or more Who's On First style repositories to Parquet files.\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\t%s [options]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Valid options are:\n")
		fs.PrintDefaults()
	}

	flagset.Parse(fs)

	if verbose {
		slog.SetLogLoggerLevel(slog.LevelDebug)
		slog.Debug("Verbose logging enabled")
	}

	ctx := context.Background()

	abs_root, err := filepath.Abs(writer_root)

	if err != nil {
		log.Fatalf("Failed to derive absolute path for writer root, %v", err)
	}

	iterator_sources := make(map[string]string)

	iter_u, err := url.Parse(iterator_uri)

	if err != nil {
		log.Fatalf("Failed to parse iterator URI, %v", err)
	}

	switch iter_u.Scheme {
	case "git":

		if len(git_repos) == 0 {

			if len(git_prefixes) == 0 {
				git_prefixes = []string{
					"whosonfirst-data-admin",
				}
			}

			if len(git_exclude) == 0 {
				git_exclude = []string{
					"whosonfirst-data-admin-alt",
				}
			}

			list_opts := &organizations.ListOptions{
				Prefix:          git_prefixes,
				Exclude:         git_exclude,
				ExcludeArchived: true,
			}

			list_repos, err := organizations.ListRepos(git_org, list_opts)

			if err != nil {
				log.Fatalf("Failed to list repos, %v", err)
			}

			git_repos = list_repos
			sort.Strings(git_repos)
		}

		for _, repo := range git_repos {
			iterator_sources[repo] = fmt.Sprintf("https://github.com/%s/%s.git", git_org, repo)
		}

	case "repo":

		for _, uri := range fs.Args() {
			repo := filepath.Base(uri)
			iterator_sources[repo] = uri
		}

	default:
		log.Fatalf("Unsupported iterator scheme.")
	}

	for repo, iterator_source := range iterator_sources {

		writer_fname := fmt.Sprintf("%s.parquet", repo)
		writer_uri := filepath.Join(abs_root, writer_fname)

		logger := slog.Default()
		logger = logger.With("source", iterator_source)
		logger = logger.With("target", writer_uri)

		logger.Info("Start export")

		wr, err := parquet.NewWriter(ctx, writer_uri)

		if err != nil {
			log.Fatalf("Failed to create new writer, %v", err)
		}

		iter, err := iterate.NewIterator(ctx, iterator_uri)

		if err != nil {
			log.Fatalf("Failed to create new iterator, %v", err)
		}

		for rec, err := range iter.Iterate(ctx, iterator_source) {

			if err != nil {
				log.Fatalf("Iterator yielded an error, %v", err)
			}

			_, err = wr.WriteFromReader(rec.Body)

			rec.Body.Close()

			if err != nil {
				log.Fatalf("Failed to write row %s, %v", rec.Path, err)
			}
		}

		err = wr.Close()

		if err != nil {
			log.Fatalf("Failed to close after writing, %v", err)
		}

		logger.Info("Completed export")
	}
}
