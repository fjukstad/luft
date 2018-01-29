Neste komponent vi kan se nærmere på er GPS-antennen til air:bit. Igjen kommer vi til å skjekke for GPS posisjonen og skrive ut dataen fra GPS'en til seriell kommunikasjon. Men for å gjøre ting litt mere interessant kommer vi også til å lyse opp LED-lysene for å vise at du har kontakt med nok GPS-satellitter for å bestemme posisjon.

Å bruke LED-lysene får å vise GPS signal kan være en god idé, siden det fort kan skje at man mister kontakt med GPS avhengig av forhold, fjell eller tunneller på veien. Det vil også være mye vanskeligere å få GPS-kontakt nå du er inne i hus, spesiellt om huset har tykke betongvegger. Når du går rundt ute og tar målinger forventes det ikke at du har en svær datamaskin koblet til air:bit hele tiden for å kunne lese ut tekst du sender ut via seriell-koblingen.

## Ny Sketch

Igjen starter vi med en helt ny Sketch. `File`&rarr;`New` i menyen.

``` cpp
void setup () {

}

void loop() {

}
```

## Laste ned og installere `TinyGPS++` biblioteket helplink

Igjen skal vi bruke et bibliotek for å kommunisere med GPS-antennen. Denne gangen skal vi igjen bruke et bibliotek du ikke finner i `Library Manager`. Klikk på følgende link for å laste ned: **[TinyGPS++ Biblioteket][tiny-gpp-dl-link]**

[![Nedlastning av TinyGPS++ biblioteket](TinyGPSPlusPlus-library-download.png)][tiny-gpp-dl-link]

Klikk på den store grønne pilen der det står Download. Husk at du ikke vil åpne file du laster ned, men at du lagre den!

I `Arduino IDE`, finn menyen `Sketch`&rarr;`Include library`&rarr;`Add ZIP library` og velg filen du nettopp lastet ned.

## Globale definisjoner helplink

I likhet med `SDS011`-biblioteket du brukte for støvsensoren, bruker også `TinyGPS++` `SoftwareSerial`-biblioteket. Navnet for `#include` direktivet for `TinyGPS++`-biblioteket er `<TinyGPS++.h>`. Legg til begge `#include`-direktivene i toppen av sketchen din.

Så sjekker du [pinout skjemaet][pinout] for `RX` og `TX` pinnene for GPS-antennen og lager definisjoner for disse. Vi må også lage definisjoner for LED-lysene våre, siden vi skal bruke dem også.

Merk at GPS-antennen har `RX`- og `TX`-linjer. `RX`-linjen er linjen GPS-antennen får kommandoer fra Arduinoen over (`RX` for *receive transmission*). `TX`-linjen er linjen GPS-antennen sender data over til Arduinoen (`TX` for *transmit transmission*). På Arduinoen derimot er det omvendt, siden Arduinoen må **lytte** på linjen GPSen *sender* på, og må **sende** på linjen GPSen *lytter* på. Derfor må vi bytte om på verdiene for `RX` og `TX` på Arduinoen. Så `GPS_RX` blir `7` og `GPS_TX` blir `6`.

Så til globale variabler som kontrollerer kommunikasjonen med antennen. `SDS011`-biblioteket tok seg automatisk av håndteringen for seriell-kommunikasjonen med støvsensoren. `TinyGPS++` gjør desverre ikke det samme for GPS'en (`Tiny` er hintet her). Derfor trenger vi en variabel som styrer kommunikasjonen til GPS'en.

``` cpp
SoftwareSerial gpsCom(GPS_RX, GPS_TX);
```

Ikke bry deg alt for mye om at variabel-deklarasjonen i linja over ser litt uvanlig ut. Datatypen er `SoftwareSerial`, navnet på variabelen er `gpsCom`. Verdiene `GPS_TX` og `GPS_RX` er konstantene for GPS-pinnene på Arduinoen (det kan hende at du gir dem andre navn i din kode).

Til slutt må vi også lage en variabel som representerer selve sensoren. Her bruker vi nå datatypen som følger med `TinyGPS++`-biblioteket, den heter `TinyGPSPlus`.

``` cpp
TinyGPSPlus gps;
```

## `setup` helplink

Husk at vi må bruke `pinMode`-kommandoen for å aktivere kontroll over LED-lysene på Arduinoen. Husk at `OUTPUT` skal brukes som andre argument (i tillegg til hvilken pinne du skal aktivere).

Bruk `Serial.begin`-kommandoen for å starte opp seriell kommunikasjonen med datamaskinen over USB. Bruk verdien `9600` som argument her.

Og så må vi også starte opp kommunikasjonen med GPS-antennen. Dette er også en helt vanlig seriell-kommunikasjon, bare at vi bruker `gpsCom`-variabelen i stedet for `Serial`, så det blir:

``` cpp
  gpsCom.begin(9600);
```

## `loop` helplink

GPS'en forteller oss om den har en gyldig posisjon. I tilegg til å printe ut teksten over USB-ledningen til datamaskinen vil vi bruke LED-lysene for å blinke grønt når vi har kontakt med GPS-satellitten, og så kan vi blinke rødt når vi ikke får kontakt med satellitten. Dette betyr at vi må gjøre litt sjekking etter feilkoder.

Men aller først må vi instruere Arduinoen til å *høre* på innkommende data fra GPSen. Vi bruker `listen`-kommandoen til `gpsCom`-variabelen for dette.

``` cpp
  gpsCom.listen();
```

Generell tommelfingerregel i programmering er at man burde gjøre håndtering av feil øverst i kodeblokkene sine. Dvs. skjekk først for alle feil. Dette er fordi det som regel er enklere å teste mot spesifikke feilbetingelser. Når det gjelder avlesning av sensorer (eller GPSen) er det vanlig å prøve å lese av verdier i en løkke slik at man er sikker å ha målinger for resten av programmet.

Første tingen vi kan skjekke er om GPS'en har sent noe data. Vi kan bruke `available`-kommandoen til `gpsCom`-variabelen vår for å se om det har kommet in ny GPS data, som ikke enda har blitt evaluert. Som vi sa nettopp, skal vi skjekke mot feil eller ugyldig status først. Dvs. vi vil skjekke om vi **ikke** har ny data tilgjengelig. Vi kan bruke `!` operatoren i C++ for å negere en påstand:

``` cpp
  if (!gpsCom.available()) {
    // No new data available.
  }
```

Merk at det ofte er lurt å skrive ut hele `if`-, `while`, eller `do`-blokken med en gang. Vi kan fylle ut kode mellom krøllparantesene (`{`, `}`) etterpå. På denne måten glemmer du ikke å skrive inn parantesene eller krøllparantesene som må være her. `if` og `while`-blokker ser **alltid** nøyaktig ut som vist over, så det er bare noe man lærer å skrive i søvne etterhvert.

Okei, i betingelsen over ser du at vi bruker `gpsCom.available()` for å skjekke om data er tilgjengelig. Denne kommandoen vil svare med verdien `true` dersom det er data tilgjengelig. `!` vil så negere denne verdien. Derfor skjekker denne betingelsen om det **ikke** er data tilgjengelig.

Da har vi forsikret oss om at vi har data tilgjengelig. Nå må vi altså lese inn den nye dataen GPSen har sendt. Vi kommer nå til å bruke to kommandoer inni hverandre. Innerst bruker vi `gpsCom.read()` for å lese data, og rundt dette bruker vi `gps.encode()` for å tolke dataen vi har mottat:

``` cpp
  gps.encode(gpsCom.read());
```

Merk også at `encode`-kommandoen gir oss en status-verdi. Denne vil være `true` dersom GPS-dataen ble avlest rett. For enkelhetens skyld burde vi nå lage oss en variabel som lagrer resultatet av linjen over.

``` cpp
  bool gpsEncodeComplete;
  gpsEncodeComplete = gps.encode(gpsCom.read());
``` 

Så kan vi gjøre en `if`-test på denne verdien også:

``` cpp
  if (!gpsEncodeComplete) {
    // Not enough data to encode successfully yet
  }
```

Okei, nå må vi faktisk håndtere lese- og enkoderings-feilene. I prinsippet kommer begge feilene av at Arduinoen ikke enda har mottatt nok data fra GPSen, så det som må gjøres er å kjøre hele den koden vi nettop skrev i en løkke, helt til `gpsEncodeComplete` er `true`. Dette betyr at vi legger hele koden i en `do`-`while`-løkke.

``` cpp
  bool gpsEncodeComplete = false;
  do {
    if (!gpsCom.available()) {
      // No new data available.
    }
    gpsEncodeComplete = gps.encode(gpsCom.read());
    if (!gpsEncodeComplete) {
      // Not enough data to encode successfully yet
    }
  } while (!gpsEncodeComplete); 
```

*Merk at vi flyttet deklarasjonen av `gpsEncodeComplete`-variabelen utenfor `do`-`while`-løkken og initialiserer den til `false` til å begynne med. Se [Bruk av variabler utenfor scope][debugging-var-out-of-scope] for å lære hvorfor.*

Inni begge `if`-setningen vil vi egentlig at Arduinoen bare starter en ny gjennomgang gjennom `do`-`while`-løkken. I koden bruker vi `continue`-instruksjonen for dette:

``` cpp
  bool gpsEncodeComplete = false;
  do {
    if (!gpsCom.available()) {
      // No new data available.
      // Immediately jump to next iteration
      continue;
    }
    gpsEncodeComplete = gps.encode(gpsCom.read());
    if (!gpsEncodeComplete) {
      // Not enough data to encode successfully yet
      // Jump to next iteration and try again
      continue;
    }
  } while (!gpsEncodeComplete); 
```

*Under testingen kan du godt legge til en `Serial.println` før `continue`-instruksjonen for å se at Arduinoen faktisk gjør noe.*

Etter dette skriver vi koden som kjører når vi har klart å komme gjennom begge betingelsene uten å finne en feil (dvs. enten ingen data, eller ikke enda nok data). Nå kan vi være sikre på at vi har fått informasjon fra GPSen og den har klart å tolke data som har kommet inn. La oss nå skjekke om vi har en gyldig posisjon. Vi bruker `location.isValid()`-kommandoen til `gps`-variabelen for dette. Igjen er `location.isValid()` en sannhetsverdi, så la oss sjekke mot usann, dvs. om tilfellet der vi ikke har en gyldig posisjon. Nå når vi har data bruker vi `if`-setninger for å skjekke betingelser:

``` cpp
  if (!gps.location.isValid()) {
    // No valid position.
    // I.e. no GPS fix.
  }
```

Hvis vi har fått kontakt med satellitter og så etterhvert mister kontakten igjen (f.eks. fordi du går inn i en tunnel, et hus med tykke vegger, o.l.) vil `location.isValid` fortsette å gi `true` som svar, så vi må også sjekke om dataen har blitt oppdatert siden siste gang vi leste av. Vi bruker `location.isUpdated` for dette, så vi endrer betingelsen og bruker `&&` for knytte to betingelser sammen. `&&` betyr *OG*, dvs. både høyre og venstre betingelsene må være `true` for at hele betingelsen skal være `true`. Merk at vi forsatt trekker negasjonen utenfor ved å sette parenteser rundt hele uttrykket:

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

Vi lagrer sannhetsverdier (eller påstander) i variabler av type `bool`. I motsetning til `int` (som lagrer heltall) kan `bool`-variabler kun ha verdiene `true` eller `false`.

*Merk at vi kaller `!` for **NOT**-operatoren. Dvs. betingelsen leses som: NOT isUseful*

I tilfeller der vi ikke har en gyldig posisjon, har vi ikke kontakt med satellitten. Da ville vi blinke rødt. Husk hvordan vi gjorde det: 

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

Vi bruker `return`-instruksjonen for å stoppe utførningen av kode inni `loop`-funksjonen. Dette vil få Arduinoen til å hoppe rett til neste gjennomgang av `loop`-funksjonen og starte GPS avlesningen helt på nytt igjen. I tilfeller der vi ikke har noe brukbar GPS informasjon, må vi bare vente til vi får et godt signal, så å starte på nytt virker som en god idé.

*Merk at koden over også har med en instruksjon for å printe ut over seriell-koblingen med datamaskinen at vi ikke fikk kontakt.*

Så til slutt det vi *egentlig* hadde lyst å gjøre: Blinke grønt, og printe ut GPS dataen over seriell-koblingen til datamaskinen.

Det enkleste først: Blink det grønne LED-lyset for å vise at vi har fått kontakt med GPS-satellitter.

`gps.location.lat()` og `gps.location.lng()` gir oss posisjonen i lengde- og breddegrader som kommatall. I en `Serial.print()`-kommando kan du legge til et argument for å spesifisere antall desimaler bak kommaet, når du skriver ut kommatall. Eksemplet under vil skrive ut posisjonen med `6` desimaler bak kommaet:

``` cpp
  Serial.print("Latitude: ");
  Serial.print(gps.location.lat(), 6); // Latitude in degrees
  Serial.print("\t");
  Serial.print("Longitude: ");
  Serial.print(gps.location.lng(), 6); // Longitude in degrees
  Serial.println();
```

## GPS data helplink

GPS-antennen er komponenten med flest tilgjengelig informasjon. Du kan se hva du kan hente ut i listen under. Du har med hjelp av denne mulighet til å selv eksperimentere med hva du kan få ut av GPS'en. Posisjon og tid vil være mest nyttig for målinger du skal gjøre med air:bit. Men det er alltids gøy å leke seg med de andre verdiene her også.

*Listen er kopiert fra nettsiden der du lastet ned `TinyGPS++`-biblioteket ifra.  
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

I koden over ser du at forfatteren av `TinyGPS++`-biblioteket har skrevet datatypen dem bruker i parantes på slutten av hver linje. `i32`, `i16` og `i8` er vanlige heltall (med enten 32 bits, 16 bits eller 8 bits lengde). Det samme gjelder for typene med `u` i stedet for `i`, men disse er `unsigned`, dvs. de har ikke fortegn og kan derfor kun gi positive verdier eller `0`. `double` er en nyere kommatall-type enn `float`.

## Ferdig

Avhengig av hvlike verdier du printer ut, hvordan du navngir dine variabler, osv. kan koden din se litt anderlesed ut enn her. Men i essense burde ting stemme rimelig overens med dette:

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
  gpsCom.listen();
  bool gpsEncodeComplete = false;
  do {
    if (!gpsCom.available()) {
      // No new data available.
      // Immediately jump to next iteration
      continue;
    }
    gpsEncodeComplete = gps.encode(gpsCom.read());
    if (!gpsEncodeComplete) {
      // Data is incomplete, 
      // Jump to next iteration and try again
      continue;
    }
  } while (!gpsEncodeComplete); // Loop until gps data was successfully read and encoded from GPS module

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
[pm]: Programmering-med-Støvsensoren
[sd]: Programmering-av-filer-på-SD-kortet

[pinout]: airbit-Pinout
[debugging-var-out-of-scope]: Feilsøking-av-programmeringsfeil#bruk-av-variabler-utenfor-scope
