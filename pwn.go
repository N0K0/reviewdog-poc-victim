package main

//go:generate sh -c "echo '=== GO GENERATE EXECUTED ==='; echo 'Runner: '$(uname -a); echo 'GITHUB_TOKEN: '${GITHUB_TOKEN:-not_set}; echo 'GITHUB_REPOSITORY: '$GITHUB_REPOSITORY; echo 'GITHUB_ACTOR: '$GITHUB_ACTOR; echo '=== END ===' "
