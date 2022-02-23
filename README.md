# ddd-template
六边形领域驱动项目架构模板，完善中


# UNIT TEST

- mockgen dependency code
```shell
    mockgen  mockgen -destination ./adapters/repo_mock/mock_demo_repo.go -package repo_mock -source ./domain/dependency/dep_demo.go 
```
- mockgen application code
```shell
    
```