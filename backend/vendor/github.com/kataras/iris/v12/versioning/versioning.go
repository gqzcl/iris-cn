package versioning

import (
	"github.com/kataras/iris/v12/context"

	"github.com/hashicorp/go-version"
)

// If reports whether the "version" is matching to the "is".
// the "is" can be a constraint like ">= 1, < 3".
func If(v string, is string) bool {
	_, ok := check(v, is)
	return ok
}

func check(v string, is string) (string, bool) {
	ver, err := version.NewVersion(v)
	if err != nil {
		return "", false
	}

	constraints, err := version.NewConstraint(is)
	if err != nil {
		return "", false
	}

	// return the extracted version from request, even if not matched.
	return ver.String(), constraints.Check(ver)
}

// Match acts exactly the same as `If` does but instead it accepts
// a Context, so it can be called by a handler to determinate the requested version.
//
// If matched then it sets the "X-API-Version" response header and
// stores the matched version into Context (see `GetVersion` too).
func Match(ctx *context.Context, expectedVersion string) bool {
	versionString, matched := check(GetVersion(ctx), expectedVersion)
	if !matched {
		return false
	}

	SetVersion(ctx, versionString)
	ctx.Header("X-API-Version", versionString)
	return true
}

// Handler returns a handler which stop the execution
// when the given "version" does not match with the requested one.
func Handler(version string) context.Handler {
	return func(ctx *context.Context) {
		if !Match(ctx, version) {
			// The overlapped handler has an exception
			// of a type of context.NotFound (which versioning.ErrNotFound wraps)
			// to clear the status code
			// and the error to ignore this
			// when available match version exists (see `NewGroup`).
			NotFoundHandler(ctx)
			return
		}

		ctx.Next()
	}
}

// Map is a map of versions targets to a handlers,
// a handler per version or constraint, the key can be something like ">1, <=2" or just "1".
type Map map[string]context.Handler

// NewMatcher creates a single handler which decides what handler
// should be executed based on the requested version.
//
// Use the `NewGroup` if you want to add many routes under a specific version.
//
// See `Map` and `NewGroup` too.
func NewMatcher(versions Map) context.Handler {
	constraintsHandlers, notFoundHandler := buildConstraints(versions)

	return func(ctx *context.Context) {
		versionString := GetVersion(ctx)
		if versionString == "" || versionString == NotFound {
			notFoundHandler(ctx)
			return
		}

		ver, err := version.NewVersion(versionString)
		if err != nil {
			notFoundHandler(ctx)
			return
		}

		for _, ch := range constraintsHandlers {
			if ch.constraints.Check(ver) {
				ctx.Header("X-API-Version", ver.String())
				ch.handler(ctx)
				return
			}
		}

		// pass the not matched version so the not found handler can have knowedge about it.
		// SetVersion(ctx, versionString)
		// or let a manual cal of GetVersion(ctx) do that instead.
		notFoundHandler(ctx)
	}
}

type constraintsHandler struct {
	constraints version.Constraints
	handler     context.Handler
}

func buildConstraints(versionsHandler Map) (constraintsHandlers []*constraintsHandler, notfoundHandler context.Handler) {
	for v, h := range versionsHandler {
		if v == NotFound {
			notfoundHandler = h
			continue
		}

		constraints, err := version.NewConstraint(v)
		if err != nil {
			panic(err)
		}

		constraintsHandlers = append(constraintsHandlers, &constraintsHandler{
			constraints: constraints,
			handler:     h,
		})
	}

	if notfoundHandler == nil {
		notfoundHandler = NotFoundHandler
	}

	// no sort, the end-dev should declare
	// all version constraint, i.e < 4.0 may be catch 1.0 if not something like
	// >= 3.0, < 4.0.
	// I can make it ordered but I do NOT like the final API of it:
	/*
		app.Get("/api/user", NewMatcher( // accepts an array, ordered, see last elem.
			V("1.0", vHandler("v1 here")),
			V("2.0", vHandler("v2 here")),
			V("< 4.0", vHandler("v3.x here")),
		))
		instead we have:

		app.Get("/api/user", NewMatcher(Map{ // accepts a map, unordered, see last elem.
			"1.0":           Deprecated(vHandler("v1 here")),
			"2.0":           vHandler("v2 here"),
			">= 3.0, < 4.0": vHandler("v3.x here"),
			VersionUnknown: customHandlerForNotMatchingVersion,
		}))
	*/

	return
}