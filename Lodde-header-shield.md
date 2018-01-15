I stedet for å få ett virvarr av ledninger, vil vi gjøre livet enklere ved å bruke et header shield mellom Arduino'en og de forskjellige komponentene.

Vi skal lodde header shieldet oppå Arduino'en. Så kan vi lodde sensorene til de forskjellige områdene på shieldet. Inni plastplaten som utgjør shieldet finnes det en rekke koblinger som vil koble sensorene til de rette pinnene på Arduino'en.

## Dette trenger du

Du trenger:

* Arduino Uno
* Header shieldet
* ZipLock-posen med smådeler
* Avbitertang

## Sette pinner i Arduinoen helplink

Med Arduino'en følger det en (eller to) lang rad med *male* header pinner. Det lange endestykke til pinnene skal inn pluggene til Arduinoen. Bruk en avbitertang for å klippe av passende lengde med pinner for hver seksjon med plugger på Arduinoen, slik som vist i bildet under.

![Avbitertang kutter header-pinner som plugges i Arduinoen][header-pins-cut-arduino]

Plugg i header-pinner for alle fire seksjonene på Arduinoen. Pass på at du fyller alle pluggene. Du trenger ikke å tenke på de seks pinnene som stikker opp fra Arduino'en på midten av høyrekanten.

Til slutt vil Arduino'en din se ut som på bildet under.

![Header-pinner plassert i alle pluggene på Arduinoen][header-pins-arduino]

## Plassér shieldet helplink

Legg shieldet oppå Arduinoen med skriften opp og plassér det slikt at de korte endestykkene av header pinnene stikker gjennom de ytre hullene på shieldet på begge sider. Shieldet passer bare én vei, så snu det det rundt dersom det ikke passer.

![Header pinnene stikker opp gjennom shieldet][header-pins-shield]

## Lodde fast pinnene helplink

Varm opp din loddebolt til omtrent 370&deg;C. Avhengig av hvilken type loddebolt og loddetinn du bruker kan den anbefalte temperaturen variere. Gjerne dobbeltskjekk hvilken temperatur du burde bruke.

Lodding kan føre til dannelse av gasser som kan være heleskadelige i store mengder! Bruk alltids en avtrekksvifte der du lodder, slik at du ikke puster inn utslippet av loddingen. Før du begynner bør du også forsikre deg at avtrekket er skrudd på.

Ta opp den varme loddebolten og før den fra siden mot pinnen som stikker opp av hullet. Tuppen av bolten skal varme opp pinnen og metallet på shieldet rundt hulle som pinnen stikker ut på.  
Før så til loddetinn fra den andre siden. Når tinnet kommer i kontakt med den varme bolten, pinnen og metallringen på shieldet, vil det smelte og legge seg rundt pinnen.

![Loddetinn føres til fra venstre mot loddebolten fra høyre][soldering-begin]

Vri tinnet og bolten rundt pinnen, slik at tinnet legger seg jevnt runt pinnen fra alle sider og dekker hullet den stikker opp fra. Etterpå vil du ende med en liten haug der hullet var før med resten av pinnen stikkende ut på toppen.

![Ferdig loddet header-pinne][soldering-complete]

Når alle pinnene er loddet fast, slik som vist i bildet over, burde du kunne løfte av header shieldet fra Arduinoen igjen med alle header pinnene hengende fast i shieldet. Gratulerer! Du har nå et header shield som du enkelt kan plugge på hvilken som helst Arduino Uno.

## Lodde fast mostandene helplink

I midten av shieldet finner du to små firkanter under hverandre som er merket `R1 220` og `R2 220`. Her skal vi plassere de to elektriske motstanderene som du finner i ZipLock-posen med smådeler.

Dra av papirstripen som motstandene henger fast i, og bøy til de lange endestykkene til motstandene slik at de passer inn i hullene på shieldet på høyre- og venstresiden av firkantene.

![Plassering av motstandene på shieldet][resistors-placement]

Lodd fast endestykkene til begge motstandene på shieldet, akkurat som du gjorde tidligere med header-pinnene til Arduino'en. Når du har loddet fast motstandene klipper du av de lange endestykkene som stikker ut på undersiden av shieldet med en avbitertang.

## Lodde fast LED lysene helplink

Øverst i høyre hjørnet av shieldet, ved siden av firkanten der det står GPS, skal LED'ene plasseres. 

`D` i LED står for *Diode*. Dette vil si at strømretningen er veldig viktig å passe på, strømmen kan bare flyte én vei gjennom en diode, så du må passe på å montere den rett vei på shieldet.

![Plassering av LED][led-placement]

Se nøye på bildet over. Legg merke til at diodene har en flat kant mot den ene siden. I bildet vises det best på den grønne dioden som har sin flate kant pekende mot venstresiden av bildet. Den flate kanten av dioden skal peke mot den flate siden av tegningen på shieldet (der det står `LED1` og `LED2`).

Monter det blanke LED-lyset der det er merket `LED1 Rød` og det grønne LED-lyset der det er merket `LED2 Blå` på kretskortet. *Det blanke LED-lyset lyser rødt, og tidligere brukte air:bit en blå LED i stedet for grønn.*

Lodd fast diodene slik du gjorde med motstandene. Klipp så av endestykkene med avbitertang.

## Ferdig

Du skal nå har et helt funksjonellt shield klart for bruk. Nå er vi kommet godt i gang med loddingen, så neste steg er å lodde på sensorene.

## Gå videre

&uarr; [Gå til **innholdsfortegnelsen**][home]  
&larr; [Gå tilbake til forrige steg: **Bygge treboksen**][casing]  
&rarr; [Gå til neste steg: **Lodde sensorene**][sensors]  

*Dersom du vil hoppe over montering av sensorer kan du i teorien nå begynne å programmere Arduino'en til å blinke LED lysene. Du vil finne tilsvarende kodeeksempler i høyremenyen.*

[home]: Guide-Bygging-og-Lodding
[casing]: Sette-sammen-treboksen
[sensors]: Lodde-sensorene

[header-pins-cut-arduino]: 20171019_113609.jpg
[header-pins-arduino]: 20171019_113707.jpg
[header-pins-shield]: 20171019_113916.jpg
[soldering-begin]: 20171019_114336.jpg
[soldering-complete]: 20171019_114915.jpg
[resistors-placement]: 20171019_115039.jpg
[led-placement]: 20171019_115344.jpg
