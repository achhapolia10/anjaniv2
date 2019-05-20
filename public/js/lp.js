
days=[]
function onSubmitClick(){
    to = document.getElementById("to-date").value
    from = document.getElementById("from-date").value
    if(from !=="" && to != ""){
    $.ajax({
        type: "POST",
        url: "/labourpayment",
        data: {day1: days[0],day2: days[1],day3: days[2],day4: days[3],day5: days[4],day6: days[5],day7: days[6]},
        success: function (response) {
            res=JSON.parse(response)
            createLabourPaymentTable(res.data)
        }
    }); } else {
        window.alert("Pick From and to Date")
    }
}

function onFromDateChange(){
    days=[]
    fromDate = document.getElementById("from-date").value
    from = new Date(fromDate)
    for (i=0;i<7;i++){
        days = days.concat(from.toISOString().slice(0,10))
        if (i<6){
        from.setDate(from.getDate()+1)
        }
    }
    toDate = document.getElementById("to-date")
    toDate.value=from.toISOString().slice(0,10);
}

function onToDateChange(){
    days=[]
    fromDate = document.getElementById("to-date").value
    from = new Date(fromDate)
    for (i=0;i<7;i++){
        days = days.concat(from.toISOString().slice(0,10))
        if(i<6){
        from.setDate(from.getDate()-1)
        }
    }
    days= days.reverse()
    toDate = document.getElementById("from-date")
    toDate.value=from.toISOString().slice(0,10);
}

function createLabourPaymentTable(arr){
}