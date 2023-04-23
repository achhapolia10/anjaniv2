function onSubmitClick() {
    date = document.getElementById("sdate")
    if (date.value) {
        $.ajax({
            type: "POST",
            url: "/report",
            data: { fdate: date.value },
            success: function (response) {
                reports = JSON.parse(response)
                reports.sort(function (a, b) {
                    if (a.product.name > b.product.name) {
                        return 1
                    } else if (a.product.name < b.product.name) {
                        return -1
                    } else {
                        return 0
                    }
                })
                createReportTable(reports)
            }
        })
    } else {

    }
}

function createReportTable(reports) {
    i = 0;
    let t_box = 0
    let t_raw = 0
    let t_plastic = 0
    $("#tbody").html("")
    reports.forEach(r => {
        t_box += r.boxin
        t_raw += r.raw
        t_plastic += r.plastic
        i += 1
        $("#tbody").append(
            "<tr>"
            + "<td>" + i + ". </td>"
            + "<td>" + r.product.name + "</td>"
            + "<td>" + r.boxin + "</td>"
            + "<td>" + r.raw + " Kg</td>"
            + "<td>" + r.plastic + " Kg</td>"
            + "</tr>"
        )
    });
    $("#tbody").append(
        "<tr>"
        + "<td colspan=2><b>Total</b></td>"
        + "<td>" + t_box + "</td>"
        + "<td>" + t_raw.toFixed(3) + " Kg</td>"
        + "<td>" + t_plastic.toFixed(3) + " Kg</td>"
        + "</tr>"
    )
}