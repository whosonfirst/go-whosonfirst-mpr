# go-whosonfirst-mpr

Go package for defining a minimal places response for Who's On First (WOF) related services.

## Important

This package is still considered experimental. It is meant to work in tandem with the [standard places response](https://github.com/whosonfirst/go-whosonfirst-spr) (SPR) interfaces for Who's On First services. Specifically, the SPR interfaces should implement the "minimal places response" (MPR) interfaces.

The MPR exists to ensure a common interface for services where storing the data necessary to conform to the SPR is unnecessary or impractical. For exaple, packages like:

* https://github.com/whosonfirst/go-whosonfirst-search-bleve

The MPR has the same interface as the [go-whosonfirst-findingaid](https://github.com/whosonfirst/go-whosonfirst-findingaid) `Repo` interface:

* https://github.com/whosonfirst/go-whosonfirst-findingaid/blob/main/repo/repo.go

Assuming the MPR is considered a viable approach then the `Repo` interface will be retired.

## Known issues

The `Id()` method for the SPR interface returns a string. This was done to accomodate non-WOF GeoJSON records. The MPR, as written, requires the `Id`()` method return an int64. This disconnect needs to be resolved one way or another.

## See also

* https://github.com/whosonfirst/go-whosonfirst-spr
* https://github.com/whosonfirst/go-whosonfirst-findingaid