<!-- --- title: Guide: Lodde header shieldet -->

I stedet for å få ett virvarr av ledninger, vil vi gjøre livet enklere ved å
bruke et header shield mellom Arduino'en og de forskjellige komponentene.

Vi skal lodde header shieldet oppå Arduino'en. Så kan vi lodde sensorene til
de forskjellige områdene på shieldet. Inni plastplaten som utgjør shieldet
finnes det en rekke koblinger som vil koble sensorene til de rette pinnene på
Arduino'en.

## Dette trenger du

Du trenger:

* Arduino Uno
* Header shieldet
* ZipLock-posen med smådeler
* Avbitertang

## Sette pinner i Arduinoen

Med Arduino'em følger det en (eller to) lang rad med *male* header pinner. Det
lange endestykke til pinnene skal inn pluggene til Arduinoen. Bruk en
avbitertang for å klippe av passende lengde med pinner for hver seksjon med
plugger på Arduinoen, slik som vist i bildet under.

![Avbitertang kutter header-pinner som plugges i Arduinoen][header-pins-cut-arduino]

Plugg i header-pinner for alle fire seksjonene på Arduinoen. Pass på at du
fyller alle pluggene. Du trenger ikke å tenke på de seks pinnene som stikker
opp fra Arduino'en på midten av høyrekanten.

Til slutt vil Arduino'en din se ut som på bildet under.

![Header-pinner plassert i alle pluggene på Arduinoen][header-pins-arduino]

## Plassér shieldet

Legg shieldet oppå Arduinoen med skriften opp og plassér det slikt at de korte
endestykkene av header pinnene stikker gjennom de ytre hullene på shieldet på
begge sider. Shieldet passer bare én vei, så snu det det rundt dersom det ikke
passer.

![Header pinnene stikker opp gjennom shieldet][header-pins-shield]

## Lodde fast pinnene

Varm opp din loddebolt til omtrent 370&deg;C. Avhengig av hvilken type loddebolt
og loddetinn du bruker kan den anbefalte temperaturen variere. Gjerne
dobbeltskjekk hvilken temperatur du burde bruke.

Lodding kan føre til dannelse av gasser som kan være heleskadelige i store
mengder! Bruk alltids en avtrekksvifte der du lodder, slik at du ikke puster
inn utslippet av loddingen. Før du begynner bør du også forsikre deg at
avtrekket er skrudd på.

Ta opp den varme loddebolten og før den fra siden mot pinnen som stikker opp av
hullet. Tuppen av bolten skal varme opp pinnen og metallet på shieldet rundt
hulle som pinnen stikker ut på.  
Før så till loddetinn fra den andre siden. Når tinnet kommer i kontakt med den
varme bolten, pinnen og metallringen på shieldet, vil det smelte og legge seg
rundt pinnen.

![Loddetinn føres til fra venstre mot loddebolten fra høyre][soldering-begin]

## Gå videre

&uarr; [Gå til **innholdsfortegnelsen**][home]  
&larr; [Gå tilbake til forrige steg: **Bygge treboksen**][casing]  
&rarr; [Gå til neste steg: **Lodde sensorene**][sensors]  

[home]: guides-build-home
[casing]: guides-build-casing
[sensors]: guides-build-sensors

[header-pins-cut-arduino]: 20171019_113609.jpg
[header-pins-arduino]: 20171019_113707.jpg
[header-pins-shield]: 20171019_113916.jpg
[soldering-begin]: 20171019_114336.jpg
