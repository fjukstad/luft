// Create new map that we can add data points to. 
function newMap(id, coordinates){
    map = L.map(id, {editable:true})
    map.setView(coordinates, 9.6);

    var accessToken = 'pk.eyJ1IjoiZmp1a3N0YWQiLCJhIjoiY2l2Mnh3azRvMDBrYTJ5bnYxcDAzZ3Z0biJ9.RHb5ENfbmzN65gjiB-L_wg';

    L.tileLayer(
        "https://api.mapbox.com/styles/v1/mapbox/dark-v9/tiles/256/{z}/{x}/{y}?access_token="+accessToken,
        {
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
        if (feature.properties && feature.properties.name) {
            content = "<b>"+feature.properties.name+"</b></br>"
            if(feature.properties.component) {
                content += feature.properties.component+": "+feature.properties.value + "</br>"
            } 
            if(feature.properties.dust){ 
                content += "Dust: " + feature.properties.dust + "</br>" 
                content += "Temperature: "+feature.properties.temperature + "</br>"
                content += "Humidity: "+feature.properties.humidity + "</br>"
            }
            content += feature.properties.date
            layer.bindPopup(content);
        }
    }

    var colorScheme = d3.scaleOrdinal(d3.schemeCategory20);

    $.ajax({
        dataType: "json",
        url: function(customGPS) {
            if(customGPS == true) {
                return "/"+ provider+"aqis?within="+area+"&"+datestring+"&component="+component
            }
            else {
                return "/"+ provider+"aqis?area="+area+"&"+datestring+"&component="+component
            }
        },
        success: function(data) {
            var layer = L.geoJSON(data.features, {
                pointToLayer: function(feature, latlng){

                    var color = "" ; 
                    if(!feature.properties.color){
                        color = colorScheme(feature.properties.name)
                    } else {
                        color = "#" + feature.properties.color
                    }

                    var geojsonMarkerOptions = {
                        color: color,
                        weight: feature.properties.weight,
                        opacity: 1,
                        fillOpacity: 0.8
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
    
    var parseTime = d3.utcParse("%Y-%m-%dT%H:%M:%S.%LZ");
    var charts = []

    for(var i = 0; i < components.length; i++) {
        var container = "chart-"+ components[i]
        var element = "svg#chart-" + components[i]
        var svg = document.querySelector(element); 
        svg.setAttribute("width", document.getElementById(container).clientWidth) 
        var svg = d3.select(element),
        margin = {top: 20, right: 30, bottom: 20, left: 30},
        width = +svg.attr("width") - margin.left - margin.right,
        height = +svg.attr("height") - margin.top - margin.bottom,
        g = svg.append("g").attr("transform", "translate(" + margin.left + "," + margin.top + ")");
        charts.push(g);
    }

    var x = d3.scaleTime()
        .rangeRound([0, width]);

    var y = d3.scaleLinear()
        .rangeRound([height, 0]);

    var z = d3.scaleOrdinal(d3.schemeCategory20);

    var url = getStudentUrl(area, customGPS, datestring) 

    var components = ["dust", "humidity", "temperature"];
    var units = [];
    var stations = {};
    
    d3.csv(url,
        function(d) { 
            if(!stations[d.station]){
                stations[d.station] = []
            }
            console.log(stations.length)
            d.timestamp = parseTime(d.timestamp);
            d.values = [[parseFloat(d.pmTwoFive), parseFloat(d.pmTen)], parseFloat(d.humidity), parseFloat(d.temperature)]; 
            units = [d.unitDust, d.unitHum, d.unitTemp];
            stations[d.station].push(d) 
            return d; 
        },
        function(error, data){
            var line;
            if(data == 0) {
                drawNoData(components);
            }
            for(var v = 0; v < units.length; v++) {
                var component = components[v];
                x.domain(d3.extent(data, function(d) { return d.timestamp; }));
                
                if(v == 0) {
                    var min = d3.min([d3.min(data, function(d) { return d.values[v][0]}), d3.min(data, function(d) { return d.values[v][1]})])
                    var max = d3.max([d3.max(data, function(d) { return d.values[v][0]}), d3.max(data, function(d) { return d.values[v][1]})])
                    y.domain([min, max]);
                    
                } else {
                    y.domain(d3.extent(data, function(d) { return d.values[v]; }));
                    line = d3.line()
                        .curve(d3.curveBasis)
                        .x(function(d) { return x(d.timestamp); })
                        .y(function(d) { return y(d.values[v]); });
                } 
                
                charts[v].append("g")
                .attr("transform", "translate(0," + height + ")")
                .call(d3.axisBottom(x))
                .select(".domain")
                .remove();

                charts[v].append("g")
                  .call(d3.axisLeft(y))
                  .append("text")
                  .attr("fill", "#000")
                  .attr("transform", "rotate(-90)")
                  .attr("y", 6)
                  .attr("dy", "0.71em")
                  .attr("text-anchor", "end")
                  .text(component + "("+ units[v]+")"); 

                label_offset = width/2
                var component_selector = component.replace("/","") 
                
                charts[v].append("g")
                    .append("text")
                    .attr("id",component_selector+"-label")
                    .attr("transform", "translate("+label_offset+",0)")
                    .attr("fill", "#000")
                    .text("")
                    
                for(var station in stations){ 
                    var id = station.replace("\ ", "")
                        id = id.replace(",","")
                        id = id.replace(".","")

                    if(v == 0) {
                        for(var s = 0; s < 2; s++ ) {
                            line = d3.line()
                                .curve(d3.curveBasis)
                                .x(function(d) { return x(d.timestamp); })
                                .y(function(d) { return y(d.values[v][s]); });
                            path = charts[v].append("path")
                              .datum(stations[station])
                              .attr("fill", "none")
                              .style("stroke", z(station))
                              .attr("stroke", "steelblue")
                              .attr("stroke-linejoin", "round")
                              .attr("stroke-linecap", "round")
                              .attr("stroke-width", 1.5)
                              .attr("d", line)
                              .attr("id", id+"-"+component_selector)
                        }
                    }    
                    else {
                        path = charts[v].append("path")
                          .datum(stations[station])
                          .attr("fill", "none")
                          .style("stroke", z(station))
                          .attr("stroke", "steelblue")
                          .attr("stroke-linejoin", "round")
                          .attr("stroke-linecap", "round")
                          .attr("stroke-width", 1.5)
                          .attr("d", line)
                          .attr("id", id+"-"+component_selector)
                    }
                    d3.select("path#"+id+"-"+component_selector).on("mouseover", function(){
                        d3.select(this).style("stroke-width", 5); 
                        label = d3.select(this).data()[0][0].station
                        d3.select("text#"+component_selector+"-label").text(label)
                     })

                    .on("mouseout", function(){
                        d3.select(this).style("stroke-width", 1.5); 
                        d3.select("text#"+component_selector+"-label").text("")
                    })
                }
            }

        }
    )
}

function drawNoData(components) {
    for(var i = 0; i < components.length; i++) {
        $('svg#chart-' + components[i]).empty();
        var container = "chart-"+ components[i]
        console.log(container)
        var element = "svg#chart-" + components[i]
        var svg = document.querySelector(element); 
        svg.setAttribute("width", document.getElementById(container).clientWidth) 
        var svg = d3.select(element),
        margin = {top: 20, right: 30, bottom: 20, left: 0},
        width = +svg.attr("width") - margin.left - margin.right,
        height = +svg.attr("height") - margin.top - margin.bottom,
        g = svg.append("g").attr("transform", "translate(" + margin.left + "," + margin.top + ")");
        g.append("text")
            .text("Ingen data tilgjengelig for denne tidsperdioden")
    }
}
function barChartNilu(area, component, datestring, container, element) { 
    
    var parseTime = d3.utcParse("%Y-%m-%dT%H:%M:%S.%LZ");

    var svg = document.querySelector(element); 
    svg.setAttribute("width", document.getElementById(container).clientWidth) 

    var svg = d3.select(element),
    margin = {top: 20, right: 30, bottom: 20, left: 30},
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
        .x(function(d) { return x(d.from); })
        .y(function(d) { return y(d.value); });

    var unit = "" ; 
        
    var stations = {}; 
    var url; 
    url = getHistoricalUrl(area, datestring, component);

    d3.csv(url,function(d) { 
        if(!stations[d.station]){
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
    function(error, data){
        console.log(data);
        if(data == 0) {
            drawNoData(component);
        }
        x.domain(d3.extent(data, function(d) { return d.from; }));
        y.domain(d3.extent(data, function(d) { return d.value; }));

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
          .text(component + "("+unit+")"); 

        label_offset = width/2
        var component_selector = component.replace("/","") 
        
        g.append("g")
            .append("text")
            .attr("id",component_selector+"-label")
            .attr("transform", "translate("+label_offset+",0)")
            .attr("fill", "#000")
            .text("")
            

        for(var station in stations){ 
            var id = station.replace("\ ", "")
                id = id.replace(",","")
                id = id.replace(".","")


            path = g.append("path")
              .datum(stations[station])
              .attr("fill", "none")
              .style("stroke", z(station))
              .attr("stroke", "steelblue")
              .attr("stroke-linejoin", "round")
              .attr("stroke-linecap", "round")
              .attr("stroke-width", 1.5)
              .attr("d", line)
              .attr("id", id+"-"+component_selector)
              
            d3.select("path#"+id+"-"+component_selector).on("mouseover", function(){
                d3.select(this).style("stroke-width", 5); 
                label = d3.select(this).data()[0][0].station
                d3.select("text#"+component_selector+"-label").text(label)
             })

            .on("mouseout", function(){
                d3.select(this).style("stroke-width", 1.5); 
                d3.select("text#"+component_selector+"-label").text("")
            })
        }
    })
}


function getHistoricalUrl(area, datestring, component) {
    return "/historical?area="+area+"&"+datestring+"&component="+component
}

function getStudentUrl(area, customGPS, datestring, component) {
    if(customGPS == true) {
        return "/student?within="+area+"&"+datestring
    }
    else {
        return "/student?area="+area+"&"+datestring
    }
} 


function clearVis(element, map) { 
    $(element).html("") 
    if(map != undefined) { 
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
        .addClass("table-bordered")
        ;

    // Blockquote tags need to have the blockquote class in bootstrap
    $("blockquote", markdownTags)
        .addClass("blockquote")
        ;
});

