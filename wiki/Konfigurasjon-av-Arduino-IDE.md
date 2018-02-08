Nå må vi sette opp `Arduino IDE` til å snakke med mikrokontrolleren vår. Vi vil også sette noen innstillinger som vil gjøre det enkere å programmere og finne mulige feil i koden senere.

## Valg av Port for din Arduino helplink

Nå er det på tide å bruke USB-ledningen til Arduino'en og plugge den inn i PCen på en helt vanlig USB-port. Om du har skrudd på lyden på maskinen din, vil du høre den vanlige lyden PCer gir fra seg når du plugger inn nye enheter.

Når du plugger inn en Arduino vil Operativsystemet sette opp noe som kalles en seriell port, også kallt `COM`. På Windows er disse portene nummerert med et tall og tallet blir større for hver ny COM-enhet du kobler til maskinen din.

**Når du nettopp har plugget inn Arduino kontrolleren din vil den høyeste tilgjengelige porten være porten som er tildelt til din Arduino!**

I `Arduino IDE` i menylinjen øvert velg `Tools` (eller `Verktøy`) så velg `Port` og velg den COM-porten med det høyeste tallet.

![Velg Arduino COM Port i Arduino IDE][select-arduino-ide-port]

*Merk at tallet du ser i listen kan være forskjellig fra det du ser i bildet over*

I den samme menyen, i punket over Porten, står det også hvilken type Arduino systemet er valg å programmere imot (det finner mange forskjellige Arduinoer). Vanligvis vil `Arduino/Genuino Uno` være valgt som er den rette. Skjekk at den rette enhetstypen er valgt. Du vil også se enhetstypen stående nederst i høyre hjørnet av vinduet.

På Windows maskiner vil operativsystemet huske din Arduino. Så lenge du bruker samme Arduino på din PC, vil COM-porten ikke endre seg om du plugger inn Arduinoen på nytt. Merk at port vil forsvinne som valgmulighet fra menyen dersom du kobler fra Arduinoen fra datamaskinen.

## Innstillinger helplink

Vi kommer til å bruke mye tid med leke oss med Arduino-kode i `Arduino IDE`. For å være mest produktiv, er det lurt å sette opp noen innstillinger for programmet nå, slik at ting blir litt mere oversiktlig og enklere å håndtere.

I menylinjen øverst, klikk på `File` (eller `Fil`) og velg så `Preferences` (`Preferanser`).  
![Menypunkt File -> Preferences][menu-file-preferences]

Når du holder på med programmering anbefales det på det sterkeste å skru på linjenummering i editoren. Det vil gjøre det enklere å lese feilmeldinger senere siden de alle kommer med et linjenummer der feilen oppsto.  
![Innstillingen for linjenummerering][preferences-linenumbers]

I det samme vinduet kan du også endre skriftsørrelsen på teksten i editoren (`Editor font size`), skru på flere feil- og varselmeldinger fra programmet som tolker koden din (`Compiler warnings`) og du kan bytte språk på `Arduino IDE` (`Editor Language`).

Siden alle guidene, og bilder i disse guidene bruker Engelsk språk i Arduino IDE anbefales det at du velger Engelsk som språk. Dette vil gjøre det enklere å finne rette menypunkt. I tillegg vil feilmeldinger være enklere å 'google' når dem er på engelsk, siden det er så få som snakker norsk og enda ferre som programmerer Arduinoer.

Du kan senere alltids gå tilbake til innstillingene senere og endre på dem om du ønsker. Ignorér innstillinger som du ikke skjønner hva dem gjør, i disse tilfellene er det best å la dem være på standard-innstillingene.

## Gå videre

&uarr; [Gå til **innholdsfortegnelsen**][setup-home]  
&larr; [Gå tilbake til forrige steg: **Start Arduino IDE**][start-arduino-ide]  
&rarr; [Gå til neste steg: **Laste opp tom sketch**][upload-empty-sketch]  

[setup-home]: Guide-Oppsett-for-programmering
[start-arduino-ide]: Start-Arduino-IDE
[upload-empty-sketch]: Laste-opp-tom-sketch-til-Arduinoen

[select-arduino-ide-port]: Arduino-IDE-Select-Port-Screen.png
[menu-file-preferences]: Arduino-IDE-Menu-File-Preferences.png
[preferences-linenumbers]: Arduino-IDE-Preferences-DisplayLineNumbers.png
