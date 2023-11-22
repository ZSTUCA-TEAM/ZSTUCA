/**
 * 向服务器动态POST管理员信息,以注册
 * @returns 阻止表单提交
 */
function registerAdmin() {
    // 获取表单数据
    let admin = {};
    $.each($("#adminInfo").serializeArray(), function (index, item) {
        let value = item.value;
        admin[item.name] = value;
    });

    // 表单验证

    // 转为JSON格式
    let adminJson = JSON.stringify(admin);
    console.log(adminJson);

    // post请求,提交报修信息
    $.ajax({
        url: "/repair/bs/admin",
        type: "POST",
        async: true,
        data: adminJson,
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

    // 阻止表单提交
    return false;
}