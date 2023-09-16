/**
 * 向服务器动态POST预约信息
 * @returns 组织表单提交
 */
function postRepair() {
    // 获取表单数据
    let apply = {};
    $.each($("#applyInfo").serializeArray(), function (index, item) {
        let value = item.value;
        apply[item.name] = value;
    });

    // 表单验证

    // 转为JSON格式
    let applyJson = JSON.stringify(apply);
    console.log(applyJson);

    // post请求,提交报修信息
    $.ajax({
        url: "/repair/apply",
        type: "POST",
        async: true,
        data: applyJson,
        dataType: "text",
        contentType: "application/json",
        success: function (result) {
            console.log("报修申请提交成功");
            alert("提交成功,已向您的邮箱发送了一份邮件,后续消息将通过邮箱通知您.如果您的邮箱自动拦截了这份邮件,请到垃圾箱中查看并设置正常接收来自我们的邮件.");
            // 刷新网页,防止表单重复提交
            window.location.reload();
        },
        error: ajaxErr,
    });

    // 阻止表单提交
    return false;
}