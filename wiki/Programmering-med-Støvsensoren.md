Støvsensoren til air:bit er litt mere komplisert enn temperatursensoren, men igjen finnes det et Arduino bibliotekt for å lese av målinger fra sensoren.

Som i temperatursensor-eksemplet kommer vi til å demonstrere hvordan vi tar målinger fra støvsensoren, og skrive dem ut til seriell-forbindelsen til PC'en.

## Ny Sketch

Igjen starter vi med en helt ny Sketch. `File`&rarr;`New` i menyen.

``` cpp
void setup () {

}

void loop() {

}
```

## Laste ned bibliotek helplink

Akkurat som DHT-biblioteket vi brukte for temperatursensoren, finnes det et bibliotek for støvsensoren. Desverre oppstår det komplikasjoner med de andre sensorene i air:bit når vi bruker denne. Derfor har air:bit teamet laget sin egen versjon for dette biblioteket. Last ned den nyeste, dvs. den øverste, versjonen i listen du finner når du klikker på denne linken: (klikk der det står `zip`) **[SDS011 biblioteket](https://github.com/skolelab/SDS011/releases)**  
Det er viktig at du lagrer (**ikke åpner**) filen. Husk hvor du lagrer filen, vi må finne frem den filen i neste steg.

Etter du har lastet ned biblioteket og lagret filen må vi installere biblioteket i `Arduino IDE`. Klikk i menyen på `Sketch`&rarr;`Include library`&rarr;`Add ZIP library` (punktet under `Manage libraries`). Velg filen du nettopp lastet ned. Om nettleseren din ikke ba deg om å velge hvor filen skulle lagres, vil du mest sannsynligvis finne den under `Downloads` (`Nedlastinger`).

## Globale definisjoner helplink

Arduino IDE kommer med en del små hjelper-bibliotek som allerede er installert på forhånd. Ett av disse er `SoftwareSerial`-biblioteket. Støvsensoren sender faktisk målingene over ledningen til Arduinoen på lignende måte som Arduinoen 'snakker' med PCen over USB-ledningen. `SoftwareSerial` tillater å lage en seriell forbindelse over pinner på Arduinoen i stedet for USB-ledningen, så derfor bruker vi følgende `#include` direktiv:

``` cpp
#include <SoftwareSerial.h>
```

Først må vi inkludere `SDS011`-biblioteket. Bruk det følgende `#include` direktivet:

``` cpp
#include <SDS011.h>
```

Så er det tid å konsultere [pinout skjemaet][pinout]. Som vanlig definerer vi konstanter for pinnene som brukes for kommunikasjon med Arduinoen.

``` cpp
#define PM_TX 2
#define PM_RX 3
```

Så må vi lage en global variabel for tilkoblingen til Støvsensoren. Denne gang bruker vi `SDS011` datatypen. *Igjen kan variabler navngis som du synes best.*

``` cpp
SDS011 sds;
```

## `setup` helplink

Akkurat som kommunikasjonen med PC'en din, må vi også initialisere kommunikasjonen med støvsensoren. Vi bruker `begin`-kommandoen som hører til `sds` variabelen. Den tar imot TX-pinnen og så RX-pinnen som sensoren er koblet til med som argument.

``` cpp
  sds.begin(PM_TX, PM_RX);
```

Vi skal også skrive ut målingene over seriell-tilkoblingen til datamaskinen.

``` cpp
  Serial.begin(9600);
```

## `loop` helplink

Som med temperatursensoren, skal vi lese ut målingene fra støvsensoren og så printe dem ut. Støvsensoren tar målinger for to forskjellige partikelstørrelser: 2.5µm og 10µm. Begge tall gir konsentrasjoner som desimaltall. Enheten for verdiene er i `µg/m³` (mikrogramm per kubikkmeter)

``` cpp
  float pm25, pm10;
```

*Merk at du kan deklarere flere variabler av samme type ved å skrive komma mellom variablene.*

`read`-kommandoen til `sds`-variabelen brukes for å lese ut data fra sensoren. Kommandoen kan feile, så kommandoen returnerer en feilverdi som vi lagrer i en variabel. Kommadoen tar plassering først for 2.5µm og så for 10µm konsentrasjonen som argument.

``` cpp
  int error = sds.read(&pm25, &pm10);
```

I koden over ser du at feilverdier som regel lagres av datamaskinen som heltall (derfor har `error` variabelen datatypen `int`).

I den tekniske spesifikasjonen til Støvsensoren står det at den kun kan leses av én gang hvert sekund. `read` kommandoen vil derfor feile om vi spør etter målinger oftere enn det. Dersom `read`-kommandoen feilet vil `error` ha en verdi ulik `0`. Den beste måten å håndtere denne situasjonen på er å prøve på nytt helt til vi har fått inn ny data fra støvsensoren.

Får å forsikre at vi faktisk har en glydig måleverdi fra sensoren, må vi altså prøve å kjøre `read`-kommandoen om og om igjen helt til feilverdien er lik `0` (dvs. ingen feil). Når vi vil be Arduinoen om å kjøre den samme koden flere ganger bruker vi i programmering en løkke (*loop* på engelsk).

I `C++`, kodespråket til Arduino, finnes det tre typer løkker. I dette tilfellet skal vi benytte oss av *do-while*-løkken. Denne løkken kjører en kommando (eller flere) og repeterer dersom en betinglse er sann. Dette ser slikt ut:

``` cpp
  do {
    // Code to repeat
  } while (error != 0); // Stop repetition once error is 0.
```

I kodebiten over ser du at vi skjekker om `error` er ulik 0. Med én gang `error` er `0`, er vi sikre på at vi har fått gyldige målinger fra Støvsensoren og kan fortsette med å skrive dem ut. Så la oss fylle inn `read`-kommandoen mellom krøllparentesene til `do`-`while`-løkken vår.

``` cpp
  int error;
  do {
    error = sds.read(&pm25, &pm10);
  } while (error != 0);
```

**Merk** at vi deklarer utenfor løkken! [Bruk av variabler utenfor scope][debugging-scopes] er en vanlig feil som lett er gjort når du skriver kode. Gå til [Feilsøking av programmeringsfeil][debugging-scopes] får å vite mer om hvorfor vi flyttet deklarasjonen av `error` utenfor løkken.

Siden Støvsensoren kun kan avleses én gang per sekund er det lurt å bruke `delay`-kommandoen til slutt. Få Arduinoen til å vente for minst ett sekund.

## Ferdig

Under ser du et eksempel for koden når alt er på plass.

``` cpp
#include <SoftwareSerial.h>
#include <SDS011.h>

#define PM_TX 2
#define PM_RX 3

SDS011 sds;

void setup() {
  sds.begin(PM_TX, PM_RX);
  Serial.begin(9600);
}

void loop() {
  float pm25, pm10;
  int error;
  do {
    error = sds.read(&pm25, &pm10);
  } while (error != 0);

  Serial.print(pm25);
  Serial.print("\t");
  Serial.print(pm10);
  Serial.println();

  delay(1000);
}
```

## Gå videre

&uarr; [Gå til **innholdsfortegnelsen**][home]  
&larr; [Gå tilbake forrige neste steg: **Temperatursensoren**][dht]  
&rarr; [Gå til neste steg: **GPS-antenna**][gps]  

[home]: airbit-Programmering
[dht]: Programmering-med-Temperatursensoren
[gps]: Programmering-med-GPS-antenna

[pinout]: airbit-Pinout

[debugging-scopes]: Feilsøking-av-programmeringsfeil#bruk-av-variabler-utenfor-scope
