package edit

import (
	"crypto/md5"  //nolint:gosec // md5 is reachable only as an explicit hash:'md5' choice, for matching digests other systems already produced
	"crypto/sha1" //nolint:gosec // sha1 is reachable only as an explicit hash:'sha1' choice, for matching digests other systems already produced
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameHash is the template name for the Hash modifier.
const ModifierNameHash functions.ModifierName = "hash"

// Hash returns the lowercase hex digest of the value under the named algorithm,
// one of sha256, sha512, sha1, or md5. The algorithm is required and there is no
// default, so a template always names the algorithm it was written against and
// nothing silently moves under it later. An unrecognized name is an error rather
// than a fallback to some other algorithm.
//
// The bytes are hashed exactly as they arrive, with no trimming, case folding,
// or other normalization, so a digest computed here matches one computed
// elsewhere over the same input.
//
// sha1 and md5 are offered because a template often has to reproduce a digest
// another system already published, a Gravatar URL or a stored checksum among
// them. Neither is safe where a reader must not be able to forge the input.
func Hash(s, algorithm string) (string, error) {
	switch algorithm {
	case "sha256":
		sum := sha256.Sum256([]byte(s))
		return hex.EncodeToString(sum[:]), nil
	case "sha512":
		sum := sha512.Sum512([]byte(s))
		return hex.EncodeToString(sum[:]), nil
	case "sha1":
		sum := sha1.Sum([]byte(s)) //nolint:gosec // the template asked for sha1 by name, see the algorithm note above
		return hex.EncodeToString(sum[:]), nil
	case "md5":
		sum := md5.Sum([]byte(s)) //nolint:gosec // the template asked for md5 by name, see the algorithm note above
		return hex.EncodeToString(sum[:]), nil
	default:
		return "", fmt.Errorf("unknown hash algorithm %q: %w", algorithm, functions.ErrInvalidParamValue)
	}
}
