/**
 * 通用ajax错误处理
 * @param {XMLHttpRequest} xhr 
 */
function ajaxErr(xhr) {
    console.log("请求发送失败,error:" + xhr.responseText);
    if (xhr.status === 400) {
        window.location.href = "/errorPage/400.html";
    } else if (xhr.status === 401) {
        window.location.href = "/errorPage/401.html";
    } else if (xhr.status === 404) {
        window.location.href = "/errorPage/404.html";
    } else if (xhr.status == 409) {
        alert("请勿重复提交,如果你没有重复提交,请无视此条提醒.");
        window.location.reload();
    }
    else if (xhr.status === 500) {
        window.location.href = "/errorPage/500.html";
    } else if (xhr.status === 502) {
        window.location.href = "/errorPage/502.html";
    } else {
        alert("请求发送失败,请检查您的浏览器状态与网络状态.");
        window.location.reload();
    }
}