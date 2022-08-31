echo "开始"
git pull
git add .
time=$(date "+%Y-%m-%d %H:%M:%S")
git commit -m "${time}"
git push
git push origin-github master
echo "结束"
