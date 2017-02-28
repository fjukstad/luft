// Create new map that we can add data points to. 
function newMap(id){
    map = L.map(id, {editable:true})
    map.setView([69.680, 18.951], 9.6);


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

function addToMap(map, area, provider, component, datestring) { 

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

    $.ajax({
        dataType: "json",
        url: "/"+ provider+"aqis?area="+area+"&"+datestring+"&component="+component,
        success: function(data) {
            var layer = L.geoJSON(data.features, {
                pointToLayer: function(feature, latlng){
                    var geojsonMarkerOptions = {
                        color: "#"+feature.properties.color,
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

function barChart(area, component, datestring, container, element) { 
    
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

    var z = d3.scaleOrdinal(d3.schemeCategory10);

    var line = d3.line()
        .curve(d3.curveBasis)
        .x(function(d) { return x(d.from); })
        .y(function(d) { return y(d.value); });

    var unit = "" ; 
        
    var stations = {}; 
    
    var url; 
    if(component == "NO2" || component == "PM10"){ 
        url = getHistoricalUrl(area, datestring, component);
    }  else if(component == "dust" || component == "humidity" || component == "temperature"){
        url = getStudentUrl(datestring, component) 
    }


    d3.csv(url,
            function(d) { 
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


    //$.ajax({
    //    dataType: "json",
    //    url: "/logs?"+datestring,
    //    success: function(data) {
    //        L.geoJSON(data.features, {
    //            pointToLayer: function(feature, latlng){
    //                var geojsonMarkerOptions = {
    //                    weight: 1,
    //                    opacity: 1,
    //                    fillOpacity: 0.8
    //                };

    //                return L.circle(latlng, geojsonMarkerOptions)
    //                },
    //                onEachFeature: onEachFeature
    //            }) .addTo(map);
    //        }
    //    });

}

function getHistoricalUrl(area, datestring, component) {
    return "/historical?area="+area+"&"+datestring+"&component="+component
}

function getStudentUrl(datestring, component) {
    return "/student?"+datestring+"&component="+component
} 


function clearVis(element, map) { 
    $(element).html("") 
    if(map != undefined) { 
        map.remove(); 
    }
}

