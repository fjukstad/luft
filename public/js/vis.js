// Create new map that we can add data points to. 
function newMap(id, coordinates) {
    map = L.map(id, {
        editable: true
    })
    map.setView(coordinates, 9.6);

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
            url = "/" + provider + "aqis?within=" + areaUri + "&" + datestring
        } else {
            url = "/" + provider + "aqis?area=" + areaUri + "&" + datestring
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

function barChartStudent(area, customGPS, components, datestring) {
  var dates = datestring.split("&")

  var start = moment(dates[0].split("=")[1])
  var end = moment(dates[1].split("=")[1])

  var duration = moment.duration(end.diff(start))
  var timespan = duration.asHours()
  
  var url = getStudentUrl(area, customGPS, datestring)
  Plotly.d3.csv(url, function(err, rows){
    
    var groupedBy = rows.groupBy('timestamp', timespan)
    var averagePmTen = calculateAvg(groupedBy, "pmTen")
    var averagePmTwoFive = calculateAvg(groupedBy, "pmTwoFive")
    var averageTemperature = calculateAvg(groupedBy, "temperature")
    var averageHumidity = calculateAvg(groupedBy, "humidity")

    function unpack(rows) {
      return Object.keys(rows).map(function(key){ return rows[key][0]; });
    }
  
    var pm10 = {
      type: "scatter",
      mode: "lines+markers",
      name: 'PM10',
      x: Object.keys(averagePmTen),
      y: unpack(averagePmTen),
      line: {color: '#17BECF'}
    }


    var pm25 = {
      type: "scatter",
      mode: "lines+markers",
      name: 'PM2.5',
      x: Object.keys(averagePmTwoFive),
      y: unpack(averagePmTwoFive),
      line: {color: '#7F7F7F'}
    }

    var dataDust = [pm10,pm25];
    var layoutDust = {
      title: 'Støvkonsentrasjon',
      height: 500,
      width: 600,
    };

    Plotly.newPlot('chart-dust', dataDust, layoutDust);

    var temperature = {
      type: "scatter",
      mode: "lines+markers",
      name: "Temperatur",
      x: Object.keys(averageTemperature),
      y: unpack(averageTemperature),
      line: {color: '#17BECF'}
    }

    var dataTemperature = [temperature];
    var layoutTemperature = {
      title: 'Temperatur',
      height: 500,
      width: 600,
      yaxis: {title: "Celcius"},
    };

    Plotly.newPlot('chart-temperature', dataTemperature, layoutTemperature);

    var humidity = {
      type: "scatter",
      mode: "lines+markers",
      name: "Luftfuktighet",
      x: Object.keys(averageHumidity),
      y: unpack(averageHumidity),
      line: {color: '#17BECF'}
    }

    var dataHumidity = [humidity];
    var layoutHumidity = {
      title: 'Luftfuktighet',
      height: 500,
      width: 600,
      yaxis: {title: "%"},      
    };

    Plotly.newPlot('chart-humidity', dataHumidity, layoutHumidity);
  })
}

Array.prototype.groupBy = function(prop, timespan) {
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


function barChartNilu(area, component, datestring, container, element) {

    var parseTime = d3.utcParse("%Y-%m-%dT%H:%M:%S.%LZ");

    var svg = document.querySelector(element);
    svg.setAttribute("width", document.getElementById(container).clientWidth)

    var svg = d3.select(element),
        margin = {
            top: 20,
            right: 30,
            bottom: 20,
            left: 30
        },
        width = +svg.attr("width") - margin.left - margin.right,
        height = +svg.attr("height") - margin.top - margin.bottom,
        g = svg.append("g").attr("transform", "translate(" + margin.left + "," + margin.top + ")");

    var x = d3.scaleTime()
        .rangeRound([0, width]);

    var y = d3.scaleLinear()
        .rangeRound([height, 0]);

    var z = d3.scaleOrdinal(d3.schemeCategory20);

    var line = d3.line()
        .curve(d3.curveBasis)
        .x(function (d) {
            return x(d.from);
        })
        .y(function (d) {
            return y(d.value);
        });

    var unit = "";

    var stations = {};
    var url;
    url = getHistoricalUrl(area, datestring, component);

    d3.csv(url, function (d) {
            if (!stations[d.station]) {
                stations[d.station] = []
            }
            d.from = parseTime(d.from);
            d.to = parseTime(d.to);
            d.value = parseFloat(d.value)
            component = d.component
            unit = d.unit
            stations[d.station].push(d)
            return d;
        },
        function (error, data) {
            if (data == 0) {
                drawNoData(component);
            }
            x.domain(d3.extent(data, function (d) {
                return d.from;
            }));
            y.domain(d3.extent(data, function (d) {
                return d.value;
            }));

            g.append("g")
                .attr("transform", "translate(0," + height + ")")
                .call(d3.axisBottom(x))
                .select(".domain")
                .remove();

            g.append("g")
                .call(d3.axisLeft(y))
                .append("text")
                .attr("fill", "#000")
                .attr("transform", "rotate(-90)")
                .attr("y", 6)
                .attr("dy", "0.71em")
                .attr("text-anchor", "end")
                .text(component + "(" + unit + ")");

            label_offset = width / 2
            var component_selector = component.replace("/", "")

            g.append("g")
                .append("text")
                .attr("id", component_selector + "-label")
                .attr("transform", "translate(" + label_offset + ",0)")
                .attr("fill", "#000")
                .text("")


            for (var station in stations) {
                var id = station.replace("\ ", "")
                id = id.replace(",", "")
                id = id.replace(".", "")


                path = g.append("path")
                    .datum(stations[station])
                    .attr("fill", "none")
                    .style("stroke", z(station))
                    .attr("stroke", "steelblue")
                    .attr("stroke-linejoin", "round")
                    .attr("stroke-linecap", "round")
                    .attr("stroke-width", 1.5)
                    .attr("d", line)
                    .attr("id", id + "-" + component_selector)

                d3.select("path#" + id + "-" + component_selector).on("mouseover", function () {
                        d3.select(this).style("stroke-width", 5);
                        label = d3.select(this).data()[0][0].station
                        d3.select("text#" + component_selector + "-label").text(label)
                    })

                    .on("mouseout", function () {
                        d3.select(this).style("stroke-width", 1.5);
                        d3.select("text#" + component_selector + "-label").text("")
                    })
            }
        })
}


function getHistoricalUrl(area, datestring, component) {
    area = encodeURIComponent(area);
    return "/historical?area=" + area + "&" + datestring + "&component=" + component
}

function getStudentUrl(area, customGPS, datestring, component) {
    area = encodeURIComponent(area);
    if (customGPS == true) {
        return "/student?within=" + area + "&" + datestring
    } else {
        return "/student?area=" + area + "&" + datestring
    }
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