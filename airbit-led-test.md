# Example code to test the two LEDs on the air:bit

```
#define LED_WHITE A1
#define LED_GREEN A0

void setup() {
  // put your setup code here, to run once:
  pinMode(LED_WHITE, OUTPUT); 
  pinMode(LED_GREEN, OUTPUT); 

}

void loop() {
  // put your main code here, to run repeatedly:
  digitalWrite(LED_WHITE, HIGH);
  digitalWrite(LED_GREEN, LOW);

  delay(1000);
  digitalWrite(LED_WHITE, LOW);
  digitalWrite(LED_GREEN, HIGH);

  delay(1000);
}
```
