For å teste at datamaskinen din kan kommunisere med Arduinoen, og at vi faktisk kan sende opp kode til den, skal vi nå prøve å laste opp et "tomt" program til Arduinoen din.

Først, sørg for at koden du ser i din editor er lik den du ser under.

``` cpp
void setup() {
  // put your setup code here, to run once:

}

void loop() {
  // put your main code here, to run repeatedly:

}
```

Ikke vær redd om du ikke skjønner hva som foregår her og hva `void` og alle parantesene betyr. Vi kommer til å forklare det i neste Guide. Nå skal vi bare se om vi klarer å laste koden opp til Arduinoen uten feilmeldinger.

## Laste opp helplink

Vær sikker at Arduinoen er plugget inn. Kill så på `Upload` knappen (se bildet under) for å laste opp.  
![Upload knappen, for å laste opp kode til Arduinoen][upload-empty]

Når du klikker på `Upload`, så vil det dukke opp en *progress-bar* på den nedre delen av vinduet. Du vil se at den først viser teksten `Compiling` (`Kompilerer`) og så går over til `Done uploading.` når den er ferdig.

![Skjermen etter opplastingen er fullført][upload-empty-complete]

## Gå videre

Du har nå fullført oppsett av din datamaskin og `Arduino IDE`. Du er nå klar for å lære deg hvordan vi programmerer en Arduino. Gå tilbake til Guide-menyen og gå videre til neste guide.

&uarr; [Gå til **innholdsfortegnelsen**][setup-home]  
&larr; [Gå tilbake til forrige steg: **Konfigurasjon av Arduino IDE**][config-arduino-ide]  
&darr; [Gå til: **Guider**][guides-home]  

[setup-home]: Guide-Oppsett-for-programmering
[config-arduino-ide]: Konfigurasjon-av-Arduino-IDE
[guides-home]: airbit-Guider
[upload-empty]: Arduino-IDE-Upload-empty.png
[upload-empty-complete]: Arduino-IDE-Upload-empty-complete.png
