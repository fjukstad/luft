Når vi programmerer, skriver vi *kode* i et programmeringsspråk. *Kode* er egentlig bare ren tekst slik du ville skrive på en notatblokk, eller i matteleksen. Men i stedet for vanlig norsk, skriver vi et *språk* som datamaskinen kan lese og forstå. På mange måter funker dette veldig liknende til hvordan en kalkulator forstår enkle regnestykker som inneholder symboler som f.eks. `+`, `-`, `/` og `*`. Poenget er: Akkurat som du ikke kan skrive vanlig norsk inn i en kalkulator, må vi bruke en korrekt syntaks for det rette språket når vi skriver programkode.

Siden program som `Microsoft Word` gjør mye mer enn bare skrive tekst (du får mange muligheter for å endre skrift, størrelse, fettskrift, kursiv, osv.), må bruke et eget program for tekstredigering som er spesialisert for Arduino-kode. Det er blant annet derfor vi bruker `Arduino IDE`.

Innen `Arduino IDE` kalles filen du skriver kode i, en *Sketch*. La oss se en gang til på koden i en *tom* Sketch:

``` cpp
void setup() {
  // put your setup code here, to run once:

}

void loop() {
  // put your main code here, to run repeatedly:

}
```

I koden over ser vi to *funksjoner* (også kalt *sub-rutiner*). De er kalt `setup` og `loop`. Du kan se navnene etter kodeordet `void`.

### Kodeblokker mellom `{` og `}` helplink

Koden som hører til funksjonen ligger blant et par med `{` og `}` (krøllparanteser). Du skriver dem ved å trykke `AltGr` og `7` eller `0` samtidig på et vanlig nordisk tastatur. I programmering kaller vi alt som står mellom et par med krøllparanteser for en *kodeblokk*.

### Kommentarer `//` helplink

Når vi ser i koden som står mellom krøllparantesene til `setup` og `loop` ser vi at begge blokkene hver inneholder en linje som starter med to skråstreker (`//`).

Uansett hvor i koden, så vil to skråstreker `//` alltids starte en linje-kommentar. En kommentar er villkårlig tekst som ikke tolkes av datamaskinen, og som derfor bare er forklarende for oss mennesker. Det er en god idé å skrive kommentarer i koden din, slik at du etterpå fortsatt kan forstå hva du gjorde.

### Funksjonen `setup` helplink

Funksjonen `setup` er funksjonen vi bruker for å starte opp Arduinoen og sette den opp til å gjøre det vi vil. Dette er et steg som kjøres én gang når vi kobler til strømmen til Arduinoen. Koden vi plasserer her kan sammenlignes med det din datamaskin gjør når den starter opp, før du blir bedt om å skrive inn passordet.

Akkurat nå ser vi at det ikke er noen instruksjoner i blokken som utgjør denne funksjonen. Det eneste som står mellom krøllparantesene er en kommentar som sier

> put your setup code here, to run once:

### Funksjonen `loop` helplink

Etter Arduinoen har startet opp vil den starte å kjøre instruksjonene i `loop`-funksjonen. Disse vil bli kjørt om og om igjen i en evig loop.

Som funksjonen `setup` har også denne kodeblokken ingen intruksjoner i seg, siden den bare inneholder en kommentar.

----

Siden vi ikke har skrevet noen instruksjoner i hverken `setup` og `loop`, gjør programmet i denne Sketchen foreløpig ingenting. Neste steg er å fylle sketchen med instruksjoner som vil få Arduinoen til å gjøre det vi vil.

## Gå videre

&uarr; [Gå til **Introduksjon til Arduino**][intro-prog-home]  
&rarr; [Gå til neste steg: **Hello World**][hello-world]

[intro-prog-home]: Introduksjon-til-Arduino-programmering
[hello-world]: Arduino-varianten-av-Hello-World
