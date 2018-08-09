#/usr/bin/env sh

# $2 should be the github url
# meaning options: "https://github.com/etc"

echo "build called "  >> log

GITHUB_URL=$(echo $2 | awk '{ print $1 }')
CLEANUP=$(echo $2 | awk '{ print $2 }')

echo "GITHUB IS $GITHUB_URL" >> log
echo "CLEANUP IS $CLEANUP" >> log

git clone $GITHUB_URL
sleep 20
rm -r $CLEANUP
