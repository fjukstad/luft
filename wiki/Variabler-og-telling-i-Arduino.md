I forrige steg skrev vi ut teksten `Hello World!`. Vi så at Arduinoen kjørte gjennom instruksjonene i `loop` funksjonen om og om igjen. La oss nå endre Hello-World-programmet til å også skrive ut en teller verdi, som vi øker for hver linje.

For dette må vi lage en variabel i programmet vårt. En variabel er en verdi med et navn som kan endre verdi.

### Deklarere en ny variabel helplink

``` cpp
int counter;
```

Koden over skal legges helt øverst i Sketchen, dvs. **før** begynnelsen av `setup`-funksjonen. Legg gjerne til en tom linje mellom også, for å gjøre det enklere å se de seperate seksjonene.

Koden lager en ny variabel som heter `counter`. Linjen sier bare at `counter` eksisterer, dette kalles en *deklarasjon*. Vi kommer til å bruke verdien resten av programmet.

Foran navnet finner vi kodeordet `int`. Kort for `integer`, som er det engelske ordet for heltall. Dette beskriver hva slags verdier `counter` kan ha. Siden koden sier `int` (i stedet for noen annen type) kan `counter` bare lagre heltallige verdier, både negative og positive.

### Initialisering helplink

Med deklerasjonen av `counter` variablen må vi nå sette verdien til variablen til en nyttig verdi. Dette må vi gjøre én gang når Arduinoen starter opp. Derfor legger vi følgende til `setup`-funksjonen: (*dvs. før den lukkende krøllparantesen*)

``` cpp
  counter = 0;
```

Et enkel likhetstegn er i Arduinokode operatoren for å lagre en verdi i en variabel. Verdien på høyresiden av likhetstegnet plasseres i verdien på venstresiden. Kommandoen over oversettes altså til: plasser verdien `0` i variablen `counter`. Etter denne instruksjonen vil verdien til `counter` altså være `0`, helt til vi endrer den igjen med en annen instruksjon.

Kode som denne, som setter start-verdier i variabler, kaller vi *initialisering*.

### Printe ut variabler helplink

I forrige steg brukte vi `Serial.println` for å skrive ut `Hello World!`. La oss først endre teksten til `Hello World! Counter:`. Vi vil også forsette å skrive tekst i en ny kommando på samme linje. Derfor vil vi heller bruke `print` i stedet for `println`.

``` cpp
  Serial.print("Hello World! Counter: ");
```

*Merk at vi har lagt til et tomrom på slutten av teksten i koden over.*

La oss nå skrive ut verdien til `counter`. Vi legger bare til en kommando til `Serial.print` en gang til, bare med verdien `counter` i stedet for tekst.

``` cpp
  Serial.print(counter);
```

Nå er vi ferdig med å skrive linjen, så nå er rette tiden for å skrive videre på en ny linje. Dette kan vi gjøre ved å bruke `println` kommandoen uten noen argument.

``` cpp
  Serial.println();
```

*Merk at det ikke står noe mellom parantesene, men parantesene må fortsatt være der.*

### Øke verdien til variablen helplink

Etter vi ha skrevet ut verdien til `counter` som vi gjorde over, burde vi nå øke verdien `counter` variablen slik at vi får en ny verdi for hver gang koden kjøres. Å øke en verdi kan gjøres ved å legge til `1` til `counter`. Siden den økte verdien må lagres i `counter` igjen, trenger vi forsatt et likhetstegn.  
Operatoren for å øke er `+=`. Verdien på venstresiden vil bil økt med verdien på høyresiden.

Vi legger til følgende kode etter kommandoene der vi skriver ut tekst.

``` cpp
  counter += 1; // Increase counter by 1.
```

### Ferdig helplink

``` cpp
int counter; // Counting-variable, int -> integer value, whole number

void setup() {
  // initialize serial communication at 9600 bauds per second:
  Serial.begin(9600);

  counter = 0;
}

void loop() {
  // print out Hello World and the counter:
  Serial.print("Hello World! Counter: ");
  Serial.print(counter);
  Serial.println();

  counter += 1; // Increase counter by 1.

  delay(1000); // delay for one second
}
```

Etter den endrete koden har blitt sendt opp til Arduinoen, vil du se følgende i `Serial Monitor`:

``` txt
Hello World! Counter: 1
Hello World! Counter: 2
Hello World! Counter: 3
Hello World! Counter: 4
...
```

## Gå videre

&uarr; [Gå til **Introduksjon til Arduino**][intro-prog-home]  
&larr; [Gå tilbake til forrige steg: **Hello World**][hello-world]  
&darr; [Gå til: **Guider**][guides]  

[intro-prog-home]: Introduksjon-til-Arduino-programmering
[hello-world]: Arduino-varianten-av-Hello-World
[guides]: airbit-Guider
