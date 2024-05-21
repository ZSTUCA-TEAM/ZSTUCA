/**
 * 控制div内每行元素一样长
 * 
 * @param {string} div div名称
 */
function controlWidth(div) {
    let that = $("#" + div);
    // 所有子div
    let cards = that.children();
    // 转换为jQuery对象
    $.each(cards, function (index, card) {
        cards[index] = $(card);
    });

    // 计算一行有几个div
    let lineNumber = parseInt(that.width() / cards[0].width());
    let maxHeight = 0;
    for (let i = 0; i < cards.length; i += lineNumber) {
        // 取当前行最长元素长度
        let j = i;
        do {
            maxHeight = Math.max(maxHeight, cards[j++].height());
        } while (j < cards.length && j % lineNumber != 0);

        // 将当前行元素都设置成最长元素的长度
        j = i;
        do {
            cards[j++].height(maxHeight);
        } while (j < cards.length && j % lineNumber != 0);
        maxHeight = 0;
    }
}

$(function () {
    $(window).resize(function () {
        controlWidth("main");
    });
})