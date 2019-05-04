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
    clearEntryForm();
    var res=isDateProductPicked()
    if(!res){
        alert("Pick Date and Product")
    } else {
        console.log("Entry to be Made")
        $.ajax({
            type:"POST",
            url:"/entry/new?name=anshu",
            success:function(s){console.log("success from ");
        console.log(s)},
            error:function(){console.log("faluire from server")}
        })
    }   
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