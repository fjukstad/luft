Programmering er å skrive oppskrifter med stegvise instruksjoner til en datamaskin for å få den til å utføre spesifikke oppgaver. Alle datamaskiner bruker slike oppskrifter, vi som brukere av datamaskiner kjenner til en type oppskrift som kalles *Applikasjoner* eller *Apper*. *Apper* er gjerne fullverdige programmer som utfører oppgaver som vi brukere kan dra nytte av. Noen eksempler på programmer som kan kalles *Apper* er: *Nettlesere, Spill, Mediaprogrammer, Tekstbehandling osv.* Disse programmene er veldig synlige for brukere fordi det er gjerne disse vi interagerer med når vi bruker en datamaskin.

Det mange ikke er klar over i dag, er at det finnes store mengder andre programmer enn *Applikasjoner*. Mange av disse (men ikke alle) er mye mindre synlig for en bruker, ofte fordi de enten kjører i bakgrunnen eller fordi det er bygd andre programmer på toppen av disse slik at brukere ikke skal trenge å bekymre seg for hvordan de fungerer. Slike typer programmer kan være:

- **System Software**
  - **Operativsystem** - Fungerer som et mellomlag mellom en bruker (eller en applikasjon) og enhetens *Hardware* slik som skjerm, tastatur, harddisker, minne, kamera osv. Noen kjente eksempler på disse er Windows 10, OSX, Ubuntu, Android, iOS etc.
  - **Sikkerhet** - Slik som en brannmur eller antivirus.
- **Verktøy (Utilities)**
  - **Drivere** - Programmer som kontrollerer kommunikasjon og styring av enheter.
  - **Backup** - Programmer som kopierer og sikkerhetslagrer informasjon og filer.
  - **Defrag** - Programmer som flytter på ting som ligger på en harddisk, slik at informasjon ligger lagret mer kompakt. Dette er gjerne ikke et problem lengre med SSDer (Solid State Drives).

- **Interfaces**
  - **GUI (Graphical User Interface)** - Grafiske bruker grensesnitt er den vanligste måten å bruke datamaskiner på. Et vindu i Windows er eksempel på et grafisk brukergrensesnitt.
  - **CLI (Command Line Interface)** - Gjerne en svart skjerm med hvit skrift hvor man skriver kommandoer til datamaskinen. Kommandolinjen er ofte litt vanskeligere å bruke enn et grafisk grensesnitt, men man har ofte mer kontroll over forskjellige programmer.

Man kan lese mer om de forskjellige typene programmer [her](https://en.wikiversity.org/wiki/Types_of_Computer_Software)

## Arduino Programmering

Et Arduino program er nok mest relatert til en *Applikasjon*, men den kjører gjerne ikke sammen med et fullverdig operativsystem. Noen av ulempene med dette er at det ikke er mulig å kjøre mange programmer samtidig på en Arduino, men det er kun et som kan kjøre om gangen, man har ikke like mye muligheter på en Arduino som en vanlig datamaskin i forhold til skjermer, internett ol. Noen av fordelene er at enkle programmer bruker gjerne lite strøm, det er forholdsvis lett å forstå oppbygningen av programmer og man kan koble sensorer direkte inn i en Arduino som gjør den perfekt til å lage små enheter man kan bære med seg.

Arduino benytter seg av små programmer som kalles *Sketches*. Når man programmerer en Arduino så laster man opp *Sketchen* til Arduinoens flash minne. Når man da kobler strøm til arduinoen så vil den laste inn og kjøre *Sketchen* man lastet opp. For å programmere Arduinoen på nytt så laster man bare opp en ny *Sketch* så vil denne overskrive den forrige. Arduino Uno kan bare programmeres med **én** *Sketch* om gangen.
