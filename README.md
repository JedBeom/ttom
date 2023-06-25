# 밀리시타 트위터 계정 미러링 봇

아이돌마스터 밀리언 라이브! 시어터 데이즈 트위터 공식 계정의 트윗을 마스토돈에 미러링해서 툿합니다.

```
cp config.json.example config.json
vi config.json
go build
./ttom
```

## TODO
- [ ] 미디어 가져오거나 업로드 중 에러가 발생하면 재시도
- [ ] Tweet-Toot 연동 테이블
- [ ] 트위터에서 자신의 트윗을 리트윗 한 경우에 봇도 자신의 툿을 부스트하도록
