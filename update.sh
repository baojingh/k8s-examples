#/bin/bash

echo "====== $(date '+%Y-%m-%d %H:%M:%S') ================"
echo " "
echo "====== git pull ================"
git pull
echo  ""
echo  ""

echo "====== git status ================"
git status
echo  ""
echo  ""

echo "====== git add ================"
git add .
echo  ""
echo  ""

echo "====== git commit ================"
git commit -m "1"
echo  ""
echo  ""

echo "====== git push origin main ================"        
git push origin main
echo  ""
echo  ""
echo "====== $(date '+%Y-%m-%d %H:%M:%S') ================"
