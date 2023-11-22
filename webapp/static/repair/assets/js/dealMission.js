/**
 * GET请求服务器,接取预约信息
 * (此处使用GET并不复合REST规范,但是由于服务器端使用了Basic Auth验证,直接用ajax进行PUT请求会导致404错误,所以只能用GET了)
 * @param {*} that 按钮对象 
*/
function takeMission(that) {
    // 取出ID
    let apply = {
        "id": parseInt($(that).attr("aid"))
    };
    console.log("管理员接取ID为" + apply.id + "的预约");

    // 转为JSON格式
    let applyJson = JSON.stringify(apply);
    console.log(applyJson);

    // post请求,提交报修信息
    $.ajax({
        url: "/repair/bs/receive",
        type: "POST",
        async: true,
        data: applyJson,
        dataType: "text",
        headers: {
            "Authorization": "Basic aGFyZHdhcmU6WmpsZ2R4MjMxMg==",
        },
        contentType: "application/json",
        success: function (result) {
            alert(result);
            switchTo("pendingTasks");
        },
        error: ajaxErr,
    });
}

function abandonMission(that) {
    // 二次确认是否放弃
    if (!confirm("你确认要放弃这个委托吗?放弃后这个委托将移除计协硬件部,所有人都不可再接取这个委托.请确认与对方进行过充分交流.")) {
        return;
    }

    // 取出ID
    let apply = {
        "id": parseInt($(that).attr("aid"))
    };
    console.log("管理员放弃ID为" + apply.id + "的预约");

    // 转为JSON格式
    let applyJson = JSON.stringify(apply);
    console.log(applyJson);

    // post请求,提交报修信息
    $.ajax({
        url: "/repair/bs/abandoned",
        type: "POST",
        async: true,
        data: applyJson,
        dataType: "text",
        headers: {
            "Authorization": "Basic aGFyZHdhcmU6WmpsZ2R4MjMxMg==",
        },
        contentType: "application/json",
        success: function (result) {
            alert(result);
            switchTo("receiptedTasks");
        },
        error: ajaxErr,
    });
}

/**
 * GET请求服务器,申报当前任务已完成
 * @param {*} that 按钮对象 
*/
function finishMission(that) {
    // 取出ID
    const ID = $(that).attr("aid");
    console.log("管理员完成ID为" + ID + "的预约");

    window.location.href = "/repair/bs/finish/" + ID;
}

/**
 * GET请求服务器,申报当前任务已完成
 *
 * 不阻止表单验证,以跳转页面
*/
function submitMissionInfo() {
    // 取出ID
    let worklist = {
        "applyId": parseInt($("#completeDeclaration").attr("aid"))
    };
    console.log("管理员完成ID为" + worklist.applyId + "的预约");

    // 获取表单数据
    $.each($("#completeDeclaration").serializeArray(), function (index, item) {
        let value = item.value;
        worklist[item.name] = value;
    });

    // 表单验证

    // 转为JSON格式
    let worklistJson = JSON.stringify(worklist);
    console.log(worklistJson);

    // post请求,提交报修信息
    $.ajax({
        url: "/repair/bs/finish",
        type: "POST",
        async: true,
        data: worklistJson,
        dataType: "text",
        headers: {
            "Authorization": "Basic aGFyZHdhcmU6WmpsZ2R4MjMxMg==",
        },
        contentType: "application/json",
        success: function (result) {
            alert(result);
        },
        error: ajaxErr,
    });
}