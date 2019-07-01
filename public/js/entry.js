var labours;
function onDateChange() {
  date = document.getElementById("date-picker");
  if (isDateProductPicked()) createEntryTable();
}

function onProductChange() {
  product = document.getElementById("product-picker");
  if (isDateProductPicked()) createEntryTable();
}

function isDateProductPicked() {
  product = document.getElementById("product-picker");
  date = document.getElementById("date-picker");
  return Boolean(date.value && product.value);
}

function onEntryFormSubmit() {
  var labour = document
    .getElementById("labour-name")
    .value.toUpperCase()
    .trim();
  var box = document.getElementById("box-no").value;
  var packet = document.getElementById("packet-no").value;
  var date = document.getElementById("date-picker").value;
  var product = document.getElementById("product-picker").value;
  var res = isDateProductPicked();
  var i = labours.indexOf(labour);
  if (i < 0) {
    if (!res) {
      alert("Pick Date and Product");
    } else {
      console.log("Entry to be Made");
      $.ajax({
        type: "POST",
        url:
          "/entry/new?labour=" +
          labour +
          "&box=" +
          box +
          "&packet=" +
          packet +
          "&date=" +
          date +
          "&product=" +
          product,
        success: function(s) {
          createEntryTable();
        },
        error: function() {
          console.log("faluire from server");
        }
      });
    }
  } else {
    alert("Name Already exist");
  }
  clearEntryForm();
}

function clearEntryForm() {
  nameControl = document.getElementById("labour-name");
  boxControl = document.getElementById("box-no");
  packetControl = document.getElementById("packet-no");
  nameControl.value = "";
  boxControl.value = "";
  packetControl.value = "";
  nameControl.focus();
}

function createEntryTable() {
  $("#journal-table").html("");
  var date = document.getElementById("date-picker").value;
  var product = document.getElementById("product-picker").value;
  $.ajax({
    type: "GET",
    url: "/entry/getall?date=" + date + "&id=" + product,
    success: function(p) {
      labours = [];
      let response = JSON.parse(p);
      
      names = Object.keys(response.labours)
      autocomplete(document.getElementById("labour-name"), names);


      entries = response.entries
      if (entries)
        entries.forEach(entry => {
          $("#journal-table").append(
            "<tr><td>" +
              entry.labour +
              "</td><td>" +
              entry.box +
              "</td>" +
              "<td>" +
              entry.packet +
              '</td><td><button onclick="RemoveEntry(' +
              entry.id +
              "," +
              entry.product +
              ')" ' +
              'class="btn btn-danger">Remove</button>'
          );
          labours = labours.concat(entry.labour);
        });
        totalmark = document.getElementById("total-entries")
        let box = response.tbox;
        let packet= response.tpacket
        if (box ==0 && packet == 0){
          totalmark.innerHTML = ""
        } else {
          totalmark.innerHTML = "Total: "+ box +" Boxes "+ packet+" Packets"
        }
    },
    error: function() {
      console.log("error in getiing jounral data");
    }
  });
}

function RemoveEntry(id, productID) {
  console.log(id, productID);
  $.ajax({
    type: "post",
    url: "/entry/delete?productID=" + productID + "&id=" + id,
    success: function(p) {
      console.log("product Deleted");
      createEntryTable();
    },
    error: function() {
      console.log("error in getiing jounral data");
    }
  });
}




function autocomplete(inp, arr) {
  /*the autocomplete function takes two arguments,
  the text field element and an array of possible autocompleted values:*/
  var currentFocus;
  /*execute a function when someone writes in the text field:*/
  inp.addEventListener("input", function(e) {
      var a, b, i, val = this.value;
      /*close any already open lists of autocompleted values*/
      closeAllLists();
      if (!val) { return false;}
      currentFocus = -1;
      /*create a DIV element that will contain the items (values):*/
      a = document.createElement("DIV");
      a.setAttribute("id", this.id + "autocomplete-list");
      a.setAttribute("class", "autocomplete-items");
      /*append the DIV element as a child of the autocomplete container:*/
      this.parentNode.appendChild(a);
      /*for each item in the array...*/
      for (i = 0; i < arr.length; i++) {
        /*check if the item starts with the same letters as the text field value:*/
        if (arr[i].substr(0, val.length).toUpperCase() == val.toUpperCase()) {
          /*create a DIV element for each matching element:*/
          b = document.createElement("DIV");
          /*make the matching letters bold:*/
          b.innerHTML = "<strong>" + arr[i].substr(0, val.length) + "</strong>";
          b.innerHTML += arr[i].substr(val.length);
          /*insert a input field that will hold the current array item's value:*/
          b.innerHTML += "<input type='hidden' value='" + arr[i] + "'>";
          /*execute a function when someone clicks on the item value (DIV element):*/
              b.addEventListener("click", function(e) {
              /*insert the value for the autocomplete text field:*/
              inp.value = this.getElementsByTagName("input")[0].value;
              /*close the list of autocompleted values,
              (or any other open lists of autocompleted values:*/
              closeAllLists();
          });
          a.appendChild(b);
        }
      }
  });
  /*execute a function presses a key on the keyboard:*/
  inp.addEventListener("keydown", function(e) {
      var x = document.getElementById(this.id + "autocomplete-list");
      if (x) x = x.getElementsByTagName("div");
      if (e.keyCode == 40) {
        /*If the arrow DOWN key is pressed,
        increase the currentFocus variable:*/
        currentFocus++;
        /*and and make the current item more visible:*/
        addActive(x);
      } else if (e.keyCode == 38) { //up
        /*If the arrow UP key is pressed,
        decrease the currentFocus variable:*/
        currentFocus--;
        /*and and make the current item more visible:*/
        addActive(x);
      } else if (e.keyCode == 13) {
        /*If the ENTER key is pressed, prevent the form from being submitted,*/
        e.preventDefault();
        if (currentFocus > -1) {
          /*and simulate a click on the "active" item:*/
          if (x) x[currentFocus].click();
        }
        document.getElementById("box-no").focus()
        closeAllLists()
      }
  });
  function addActive(x) {
    /*a function to classify an item as "active":*/
    if (!x) return false;
    /*start by removing the "active" class on all items:*/
    removeActive(x);
    if (currentFocus >= x.length) currentFocus = 0;
    if (currentFocus < 0) currentFocus = (x.length - 1);
    /*add class "autocomplete-active":*/
    x[currentFocus].classList.add("autocomplete-active");
  }
  function removeActive(x) {
    /*a function to remove the "active" class from all autocomplete items:*/
    for (var i = 0; i < x.length; i++) {
      x[i].classList.remove("autocomplete-active");
    }
  }
  function closeAllLists(elmnt) {
    /*close all autocomplete lists in the document,
    except the one passed as an argument:*/
    var x = document.getElementsByClassName("autocomplete-items");
    for (var i = 0; i < x.length; i++) {
      if (elmnt != x[i] && elmnt != inp) {
      x[i].parentNode.removeChild(x[i]);
    }
  }
}
/*execute a function when someone clicks in the document:*/
document.addEventListener("click", function (e) {
    closeAllLists(e.target);
});
}