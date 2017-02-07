# Luftkvalitet
Go package to get air quality data from [api.nilu.no](https://api.nilu.no/docs/). 
It supports retrieval of: up to date data (latest measurements); historical data; 
available areas, components, stations and aqis; and air quality forecasts. 


# Example
For more examples have a look at [gonum-plot-example](https://github.com/fjukstad/gonum-plot-example),
[polluteman](https://github.com/fjukstad/polluteman) or in [example/](example).

Get current air quality measurements from Tromsø: 

```go
package main

import (
	"fmt"

	"github.com/fjukstad/luftkvalitet"
)

func main() {

	areas := []string{"Tromsø"}

	m, err := luftkvalitet.GetMeasurements(luftkvalitet.Filter{Areas: areas})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(m)
}
```

produces (on 4.12.2016 19:44 EST): 

```
[{{{Troms og Finnmark Tromsø Tromsø} {69.67957 18.95402} Tverrforbindelsen} NO0085A PM10 2016-12-05 00:00:00 +0100 +0100 2016-12-05 01:00:00 +0100 +0100 10.34 µg/m³ 1 6ee86e} {{{Troms og Finnmark Tromsø Tromsø} {69.65625 18.96372} Hansjordnesbukta} NO0079A PM10 2016-12-05 00:00:00 +0100 +0100 2016-12-05 01:00:00 +0100 +0100 11.44 µg/m³ 1 6ee86e} {{{Troms og Finnmark Tromsø Tromsø} {69.65625 18.96372} Hansjordnesbukta} NO0079A PM2.5 2016-12-05 00:00:00 +0100 +0100 2016-12-05 01:00:00 +0100 +0100 5.1 µg/m³ 1 6ee86e} {{{Troms og Finnmark Tromsø Tromsø} {69.65625 18.96372} Hansjordnesbukta} NO0079A NO2 2016-12-05 00:00:00 +0100 +0100 2016-12-05 01:00:00 +0100 +0100 -0.0864654577 µg/m³ 1 6ee86e}]
```

you get the idea? 


# Acknowledgements
The data belongs to The Norwegian Institute for Air Research (NILU), see
[luftkvalitet.info](http://www.luftkvalitet.info) and
[nilu.no](http://www.nilu.no) for more information.  
