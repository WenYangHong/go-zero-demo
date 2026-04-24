tables=$2
modeldir=./genModel

host=127.0.0.1
port=3306
dbname=$1
username=root
password=123456

echo "开始创建库:$dbname 的表:$2"
goctl model mysql datasource -url="${username}:${password}@tcp(${host}:${port})/${dbname}" -table="${tables}" -dir="${modeldir}" --style=goZero