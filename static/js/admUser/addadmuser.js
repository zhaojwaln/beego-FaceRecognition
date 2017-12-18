$(function () {
    addAdmUserObj = {
        search: function () {
            $('#admusergroup').datagrid('load', {
                groupName: $('input[name="addAdmUser_UserGroupName"]').val()         
            });
        }
    }

    //datagrid初始化
    $('#admusergroup').datagrid({
        url: 'admusergroup/gridlist',
        // queryParams: { roleid: 0 },
        iconCls: 'icon-edit',//图标
        width: 700,
        height: 600,
        nowrap: false,
        striped: true,
        border: true,
        collapsible: false,//是否可折叠的
        fit: true,//自动大小
        //sortName: 'code',
        //sortOrder: 'desc',
        remoteSort: false,
        idField: 'id',
        singleSelect: false,//是否单选
        pagination: true,//分页控件
        rownumbers: true,//行号
        fitColumns: true,//列宽自适应（列设置width=100）
        frozenColumns: [[
            { field: 'ck', checkbox: true }
        ]],//设置表单复选框
        toolbar: addAdmUser_toolbar
    });
})

function submitAddAmdUserForm() {
    var selections = $('#admusergroup').datagrid('getSelections')
    if (selections.length == 0) {
        $.messager.alert('操作提示', "请至少选择一个组", 'info');
        return false
    }

    var idArray = new Array(selections.length)
    for (var i = 0; i < selections.length; i++) {
        idArray[i] = selections[i].id
    }
    ids = idArray.join(",")

    url = "/admuser/addadmuser"
    var data = {
        ids: ids,
        userId: $("input[name='admUserId']").val(),
        userName: $("input[name='admUserName']").val(),
    };

    if (data.userId.length < 1 || data.userName.length < 1) {
        $.messager.alert('操作提示', "信息填写不完整,请补充后重新提交", 'info');
        return
    }

    $.post(url, data, function (result) {
        if (result == "success") {
            $('#addAdmUser').window("close")
            $.messager.alert('操作提示', "添加成功", 'info');
            loadAdmUserGrid()
        } else {
            $.messager.alert('操作提示', result, 'info');
        }
    });
}

function clearAddAmdUserForm() {
    $('#addAdmUser').form('clear');
}

function loadAdmUserGrid() {
    $('#admUser_list').datagrid('load', {
    });
}