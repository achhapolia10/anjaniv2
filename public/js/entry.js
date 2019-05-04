function onDateChange(){
    date = document.getElementById("date-picker")
    if(isDateProductPicked())
    createEntryTable()
}

function onProductChange(){
    product  =document.getElementById("product-picker")
    if(isDateProductPicked())
    createEntryTable()
}

function isDateProductPicked(){
    product = document.getElementById("product-picker")
    date = document.getElementById("date-picker")
    return (Boolean)(date.value && product.value)
}

function onEntryFormSubmit(){
    var labour = document.getElementById("labour-name").value.toUpperCase()
    var box = document.getElementById("box-no").value
    var packet=document.getElementById("packet-no").value
    var date = document.getElementById("date-picker").value
    var product = document.getElementById("product-picker").value
    var res=isDateProductPicked()
    if(!res){
        alert("Pick Date and Product")
    } else {
        console.log("Entry to be Made")
        $.ajax({
            type:"POST",
            url:"/entry/new?labour="+labour+"\&box="+box+"\&packet="+packet+
            "\&date="+date+"&product="+ product,
            success:function(s){console.log("success from ");
        console.log(s)},
            error:function(){console.log("faluire from server")}
        })
    }   
    clearEntryForm();
}

function clearEntryForm(){
    nameControl = document.getElementById("labour-name")
    boxControl = document.getElementById("box-no")            
    packetControl = document.getElementById("packet-no")            
    nameControl.value=""
    boxControl.value=""
    packetControl.value=""
    nameControl.focus()
}

function createEntryTable(){
    console.log("Entry table will be created ")
}