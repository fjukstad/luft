Den mest vanlige opplevelsen mens man skriver kode er desverre at ting ikke
funker fordi man har gjort en feil ett eller annet sted. Dette er omtrent som
å skrive en stil i norsk, du kommer alltids til å en eller annen skrivefeil her
eller der.  Akkurat som i norsk, er det viktig å huske på en ting: Dette er
helt vanlig! Selv om du har norsk som morsmål så vil du fortsatt av og stave noe
feil. Selv om du har programmert i flere tiår, vi du forsatt gjøre små feil.
Slik er det bare.

## Innhold

* [Feilmarkeringer i Arduino IDE](#Feilmarkeringer-i-Arduino-IDE)
* [En feil om gangen](#En-feil-om-gangen)
* [Vanlige feil](#Vanlige-feil)
  * [Manglende Semikolon](#Manglende-Semikolon)
  * [Manglende lukkende parentes](#Manglende-lukkende-parentes)
  * [Feilstavelser i navn](#Feilstavelser-i-navn)

## Feilmarkeringer i Arduino IDE

Akkurat som læreren din i språkfag er også datamaskinen veldig nøye med
hvordan du skriver koden din. Slike feil er forholdsvis enkle feil, siden
datamaskinen oppdager feilstavelser når du laster opp programmet ditt til
Arduinoen. Så selv om dette kanskje er plagsomt, så vil datamaskinen i det
minste ikke får Arduinoen til gjøre noe helt fullstendig galt. I slike tilfeller
vil Arduino IDE også vise hvor feilen ligger og si hva som er feil.

Ta for eksempel denne koden:

``` cpp
int counter;

void setup() {
  counter = 0
}
```

Ser du feilen? Om du klikker på `Verify`-knappen til venstre over teksten kan
du verifisere at koden er rett og vise feil som datamaskinen kan oppdage.

![Verifiser kode][missing-semicolon]

I bildet over ser du også at Arduino IDE fremhever linjen der feilen oppsto med
rød bakgrunn. Feilen oppsto ifølge Arduino IDE i linje 5.

I feilmeldingen under teksteditoren skriver Arduino IDE dessuten også ut en
forklaring av feilen:

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

Første linjen i feilmeldingen viser versjonen til Arduinoe IDE og hvilken type
Arduino vi programmerer med.

Så kommer filnavnet og i hvilken del av programmet feilen oppsto. I dette
tilfellet oppsto feilen i filen `C:\Users\Fredrik\Documents\Arduino\sketch_nov14a\sketch_nov14a.ino`
og der i funksjonen `void setup()`, altså i `setup`-funksjonen vår.

Så kommer en liste med alle feil i denne funksjonen. I dette tilfellet er det
bare én feil. Først kommer navnet av filen, så et kolon, og så linjenummeret
(*5*). Etter dette kommer selve feilmeldingen:

> error: expected ';' before '}' token

Arduino IDE klager på at det mangler ett semikolon (`;`) ett eller annet sted
**før** den lukkende klammerparantesen i linje 5. Og det er jo helt rett, fordi
vi har glemt å avslutte instruksjonen i linje 4 med et semikolon.

``` cpp
int counter;

void setup() {
  counter = 0;
}
```

**Så husk**: Ofte ligger feilen ikke nøyaktig der hvor Arduino IDE sier at den
oppsto, men gjerne i linjen over. Dette kommer av at datamaskinen prøver så godt
den kan og skjønne hva du vil gjøre. Derfor oppstår feilen først der hvor
datamaskinen må gi opp og si at det er noe feil.

## En feil om gangen

Nå du skriver et større programm med mye kode, kan det forst skje at du får
mange feil, og noen av disse kan også påvirke hverandre. Da blir feilsøking
gjerne litt vanskelig og feilmeldingene blir mer og mer kryptisk. En
tommelfingerregel er å verifisere koden så ofte som mulig etter hvert småsteg du
tar i programmeringen.

Når du finner feil koden din, lønner det seg å alltids rette opp en og en feil
og verifisere koden mellom hver rettelse. Feil lager ofte følgefeil, som da
plutselig forsvinner. *Se for deg et stort regnestykke. Om du har ett tall feil
helt i begynnelsen blir alt etter det feil, men om du bare retter opp feilen
helt i begynnelsen, så blir gjerne alle følgende regnestykkene rett igjen.*

## Vanlige feil

Akkurat som små stavefeil i norsk, gjør vi også veldig vanlige feil som går
igjen hele tiden når vi skriver kode. Her er noen av de mest vanlige feilene:

### Manglende Semikolon

Instruksjoner i C++ må avsluttes med semikolon. Om du glemmer semikolon vil det
oppstå en feil i linjen **etter** der du glemte semikolon. Feilmeldingen ser da
slik ut:

> expected ';' before &lt;tegn&gt; token

*&lt;tegn&gt; kan være hva som helst som kommer etter linjen der du glemte
semikolon.

I noen tilfeller vil feilmeldingen si at den forventet noe annet enn ';' som
her:

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

Parenteser kommer alltid i par. For hver parentes du åpner må du også lukke den
igjen. Dette gjelder for vanlige parenteser (`(`, `)`), krøllparenteser (`{`, `}`)
og firkant-parenteser (`[`, `]`).

For manglende lukkende krøllparanteser `}` vil du mange rare feilmeldinger og
helt til slutt denne:

> error: expected '}' at end of input

For manglende lukkende parenteser (`)`) vil du mange rare feilmeldinger og blant
dem denne. Det samme gjelder for manglende lukkende firkant-parenteser (`]`):

> error: expected ')' before ...

### Feilstavelser i navn

Datamaskinen er **veldig** nøye når det gjelder navn på funksjoner og variabler.
Husk at stor- og små-bokstaver er forskjellige for datamaskinen:

``` cpp
void setup() {
  int intger;
  integer = 42;
}
```

Feilmeldingen ser da slik ut:

> sketch.ino:3: error: 'integer' was not declared in this scope

-----

## Gå tilbake

&uarr; [Gå til: **Guider**][guides]

[guides]: airbit-Guider

[missing-semicolon]: Arduino-IDE-Missing-Semicolon.png
