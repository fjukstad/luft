// Create new map that we can add data points to. 
function newMap(id, coordinates) {
    map = L.map(id, {
      editable: true,
    })
    map.setView(coordinates, 11.0);

    var accessToken = 'pk.eyJ1IjoiZmp1a3N0YWQiLCJhIjoiY2l2Mnh3azRvMDBrYTJ5bnYxcDAzZ3Z0biJ9.RHb5ENfbmzN65gjiB-L_wg';

    L.tileLayer(
        "https://api.mapbox.com/styles/v1/mapbox/dark-v9/tiles/256/{z}/{x}/{y}?access_token=" + accessToken, {
            attribution: 'Map data &copy; <a href="http://openstreetmap.org">OpenStreetMap</a> contributors, <a href="http://creativecommons.org/licenses/by-sa/2.0/">CC-BY-SA</a>, Imagery © <a href="http://mapbox.com">Mapbox</a>',
            maxZoom: 18,
            id: 'fjukstad.2148odo2',
            accessToken: accessToken
        }).addTo(map);

    return map;
}

function addToMap(map, area, customGPS, provider, component, datestring) {

    function onEachFeature(feature, layer) {
        // does this feature have a property named popupContent?
        if (feature.properties) {
            var content = ""
            if (feature.properties.name) {
                content = "<b>" + feature.properties.name + "</b><br />"
            }
            if (feature.properties.component) {
                content += feature.properties.component + ": " + feature.properties.value + "<br />"
            }
            if (feature.properties.pmTen) {
                content += "PM10: " + feature.properties.pmTen + "<br />"
                content += "PM2.5: " + feature.properties.pmTwoFive + "<br />"
                content += "Temperature: " + feature.properties.temperature + "<br />"
                content += "Humidity: " + feature.properties.humidity + "<br />"
            }
            content += feature.properties.date
            layer.bindPopup(content);
        }
    }

    var colorScheme = d3.scaleOrdinal(d3.schemeCategory20);
    var areaUri = encodeURIComponent(area);
    var url = ""
    if (provider == "nilu") {
        url = "/" + provider + "aqis?area=" + areaUri + "&" + datestring + "&component=" + component
    } else {
        if (customGPS == true) {
            url = "/" + provider + "aqis?within=" + areaUri + "&" + datestring + "&plotmap=true" 
        } else {
            url = "/" + provider + "aqis?area=" + areaUri + "&" + datestring + "&plotmap=true"
        }
    }
    $.ajax({
        dataType: "json",
        url: url,
        success: function (data) {
            var layer = L.geoJSON(data.features, {
                pointToLayer: function (feature, latlng) {
                    var color = ""
                    if (!feature.properties.color) {
                        color = colorScheme(feature.properties.name)
                    } else {
                        color = "#" + feature.properties.color
                    }
                    var geojsonMarkerOptions = {
                        color: color,
                        weight: feature.properties.weight,
                        opacity: 0.2,
                        fillOpacity: 0.2
                    };
                    return L.circle(latlng, geojsonMarkerOptions)
                },
                onEachFeature: onEachFeature
            })

            layer.addTo(map);
        }
    });
}

function barChartStudent(area, customGPS, datestring) {
  var dates = datestring.split("&")

  var start = moment(dates[0].split("=")[1])
  var end = moment(dates[1].split("=")[1])

  var duration = moment.duration(end.diff(start))
  var timespan = duration.asHours()
  
  var url = getStudentUrl(area, customGPS, datestring, true)
  Plotly.d3.csv(url, function(err, rows){
    if ( rows.length == 0 ) {
      $("#infobox-student").show()
      return; 
    }

    function unpack(rows, key) {
      return rows.map(function(row){ return row[key]; });
    }
    var pm10 = {
      type: "scatter",
      mode: "lines+markers",
      name: 'PM10',
      x: unpack(rows, "timestamp"),
      y: unpack(rows, "pmTen"),
      line: {color: '#17BECF'}
    }


    var pm25 = {
      type: "scatter",
      mode: "lines+markers",
      name: 'PM2.5',
      x: unpack(rows, "timestamp"),
      y: unpack(rows, "pmTwoFive"),
      line: {color: '#7F7F7F'}
    }

    var dataDust = [pm10,pm25];
    var layoutDust = {
      title: 'Støvkonsentrasjon',
      height: 470,
      width: 570,
      yaxis: {title: '\u03BC'+"g/m3"},      
    };

    Plotly.newPlot('chart-dust', dataDust, layoutDust);

    var temperature = {
      type: "scatter",
      mode: "lines+markers",
      name: "Temperatur",
      x: unpack(rows, "timestamp"),
      y: unpack(rows, "temperature"),
      line: {color: '#17BECF'}
    }

    var dataTemperature = [temperature];
    var layoutTemperature = {
      title: 'Temperatur',
      height: 470,
      width: 570,
      yaxis: {title: "Celcius"},
    };

    Plotly.newPlot('chart-temperature', dataTemperature, layoutTemperature);

    var humidity = {
      type: "scatter",
      mode: "lines+markers",
      name: "Luftfuktighet",
      x: unpack(rows, "timestamp"),
      y: unpack(rows, "humidity"),
      line: {color: '#17BECF'}
    }

    var dataHumidity = [humidity];
    var layoutHumidity = {
      title: 'Luftfuktighet',
      height: 470,
      width: 570,
      yaxis: {title: "%"},      
    };

    Plotly.newPlot('chart-humidity', dataHumidity, layoutHumidity);
  })
}

function barChartNilu(area, component, datestring) {
  layoutColors = ['#17BECF', '#7F7F7F']

  var url = getHistoricalUrl(area, datestring, component)

  Plotly.d3.csv(url, function(err, rows){
    function unpack(rows, key) {
      return rows.map(function(row){ return row[key]; });
    }
    if ( rows.length == 0 ) {
      $("#infobox-"+component).show()
      return; 
    }

    var stations = rows.groupBy("station")
    var keys = Object.keys(stations)

    var data = []
    for ( i = 0; i < keys.length; i++ ) {
      var trace = {
        type: "scatter",
        mode: "lines+markers",
        name: keys[i],
        x: unpack(stations[keys[i]], "to"),
        y: unpack(stations[keys[i]], "value"),
        line: {color: layoutColors[i]}
      }

      data.push(trace);
    };
       
    var layout = {
      title: component,
      height: 500,
      width: 600,
      yaxis: {title: '\u03BC'+"g/m3"},     
      showlegend: true, 
    };

    Plotly.newPlot('chart-'+ component, data, layout);

  })
   
}


function getHistoricalUrl(area, datestring, component) {
    area = encodeURIComponent(area);
    return "/historical?area=" + area + "&" + datestring + "&component=" + component
}

function getStudentUrl(area, customGPS, datestring, plotChart) {
    area = encodeURIComponent(area);
    url = ""
    if (customGPS == true) {
        url = "/student?within=" + area + "&" + datestring
    } else {
        url = "/student?area=" + area + "&" + datestring
    }
    if (plotChart == true) {
      url += "&plotchart=true"
    }

    return url
}

Array.prototype.groupByTime = function(prop, timespan) {
    return this.reduce(function(groups, item) {
      var val;
      if (timespan <= 1) {val = item[prop].slice[0,16]}
      if (timespan <= 24) { val = item[prop].slice(0,13)}
      else if (timespan <= 744) { val = item[prop].slice(0,10) }
      else { val = item[prop].slice(0,7)}

      groups[val] = groups[val] || []
      groups[val].push(item)
      return groups
    }, {})
}

Array.prototype.groupBy = function(prop) {
    return this.reduce(function(groups, item) {
      var val = item[prop]; 
      groups[val] = groups[val] || []
      groups[val].push(item)
      return groups
    }, {})
}

function calculateAvg(groups, key) {
  var newGroups = {}
  for (var index in groups) {
    newGroups[index] = newGroups[index] || []
    newGroups[index].push(getAvg(groups[index].map(function(row) {return row[key]})))
  }
  return newGroups
}

function getAvg(values) {
  return String(values.reduce(function (p, c) {
    return Number(p) + Number(c);
  }) / values.length);
}


function clearVis(element, map) {
    $(element).html("")
    if (map != undefined) {
        map.remove();
    }
}

$(document).ready(function () {
    // Apply bootstrap style classes to markdown tags
    var markdownTags = $("div.markdown");

    // Tables need to have the table class in bootstrap
    $("table", markdownTags)
        .addClass("table")
        .addClass("table-responsive")
        .addClass("table-sm")
        .addClass("table-striped")
        .addClass("table-bordered");

    // Blockquote tags need to have the blockquote class in bootstrap
    $("blockquote", markdownTags)
        .addClass("blockquote");

    $(".wiki-helplink")
        .addClass("pull-right")
        .addClass("btn")
        .addClass("btn-link");
});