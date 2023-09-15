// 各div名称
const divNames = ["pendingTasks", "acceptedTasks", "hardwareMemberRegist"];

/**
 * 切换到指定的div
 * @param {String} targetName 指定div名称
 */
function switchTo(targetName) {
    $.each(divNames, function (index, value) {
        // 当前div
        let div = $("#" + value);
        if (value == targetName) {
            // 显示当前div
            div.attr("style", "");
            console.log("显示" + targetName);
        } else {
            // 隐藏当前div
            div.attr("style", "display:none!important");
        }
    });
}

$(function () {
    switchTo("pendingTasks");
});