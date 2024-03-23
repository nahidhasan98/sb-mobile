$(document).ready(() => {
    $('#counterId, #stationId').select2({
        autoWidth: true,
    });
    $('#journeyDate').datepicker({
        minDate: 0,
    });
    $('#journeyDate').val(getFormattedDate(new Date()));

    $('#counterId').on('change', (event) => {
        $('#counterIdError').text("");
        getStations($(event.target).val());
    })

    $('#stationId').on('change', (event) => {
        $('#stationIdError').text("");
    })

    // Attach input event listener to the input field
    $('#journeyDate').on('input', (event) => {
        // Clear error message when input field is not empty
        if ($(event.target).val().trim() != "") {
            $('#journeyDateError').text("");
        }
    });

    // search form submit function
    $('.searchForm').on('submit', (event) => {
        event.preventDefault(); // Prevent default form submission

        if (!formValidate()) return false;

        let formData = {};
        $(event.target).serializeArray().forEach((field) => {
            formData[field.name] = field.value;
        });
        let jsonData = JSON.stringify(formData);
        // console.log(jsonData);

        getSchedule(jsonData);
        return false;
    });
});

function formValidate() {
    let isValidate = true;

    // taking care of counter name
    if ($('#counterId').val() == null || $('#counterId').val() == "") {
        $('#counterIdError').text("Please select a counter");
        isValidate = false;
    }
    // taking care of destination name
    if ($('#stationId').val() == null || $('#stationId').val() == "") {
        $('#stationIdError').text("Please select a destination");
        isValidate = false;
    }
    // taking care of date
    if ($('#journeyDate').val().trim() == "") {
        $('#journeyDateError').text("Please choose a date");
        isValidate = false;
    } else {
        let date = new Date($('#journeyDate').val().trim());
        if (isNaN(date.getTime())) {
            $('#journeyDateError').text("invalid date");
            isValidate = false;
        }
    }

    return isValidate
}

function getStations(counterID) {
    $('#stationId').prop('disabled', true);
    $('#stationIdLoading').show();

    // sending ajax post request
    let request = $.ajax({
        async: true,
        type: "GET",
        url: "/api/getStations/" + counterID,
    });
    request.done(function (response) {
        // console.log(response);

        response.data.unshift({ StationId: "", StationName: "Where To?" });
        // Remove existing options
        removeOptions($('#stationId'));

        // Add new options obtained from API
        addOptions($('#stationId'), response.data);
    });
    request.fail(function (response) {
        console.log(response.responseJSON.message)
    });
    request.always(function () {
        // console.log("always")
        $('#stationId').prop('disabled', false);
        $('#stationIdLoading').hide();
        $('#stationIdError').text("");
    });
}

function getSchedule(postData) {
    $('#findBtn').prop('disabled', true);
    $('#findBtnLoading').show();
    $("#scheduleTable tbody").css("opacity", "0.4");
    $('.loadingIconSection').show();

    // sending ajax post request
    let request = $.ajax({
        type: "POST",
        url: "/api/getSchedule",
        data: postData,
        dataType: 'json',
    });
    request.done(function (response) {
        // console.log(response);
        displayList(response.data);
    });
    request.fail(function (response) {
        console.log(response.responseJSON.message)
        // Clear existing table rows
        $("#scheduleTable tbody").empty();

        let newRow = $("<tr>").append(
            $("<td colspan='100%'>").text(response.responseJSON.message).addClass("noSchedule")
        );
        $("#scheduleTable tbody").append(newRow);
    });
    request.always(function () {
        // console.log("always")
        $('#findBtn').prop('disabled', false);
        $('#findBtnLoading').hide();
        $("#scheduleTable tbody").css("opacity", "");
        $('.loadingIconSection').hide();
    });
}

// Function to remove existing options
function removeOptions(selectElement) {
    $(selectElement).empty();
}

// Function to add new options obtained from API
function addOptions(selectElement, options) {
    $.each(options, (index, option) => {
        let selected = (option.StationId == 0) ? true : false;
        $(selectElement).append($('<option>', {
            value: option.StationId,
            text: option.StationName,
            selected: selected,
            disabled: selected
        }));
    });
}

function displayList(list) {
    // Clear existing table rows
    $("#scheduleTable tbody").empty();

    if (list == null) {
        let newRow = $("<tr>").append(
            $("<td colspan='100%'>").text("Sorry, No Schedule Found.").addClass("noSchedule")
        );
        $("#scheduleTable tbody").append(newRow);

        return;
    }

    // Iterate over each object in the response and append a row to the table
    $.each(list, (index, row) => {
        let newRow = $("<tr>").append(
            $("<td>").text(row.ScheduleName).css("border-right", "1px solid rgb(212 218 223)"),
            $("<td>").html("<span>" + formatTime(row.Time) + "</span><br><span class='smallText'>" + formatDate(row.DDate) + "</span>").css("border-right", "1px solid rgb(212 218 223)"),
            $("<td>").text(row.NumberOfSeat).css("border-right", "1px solid rgb(212 218 223)"),
            $("<td>").text(row.SeatUpdate).css("border-right", "1px solid rgb(212 218 223)"),
            $("<td>").html("<span>" + row.RouteName + "</span><br><span class='smallText'>" + row.BusDescription + " [" + row.NumberOfSeat + "]" + "</span>").css("border-right", "1px solid rgb(212 218 223)").css("min-width", "300px"),
            $("<td>").text(row.SeatFare).css("border-right", "1px solid rgb(212 218 223)"),
            $("<td>").append($("<button disabled>").text("View Seats").addClass("btn viewBtn").css("font-size", "12px"))
        );

        $("#scheduleTable tbody").append(newRow);
    });

    if (list && list.length > 0) $('#dateCol').text(formatDate(list[0].DDate)).css({ "min-width": "116px", "font-size": "16px", "font-weight": "bold", "background": "rgb(209 49 47)" });
}

function formatDate(inputDate) {
    // Split the input date into parts (month, day, year)
    let parts = inputDate.split('/');

    // Convert month to a name
    let monthNames = ["Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"];
    let monthIndex = parseInt(parts[0], 10) - 1; // Subtract 1 because monthNames are 0-based
    let monthName = monthNames[monthIndex];

    // Format the output date
    let outputDate = parts[1] + ' ' + monthName + ' ' + parts[2];

    return outputDate;
}

function formatTime(inputTime) {
    // Split the input time into parts (hours, minutes, seconds)
    let parts = inputTime.split(':');

    // Convert hours to 12-hour format
    let hours = parseInt(parts[0], 10);
    let ampm = hours >= 12 ? 'PM' : 'AM';
    hours = hours % 12;
    hours = hours ? hours : 12; // 0 should be treated as 12
    let minutes = parts[1];

    // Format the output time
    let outputTime = hours + ':' + minutes + ' ' + ampm;

    return outputTime;
}

function getFormattedDate(date) {
    var dd = String(date.getDate()).padStart(2, '0'); // Get the day and pad with leading zero if needed
    var mm = String(date.getMonth() + 1).padStart(2, '0'); // Get the month (January is 0) and pad with leading zero if needed
    var yyyy = date.getFullYear(); // Get the full year

    return mm + '/' + dd + '/' + yyyy;
}
