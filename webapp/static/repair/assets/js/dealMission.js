/**
 * GET请求服务器,接取预约信息
 * (此处使用GET并不复合REST规范,但是由于服务器端使用了Basic Auth验证,直接用ajax进行PUT请求会导致404错误,所以只能用GET了)
 * @param {*} that 按钮对象 
*/
function takeMission(that) {
    // 取出ID
    const ID = getMissionId(that);
    console.log("管理员接取ID为" + ID + "的预约");

    // 从当前URI参数中取出用户名和密码
    const username = getUrlParam("username");
    const password = getUrlParam("password");
    window.location.href = "./receive?ID=" + ID + "&username=" + username + "&password=" + password;
}

/**
 * GET请求服务器,申报当前任务已完成
 * @param {*} that 按钮对象 
*/
function finishMission(that) {
    // 取出ID
    const ID = getMissionId(that);
    console.log("管理员完成ID为" + ID + "的预约");

    // 从当前URI参数中取出用户名和密码
    const username = getUrlParam("username");
    const password = getUrlParam("password");
    window.location.href = "./complete?ID=" + ID + "&username=" + username + "&password=" + password;
}

/**
 * 获取当前行ID
 * @param {*} that 单元格内的按钮对象
 * @returns 当前行ID
 */
function getMissionId(that) {
    // 获取当前行对象
    const tr = that.parentElement.parentElement;

    // 取出ID
    return tr.cells[0].textContent;
}

function getUrlParam(name) {
    var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)"); //构造一个含有目标参数的正则表达式对象
    var r = window.location.search.substr(1).match(reg);  //匹配目标参数
    if (r != null) return unescape(r[2]); return null; //返回参数值
}

