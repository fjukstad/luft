# Dataformat

Husk at det er viktig at fila dere skriver er en komma-separert fil. Dette betyr
at den skal starte med en linje som ser slik ut:

``` csv
Time, Latitude, Longitude, PM10, PM25, Humidity, Temperature
```

og alle de påfølgende linjene i fila skal følge samme mønster, f.eks

``` csv
2000-00-00T00:00:00.000Z,0.000000,0.000000,25.50,4.70,48.60,18.70
2017-12-13T13:34:35.000Z,69.680770,18.974775,2.50,1.40,23.60,20.20
2017-12-13T13:34:38.000Z,69.680946,18.974555,2.70,1.50,23.30,20.10
2017-12-13T13:34:41.000Z,69.681022,18.974626,2.80,1.60,22.90,19.90
2017-12-13T13:34:41.000Z,69.681022,18.974626,2.80,1.70,23.00,19.80
2017-12-13T13:34:46.000Z,69.681022,18.974748,2.80,1.80,23.20,19.90
2017-12-13T13:34:49.000Z,69.681007,18.974828,2.90,1.80,23.20,19.80
2017-12-13T13:34:52.000Z,69.680992,18.974906,2.90,1.90,23.00,19.70
2017-12-13T13:34:55.000Z,69.680969,18.974990,2.90,1.90,22.70,19.50
2017-12-13T13:34:58.000Z,69.680946,18.975065,2.90,2.00,0.00,0.00
2017-12-13T13:34:58.000Z,69.680946,18.975065,3.00,2.00,22.80,19.40
2017-12-13T13:35:03.000Z,69.680915,18.975215,3.10,2.00,22.80,19.30
2017-12-13T13:35:03.000Z,69.680915,18.975215,3.00,2.10,22.60,19.20
2017-12-13T13:35:08.000Z,69.680892,18.975383,2.80,2.10,22.30,19.10
2017-12-13T13:35:08.000Z,69.680892,18.975383,2.70,2.10,22.20,19.00
2017-12-13T13:35:13.000Z,69.680892,18.975561,2.70,2.10,22.00,18.80
2017-12-13T13:35:16.000Z,69.680885,18.975656,2.70,2.10,21.70,18.70
2017-12-13T13:35:19.000Z,69.680885,18.975719,2.30,2.20,21.60,18.60
2017-12-13T13:35:19.000Z,69.680885,18.975719,2.20,2.10,21.40,18.50
2017-12-13T13:35:24.000Z,69.680892,18.975727,2.20,2.10,21.50,18.50
2017-12-13T13:35:27.000Z,69.680877,18.975725,2.10,2.00,21.60,18.40
2017-12-13T13:35:30.000Z,69.680862,18.975769,2.10,2.00,21.70,18.40
2017-12-13T13:35:33.000Z,69.680847,18.975776,2.10,2.00,21.80,18.30
```

## Desimaltall

Som du ser i eksemplet over, brukes det komma `,` for å separere kolonner, og punktum `.` for å skille
mellom siffrer foran og bak desimalen i et desimaltall (f.eks. `69.680770`).

`print`- og `println`-kommandoene i Arduinoen tar imot desimaltall som argument. Du må da angi antall desimaler bak desimal-separatoren som du vil ha printet ut. Om du angir `0` vil desimaltallet blir avrundet til et heltall uten komma. For eksempel:

``` cpp
double tall = 42.427;
Serial.println(tall, 0); // -> 42
Serial.println(tall, 1); // -> 42.4
Serial.println(tall, 2); // -> 42.43
Serial.println(tall, 5); // -> 42.42700
```

## Format på tid og dato

Du ser at tid og dato vises på en litt uvanlig måte her. Dette er dato- og tidsrepresentasjonen som følger det vi kaller standard [ISO 8601][iso8601-wiki]. Det er en av de mest brukte og uproblematiske måtene å representere et tidspunkt på slik at datamaskinen kan fortstå det:

1. Årstall, **4** siffrer
1. En bindestrek, `-`
1. Månedstall, **2** siffrer
1. En bindestrek, `-`
1. Dag, **2** siffrer
1. En stor bokstav T, `T`
1. Timer i 24-timers klokke, **2** siffrer (dvs. `04` og `17`, men ikke `5 p.m.`)
1. Et kolon, `:`
1. Minutter, **2** siffrer
1. Et kolon, `:`
1. Sekunder, **2** siffrer. Om du har desimal-sekunder (altså sekunder angitt som kommatall), bruker du punktum, `.` for å skille mellom siffrer før og etter desimalen.
1. Helt til slutt kommer tidssonen. GPS bruker det vi kaller [UTC][utc-wiki] som ligger **én time bak norsk tid** om vinteren, og **to timer bak** om sommeren. *I militæret kaller man UTC-tid også for Z- eller Zulu-tid.* Bruk en stor bokstav Z for å angi UTC-tidssonen, `Z`

Se før deg at du har følgende variabler i koden din for å lagre tiden i din air:bit sketch:

``` cpp
int year, month, day, hour, minute;
double second;
```

Da kan du bruke følgende kode for å skrive ut tid og dato i korrekt format:

``` cpp
// YEAR
if (year < 10) {
  Serial.print("000");
} else if (year < 100) {
  Serial.print("00");
} else if (year < 1000) {
  Serial.print("0");
}
Serial.print(year);
// SEPARATOR
Serial.print("-");
// MONTH
if (month < 10) {
  Serial.print("0");
}
Serial.print(month);
// SEPARATOR
Serial.print("-");
// DAY
if (day < 10) {
  Serial.print("0");
}
Serial.print(day);
// SEPARATOR
Serial.print("T");
// HOUR
if (hour < 10) {
  Serial.print("0");
}
Serial.print(hour);
// SEPARATOR
Serial.print(":");
// MINUTE
if (minute < 10) {
  Serial.print("0");
}
Serial.print(minute);
// SEPARATOR
Serial.print(":");
// SECOND
if (second < 10) {
  Serial.print("0");
}
Serial.print(second, 3); // three decimals of precision -> 1ms precision
// TIMEZONE
Serial.print("Z");
```

*Merk at koden over bruker printing til Seriell, men det vil være veldig likt til hvordan det gjøres for en fil, kommandoene er de samme.*

[utc-wiki]: https://no.wikipedia.org/wiki/UTC
[iso8601-wiki]: https://no.wikipedia.org/wiki/ISO_8601#Kombinerte_presentasjoner
