Støvsensoren til air:bit er litt mere komplisert enn temperatursensoren, men
igjen finnes det et Arduino bibliotekt for å lese av målinger fra sensoren.

Som i temperatursensor-eksemplet kommer vi til å demonstrere hvordan vi tar
målinger fra støvsensoren, og skrive dem ut til seriell-forbindelsen til PC'en.

## Ny Sketch

Igjen starter vi med en helt fersk ny Sketch. `File`&rarr;`New` i menyen.

``` cpp
void setup () {

}

void loop() {

}
```

## Laste ned bibliotek

Akkurat som DHT-biblioteket vi brukte for temperatursensoren, finnes det et
bibliotek for støvsensoren. Desverre oppstår det komplikasjoner med de andre
sensorene i air:bit når vi bruker denne. Derfor har air:bit teamet laget sin
egen versjon for dette biblioteket. Last ned den nyeste, dvs. den øverste,
versjonen i listen du finner når du klikker på denne linken:
(klikk det står `zip`)
**[SDS011 biblioteket](https://github.com/skolelab/SDS011/releases)**  
Det er viktig at du lagrer (**ikke åpner**) filen. Husk hvor du lagrer filen,
vi må finne den frem filen i neste steg.

Etter du har lastet ned biblioteket og lagret filen må vi installere
biblioteket i `Arduino IDE`. Klikk i menyen på 
`Sketch`&rarr;`Include library`&rarr;`Add ZIP library`
(punktet under `Manage libraries`). Velg filen du nettopp lastet ned. Om
nettleseren din ikke ba deg om å velge hvor filen skulle lagres, vil du mest
sannsynligvis finne den under `Downloads` (`Nedlastninger`).

## Globale definisjoner

Arduino IDE kommer med del små hjelper-bibliotek som allerede er installert på
forhånd. Ett av disse er `SoftwareSerial`-biblioteket. Støvsensoren sender
faktisk målingene over ledningen til Arduinoen på liknende måte som Arduinoen
'snakker' med PCen over USB-ledningen. `SoftwareSerial` tillater å lage en
seriell forbindelse over pinner på Arduinoen i stedet for USB-ledningen. Så
derfor bruker vi følgende `#include` direktiv:

``` cpp
#include <SoftwareSerial.h>
```

Først må vi inkludere `SDS011`-biblioteket. Bruk det følgende `#include`
direktivet:

``` cpp
#include <SDS011.h>
```

Så er det tid å konsultere [pinout skjemaet][pinout]. Som vanlig definerer vi
kontanter for pinnene som brukes for kommunikasjon med Arduinoen.

``` cpp
#define PM_TX 2
#define PM_RX 3
```

Så må lage en global variabel for tilkoblingen til Støvsensoren. Denne gang
bruker vi `SDS011` datatypen. *Igjen kan variabler navngis som du synes best.*

``` cpp
SDS011 sds;
```

## `setup`

Akkurat som kommunikasjonen med PC'en din, må vi også initialisere
kommunikasjonen med støvsensoren. Vi bruker `begin`-kommandoen som hører til
`sds` variablen. Den tar imot TX-pinnen og så RX-pinnen som sensoren er koblet
til med som argument.

``` cpp
  sds.begin(PM_TX, PM_RX);
```

Vi skal også skrive ut målingene over seriell tilkoblingen til datamaskinen.

``` cpp
  Serial.begin(9600);
```

## `loop`

Som med temperatursensoren, skal vi lese ut målingene fra støvsensoren og så
printe dem ut. Støvsensoren tar målinger for to forskjellige partikelstørrelser:
2.5µm og 10µm. Begge tall gir konsentrasjoner som desimaltall.

``` cpp
  float pm25, pm10;
```

*Merk at du kan deklarere flere variabler av samme type ved å skrive komma
mellom variablene.*

`read`-kommandoen til `sds`-variablen brukes for å lese ut data fra sensoren.
Kommandoen kan feile, så kommandoen returnerer en feilverdi som vi lagrer i en
variabel. Kommadoen tar plassering først for 2.5µm og så for 10µm
konsentrasjonen som argument.

``` cpp
  int error = sds.read(&pm25, &pm10);
```

I koden over ser du at feilverdier som regel lagres av datamaskinen som heltall
(derfor har `error` variablen datatypen `int`).

Dersom `read`-kommandoen feilet vil `error` ha en verdi ulik `0`. Så vi bruker
en `if`-block for å skjekke mot dette. Etter `if`-blokker følger som regel en
`else`-blokk. Koden i `if`-blokken kjøres dersom en betingelse evalueres til
`true` (sant), `else`-blokken kjøres når betingelsen er `false` (usann). Koden
ser slik ut:

``` cpp
  if (error == 0) { // Check agains 0.
    // 0 means no error.
  } else { // Error is not equal to 0.
    // Not equal to 0 means there is an error.
  }
```

Nå må vi fylle blokkene med nyttig kode. Først, dersom `read`-kommandoen gikk
bra, må vi skrive ut målingene som vi har gjort tidligere. Koden for dette hører
til i `if`-blokken.

I `else`-blokken kan vi skrive kode for å skrive en feilmelding. Det er ofte
lurt å skrive ut verdien til `error`-variablen, slik at feilverdien kan skjekkes
i dokumentasjonen.

Den vanligste feilen for avlesning med støvsensoren er når du prøver å lese av
før det er noe data tilgjengelig. Støvsensoren kan kun avleses én gang per
sekund. Derfor er det lurt å bruke `delay`-kommandoen til slutt. Få Arduinoen
til å vente for minst ett sekund.

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
  int error = sds.read(&pm25, &pm10);

  if (error == 0) {
    Serial.print(pm25);
    Serial.print("\t");
    Serial.print(pm10);
    Serial.println();
  } else {
    Serial.print("Could not read air data. Error code: ");
    Serial.print(error);
    Serial.println();
  }

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
