<!-- --- title: Guide: Oppsett for programmering -->

## Steg 1: <small>Laste ned Arduino programmvare</small>

For å starte å programmere med Arduino, må man innstallere
[Arduino IDE](https://www.arduino.cc/en/Main/Software) (Integrated Development Environment).
Det er også mulig å registrere seg med en bruker og få tilgang til en editor i
nettleseren.

### Windows

Klikk på `Windows Installer` (den øverste linken) for en helt vanlig
installasjon på en Windows PC.

Før nedlastningen starter, kan du velge å donere penger for Arduino. Siden
dette er et undervisningsprosjekt som vi ikke skal tjene noe penger på, kan vi
med god samvittighet velge å trykke på `Just Download` uten å gi penger.

Nedlastningen vil starte og nettleseren din vil enten lagre Installasjonsfilen
på det vanlige stedet for nedlastninger på harddisken din, eller be deg om å
velge hvor filen skal lagres.

Etter Installasjonsfilen er lastet ned, må vi kjøre Installasjonsprogrammet.
I dette tilfellet kan du la alle instillingene som programmet foreslår. Trykk
`Next` (eller `Neste`) flere ganger og så `Finish` (eller `Fullfør`). Under
installasjonen vil PCen be deg om å godkjenne installasjon av Drivere som er
utgitt av `Adafruit`, `Arduino LLC` m.m. Godkjenn installasjonen for alle disse,
dette er programmvare som er nødvendig for at PCen din kan 'snakke' med Arduino
kontrolleren. Denne meldingen ser omtrent slik ut (språket kan variere
avhenging av dine språkinnstillinger).

![Arduino driver installasjon][Arduino-Driver-Install-Windows-Security]

## Steg 2: <small>Starte Arduino IDE for første gang</small>

Etter installasjonen vil du kunne finne `Arduino` eller `Arduino IDE` (navnet
kan variere) i start-menyen og også på Skrivebordet. Klikk på denne for å starte
opp Arduino IDE programmet for første gang.

Du vil få opp et vidu som viser følgende kode (eller noe veldig liknende)

``` cpp
void setup() {
  // put your setup code here, to run once:

}

void loop() {
  // put your main code here, to run repeatedly:

}
```

*I noen eldre versjoner av Arduino IDE, kan det hende at du starter med en helt
tom fil uten noen tekst. Du kan da bare kopiere koden over og lime den inn i
editoren din, slik at det er enklere å følge de neste stegene.*

Nå må vi **én** gang sette opp programmet til å snakke med mikrokontrolleren
vår. Klikk [her](Arduino-IDE-Select-Port) for å fortsette med det:  
**Gå til [Velg Arduino IDE Port](Arduino-IDE-Select-Port)**

[Arduino-Driver-Install-Windows-Security]: Arduino-Driver-Install-Windows-Security.PNG
