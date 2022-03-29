# EshelonTestTask
 Test task

Запуск:
go run server.go

Запросы:
curl --header "Content-Type: application/json"   --request POST   --data '{"cmd": "df -h", "os": "linux"}' -k   https://localhost:8080/api/v1/remote-execution
curl --header "Content-Type: application/json"   --request POST   --data '{"cmd": "ping.exe https://google.com", "os": "windows", "stdin":""}' -k   https://localhost:8080/api/v1/remote-execution

