function onSubmitClick(){
    date = document.getElementById("sdate")
    if(date.value){
    $.ajax({
            type: "POST",
            url: "/report",
            data:{fdate:date.value},
            success: function (response) {
                reports = JSON.parse(response)
                reports.sort(function(a,b){
                    if(a.product.name>b.product.name){
                        return 1
                    } else if (a.product.name < b.product.name){
                        return -1
                    } else {
                        return 0
                    }
                })
                createReportTable(reports)
            }
        })
    } else{

    }
}

function createReportTable(reports){
    i=1;
    $("#tbody").html("")
    reports.forEach(r => {
        $("#tbody").append(
                    "<tr>"
                +       "<td>"+i+". </td>"
                +       "<td>"+r.product.name+"</td>"
                +       "<td>"+r.boxin+"</td>"
                +       "<td>"+r.raw +" Kg</td>"
                +       "<td>"+r.plastic+" Kg</td>"
                +   "</tr>"
        )
    });
}