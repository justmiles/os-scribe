default:
	goreleaser --snapshot --skip-publish --rm-dist

release:
	goreleaser --rm-dist