import QtQml
import QtQml.WorkerScript
import QtQuick
import QtQuick.Controls
import QtQuick.Controls.Material
import QtQuick.Layouts
import QtQuick.Window


// mobile mode, 500, 650
// desktop mode, 800, 650

ApplicationWindow {

    ////////////
    id: appwin
    objectName: "appwin"
    // minimumWidth: 300
    // minimumHeight: parent.height
    // width: 500
    // height: 650
    width: Qt.platform.os=='android' ? Screen.width : 880;
    height: Qt.platform.os=='android' ? Screen.height : 650;
    // maximumHeight: 3800
    // height: parent.height // not work
    visible: true
    // color: "#010101"

    Material.theme: Material.Dark
    // Material.accent: Material.Purple
    // Material.foreground : "red"


}
