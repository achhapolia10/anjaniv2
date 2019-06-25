function onFromDateChange(event) {
    fromDate= document.getElementById("from-date")
    toDate  = document.getElementById("to-date")
    toDate.min=fromDate.value
}

let q =0
function onToDateChange(event)  {
    fromDate= document.getElementById("from-date")
    toDate= document.getElementById("to-date")
    fromDate.max=toDate.value
}

function onSubmitClick(event){
    fromDate= document.getElementById("from-date")
    toDate= document.getElementById("to-date")
    product = document.getElementById("product")
    if (fromDate.value || toDate.value){
    $.ajax({
        type: "POST",
        url: "/stock/product?id="+product.value+"&fdate="+fromDate.value+"&tdate="+toDate.value,
        success: function (response) {
            res = JSON.parse(response)
            stock= res.stock
            $("#ostock").html(stock.obox + " Box "+ stock.opacket + " Packets" )
            $("#cstock").html(stock.cbox + " Box "+ stock.cpacket + " Packets" )
            createStockTable(res.entries)
        }
    })
        
    } else {
        window.alert("Enter the dates first")
    }
}

function createStockTable(entries){
    clearStockTable()
    table = document.getElementById("stock-table-body")    
        entries.forEach((s)=>{
        $("#stock-table-body").append(
           "<tr>"
            +       "<td>"+parseDate(s.date)+"</td>"
            +       "<td>"+s.boxin +" box "+s.packetin+" packet"+"</td>"
            +       "<td>"+s.boxout +" box "+s.packetout+" packet"+"</td>"
            +   "</tr>"
 
        )
    })

}

function clearStockTable(){
    table = document.getElementById("stock-table-body")
    table.innerHTML= ""
}

$(document).ready(function(){
    let query = parseQuery(window.location.search)
    fromDate= document.getElementById("from-date")
    toDate= document.getElementById("to-date")
    product = document.getElementById("product")
    q= query
    if(q.from && q.to && q.id)
    {
        fromDate.value= q.from
        toDate.value =q.to
        product.value = q.id
        onSubmitClick()
    }

})

function parseQuery(queryString) {
    var query = {};
    var pairs = (queryString[0] === '?' ? queryString.substr(1) : queryString).split('&');
    for (var i = 0; i < pairs.length; i++) {
        var pair = pairs[i].split('=');
        query[decodeURIComponent(pair[0])] = decodeURIComponent(pair[1] || '');
    }
    return query;
}

function parseDate(date) {
    return date
}