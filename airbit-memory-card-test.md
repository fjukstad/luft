# Example code to write to a file on the memory card
```
#include <SD.h>
#define SD_PIN 10

File file;

int i = 0;

void setup() {
    pinMode(SD_PIN, OUTPUT);
    SD.begin();
    char filename[] = "file.txt";

    // If file exsist on the memory card, append to it, if it doesn't exist 
    // create it and write one line to the file before continuing. 
    
    if(SD.exists(filename)){
        file.close(); 
        file = SD.open(filename, O_WRITE | O_APPEND); 
    } else {
        file = SD.open(filename, O_WRITE | O_CREAT); 
        file.print("Dette er første linje. Den vil aldri bli overskrevet.\n"); 
        file.flush();
    } 
}

void loop() {
    
    file.print("Linje ");
    file.print(i);
    file.print("\n"); 
    file.flush();

    i = i + 1;

    delay(5000); 

}
```
