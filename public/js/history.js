var areas = {
    "Tromsø" : [69.680, 18.95],
    "Bodø"   : [67.28, 14.405],
    "Narvik" : [68.438,17.427] 

};

function toggleGPSInput(checkbox) {
  if (checkbox.checked){
    $('#gps-input').css('margin-top', '10px');
    $('#gps-input').css('display', 'flex');
    document.getElementById("area").disabled = true;
    $('#checkbox-nilu').attr('disabled', true);
  	$('#checkbox-nilu').attr('checked', false);
    $('#nilu-label').css('color', '#bfbfbf');
  } else {
    $('#gps-input').hide();
    $('#nilu-label').css('color', 'black');
    document.getElementById("area").disabled = false;
  	$('#checkbox-nilu').attr('disabled', false);
  }
}

function clearCharts() {
  $("svg#chart-PM10").empty()
  $("svg#chart-NO2").empty()
  $("svg#chart-dust").empty()
  $("svg#chart-temperature").empty()
  $("svg#chart-humidity").empty()
  // $("#nilu-data").hide();
  // $("#student-data").hide();
}

function getCoordinates() {
  var areaObject = document.getElementById('area');
  if (areaObject.disabled == true) {
    var latitude = document.getElementById('latitude').value;
    var longitude = document.getElementById('longitude').value;
    coordinates = [latitude, longitude]
    radius = document.getElementById('radius').value;
  } 
  else {
    area = areaObject.value
    coordinates = areas[area]
  }
}

function drawMap() {
  return newMap(mapid, coordinates);
}

function createDatestring() {
  var now = new Date()
  var to = now  ;
  var from = new Date();
  var selectedTime = $('input[name=time]:checked').val()
  switch(selectedTime) {
    case "hour":
      from.setHours(now.getHours()-1);
      break;
    case "24hour":
      from.setDate(now.getDate()-1);
      break;
    case "week":
      from.setDate(now.getDate()-7);
      break;
    case "month":
      from.setMonth(now.getMonth()-1);
      break;
    case "custom":
      from = new Date($('#from').val())
      to = new Date($('#to').val())
      break;
  }

  return "from=" + from.toJSON() + "&to=" + to.toJSON()
}