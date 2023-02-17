# dump config files from git repo

## play

```go
go run main.go
```
## output
```
start clone github repo: https://github.com/kubernetes-sigs/kind
Enumerating objects: 26486, done.
Counting objects: 100% (30/30), done.
Compressing objects: 100% (26/26), done.
Total 26486 (delta 4), reused 20 (delta 1), pack-reused 26456
git head ref: b3d74a382d5fe3be0f6efd56e733a74aa9a3f27a refs/heads/main
dump git files(in 'hack/build') success:
        README.md
        setup-go.sh
        goinstalldir.sh
        init-buildx.sh
start clone github repo: https://oauth2:CTfJrDFMLoW7DgUbX_sH@gitlab.qiniu.io/qbox/c-deploy.git
Counting objects: 68099, done.
Compressing objects: 100% (26362/26362), done.
Total 68099 (delta 40872), reused 66417 (delta 39621)
git head ref: 282364f909cb46c5c27caef0870b1877124a5f03 refs/heads/master
dump git files(in 'callisto/templates/dora/base/service-gate') success:
        service.yaml
        deployment.yaml
        ingress.yaml
        log-auditlog-configmap.yaml
        log-applog-configmap.yaml
```