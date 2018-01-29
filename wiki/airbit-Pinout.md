For å vite hvilke sensorer er koblet til hvilke pinner som du må bruke i programmet ditt brukes i vanlig elektronikk et koblingsskjema, eller det vi kaller en `pinout`.

I bildet under vil du se pinout-skjemaet til air:bit.

![air:bit Pinout-Skjema][airbit-pinout-skjema]  
*Klikk [her][airbit-pinout-skjema] for å vise bildet i en egen fane:
**[air:bit Pinout-Skjema][airbit-pinout-skjema]***

## Som tabell

Om du heller vi se komponentene listet opp med sin tilsvarende pinne.

| Komponent | Funskjon | Pinne |
| --------- | -------- | ----- |
| LED 1 (Rød) | Strøm | `A1` eller `15` (begge funker) |
| LED 2 (Blå eller Grønn) | Strøm | `A0` eller `14` (begge funker) |
| Temperatorsensoren | Data | `9` |
| Støvmåleren | Data `TX` | `2` |
| Støvmåleren | Data `RX` | `3` |
| GPS-antenna | Data `TX` | `7` |
| GPS-antenna | Data `RX` | `6` |
| SD-kortleser | `CS` | `10` |
| SD-kortleser | `SCK` | `13` |
| SD-kortleser | `MOSI` | `11` |
| SD-kortleser | `MISO` | `12` |

## Bruk

Guidene for de forskjellige sensorene vil fortelle deg hvordan disse verdiene for pinnene skal brukes. I utgangspunketet kan man plugge inn sensorer på hvilken som helst pinne i Arduinoen, men da må man forklare til Arduinoen fra hvilken pinne han skal lese data for hvilken sensor. Det er dette pinout skjemaet brukes til.

## Gå videre

&uarr; [Gå til **innholdsfortegnelsen**][home]  
&rarr; [Gå til neste steg: **LED-Blinking**][led]  

[airbit-pinout-skjema]: Vaerlogger_v2.1_skjema.png

[home]: airbit-Programmering
[led]: airbit-LED-Blinking
