

### Command

```
# 実行
go run .

# テスト
go test -v

# ビルド
go build -o .
```

### Usage

```
# calver取得（2023/10/22時点）
go run .
20231022.1

go run . 20231022.1
20231022.2

go run . 20231021.1
20231022.1
```
