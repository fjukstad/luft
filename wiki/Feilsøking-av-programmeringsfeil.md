Den mest vanlige opplevelsen mens man skriver kode er desverre at ting ikke funker fordi man har gjort en feil ett eller annet sted. Dette er omtrent som å skrive en stil i norsk, du kommer alltids til å ha en eller annen skrivefeil her eller der. Akkurat som i norsk, er det viktig å huske på en ting: dette er helt vanlig! Selv om du har norsk som morsmål så vil du fortsatt av og til stave noe feil. Selv om du har programmert i flere tiår, vi du forsatt gjøre små feil. Slik er det bare.

## Innhold

* [Feilmarkeringer i Arduino IDE](#feilmarkeringer-i-arduino-ide)
* [En feil om gangen](#en-feil-om-gangen)
* [Vanlige feil](#vanlige-feil)
  * [For hyppig klikk på Upload-knappen](#for-hyppig-klikk-på-upload-knappen)
  * [Bruk av annen `COM`-port](#bruk-av-annen-com-port)
  * [Manglende Semikolon](#manglende-semikolon)
  * [Manglende lukkende parentes](#manglende-lukkende-parentes)
  * [Feilstavelser i navn](#feilstavelser-i-navn)
  * [Bruk av variabler utenfor scope](#bruk-av-variabler-utenfor-scope)

## Feilmarkeringer i Arduino IDE

Akkurat som læreren din i språkfag er også datamaskinen veldig nøye med hvordan du skriver koden din. Slike feil er forholdsvis enkle feil, siden datamaskinen oppdager feilstavelser når du laster opp programmet ditt til Arduinoen. Så selv om dette kanskje er plagsomt, så vil datamaskinen i det minste ikke få Arduinoen til gjøre noe helt fullstendig galt. I slike tilfeller vil Arduino IDE også vise hvor feilen ligger og si hva som er feil.

Ta for eksempel denne koden:

``` cpp
int counter;

void setup() {
  counter = 0
}
```

Ser du feilen? Om du klikker på `Verify`-knappen til venstre over teksten kan du verifisere at koden er rett og vise feil som datamaskinen kan oppdage.

![Verifiser kode][missing-semicolon]

I bildet over ser du også at Arduino IDE fremhever linjen der feilen oppsto med rød bakgrunn. Feilen oppsto ifølge Arduino IDE i linje 5.

I feilmeldingen under teksteditoren skriver Arduino IDE dessuten også ut en forklaring av feilen:

> Arduino: 1.8.5 (Windows 10), Board: "Arduino/Genuino Uno"
> 
> C:\Users\Fredrik\Documents\Arduino\sketch_nov14a\sketch_nov14a.ino: In
> function 'void setup()':
> 
> sketch_nov14a:5: error: expected ';' before '}' token
> 
>  }
> 
>  ^
> 
> exit status 1
> expected ';' before '}' token

Første linjen i feilmeldingen viser versjonen til Arduinoe IDE og hvilken type Arduino vi programmerer med.

Så kommer filnavnet og i hvilken del av programmet feilen oppsto. I dette tilfellet oppsto feilen i filen `C:\Users\Fredrik\Documents\Arduino\sketch_nov14a\sketch_nov14a.ino` og der i funksjonen `void setup()`, altså i `setup`-funksjonen vår.

Så kommer en liste med alle feil i denne funksjonen. I dette tilfellet er det bare én feil. Først kommer navnet av filen, så et kolon, og så linjenummeret (*5*). Etter dette kommer selve feilmeldingen:

> error: expected ';' before '}' token

Arduino IDE klager på at det mangler ett semikolon (`;`) ett eller annet sted **før** den lukkende klammerparantesen i linje 5. Og det er jo helt rett, fordi vi har glemt å avslutte instruksjonen i linje 4 med et semikolon.

``` cpp
int counter;

void setup() {
  counter = 0;
}
```

**Så husk**: Ofte ligger feilen ikke nøyaktig der hvor Arduino IDE sier at den oppsto, men gjerne i linjen over. Dette kommer av at datamaskinen prøver så godt den kan og skjønne hva du vil gjøre. Derfor oppstår feilen først der hvor datamaskinen må gi opp og si at det er noe feil.

## En feil om gangen

Nå du skriver et større program med mye kode kan det fort skje at du får mange feil, og noen av disse kan også påvirke hverandre. Da blir feilsøking gjerne litt vanskelig og feilmeldingene blir mer og mer kryptisk. En tommelfingerregel er å verifisere koden så ofte som mulig etter hvert småsteg du tar i programmeringen.

Når du finner feil koden din, lønner det seg å alltids rette opp en og en feil og verifisere koden mellom hver rettelse. Feil lager ofte følgefeil, som da plutselig forsvinner. *Se for deg et stort regnestykke. Om du har ett tall feil helt i begynnelsen blir alt etter det feil, men om du bare retter opp feilen helt i begynnelsen, så blir gjerne alle følgende regnestykkene rett igjen.*

## Vanlige feil

Akkurat som små stavefeil i norsk, gjør vi også veldig vanlige feil som går igjen hele tiden når vi skriver kode. Her er noen av de mest vanlige feilene:

### For hyppig klikk på Upload-knappen

Arduinoen er litt sårbar akkurat mens du laster opp kode til den ved å trykke på `Upload`-knappen i Arduino IDE. Når du trykker på `Upload`-knappen burde du helst vente til Arduino IDE enten har fullført å laste opp koden din, eller at den gir deg en feilmelding.

Om du trykker på `Upload`-knappen mens det allerede er en opplastning underveis, kan det oppstå veldig rare feilmeldinger som kommer av at bare deler av et faktisk Arduino program havner på Arduinoen. Detter er ikke farlig, du må bare prøve igjen og neste gang vente til opplastingen er ferdig.

### Bruk av annen `COM`-Port

På siden [Konfigurasjon av Arduino IDE][config-ide] vises det hvordan du setter opp Arduino IDE til å bruke rett `COM`-port for å snakke med Arduinoen over USB-ledningen. Om du ved en senere anledning plugger inn Arduinoen i en annen USB-port på datamaskinen din, eller plugger inn en annen Arduino, må du huske å velge rett `COM`-port på nytt, slik det står beskrevet i guiden.

### Manglende Semikolon

Instruksjoner i C++ må avsluttes med semikolon. Om du glemmer semikolon vil det oppstå en feil i linjen **etter** der du glemte semikolon. Feilmeldingen ser da slik ut:

> expected ';' before &lt;tegn&gt; token

*&lt;tegn&gt; kan være hva som helst som kommer etter linjen der du glemte semikolon.

I noen tilfeller vil feilmeldingen si at den forventet noe annet enn ';' som her:

``` cpp
void setup() {
  int x;
  int c
  x = 0;
  c = 1;
}
```

Feilmelding:

> sketch.ino:4: error: expected initializer before 'x'

Igjen her mangler det noe (et semikolon) **før** `x` i linje 4.

### Manglende lukkende parentes

Parenteser kommer alltid i par. For hver parentes du åpner må du også lukke den igjen. Dette gjelder for vanlige parenteser (`(`, `)`), krøllparenteser (`{`, `}`) og firkant-parenteser (`[`, `]`).

For manglende lukkende krøllparanteser `}` vil du se mange rare feilmeldinger og helt til slutt denne:

> error: expected '}' at end of input

For manglende lukkende parenteser (`)`) vil du se mange rare feilmeldinger og blant dem denne. Det samme gjelder for manglende lukkende firkant-parenteser (`]`):

> error: expected ')' before ...

### Feilstavelser i navn

Datamaskinen er **veldig** nøye når det gjelder navn på funksjoner og variabler. Husk at stor- og små-bokstaver er forskjellige for datamaskinen:

``` cpp
void setup() {
  int intger;
  integer = 42;
}
```

Feilmeldingen ser da slik ut:

> sketch.ino:3: error: 'integer' was not declared in this scope

### Bruk av variabler utenfor scope

Når du skriver et sett med krøllparenteser (`{`, `}`) så introduserer du det vi kaller et *scope* i koden din. Dette skjer altså for hver funksjon, if-setning, eller løkke du skriver. Det som er viktig å huske på med et scope er at variabler du deklarerer inni dette scopet kun *lever* til koden møter den matchende lukkende krøllparentesen. La oss se på dette:

``` cpp
void loop() {
  do {
    int error = sensor.read();
  } while (error < 0);
}
```

I eksemplet over ser du en `do`-`while`-løkke. Inni denne deklarer vi variabelen `error` og setter dens verdi lik resultatet av en avlesning av sensor. **Men**: vi skjekker verdien til denne variabelen i betingelsen til `do`-`while`-løkken (i `(error < 0)`). Med den lukkende krøllparentesen foran `while`-kodeordet har variabelen `error` sluttet å eksistere, så `Arduino IDE` kommer til å klage med følgende feilmelding:

> sketch.ino:4: error: 'error' was not declared in this scope

For å løse dette problemet må du deklarere variabelen i scopet utenfor løkken, som vist under. Merk at vi forsatt setter variabelens verdi *inni* løkken.

``` cpp
void loop() {
  int error;
  do {
    error = sensor.read();
  } while (error < 0);
}
```

Merk også at indre scopes har kjennskap til alle overordnete scopes. Dvs. `error` kjennes til både i scopet til `loop` funksjonen og inni scopet til `do`-`while`-løkken, siden den ligger inni scopet til `loop` funksjonen. Derimot så kan yttre scopes ikke ha kjennskap til ting som ligger i et indre scope, dvs. variabler inni `do`-`while`-løkken kommer ikke til å eksistere utenfor denne løkken.

-----

## Gå tilbake

&uarr; [Gå til: **Guider**][guides]

[guides]: airbit-Guider
[config-ide]: Konfigurasjon-av-Arduino-IDE

[missing-semicolon]: Arduino-IDE-Missing-Semicolon.png
