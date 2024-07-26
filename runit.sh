
function installallqmllib() {
    lst=`find ~/.nix-profile/lib/qt-6/qml/ -name "*dylib"`
    # echo $lst
    for lib in $lst; do
        echo "$lib"
        cp -v "$lib" ./platformsthemes/
    done
}

# installallqmllib;
# exit;

export DYLD_LIBRARY_PATH=$PWD/

    # DYLD_PRINT_LIBRARIES=1 \
    # LD_LIBRARY_PATH=~/.nix-profile/lib/qt-6:./plugins/ \
    # DYLD_LIBRARY_PATH=~/.nix-profile/lib/qt-6:./plugins/ \
# export QT_PLUGIN_PATH=./plugins/  
# export QT_DEBUG_PLUGINS=1
    # DYLD_LIBRARY_PATH=$PWD/ \
export QML_IMPORT_PATH=~/.nix-profile/lib/qt-6/qml/
export QML2_IMPORT_PATH=$PWD/:$PWD/qmlpp

export QT_QUICK_CONTROLS_STYLE=Material
export QT_QUICK_CONTROLS_MATERIAL_THEME=Dark

export QV4_FORCE_INTERPRETER=1
# export QV4_SHOW_BYTECODE=1
export QML_DISABLE_DISK_CACHE=1

set -x
LD_LIBRARY_PATH=../minqtgo/srcc GOGC=10 ./main
# ./helloworld.app/Contents/MacOS/helloworld
# ./hellogo
# export QML_IMPORT_TRACE=1
# qmlprofiler  --verbose -o qmltr.qtd -p 3768 ./hellogo
# qmlprofiler --verbose  -p 3768 -attach 127.0.0.1

# the QML_IMPORT_PATH line works
# or error like this: module "QtQuick.Controls" plugin "qtquickcontrols2plugin" not found

