# Example code to test the DHT temperature and humidity sensor 

```
#include <SimpleDHT.h>

#define DHTPIN 9
SimpleDHT22 dht22;

void setup() {
  Serial.begin(9600); 
}

void loop() {
  float temperature = 0;
  float humidity = 0;
  dht22.read2(DHTPIN, &temperature, &humidity, NULL);
  Serial.print(temperature); 
  Serial.println("C");
  Serial.print(humidity);
  Serial.println("%");

  delay(2500);
}
```
