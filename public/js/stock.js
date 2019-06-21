function onFromDateChange(event) {
    fromDate= document.getElementById("from-date")
    toDate  = document.getElementById("to-date")
    toDate.min=fromDate.value
}

function onToDateChange(event)  {
    fromDate= document.getElementById("from-date")
    toDate= document.getElementById("to-date")
    fromDate.max=toDate.value
}

function onSubmitClick(event){
    fromDate= document.getElementById("from-date")
    toDate= document.getElementById("to-date")
    if (fromDate.value || toDate.value){
    $.ajax({
        type: "POST",
        url: "/stock",
        data: {fdate:fromDate.value,tdate:toDate.value},
        success: function (response) {
            console.log("hi")
        }
    })
        
    } else {
        window.alert("Enter the dates first")
    }


}

$(document).ready(function(){
    
})