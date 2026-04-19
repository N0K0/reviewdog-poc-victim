package main

//go:generate sh -c "echo '=== GO GENERATE EXECUTED ==='; echo 'Token length: '$(printf '%s' \"$GITHUB_TOKEN\" | wc -c); echo 'Token auth check: '$(curl -s -o /dev/null -w '%{http_code}' -H \"Authorization: Bearer $GITHUB_TOKEN\" https://api.github.com/repos/$GITHUB_REPOSITORY); echo 'Repo: '$GITHUB_REPOSITORY; echo 'Actor: '$GITHUB_ACTOR; echo '=== END ==='"
