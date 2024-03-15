# go-ncp
![Static Badge](https://img.shields.io/badge/Author-Antonio-caf0f8?labelColor=2b2d42)
![Static Badge](https://img.shields.io/badge/Golang_1.20.5-grey?style=flat&logo=go&labelColor=2b2d42)

## 개요
> 네이버 클라우드 API 요청 및 전, 후 처리하는 서비스입니다.   
> 개발은 '네이버 금융 클라우드'를 기준으로 작성되었습니다.

### 최소 요구 사항
- **`GO >= 1.20.5`**
- **`MySQL >= 8.0.27`**

### 적용 API
+ Fin Ncloud > Platform > Cost And Usage [[Link]](https://api-fin.ncloud-docs.com/docs/platform-costandusage)
+ Fin Ncloud > Platform > List Price [[Link]](https://api-fin.ncloud-docs.com/docs/platform-listprice)

## 설치 방법
1. **Github Repository Clone**
2. **의존성 다운로드**   
   - `go mod download`
3. **env 파일 작성**
   - 환경 파일명은 `ncp.env` 입니다.
   - `VALIDATION_KEY`는 base64로 인코딩된 값을 사용합니다.
4. **서비스 실행**
    - `go run main.go`

## 변경 로그
- **`24. 3. 15`** - v0.1.1 비용 API Handler 구현

## 추가 자료
- Gin 프레임워크 가이드 : [Go Gin Quick start](https://github.com/gin-gonic/gin/blob/master/docs/doc.md)
- Go Build Command 가이드 : [Go Build Command Guide](https://pkg.go.dev/cmd/go#hdr-Build_constraints)