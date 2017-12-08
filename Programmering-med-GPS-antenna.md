Neste komponent vi kan se nærmere på er GPS-antenna til air:bit. Igjen kommer
vi til å skjekke for GPS posisjonen og skrive ut dataen fra GPS'en til seriell
kommunikasjon. Men for å gjøre ting litt mere interessant kommer vi også til å
lyse opp LED-lysene for å vise at du har kontakt med nok GPS-satellitter for å
bestemme posisjon.

Å bruke LED-lysene får å vise GPS signal kan være en god idé, siden det fort
kan skje at man mister kontakt med GPS avhengig av forhold, fjell eller
tunneller på veien. Det vil også være mye vanskeligere å få GPS-kontakt nå du
er inne i hus, spesiellt om huset har tykke betongvegger. Når du går rundt ute
og tar målinger forventes det ikke at du har en svær datamaskin koblet til
air:bit hele tiden for å kunne lese ut tekst du sender ut via seriell-koblingen.

## Ny Sketch

Igjen starter vi med en helt fersk ny Sketch. `File`&rarr;`New` i menyen.

``` cpp
void setup () {

}

void loop() {

}
```

## Laste ned og installere `TinyGPS++` biblioteket

Igjen skal vi bruke et bibliotek for å kommunisere med GPS-antenna. Denne gangen
skal vi igjen bruke et bibliotek du ikke finner i `Library Manager`. Klikk på
følgende link for å laste ned: **[TinyGPS++ Biblioteket][tiny-gpp-dl-link]**

[![Nedlastning av TinyGPS++ biblioteket](TinyGPSPlusPlus-library-download.png)][tiny-gpp-dl-link]

Klikk på den store grønne pilen der det står Download. Husk at du ikke vil åpne
file du laster ned, men at du lagre den!

I `Arduino IDE`, finn menyen `Sketch`&rarr;`Include library`&rarr;`Add ZIP library`
og velg filen du nettopp lastet ned.

## Globale definisjoner

I likhet med `SDS011`-biblioteket du brukte for støvsensoren, bruker også
`TinyGPS++` `SoftwareSerial`-biblioteket. Navnet for `#include` direktivet for
`TinyGPS++`-biblioteket er `<TinyGPS++.h>`. Legg til begge
`#include`-direktivene i toppen av sketchen din.

Så sjekker du [pinout skjemaet][pinout] for `RX` og `TX` pinnene for
GPS-antenna og lager definisjoner for disse. Vi må også lage definisjoner for
LED-lysene våre, siden vi skal bruke dem også.

Så til globale variabler som kontrollerer kommunikasjonen med antenna.
`SDS011`-biblioteket tok seg automatisk av håndteringen for
seriell-kommunikasjonen med støvsensoren. `TinyGPS++` gjør desverre ikke det
samme for GPS'en (`Tiny` er hintet her). Derfor trenger vi en variabel som
styrer kommunikasjonen til GPS'en.

``` cpp
SoftwareSerial gpsCom(GPS_RX, GPS_TX);
```

Ikke bry deg alt for mye om at variabel-deklarasjonen i linja over ser litt
uvanlig ut. Datatypen er `SoftwareSerial`, navnet på variablen er `gpsCom`.
Verdiene `GPS_TX` og `GPS_RX` er konstantene for GPS-pinnene på Arduinoen (det
kan hende at du gir dem andre navn i din kode).

Til slutt må vi også lage en variabel som representerer selve sensoren. Her
bruker vi nå datatypen som følger med `TinyGPS++`-biblioteket, den heter
`TinyGPSPlus`.

``` cpp
TinyGPSPlus gps;
```

## `setup`

Husk at vi må bruke `pinMode`-kommandoen for å få å aktivere kontroll over
LED-lysene på Arduinoen. Husk at `OUTPUT` skal brukes som andre argument (i
tillegg til hvilken pinne du skal aktivere).

Bruk `Serial.begin`-kommandoen for å starte opp seriell kommunikasjonen med
datamaskinen over USB. Bruk verdien `9600` som argument her.

Og så må vi også starte opp kommunikasjonen med GPS-antenna. Dette er også en
helt vanlig seriell-kommunikasjon, bare at vi bruker `gpsCom`-variablen i
stedet for `Serial`, så det blir:

``` cpp
  gpsCom.begin(9600);
```

## `loop`

GPS'en forteller oss om den har en gyldig posisjon. I tilegg til å printe ut
teksten over USB-ledningen til datamaskinen vil vi bruke LED-lysene for å blinke
grønt når vi har kontakt med GPS-satellitten. Og så kan vi blinke rødt når vi
ikke får kontakt med satellitten. Dette betyr at vi må gjøre litt `if`-sjekking.

Generell tommelfingerregel i programmering er at man burde gjøre håndtering av
feil øverst i kodeblokkene sine. Dvs. skjekk først for alle feil. Dette er fordi
det som regel er enklere å teste mot spesifikke feilbetingelser. Vi kan også
skrive instruksjoner som får Arduinoen til å forlate `loop` blokken med en gang,
slik at den hopper videre til neste gjennomgang av `loop`-funksjonen.

Første tingen vi kan skjekke er om GPS'en har sent noe data. Vi kan bruke
`available`-kommandoen til `gpsCom`-variablen vår for å se om det har kommet in
ny GPS data, som ikke enda har blitt evaluert. Som vi sa nettopp, skal vi
skjekke mot feil eller ugyldig status først. Dvs. vi vil skjekke om vi **ikke**
har ny data tilgjengelig. Vi kan bruke `!` operatoren i C++ for å negere en
påstand:

``` cpp
  if (!gpsCom.available()) {
    // No new data available.
  }
```

Merk at det ofte er lurt å skrive ut hele `if`-blokken med en gang. Vi kan fylle
ut kode mellom krøllparantesene (`{`, `}`) etterpå. På denne måten glemmer du
ikke å skrive inn parantesene eller krøllparantesene som må være her.
`if`-blokker ser **altid** nøyaktig ut som vist over, så det er bare noe
man lærer å skrive i søvne etterhvert.

Okei i betingelsen over ser du at vi bruker `gpsCom.available()` for å skjekke
om data er tilgjengelig. Denne kommandoen vil svare med verdien `true` dersom
det er data tilgjengelig. `!` vil så negere denne verdien. Derfor skjekker denne
betingelsen om det **ikke** er data tilgjengelig.

Når vi ikke har ny data å lese, er det liten vits å fortsette. Det betyr
egentlig bare at GPS'en ikke har hatt tid nok til å sende ny dataen til
Arduinoen. I likhet med støvsensoren har GPS'en en viss oppdateringshastighet.
Du kan ikke spørre den om data oftere/fortere enn den klarer å håndtere.

Så, om vi ikke har noe data tilgjengelig, vil vi egentlig bare avslutte den
aktuelle gjennomgangen gjennom `loop` med en gang. Uten data er det rett og
slett ikke noe annet vi kan gjøre. Vi bruker `return` instruksjonen for dette:

``` cpp
  if (!gpsCom.available()) {
    // No new data available.
    // return immediately
    return;
  }
```

Da har vi forsikret oss at vi har data tilgjengelig, siden `gpsCom.available()`
da vil evaluere til `true`, negasjonen av dette blir `false`, og `if`-blokken
vil ikke utføres dersom betingelsen evalueres til `false`, som vil få Arduinoen
til å 'hoppe over' koden med `return` instruksjonen.  
Nå må vi altså lese inn den nye dataen GPSen har sendt. Vi bruker nå to
kommandoen inni hverandre. Innerst bruker vi `gpsCom.read()` for å lese data,
og rundt dette bruker vi `gps.encode()` for å tolke dataen vi har mottat:

``` cpp
  bool complete = gps.encode(gpsCom.read());
```

Merk også at `encode`-kommandoen gir oss en status-verdi. Denne vil være `true`
dersom GPS-dataen ble avlest rett. Vi lagrer denne verdien i en variabel av
type `bool`. `bool` er datatypen som brukes for å lagre sannheter eller det vi
kaller påstander. Den kan ha verdiene `true` eller `false`.

Av og til vil vi ende opp i en situasjon at vi har fått litt data fra GPS'en,
men at den ikke enda har blitt ferdig med sende alt enda. Se for deg en
nedlastning av en stor fil fra internettet. Du må vente til hele filen, dvs.
all dataen i den filen er lastet ned før du kan bruke den.

Så igjen har vi fin feiltest her: Dersom ikke `encode`-kommandoen ble fullført,
kan vi igjen avbryte gjennomgangen gjennom `loop` med en gang. Dvs. vi venter
til vi er sikre på at all data har blitt lest, og det var noe nyttig informasjon
der. Vi har `complete`-variablen for teste dette. Igjen vil teste mot feil-
betingelsen, så vi må igjen negere påstanden:

``` cpp
  if (!complete) {
    // Data is incomplete, 
    // nothing to do yet, either.
    return;
  }
```

Når vi har klart å komme gjennom begge betingelsene uten å finne en feil (dvs.
enten ingen data, eller ikke enda nok data) kan vi være sikre på at vi har fått
informasjon fra GPSen og den har klart å tolke data som har kommet inn. La oss
nå skjekke om vi har en gyldig posisjon. Vi bruker
`location.isValid()`-kommandoen til `gps`-variablen for dette. Igjen er 
`location.isValid()` en sannhetsverdi, så la oss sjekke mot usann, dvs. om 
tilfellet der vi ikke har en gyldig posisjon:

``` cpp
  if (!gps.location.isValid()) {
    // No valid position.
    // I.e. no GPS fix.
  }
```

Hvis vi har fått kontakt med satellitter og så etterhvert mister kontakten igjen
(f.eks. fordi du går inn i en tunnel, et hus med tykke vegger, o.l.) vil
`location.isValid` forsette å gi `true` som svar. Så vi må også sjekke om dataen
har blitt oppdatert siden siste gang vi leste av. Vi bruker `location.isUpdated`
for dette. Så vi endrer betingelsen og bruker `&&` for knytte to betingelser
sammen. `&&` betyr *OG*, dvs. både høyre og venstre betingelsene må være `true`
for at hele betingelsen skal være `true`. Merk at vi forsatt trekker negasjonen
utenfor ved å sette parenteser rundt hele uttrykket:

``` cpp
  if (!(gps.location.isValid() && gps.location.isUpdated())) {
    // No valid position.
    // I.e. no GPS fix.
  }
```

Følgende kode vil være lik den over, men kan være litt enklere å lese:

``` cpp
  bool gpsValid = gps.location.isValid();
  bool gpsUpdated = gps.location.isUpdated();
  bool isUseful = gpsValid && gpsUpdated;
  if (!isUseful) {
    // No valid position.
    // I.e. no GPS fix.
  }
```

*Merk at vi kaller `!` for **NOT**-operatoren. Dvs. betingelsen leses som:
NOT isUseful*

I tilfeller der vi ikke har en gyldig posisjon, har vi ikke kontakt med
satellitten. Da ville vi blinke rødt. Husk hvordan vi gjorde det: 

1. Skru på rødt lys. `digitalWrite` med `HIGH` som andre argument.
1. Vent litt. `delay`
1. Skru av rødt lys. `digitalWrite` med `LOW` som andre argument.

``` cpp
  bool gpsValid = gps.location.isValid();
  bool gpsUpdated = gps.location.isUpdated();
  bool isUseful = gpsValid && gpsUpdated;
  if (!isUseful) {
    // No valid position.
    // I.e. no GPS fix.
    Serial.println("No valid GPS position");
    digitalWrite(LED_RED, HIGH);
    delay(500);
    digitalWrite(LED_RED, LOW);
    return;
  }
```

*Merk at koden over også har med en instruksjon for å printe ut over
seriell-koblingen med datamaskinen at vi ikke fikk kontakt.*

Så til slutt det vi *egentlig* hadde lyst å gjøre: Blinke grønt, og printe ut
GPS dataen over seriell-koblingen til datamaskinen.

Det enkleste først: Blink det grønne LED-lyset for å vise at vi har fått kontakt
med GPS-satellitter.

`gps.location.lat()` og `gps.location.lng()` gir oss posisjonen i lengde- og
breddegrader som kommatall. I en `Serial.print()`-kommando kan du legge til et
argument for å spesifisere antall desimaler bak kommaet, når du skriver ut
kommatall. Eksemplet under vil skrive ut posisjonen med `6` desimaler bak
kommaet:

``` cpp
  Serial.print("Latitude: ");
  Serial.print(gps.location.lat(), 6); // Latitude in degrees
  Serial.print("\t");
  Serial.print("Longitude: ");
  Serial.print(gps.location.lng(), 6); // Longitude in degrees
  Serial.println();
```

## GPS data

GPS-antenna er komponenten med flest tilgjengelig informasjon. Du kan se hav du
kan hente ut i listen under. Du har med hjelp av denne mulighet til å selv
eksperimentere med hva du kan få ut av GPS'en. Posisjon og tid vil være mest
nyttig for målinger du skal gjøre med air:bit. Men det er alltids gøy å leke seg
med de andre verdiene her også.

*Listen er kopiert fra nettsiden der du lastet ned
`TinyGPS++`-biblioteket ifra.  
Den viser også eksempelvis kode for hvordan du kan printe ut ting.

``` cpp
Serial.println(gps.location.lat(), 6); // Latitude in degrees (double)
Serial.println(gps.location.lng(), 6); // Longitude in degrees (double)
Serial.print(gps.location.rawLat().negative ? "-" : "+");
Serial.println(gps.location.rawLat().deg); // Raw latitude in whole degrees
Serial.println(gps.location.rawLat().billionths);// ... and billionths (u16/u32)
Serial.print(gps.location.rawLng().negative ? "-" : "+");
Serial.println(gps.location.rawLng().deg); // Raw longitude in whole degrees
Serial.println(gps.location.rawLng().billionths);// ... and billionths (u16/u32)
Serial.println(gps.date.value()); // Raw date in DDMMYY format (u32)
Serial.println(gps.date.year()); // Year (2000+) (u16)
Serial.println(gps.date.month()); // Month (1-12) (u8)
Serial.println(gps.date.day()); // Day (1-31) (u8)
Serial.println(gps.time.value()); // Raw time in HHMMSSCC format (u32)
Serial.println(gps.time.hour()); // Hour (0-23) (u8)
Serial.println(gps.time.minute()); // Minute (0-59) (u8)
Serial.println(gps.time.second()); // Second (0-59) (u8)
Serial.println(gps.time.centisecond()); // 100ths of a second (0-99) (u8)
Serial.println(gps.speed.value()); // Raw speed in 100ths of a knot (i32)
Serial.println(gps.speed.knots()); // Speed in knots (double)
Serial.println(gps.speed.mph()); // Speed in miles per hour (double)
Serial.println(gps.speed.mps()); // Speed in meters per second (double)
Serial.println(gps.speed.kmph()); // Speed in kilometers per hour (double)
Serial.println(gps.course.value()); // Raw course in 100ths of a degree (i32)
Serial.println(gps.course.deg()); // Course in degrees (double)
Serial.println(gps.altitude.value()); // Raw altitude in centimeters (i32)
Serial.println(gps.altitude.meters()); // Altitude in meters (double)
Serial.println(gps.altitude.miles()); // Altitude in miles (double)
Serial.println(gps.altitude.kilometers()); // Altitude in kilometers (double)
Serial.println(gps.altitude.feet()); // Altitude in feet (double)
Serial.println(gps.satellites.value()); // Number of satellites in use (u32)
Serial.println(gps.hdop.value()); // Horizontal Dim. of Precision (100ths-i32)
```

I koden over ser du at forfatteren av `TinyGPS++`-biblioteket har skrevet
datatypen dem bruker i parantes på slutten av hver linje. `i32`, `i16` og `i8`
er vanlige heltall (med enten 32 bits, 16 bits eller 8 bits lengde). Det samme
gjelder for typene med `u` i stedet for `i`, men disse er `unsigned`, dvs. de
har ikke fortegn og kan derfor kun gi positive verdier eller `0`. `double` er
en nyere kommatall-type enn `float`. *Klikk på følgende link, for en liste over
datatyper i C++: [Datatyper i C++][datatypes]*

## Ferdig

Avhengig av hvlike verdier du printer ut, hvordan du navngir dine variabler,
osv. kan koden dir se litt anderlesed ut enn her. Men i essense burde ting 
stemme rimelig overens med dette:

``` cpp
#include <SoftwareSerial.h>
#include <TinyGPS++.h>

#define LED_RED A1
#define LED_GREEN A0

#define GPS_RX 7
#define GPS_TX 6

SoftwareSerial gpsCom(GPS_RX, GPS_TX);
TinyGPSPlus gps;

void setup() {
  // Activate control over LEDs
  pinMode(LED_RED, OUTPUT);
  pinMode(LED_GREEN, OUTPUT);

  // Initialize serial communication
  Serial.begin(9600); // to Computer (USB)
  gpsCom.begin(9600); // to GPS antenna
}

void loop() {
  if (!gpsCom.available()) {
    // No new data available.
    // return immediately
    return;
  }

  bool complete = gps.encode(gpsCom.read());
  if (!complete) {
    // Data is incomplete, 
    // nothing to do yet, either.
    return;
  }

  bool gpsValid = gps.location.isValid();
  bool gpsUpdated = gps.location.isUpdated();
  bool isUseful = gpsValid && gpsUpdated;
  if (!isUseful) {
    // No valid position.
    // I.e. no GPS fix.
    Serial.println("No valid GPS position");
    digitalWrite(LED_RED, HIGH);
    delay(500);
    digitalWrite(LED_RED, LOW);
    return;
  }

  digitalWrite(LED_GREEN, HIGH);
  delay(500);
  digitalWrite(LED_GREEN, LOW);

  Serial.print("Time: ");
  Serial.print(gps.date.day());
  Serial.print(".");
  Serial.print(gps.date.month());
  Serial.print(".");
  Serial.print(gps.date.year());
  Serial.print(" ");
  Serial.print(gps.time.hour());
  Serial.print(":");
  Serial.print(gps.time.minute());
  Serial.print(":");
  Serial.print(gps.time.second());
  Serial.print(".");
  Serial.print(gps.time.centisecond());
  
  Serial.print("\t");

  Serial.print("Latitude: ");
  Serial.print(gps.location.lat(), 6); // Latitude in degrees
  Serial.print("\t");
  Serial.print("Longitude: ");
  Serial.print(gps.location.lng(), 6); // Longitude in degrees
  
  Serial.println();
}
```

## Gå videre

&uarr; [Gå til **innholdsfortegnelsen**][home]  
&larr; [Gå tilbake forrige neste steg: **Støvsensoren**][pm]  
&rarr; [Gå til neste steg: **SD-kortet og filer**][sd]  

[tiny-gpp-dl-link]: http://arduiniana.org/libraries/tinygpsplus/
[tiny-ggp-dl-img]: TinyGPSPlusPlus-library-download.png

[home]: airbit-Programmering
[pm]: Programmering-med-Stovsensoren
[sd]: Programmering-av-filer-pa-SD-kortet

[datatypes]: Primitive-datatyper-i-C-programmering
[pinout]: airbit-Pinout
