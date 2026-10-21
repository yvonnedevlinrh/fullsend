package cli

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateOpenCodeSecurity(t *testing.T) {
	for _, runtimeName := range []string{"opencode", "pi", "claude-code", "codex", "dummy"} {
		for _, enabled := range []bool{true, false} {
			err := validateOpenCodeSecurity(runtimeName, enabled)
			if runtimeName == "opencode" && !enabled {
				require.ErrorContains(t, err, "security.enabled: true")
			} else {
				require.NoError(t, err)
			}
		}
	}
}
