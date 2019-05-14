function onSubmitClick(){
}

function onFromDateChange(){
    fromDate = document.getElementById("from-date").value
    from = new Date(fromDate)
    from.setDate(from.getDate()+7)
    toDate = document.getElementById("to-date")
    toDate.value=from.toISOString().slice(0,10);
    console.log("1")
}

function onToDateChange(){
    fromDate = document.getElementById("to-date").value
    from = new Date(fromDate)
    from.setDate(from.getDate()-7)
    toDate = document.getElementById("from-date")
    toDate.value=from.toISOString().slice(0,10);
}