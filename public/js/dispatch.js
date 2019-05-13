var se = []

function onDateChange() {
  date = document.getElementById("date-picker");
  createDispatchTable();
  product = document.getElementById("product-name");
  product.focus();
}

function isDatePicked() {
  date = document.getElementById("date-picker");
  return Boolean(date.value);
}

function onRemoveClick(pid,date){
  $.ajax({
    type: "get",
    url: "/dispatch/delete?date="+date+"&product="+pid,
    success: function (response) {
      createDispatchTable()
    }
  });
}

function onDispatchSubmit() {
  if (isDatePicked()) {
    product = document.getElementById("product-name").value.trim();
    i=se.indexOf(product)
    box = document.getElementById("box").value;
    packet = document.getElementById("packet").value;
    date = document.getElementById("date-picker").value;
    console.log(date)
    if(i<0){
    $.ajax({
      type: "POST",
      url: "/dispatch/new",
      data: { product: product, box: box, packet: packet, date: date },
      success: function (response) {$("#journal-table").html("")
        r = JSON.parse(response)
        if (r.status == 301) {
          createDispatchTable()
        }
      }
    });
  } else {
    alert("Already Entred for the product")
  }
    clearForm();
  } else {
    alert("Pick the date");
  }
}

function clearForm() {
  product = document.getElementById("product-name");
  box = document.getElementById("box");
  packet = document.getElementById("packet");
  product.value = 1;
  box.value = "";
  packet.value = "";
  product.focus();
}

function createDispatchTable() {
  $("#table-body").html("");
  date = document.getElementById("date-picker").value;
  tbody = document.getElementById("table-body")
  se = []
  $.ajax({
    type: "GET",
    url: "/dispatch/entries?date=" + date,
    success: function (response) {
          $("#table-body").append(response)
    }
  })
}
