jsGrid.setDefaults({
    width:"100%",
    autoload:true,
    paging:true,
    sorting:true,
    pageLoading:true,
    pageNextText:">>",
    pagePrevText:"<<",
    pageFirstText:"首页",
    pageLastText:"尾页",
    pagerFormat:"第 {pageIndex} / {pageCount} 页 &nbsp;&nbsp; {first} {prev} {pages} {next} {last} &nbsp;&nbsp; 共 {itemCount} 个记录 ",
});

var datePickerLocale={
    applyLabel: '确认',
    cancelLabel: '取消',
    daysOfWeek:["日","一","二","三","四","五","六"],
    monthNames: ["一月","二月","三月","四月","五月","六月","七月","八月","九月","十月","十一月","十二月"],
    separator : ' 至 ',
    format:'YYYY-MM-DD'
};

var recentDays=30;