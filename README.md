# EshelonTestTask
 Test task
Важно! Установить wine
Запуск:

go run server.go

Запросы:

curl --header "Content-Type: application/json"   --request POST   --data '{"cmd": "df -h", "os": "linux", "stdin":""}' -k   https://localhost:8080/api/v1/remote-execution

curl --header "Content-Type: application/json"   --request POST   --data '{"cmd": "ping.exe google.com", "os": "windows", "stdin":""}' -k   https://localhost:8080/api/v1/remote-execution

Пример выполнения команды:

Body:
{"cmd": "df -h", "os": "linux"}

Response:
{"stdout":"Filesystem       Size   Used  Avail Capacity iused      ifree %iused  Mounted on\n/dev/disk3s1s1  460Gi   23Gi  232Gi    10%  577694 2433026960    0%   /\ndevfs           204Ki  204Ki    0Bi   100%     706          0  100%   /dev\n/dev/disk3s6    460Gi  1.0Gi  232Gi     1%       1 2433026960    0%   /System/Volumes/VM\n/dev/disk3s2    460Gi  636Mi  232Gi     1%    1298 2433026960    0%   /System/Volumes/Preboot\n/dev/disk3s4    460Gi  609Mi  232Gi     1%     224 2433026960    0%   /System/Volumes/Update\n/dev/disk1s2    500Mi  6.0Mi  480Mi     2%       1    4911280    0%   /System/Volumes/xarts\n/dev/disk1s1    500Mi  7.3Mi  480Mi     2%      29    4911280    0%   /System/Volumes/iSCPreboot\n/dev/disk1s3    500Mi  2.3Mi  480Mi     1%      47    4911280    0%   /System/Volumes/Hardware\n/dev/disk3s5    460Gi  201Gi  232Gi    47%  964232 2433026960    0%   /System/Volumes/Data\nmap auto_home     0Bi    0Bi    0Bi   100%       0          0  100%   /System/Volumes/Data/home\n/dev/disk2s1    5.0Gi  1.6Gi  3.4Gi    32%      58   35541640    0%   /System/Volumes/Update/SFR/mnt1\n/dev/disk3s1    460Gi   23Gi  232Gi    10%  500634 2433026960    0%   /System/Volumes/Update/mnt1\n","stderr":""}

Body:
{"cmd": "ls", "os": "linux","stdin":""}

Response:
{"stdout":"README.md\ncert\nserver.go\n","stderr":""}

Body:
{"cmd": "ping.exe google.com", "os": "windows", "stdin":""}

Response:
{"stdout":"Pinging google.com [64.233.165.139] with 32 bytes of data:\r\nRequest timed out.\r\nRequest timed out.\r\nRequest timed out.\r\nRequest timed out.\r\n\r\nPing statistics for 64.233.165.139\r\n\tPackets: Sent = 4, Received = 0, Lost = 4 (100% loss)\r\n","stderr":""}



