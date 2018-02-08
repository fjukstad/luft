Når du er ute og tar målinger med air:bit vil du ikke hele tiden ha den tilkoblet til din datamaskin. air:bit har ikke noe som forbinder seg med internettet, så du vil trenge en måte å lagre måledata på. air:bit har en kortleser for microSD-kort. Et microSD-kort er i prinsippet bare en minnebrikke, veldig lik en vanlig USB-minnepen og faktisk nøyaktig lik minnekortet som finnes for å lagre data på nyere mobiltelefoner, digitalkamera, osv.

microSD-kortet som følger med air:bit har ganske stor lagringskapasitet. Du vil kunne lagre flere år med målingsdata på dette, så i prinsippet skal du ikke bekymre deg for at du ikke har nok plass på "harddisken" til air:bit.

Du vil legge merke til at microSD-kortet kommer med adapter. Det bitte lille microSD-kortet passer inn i den større adapteren og former slikt et vanlig SD-kort. Laptopter nå til dags har som regel porter hvor du kan plugge inn SD-kort. Ellers finnes det også veldig vanlige kortterminaler man kan kjøpe som har SD-kort plugger og som kobles i en vanlig USB-port på PC-en din.

Når du plugger inn SD-kortet inn i en PC, vil det dukke opp som en harddisk i Filutforskeren (`Windows Explorer`, eller `Min datamaskin`, eller `This PC`) akkurat som en vanlig USB-minnebrikke også ville gjort.

## Beskrivelse av eksemplet

I dette eksemplet kommer du til å først lage en ny fil på SD-kortet. Så kommer du til å skrive ned tekst inn i den filen. For hver gjennomgang gjennom `loop` skal vi øke en teller og skrive ned verdien av telleren.

Dette vil derfor ligne veldig på [kodeeksemplet om variabler og telling i Arduino][counting-example]. Du vil i prinsippet gjøre nøyaktig det samme i dette eksemplet, borsett fra at du skriver til en fil på SD-kortet, i stedet for seriell-tilkoblingen til datamaskinen.

## Ny Sketch helplink

Start med en helt ny Sketch. Menypunkt: `File`&rarr;`New`.

``` cpp
void setup () {

}

void loop() {

}
```

## Laste ned og installere bibliotek helplink

`SD`-biblioteket er allerede forhåndsinstallert sammen med `Arduino IDE`. Det vil ikke være nødvendig å laste ned noe for dette eksemplet.

## Globale definisjoner helplink

Først må du inkludere `SD`-biblioteket. Bruk navnet `<SD.h>` for `#include` direktivet.

Fra [pinout skjemaet][pinout] ser du at det er flere pinner som kobler SD-kortleseren til Arduinoen. Vi trenger en definisjon for `CS` pinnen i koden vår.

``` cpp
#define SD_CS_PIN 10
```

Så vil du lage en global variabel får å representere skrivingen til filen på SD-kortet. Vi bruker datatypen `File` for dette. Merk at å lage en variabel vil **ikke** automatisk lage en fil på SD-kortet. Variabelen her er på en måte bare den ene enden av en telefonledning som går fra din Arduino til SD-kortet. Du vil skrive kode for å lage og åpne filen, og så skrive til den (lignende til at du må først slå ett nummer for å få en telefonsamtale, og du må faktisk begynne å prate for å sende "data", dvs. din tale, over ledningen).

``` cpp
File file;
```

*I koden over navngis tilkoblingen til filen `file`. Du kan bruke hvilket som helst navn for variabelen du vil. Dette vil ikke være navnen på filen du lagrer på SD-kortet!*

Til slutt må vi lage en global teller-variabel. Dette vil være et heltall, så bruk datatypen `int`.

## `setup` helplink

Først må du aktivere kontroll over `CS`-pinnen. Dette funker akkurat som du gjør får å aktivere pinnene til LED-lysene. Bruk `pinMode` med pinnen som første og `OUTPUT` som andre argument.

``` cpp
  pinMode(SD_CS_PIN, OUTPUT);
```

Vanligvis ville du initialisere seriell-kommunikasjonen til datamaskinen nå ved å bruke `Serial.begin` kommandoen, men la oss heller starte opp SD-kortet. Bruk `SD.begin` kommandoen og spesifiser `CS`-pinnen kortleseren er koblet til.

``` cpp
  SD.begin(SD_CS_PIN);
```

Nå må du bestemme deg for et navn for filen. Filnavnet er det du vil se i Filutforskeren (`Min datamaskin`) når du viser innholdet på SD-kortet etter du har plugget det inn i datamaskinen. Et filnavn er flere bokstaver etterhverandre. I `C++` bruker vi datatypen `char` for én bokstav. Vi bruker firkantparanteser `[` og `]` for å lage en variabel som lagrer *flere* verdier av samme type etterhverandre. Akkurat det vi trenger, det ser slikt ut:

``` cpp
  char filename[] = "testfile.txt";
```

I koden over ser du deklarasjonen av variabelen `filename`. Firkantparantesene skrives rett **etter** navnet. Verdien er bokstavene `testfile.txt`, og som vanlig står tekst mellom to anførelsestegn `"`.

I koden over er filnavnet `testfile.txt` valgt, men du kan bruke hvilket navn du vil. Å la filnavnet slutte med `.txt` er veldig vanlig, og vil automatisk fortelle til datamaskinen din at du mener å lagre vanlig tekst i filen. Du har kanskje sett at Word lagrer Word-dokumenter i `.doc` (eller nyere `.docx`) filer. Arduino-bibliotekene du lastet ned i tidligere guider bruker `.zip`, osv. Nå skal vi skrive ned tekst, så `.txt` virker mest hensiktsmessig.

Nå må du faktisk lage en ny fil på SD-kortet. Dette gjøres med kommandoen `SD.open`. Kommandoen tar først imot filnavnet, og så en *modus*. Filer kan åpnes bare for lesing, eller for skriving, eller begge deler. Når du åpner ett dokument i Word, vil det vanligvis åpnes for skriving, slik at du kan redigere dokumentet. En videospiller derimot vil bare lese en fil, du vil ikke endre på en video når du bare skal spille den av. Koden ser slik ut:

``` cpp
  file = SD.open(filename, O_CREAT | O_WRITE);
```

`SD.open` kommandoen vil gi oss tilkoblingen til filen som resultat, derfor lagrer vi denne i vår globale variabel for filtilkoblingen (kalt `file`). `O_CREAT | O_WRITE` betyr: Lag ny fil (`O_CREAT`) og åpne filen for skriving (`O_WRITE`). *Du finner `|` (pipe-symbolet) til venstre for `1`-tasten i den øverste raden på et vanlig norsk tastatur.*

Når filen er åpnet, la oss skrive en velkomsttekst til den første linja i filen:

``` cpp
  file.println("Dette er den første linjen i filen.");
```

Merk at vi bruker `file.println` der vi ville brukt `Serial.println` før. `file` viser her til navnet på fil-variabelen du bruker. Den kan være forskjellig dersom du har navngitt variabelen anderledes.

Se for deg at du skriver masse tekst i ett Word-dokument. Du vil kanskje ha lagt merke til at dokumentet ikke lagres til harddisken din med én gang du skriver noe. Du må trykke på `Lagre`-knappen for å lagre dokumentet. Du har kanskje også lagt merke til at ting blir borte dersom du glemmer å lagre.

Det samme gjelder for tilkoblingen vår til filen. Det holder ikke å bare bruke `print`, eller `println` kommandoene for å skrive noe til filen. Hver gang du er ferdig med å skrive ned en viktig del informasjon bør du bruke `flush`- kommandoen for å sikre at alt blir lagret på SD-kortet. Siden vi nå er kommet til slutten av `setup`-blokken er det en god idé og gjøre dette nå, før vi fortsetter med `loop`.

``` cpp
  file.flush();
```

Helt til slutt må du også huske å initialisere teller-variabelen til `0`.

Sålangt burde koden din se omtrent slik ut:

``` cpp
#include <SD.h>

#define SD_CS_PIN 10

File file;
int counter;

void setup () {
  // Activate CS-Pin control
  pinMode(SD_CS_PIN, OUTPUT);

  // Startup SD-card reader
  SD.begin(SD_CS_PIN);

  // Define filename
  char filename[] = "testfile.txt";

  // Create new file and open for writing
  file = SD.open(filename, O_CREAT | O_WRITE);

  file.println("Dette er den første linjen i filen.");
  file.flush(); // Force saving data to SD-card

  // Start counter at 0
  counter = 0;
}

void loop() {

}
```

## `loop` helplink

For hver gjennomgang gjennom `loop` vil vi skrive en ny linje med tekst. Teksten vi skriver skal også inneholde verdien til telleren vår. Vi må også øke telleren for hver gang.

Vi kan starte med å øke telleren med én. Bruk `+=` instruksjonen til dette.

Så skriver vi ned litt tekst, som f.eks `Dette er linje nr.:`. Så skriver vi en tekst med bare et mellomrom, for å skille `:` fra verdien til telleren. Så kan vi skrive selve teller-verdien. Bruk `file.print`-kommandoen flere ganger for å skrive ned alle delene av teksten og avslutt linjen ved å bruke `file.println`-kommandoen.

Når vi er ferdige, burde vi forsikre oss at teksten blir lagret på SD-kortet. Bruk `file.flush`-kommandoen helt til slutt.

``` cpp
void loop() {
  counter += 1;

  file.print("Dette er linje nr.:");
  file.print(" ");
  file.print(counter);
  file.println();

  file.flush();
}
```

Det kan også være en god idé å legge til en `delay`-kommando etter `flush`-kommandoen på slutten, slik at du kun skriver en ny linje etter f.eks. ett sekund med venting.

## Test 1 helplink

1. Last opp koden du har skrevet så langt til air:bit.
1. Trekk så ut ledningen fra datamaskinen.
1. Sett inn microSD-kortet i kortleseren til air:bit gjennom sprekken på siden der det står `SD-kort`. Trykk inn kortet til det sier "klikk".
1. Plugg USB-ledningen i batteriet til air:bit.
1. Vent en liten stund, f.eks. i 10 sekunder.
1. Trekk så ut ledningen fra batteriet igjen.
1. Trykk igjen på SD-kortet i air:bit, kortet burde sprette ut.
1. Skyv microSD-kortet i adapteren som følgte med i emballasjen til SD-kortet og plugg det hele i en SD-kortleser på din datamaskin.
1. PC-en din burde automatisk finne SD-kortet ditt når du plugger det inn. Du høre den vanlige lyden som spilles av når du plugger inn en USB-minnebrikke.
1. Om den ikke åpnes automatisk, velg `Min datamaskin` (`This PC`) i startmenyen.  
   Du vil se SD-kortet som en egen ny harddisk. Dobbel-klikk for å vise filene på SD-kortet.
1. Du burde kun se én fil her. Filen burde ha det navnet du bruket i koden din, f.eks. `testfile.txt`. Det kan hende at du ikke ser filendelsen (`.txt`) siden Windows på noen datamaskiner skuler filendelser.  
   Om du høyre-klikker på filen, og velger `Egenskaper` (`Properties`) vil du få opp et vindu med hele filnavnet (inklusive `.txt`).
1. Dobbel-klikk på filen for å åpne den og vise innholdet.  
   ![Filinnholdet av testfile.txt i Notepad][notepad]

## Ikke overskriv filen helplink

Om du tester ut koden vi har skrevet så langt flere ganger, vil du legge merke til at air:bir skriver over filen hver gang du kobler den til strøm. Dvs. all data som lå i filen blir slettet hver gang du skrur på air:bit måleren. Dette virker ikke helt fornuftig. Det hadde kanskje vært bedre om vi bare la til ny data i en eksisterende fil dersom det allerede ligger noe der fra før.

Om du ser nøyere på `setup`-blokken, ser du problemet. Vi bruker `O_CREAT` modusen for å lage en **ny** fil, der vi bruker `SD.open`-kommandoen. Og dette gjør vi hver gang air:bit kjører `setup`, dvs. hver gang du plugger i strømmen. Det vi trenger her er en betingelse som skjekker om filen allerede finnes. Vi vil ikke bruke `O_CREAT` når det allerede finnes en fil på SD-kortet som har samme navn som vi vil bruke.

Vi kan bruke `SD.exists`-kommandoen for å spørre om en fil med et gitt navn eksisterer på SD-kortet.

La oss altså lage en `if`-`else`-kombinasjon og flytte linja vi har nå inn i `else`-blokken:

``` cpp
  if (SD.exists(filename)) {
    // Do something new here if file already exists...
  } else {
    // Create new file and open for writing
    file = SD.open(filename, O_CREAT | O_WRITE);
    file.println("Dette er den første linjen i filen.");
  }
```

I koden over ser du at `if`-blokken forsatt er tom, så hva må gjøres dersom filen allerede eksisterer? Vi må fortsatt åpne filen og vi har allerede funnet ut at vi **ikke** vil bruke `O_CREAT`. Å bare bruke `O_WRITE` vil virkelig funke, men det er et lite problem med det også: Når en fil åpnes settes skrivemarkøren i filen på begynnelsen av filen. Dette er ikke noe problem i en ny fil, hvor ellers kunne markøren være, side det ikker er noe der fra før. Men å ha markøren i begynnelsen av en fil som allerede inneholder data vil føre til at du overskriver data som er der. Det vi virkelig vil er å åpne filen, og plassere markøren helt i slutten, slik at vi bare legger til ny tekst uten å røre det som allerede står i filen. Heldigvis har vi `O_APPEND` modusen som gjør nettopp dette. *append* er engelsk for: legg til. Med dette ser hele `if`- `else`-blokken slik ut:

``` cpp
  if (SD.exists(filename)) {
    // Open existing file for writing and append
    file = SD.open(filename, O_WRITE | O_APPEND);
    file.println("--------------------");
    file.println("Filen ble åpnet på nytt.");
  } else {
    file = SD.open(filename, O_CREAT | O_WRITE);
    file.println("Dette er den første linjen i filen.");
  }
```

Merk at koden over skriver en linje med 20 bindestreker i begynnelsen når en eksisterende fil åpnes. Dette vil gjøre det enkelt å finne skillet mellom to seksjoner i filen.

## Test 2 helplink

Med de nye endringene vil du kunne slå av og på air:bit flere ganger. Filen på SD-kortet vil bare bli større og større ettersom air:bit skriver mer og mer tekst til filen.

Du kan slette filen fra SD-kortet når du har kortet plugget inn i datamaskinen. Da vil filen ikke lengre eksistere, og air:bit vil starte med en helt ny fil neste gang du setter SD-kortet inn i air:bit igjen.

``` cpp
#include <SD.h>

#define SD_CS_PIN 10

File file;
int counter;

void setup () {
  // Activate CS-Pin control
  pinMode(SD_CS_PIN, OUTPUT);

  // Startup SD-card reader
  SD.begin(SD_CS_PIN);

  // Define filename
  char filename[] = "testfile.txt";

  if (SD.exists(filename)) {
    // Open existing file for writing and append
    file = SD.open(filename, O_WRITE | O_APPEND);
    file.println("--------------------");
    file.println("Filen ble åpnet på nytt.");
  } else {
    file = SD.open(filename, O_CREAT | O_WRITE);
    file.println("Dette er den første linjen i filen.");
  }
  file.flush(); // Force saving data to SD-card

  // Start counter at 0
  counter = 0;
}

void loop() {
  counter += 1;

  file.print("Dette er linje nr.:");
  file.print(" ");
  file.print(counter);
  file.println();

  file.flush();

  delay(1000); // Wait a second.
}
```

**Merk:** SD-kortet må være plugget i air:bit når du skrur på strømmen. Du kan ikke sette i SD-kortet mens air:bit er på og kjører. Koden som starter opp SD-kortleseren og åpner filen ligger i `setup`. Dvs., den kjøres bare én gang når air:bit starter opp.

## Veien videre

Å bare skrive verdien til en teller-variabel er kanskje litt kjedelig og gir ikke så mye mening. Men du har kanskje lagt merke til at det ikke er noe særlig forskjell fra å skrive til en seriell-kobling eller å skrive til en fil på et SD-kort. Om du bare ser på koden i `loop`, og sammenligner med [teller-eksemplet som bruker seriell-kommunikasjon][counting-example], så ser du lite forskjell.

Det er mye du kan prøve å ekperimentere med nå. Du kan ta koden du skrev her, og kopiere alle `print`- og `println`-kommandoene og legge til at den også skriver samme tekst til den serielle tilkoblingen til PC-en i tillegg til at den skriver til SD-kortet. Du trenger da egentlig bare å bytte ut `file` med `Serial` og du må huske å kjøre `Serial.begin` i `setup`.

Så kan du ta koden du skrev her i `setup` for å åpne filen, og bruke den for å endre de tidligere eksemplene til å skrive til filen på SD-kortet i stedet for eller i tillegg til å skrive til den serielle tilkoblingen med datamaskinen.

## Gå videre

&uarr; [Gå til **innholdsfortegnelsen**][home]  
&larr; [Gå tilbake forrige neste steg: **GPS-antenna**][gps]  
&darr; [Gå til hovedsider for **Guider**][guides]

[home]: airbit-Programmering
[gps]: Programmering-med-GPS-antenna
[guides]: airbit-Guider

[counting-example]: Variabler-og-telling-i-Arduino
[pinout]: airbit-Pinout
[notepad]: testfile-notepad.png
