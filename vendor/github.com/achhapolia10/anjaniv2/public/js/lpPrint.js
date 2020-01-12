$(document).ready(
    ()=>{
        printDates();
    }
)


function printDates(){
    from1= $("#from-date").html();
    to1 = $("#to-date").html()
    from = new Date(from1)
    to = new Date(to1)
    console.log(from)
    $("#from-date").html(from.getDate()+"/"+from.getMonth()+"/"+from.getYear())
    $("#to-date").html(to.getDate()+"/"+to.getMonth()+"/"+to.getYear())
    for(i=0;i<7;i++){
    $("#day"+(i+1)).html(getDay(from))
    from.setDate(from.getDate()+1)

    }

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