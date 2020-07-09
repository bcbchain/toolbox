#!/usr/bin/env bash

# Determine the arch/os combos we're building for
XC_ARCH=${XC_ARCH:-"amd64"}
XC_OS=${XC_OS:-"darwin linux windows"}

TOOL0=${TOOL0:-"bcc bccrpcservice"}
TOOL1=${TOOL1:-"addr bcw txsampleV2 walv1tov2 genesis txparse"}
TOOL2=${TOOL2:-"methodid orgid sigorg smccheck smcpack"}
TOOL3=${TOOL3:-"bbm bcparser bcscan relay"}

WTOOL0=${WTOOL0:-"bcc.exe bccrpcservice.exe"}
WTOOL1=${WTOOL1:-"addr.exe bcw.exe txsampleV2.exe walv1tov2.exe genesis.exe txparse.exe"}
WTOOL2=${WTOOL2:-"methodid.exe orgid.exe sigorg.exe smccheck.exe smcpack.exe"}
WTOOL3=${WTOOL3:-"bbm.exe bcparser.exe bcscan.exe relay.exe"}

#BUILD_TAGS=-ldflags="-w -s" #todo
echo "==> Building..."
cd ..
rm -rf build
mkdir -p build/dist

IFS=' ' read -ra arch_list <<< "$XC_ARCH"
IFS=' ' read -ra os_list <<< "$XC_OS"

IFS=' ' read -ra tool_list0 <<< "$TOOL0"
IFS=' ' read -ra tool_list1 <<< "$TOOL1"
IFS=' ' read -ra tool_list2 <<< "$TOOL2"
IFS=' ' read -ra tool_list3 <<< "$TOOL3"

IFS=' ' read -ra wtool_list0 <<< "$WTOOL0"
IFS=' ' read -ra wtool_list1 <<< "$WTOOL1"
IFS=' ' read -ra wtool_list2 <<< "$WTOOL2"
IFS=' ' read -ra wtool_list3 <<< "$WTOOL3"

for arch in "${arch_list[@]}"; do
	for os in "${os_list[@]}"; do
		echo "--> $os/$arch"
		mkdir -p build/bin/"${os}_${arch}"
		cp -r bundle/.config build/bin/"${os}_${arch}"

		GOOS=${os} GOARCH=${arch} go build -tags="${BUILD_TAGS}" -o "build/bin/${os}_${arch}" ./...

		pushd "build/bin/${os}_${arch}" > /dev/null || exit 1
		if [[ $os == "windows" ]];then
      tar -zcvf ../../dist/$project_name\_bcc\_$VERSION\_${os}_${arch}.tar.gz "${wtool_list0[@]}" .config
      tar -zcvf ../../dist/$project_name\_wal\_$VERSION\_${os}_${arch}.tar.gz "${wtool_list1[@]}" .config
      tar -zcvf ../../dist/$project_name\_contract\_$VERSION\_${os}_${arch}.tar.gz "${wtool_list2[@]}" .config
      tar -zcvf ../../dist/$project_name\_scan\_$VERSION\_${os}_${arch}.tar.gz "${wtool_list3[@]}" .config
		else
      tar -zcvf ../../dist/$project_name\_bcc\_$VERSION\_${os}_${arch}.tar.gz "${tool_list0[@]}" .config
      tar -zcvf ../../dist/$project_name\_wal\_$VERSION\_${os}_${arch}.tar.gz "${tool_list1[@]}" .config
      tar -zcvf ../../dist/$project_name\_contract\_$VERSION\_${os}_${arch}.tar.gz "${tool_list2[@]}" .config
      tar -zcvf ../../dist/$project_name\_scan\_$VERSION\_${os}_${arch}.tar.gz "${tool_list3[@]}" .config
		fi

    popd > /dev/null || exit 1
	done
done