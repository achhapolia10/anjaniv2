
days=[]
data=[]
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
            printAnchorChange()
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
    updateTableDays()   
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
    updateTableDays()
    toDate.value=from.toISOString().slice(0,10);
}

function updateTableDays(){
    date1= new Date(days[0]);
    $("#day-1").html(getDay(date1));
    date2= new Date(days[1]);
    $("#day-2").html(getDay(date2));
    date3= new Date(days[2]);
    $("#day-3").html(getDay(date3));
    date4= new Date(days[3]);
    $("#day-4").html(getDay(date4));
    date5= new Date(days[4]);
    $("#day-5").html(getDay(date5));
    date6= new Date(days[5]);
    $("#day-6").html(getDay(date6));
    date7= new Date(days[6]);
    $("#day-7").html(getDay(date7));
    
}
function getDay(date){
    
    switch(date.getDay()){
        case 0:
            return "Sunday"
        break;
        case 1:
            return "Monday"
        break;
        case 2:
            return "Tuesday"
        break;
        case 3:
            return "Wednesday"
        break;
        case 4:
            return "Thursday"
        break;
        case 5:
            return "Friday"
        break;
        case 6:
            return "Saturday"
        break;
        
    }
}

function createLabourPaymentTable(arr){
    $('#lp-body').html('')
    data=arr
    console.log(arr)
    arr.sort((a,b)=> a.labour>b.labour?1:-1)
    console.log(arr)
    arr.forEach( l => {
        $("#lp-body").append("<tr>"+
                        "<td>"+l.labour+"</td>"+
                        "<td>"+l.day1+"</td>"+
                        "<td>"+l.day2+"</td>"+
                        "<td>"+l.day3+"</td>"+
                        "<td>"+l.day4+"</td>"+
                        "<td>"+l.day5+"</td>"+
                        "<td>"+l.day6+"</td>"+
                        "<td>"+l.day7+"</td>"+
                        "<td>"+l.total.toFixed(2)+"</td>"+
                            "</tr>")
    });
}

function printAnchorChange(){
    print= document.getElementById("printAnchor");
    print.href=("http://localhost:4001/labourpayment/print?day1="+days[0]+"&day2="+days[1]+"&day3="+days[2]+"&day4="+days[3]+"&day5="+days[4]+"&day6="+days[5]+"&day7="+days[6]);
}