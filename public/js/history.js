var areas = {
    "Tromsø"      : [69.680, 18.95],
    "Bodø"        : [67.28, 14.405],
    "Narvik"      : [68.438, 17.427],
    "Alta"        : [69.971, 23.303],
    "Nord-Troms"  : [69.929, 20.999],
    "Harstad"     : [68.798, 16.541],
    "Lakselv"     : [70.051, 24.971],
    "Mo i Rana"   : [66.313, 14.142]
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
  Plotly.purge("chart-dust");
  Plotly.purge("chart-humidity");
  Plotly.purge("chart-temperature");
  Plotly.purge("chart-PM10");
  Plotly.purge("chart-NO2");
  $("#student-data").hide()
  $("#student-title").hide()
  $("#nilu-title").hide()
  $("#nilu-data").hide()
  $("#infobox-student").hide()
  $("#infobox-PM10").hide()
  $("#infobox-NO2").hide()
}

function getCoordinates() {
  var areaObject = document.getElementById('area');
  area = areaObject.value
  coordinates = areas[area]
}

function drawMap() {
  return newMap(mapid, coordinates);
}

function drawCustomAreaMap() {
    var map = newMap(mapid, coordinates)
  // Initialise the FeatureGroup to store editable layers
    var drawnItems = new L.FeatureGroup();
    map.addLayer(drawnItems);

    var drawControlFull = new L.Control.Draw({
      position: 'topright',
      draw: {
        polyline: false,
        polygon: false,
        circle: { 
          shapeOptions: {
            color: '#f357a1',
            weight: 10,
            clickable: true
          },
          repeatMode: false
        }, 
        rectangle: false,
        marker: false,
      },
      edit: {
        featureGroup: drawnItems //REQUIRED!!
      }
    });

    var drawControlEditOnly = new L.Control.Draw({
      position: 'topright',
      edit: {
          featureGroup: drawnItems
      },
      draw: false
    });

    // Initialise the draw control and pass it the FeatureGroup of editable layers
    map.addControl(drawControlFull);

    map.on('draw:created', function(e) {
      layer = e.layer;
      layer.addTo(drawnItems);
      drawControlFull.remove(map);
      drawControlEditOnly.addTo(map)
      customGPS = true;
      radius = e.layer.getRadius()/1000
      coordinates = [e.layer.getLatLng()["lat"], e.layer.getLatLng()["lng"]]
    });

    map.on("draw:deleted", function(e) {
      if (drawnItems.getLayers().length === 0){
        drawControlEditOnly.remove(map);
        drawControlFull.addTo(map);
        customGPS = false;
      };
    }); 

    return map
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