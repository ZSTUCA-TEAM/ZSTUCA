let nowDiv;

/**
 * 切换到指定的div
 * @param {String} targetName 指定div名称
 */
function switchTo(divName) {
    divName = divName || nowDiv; // 设置默认值
    $.ajax({
        url: "/repair/bs/div/" + divName,
        type: "GET",
        async: true,
        dataType: "text",
        data: {
            showFinish: $('#showFinish').is(':checked'),
            showAbandoned: $('#showAbandoned').is(':checked'),
        },
        headers: {
            "Authorization": "Basic aGFyZHdhcmU6WmpsZ2R4MjMxMg==",
        },
        success: function (result) {
            if (divName == "pendingTasks" || divName == "adminRegister") {
                // 这些页面无过滤需求
                $("#filter").fadeOut(200);
            } else {
                // 这些页面有过滤需求
                $("#filter").fadeIn(200);
            }

            // 切换主界面
            $("#main").fadeOut(200, function () {
                $("#main").html(result);
                $("#main").fadeIn(200, function () {
                    try {
                        controlWidth("main");
                    } catch {
                    }
                });
            });
            nowDiv = divName;
        },
        error: ajaxErr,
    });
}

$(function () {
    if (window.location.hash != "") {
        switchTo(window.location.hash.substring(1));
    } else {
        switchTo("pendingTasks");
    }
});