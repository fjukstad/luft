Når man lærer et nytt programmeringsspråk er det vanlig å lage et program som ofte kalles `Hello World`. Det handler om å få datamaskinen, eller i dette tilfellet Arduinoen, til å skrive ut teksten `Hello World`.

I forrige steg så vi at en tom sketch inneholder to funksjoner `setup` og `loop`. Vi må skrive instruksjoner for begge funksjonene for dette eksemplet til å funke.

### Hvordan skrive ut tekst på en Arduino helplink

Når du kobler til en Arduino til datamaskinen kan vi sende data fra Arduinoen til datamaskeinen og ledningen. Dette kalles *Seriell kommunikasjon*.

I `Arduino IDE` bruker vi *Serial Monitor* (*seriell overvåker*) for å lese teksten Arduinoen sender til datamaskinen over ledningen.

![Menyknapp for Serial Monitor][serial-monitor-button]

Mens denne måten av kommunikasjon ikke alltid er særlig nyttig i en virkelig bruk av Arduinoen, er dette et godt redskap for feilsøking mens du utvikler kode.

Om du klikker på knappen nå, vil antageligvis ingenting vises i vinduet som kommer opp. Detter er fordi vi ikke ennå har sagt at Arduinoen skal gjøre noe, og derfor vil den ikke sende noe data over seriell-porten heller.

### `setup` helplink

``` cpp
void setup() {
  // initialize serial communication at 9600 bauds per second:
  Serial.begin(9600);
}
```

I koden over ser vi `setup`-koden for å aktivere seriell kommunikasjon på Arduinoen. Første linje er en kommentar (merk at linjen er farget annerledes) siden den starter med to skråstreker `//`.

Så følger en instruksjon på neste linje: `Serial.begin(9600);`  
`Serial.begin` er kommandoen for å starte seriell kommunikasjon. Så følger et par med runde (eller vanlige) paranteser `(` og `)`. Alt mellom et par med runde paranteser er argumenter for kommandoen. I dette tilfellet har vi et argument, verdien `9600`.

Dersom man leser dokumentasjonen på nettet vil man se at seriell kommunikasjon (eller `UART`) må kalibreres for en hastighet som kommunikasjonen skal foregå på. Mest vanligvis bruker man en hastighet på `9600 bauds`. Det er dette vi bruker her også. Hva, hvordan og hvorfor denne serielle kommunikasjonen foregår er ikke direkte viktig for dette eksemplet. Vend deg til å bruke `Serial.begin` med argument `9600`.

**Merk** at korrekt staving av `Serial.begin` og store og små bokstaver må skrives **nøyaktig** som vist over. Du vil legge merke til at fargen av teksten du skriver i `Arduino IDE` endrer seg når programmet skjønner hva du mener.

**Merk** også det avsluttende semi-kolonet `;`. `;` avslutter en kommando og **må** skrives. Siden datamaskinen har semi-kolonet for å finne slutten av en kommando, kan man godt bruke flere linjer for en kommando. Man kan også skrive flere kommandoer etter hverandre på én linje. *Men selv om dette er mulig, er det skjeldent en go idé å gjøre det.*

### `loop` helplink

``` cpp
void loop() {
  // print out Hello World:
  Serial.println("Hello World!");
  delay(1000); // delay for one second
}
```

Over er koden for `loop` funksjonen. Hvis vi ignorerer kommentarene, så ser vi at vi har to instruksjoner her.

Først bruker vi kommandoen `Serial.println`. `println` er en forkortelse for `print line`, så denne kommandoen vil skrive en linje med tekst over seriell-kommunikasjonen. Teksten som skrives vil vises i `Serial Monitor` i `Arduino IDE` når Arduinoen kjører. Argumentet til `Serial.println` er her verdien `"Hello World!"` (merk anførelsestegnene `"` på begynnelsen og slutten).

Neste kommando er `delay`. `delay` vil får Arduinoen til å stoppe og pause for ett gitt antall millisekunder før den fortsetter å jobbe. Argumentet er verdien `1000`, dvs. 1000 millsekunder, eller 1 sekund. Denne kommandoen vil altså få Arduinoen til å pause i ett sekund etter den har skrevet ut teksten i den forrige kommandoen.

### Upload og kjør helplink

Klikk på `Upload` (`Laste opp`) knappen til venstre over teksteditoren. Vent til `Arduino IDE` er ferdig å bygge koden og sender den opp til Arduinoen. Når den viser at alt gikk bra, klikk på `Serial Monitor` knappen.  
Om du har gjort alt rett vil du se teksten `"Hello World!` bli printet ut om og om igjen med ett sekunds mellomrom.

## Gå videre

&uarr; [Gå til **Introduksjon til Arduino**][intro-prog-home]  
&larr; [Gå tilbake til forrige steg: **Arduino Sketch**][sketch]  
&rarr; [Gå til neste steg: **Variabler og tellere**][variables]  

[intro-prog-home]: Introduksjon-til-Arduino-programmering
[sketch]: Tom-Arduino-Sketch
[variables]: Variabler-og-telling-i-Arduino

[serial-monitor-button]: Arduino-IDE-SerialMonitor-Button.png
