# Example code to print all measurements

The only thing left to do is to write them to the memory card! 

```
#include <TinyGPS++.h>
#include <SoftwareSerial.h>
#include <stdlib.h>
#include <SDS011.h>
#include <SD.h>
#include <SimpleDHT.h>

#define SD_PIN 10
#define DHTPIN 9

SimpleDHT22 dht22;
File file;

#define LED_WHITE PIN_A1
#define LED_GREEN PIN_A0
#define PM_TX 2
#define PM_RX 3

static const int GPS_RX = 7;
static const int GPS_TX = 6;
static const uint32_t GPSBaud = 9600;
TinyGPSPlus gps;
SoftwareSerial ss(GPS_RX, GPS_TX);

SDS011 sds;

void setup() {
  pinMode(LED_WHITE, OUTPUT);
  pinMode(LED_GREEN, OUTPUT);
  sds.begin(PM_TX, PM_RX);
  ss.begin(GPSBaud);
  Serial.begin(9600);

  pinMode(SD_PIN, OUTPUT);
  SD.begin();
  char filename[] = "fil.txt";

  // If file exsist on the memory card, append to it, if it doesn't exist
  // create it and write one line to the file before continuing.

  if (SD.exists(filename)) {
    file.close();
    file = SD.open(filename, O_WRITE | O_APPEND);
  } else {
    file = SD.open(filename, O_WRITE | O_CREAT);
    file.print("Time, Latitude, Longitude, PM25, PM10, Temperature, Humidity\n");
    file.flush();
  }
}

void loop() {
  while (true) {
    ss.listen();
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
          Serial.println(",");

          // put your main code here, to run repeatedly:
          float temperature = 0;
          float humidity = 0;
          while (temperature == 0 && humidity == 0) {
            dht22.read2(DHTPIN, &temperature, &humidity, NULL);
            delay(250);
          }
          Serial.print("Temperature: ");
          Serial.print(temperature);
          Serial.println("C");
          Serial.print("Humidity: ");
          Serial.print(humidity);
          Serial.println("%");


          delay(500);
          digitalWrite(LED_GREEN, LOW);
          break;
        }
        else {
          digitalWrite(LED_WHITE, HIGH);
          delay(500);
          digitalWrite(LED_WHITE, LOW);
          delay(500);

        }
      }
    }
  }

  float pm25, pm10;
  while (true) {
    int error = sds.read(&pm25, &pm10);
    if (error == 0) {
      Serial.print("PM 2.5 ");
      Serial.println(pm25);
      Serial.print("PM 10 ");
      Serial.println(pm10);
      break;
    } else {
      delay(1000);
    }
  }

}
```
