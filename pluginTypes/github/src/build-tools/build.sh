#/usr/bin/env sh

# $2 should be the github url
# meaning options: "https://github.com/etc"

echo "build called "  >> log
 
GITHUB_URL=$(echo $OPTIONS | awk '{ if ($1 == "repo") { print($2) }}')
CLEAN=$(echo $OPTIONS | awk '{ if ($1 == "clean") { print($2) }}')

echo "GITHUB IS $GITHUB_URL" >> log
echo "CLEANUP IS $CLEANUP" >> log

git clone $GITHUB_URL
sleep 5

rm -r $CLEAN

$FINISHED


