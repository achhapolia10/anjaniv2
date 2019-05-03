function onEntryFormSubmit(){
    
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