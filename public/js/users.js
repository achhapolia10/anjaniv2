function onChangePassword(username){
    p = $("#new-password").val()
    if(p!=""){
        $.ajax({
            url:"/users/edit",
            method:"POST",
            data:{username:username,password:p}
        })
    }
    $("#new-password").val("")
    window.location="http://localhost:4001/users"
}

function onDeleteClick(id,username){
        $.ajax({
            url:"/users/delete",
            method:"POST",
            data:{username:username,id:id}
        })
}