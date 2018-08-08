#/usr/bin/env sh

(
	cd ../common/hooker
	./build.sh
)

cp ../common/hooker/build/hooker ./src
