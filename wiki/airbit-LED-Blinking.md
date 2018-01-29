De enkleste komponentene på air:bit er LED lyspærene. Så for å starte med skal vi prøve å få dem til å blinke. Du kan senere bruke LED til å vise status på f.eks. GPS signalet eller blinke mens du tar målinger, osv. Det kan være en enkel måte å skjekke om alt er ok mens du går rundt og tar målinger. Husk at air:bit ikke har noen skjerm, og du vil neppe alltids ha med datamaskinen din for å kunne snakke direkte med Arduinoen.

## Komme i gang helplink

Vi starter med en helt ny Sketch. Først er det greit å bare begynne med å legge til de *tomme* funksjonene for `setup` og `loop` som vi allerede har sett på tidligere.

``` cpp
void setup() {
  
}

void loop() {
  
}
```

Så må vi ta en titt på pinout skjemaet og finne hvilke pinner vi skal bruke for LED pærene. [Klikk her for å se **Pinout-skjemaet**][pinout]

Fra skjemaet kan vi se at vi skal bruke pinne nummer `A1` for den blanke, rødt-lysende LED, og `A0` for den grønne LED. La oss definere to konstanter for dette i koden vår. Vi legger disse inn før `setup`-koden.

``` cpp
#define LED_RED   A1
#define LED_GREEN A0
```

Denne syntaksen som du ser i koden over har vi ikke brukt før. I motsetningen til å lage en variabel, deklarerer vi her to konstanter. Merk at vi gjerne bruker bare store bokstaver for konstanter. *Det kan også være fristende å bruke navn med norske tegn (f.eks. `LED_RØD`) her, men det er desverre ikke tillatt. Navn kan kun være små og store engelske bokstaver og siffrer.*

## `setup` helplink

I setup koden vår må vi nå fortelle Arduinoen at vi har lyst å bruke pinnene for LED-pærene. Kommandoen for dette er `pinMode`. Arduinoen har muligheten for å enten sende ut signaler, dvs. strøm, på en pinne, eller motta. Ved å bruke argumentet `OUTPUT` sier vi at vi ikke forventer å motta noe data på pinnen og bare skal sende strøm ut.

``` cpp
  pinMode(LED_RED, OUTPUT);
  pinMode(LED_GREEN, OUTPUT);
```

Merk at dette ikke faktisk skrur på eller av strømmen til LED enda. Det bare setter opp Arduinoen slik at den vil kunne skru av og på strømmen med en senere kommando i `loop`. Vi ser også at vi har to instruksjoner her, én for hver pinne. Om du bare skal bruke en LED ved en senere anledning kunne du bare fjernet en av de to linjene. Vi kommer til å bruke begge LED i dette eksemplet, så vi må derfor sette opp begge pinnene.

## `loop` helplink

Idéen er at vi skal blinke avvekslende med den røde og grønne LED på air:bit. Før vi skriver noen kode må vi først tenke over hvordan dette faktisk skjer, steg for steg.

1. Skru på rød LED
1. Vent i ett sekund
1. Skru av rød LED
1. Skru på grønn LED
1. Vent i ett sekund
1. Skru av grønn LED

Fra [*Hello-World*-eksemplet][hello-world] i guiden for generell Arduino-programmering, husker vi at kommandoen for å vente en viss tid er `delay` som tar tiden den skal vente i millisekunder som argument.

```cpp
  delay(1000); // Wait for one second.
```

Digitale signaler er signaler som kun har to mulige verdier: Av og På.  
Vi kan derfor tenke oss at å skru av og på strømmen til en LED er en form for et digitalt signal. La oss derfor bruke kommandoen `digitalWrite` for å skru av og på LEDene. Denne kommandoen tar pinnen den skal styre og verdien (Av eller På) som argument. Her bruker vi verdien `HIGH` for å skru på, og `LOW` for å skru av.

```cpp
  digitalWrite(LED_GREEN, HIGH); // Turn on green LED.
  digitalWrite(LED_RED, LOW); // Turn off red LED.
```

Prøv å oversett stegene vi skrev på norsk i opplistingen over, og skriv ned kommandoene i rett rekkefølge. Du kan gjerne leke litt her og se hva som skjer om du endrer rekkefølge eller vente-tid. Minker du f.eks. tallet i `delay`-kommandoene vil LEDene blinke fortere.

## Ferdig

Mye rart man kan gjøre med blinking og mange rare rytmer man kan få til her om man bare er litt kreativ, men hvis vi tar utgangspunkt i den helt kjedelige blinkingen i ett sekund mellomrom, vi du få kode som ligner på dette:

```cpp
#define LED_RED   A1
#define LED_GREEN A0

void setup() {
  pinMode(LED_RED, OUTPUT); // Enable red LED control
  pinMode(LED_GREEN, OUTPUT); // Enable green LED control
}

void loop() {
  digitalWrite(LED_RED, HIGH);
  delay(1000);
  digitalWrite(LED_RED, LOW);
  digitalWrite(LED_GREEN, HIGH);
  delay(1000);
  digitalWrite(LED_GREEN, LOW);
}
```

## Gå videre

&uarr; [Gå til **innholdsfortegnelsen**][home]  
&larr; [Gå tilbake forrige neste steg: **Pinout**][pinout]  
&rarr; [Gå til neste steg: **Temperatursensoren**][dht]  

[home]: airbit-Programmering
[pinout]: airbit-Pinout
[dht]: Programmering-med-Temperatursensoren

[hello-world]: Arduino-varianten-av-Hello-World
