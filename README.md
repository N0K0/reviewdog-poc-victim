# reviewdog-poc-victim

Private PoC reproducing the `pull_request_target` + `go generate` sink
pattern observed in `sealdice/sealdice-core` and `sealdice/sealdice-ui`.

Vulnerable workflow at `.github/workflows/reviewdog.yml` is a verbatim
copy of the upstream pattern. Test method: open a cross-fork PR from
`hackeriet/reviewdog-poc-attacker` carrying a benign `//go:generate`
directive and confirm execution in the workflow logs.
