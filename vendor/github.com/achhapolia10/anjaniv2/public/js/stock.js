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
            stocks = JSON.parse(response)
            createStockTable(stocks)
        }
    })
        
    } else {
        window.alert("Enter the dates first")
    }
}

function createStockTable(stocks){
    clearStockTable()
    fromDate= document.getElementById("from-date")
    toDate = document.getElementById("to-date")

    table = document.getElementById("stock-table-body")    
    keys = Object.keys(stocks)
    keys.sort((a,b)=>{
        p1=stocks[a].product.name
        p2=stocks[b].product.name
        if(p1>p2){
            return 1
        }   else if(p1 <p2){
            return -1
        } else return 0
    })
    var i=1
    keys.forEach((s)=>{
        $("#stock-table-body").append(
           "<tr>"
            +       "<td>"+i+".</td>"
            +       "<td>"+"<a class=\"product-anchor\" href=\"/stock/product?from="+fromDate.value+"&to="+toDate.value+"&id="+stocks[s].product.id+"\">"+stocks[s].product.name+"</a></td>"
            +       "<td>"+stocks[s].obox +" box "+stocks[s].opacket+" packet"+"</td>"
            +       "<td>"+stocks[s].inbox +" box "+stocks[s].inpacket+" packet"+"</td>"
            +       "<td>"+stocks[s].outbox +" box "+stocks[s].outpacket+" packet"+"</td>"
            +       "<td>"+stocks[s].cbox +" box "+stocks[s].cpacket+" packet"+"</td>"
            +   "</tr>"
 
        )
        i++
    })

}

function clearStockTable(){
    table = document.getElementById("stock-table-body")
    table.innerHTML= ""
}
