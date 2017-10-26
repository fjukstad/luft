# Example code to test the SDS011 PM2.5 PM10 sensor

- Download the [SDS011 library](http://github.com/skolelab/sds011) first. 

```
#include <SoftwareSerial.h>
#include <SDS011.h>
#define PM_TX 2
#define PM_RX 3

SDS011 sds; 

void setup() {
  sds.begin(PM_TX,PM_RX);
  Serial.begin(9600);
}

void loop() {
  float pm25, pm10; 
  int error = sds.read(&pm25,&pm10);
    if (! error) {
    Serial.print(pm25);
    Serial.print("\t");
    Serial.println(pm10);
  } else {
    Serial.println("Could not read air data: "+String(error));
  }
  delay(1000);
}
```
