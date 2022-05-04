- 현 프로젝트를 모듈로 만들려고 한다.

* https://www.youtube.com/watch?v=Ja-xVdcgo-s

```
// init 이후 모듈명을 적어준다.
go mod init GoLangExamle/usepkg
go: creating new go.mod: module GoLangExample/usepkg

//go.mod 생성
module GoLangExample/usepkg

go 1.13
//
go.mod tidy
//go.mod 외부 라이브러리 다운
module GoLangExample/usepkg

go 1.13

require github.com/guptarohit/asciigraph v0.5.3
// main.go가 있는 디렉토리 이동
go build
디렉토리명과 동일한 바이너리 파일 생성
```
