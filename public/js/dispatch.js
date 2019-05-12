function onDateChange(){
   date = document.getElementById("date-picker")
   createDispatchTable()
   product = document.getElementById("product-name")
   product.focus()
}

function isDatePicked(){
    date = document.getElementById("date-picker")
    return (Boolean)(date.value)
}

function onDispatchSubmit(){
  if(isDatePicked()){
      product = document.getElementById("product-name").value;
      box     = document.getElementById("box").value
      packet  = document.getElementById("packet").value
      clearForm();
      
  }else{
     alert("Pick the date")
  } 
}

function clearForm(){
   product = document.getElementById("product-name")
      box     = document.getElementById("box")
      packet  = document.getElementById("packet")
      product.value=1 
      box.value=""
      packet.value=""
      product.focus()

}

function createDispatchTable(){
   
}