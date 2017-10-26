# Example code to test the GPS

- First download the [TinyGPS++ library](http://arduiniana.org/libraries/tinygpsplus/)

``` cpp
#include <TinyGPS++.h>
#include <SoftwareSerial.h>
#include <stdlib.h>

#define LED_WHITE PIN_A1
#define LED_GREEN PIN_A0

static const int GPS_RX = 7;
static const int GPS_TX = 6;
static const uint32_t GPSBaud = 9600;
TinyGPSPlus gps;
SoftwareSerial ss(GPS_RX, GPS_TX);

void setup() {
  pinMode(LED_WHITE, OUTPUT);
  pinMode(LED_GREEN, OUTPUT);
  ss.begin(GPSBaud);
  Serial.begin(9600);
}

void loop() {
  if (ss.available()) {
    if (gps.encode(ss.read())) {
      if (gps.location.isValid()) {
        digitalWrite(LED_GREEN, HIGH);

        // Date
        Serial.print("Date: ");
        Serial.print(gps.date.day());
        Serial.print("/");
        Serial.print(gps.date.month());
        Serial.print("/");
        Serial.print(gps.date.year());
        Serial.print(" ");
        Serial.print(gps.time.hour());
        Serial.print(":");
        Serial.print(gps.time.minute());
        Serial.print(":");
        Serial.print(gps.time.second());
        Serial.print(",");

        // Latitude.
        // Note that we convert the float output to a string. This is
        // for convenience when we're later going to write it to a file
        // on a SD memory card.

        char lat[15];
        dtostrf(gps.location.lat(), 3, 6, lat);

        Serial.print(" Latitude: ");
        Serial.print(lat);
        Serial.print(",");

        // long
        char lng[15];
        dtostrf(gps.location.lng(), 3, 6, lng);
        Serial.print(" Longitude: ");
        Serial.print(lng);
        Serial.print(",");

        delay(500);
        digitalWrite(LED_GREEN, LOW);
      }
      else {
        digitalWrite(LED_WHITE, HIGH);
        Serial.print("Bad GPS signal.");
        delay(500);
        digitalWrite(LED_WHITE, LOW);
      }
      Serial.print("\n");
    }
  }
}
```
