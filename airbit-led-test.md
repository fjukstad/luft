# Example code to test the two LEDs on the air:bit

``` cpp
#define LED_WHITE A1
#define LED_GREEN A0

void setup() {
  pinMode(LED_WHITE, OUTPUT);
  pinMode(LED_GREEN, OUTPUT);
}

void loop() {
  digitalWrite(LED_WHITE, HIGH);
  digitalWrite(LED_GREEN, LOW);
  delay(1000);

  digitalWrite(LED_WHITE, LOW);
  digitalWrite(LED_GREEN, HIGH);
  delay(1000);
}
```
