#include <ESP8266WiFi.h>
#include <ESP8266HTTPClient.h>
#include <ESP8266WebServer.h>
#include <BH1750.h>
#include <Wire.h>
#include "DHT.h"
#include <EEPROM.h>

#define sensor DHT22

// Zugangsdaten zum WLAN:
char wifi_ssid_private[32];
char wifi_password_private[32];
const char* host = "ubuntu";
const int httpPort = 8080;
bool migrateMode = false;

const char* statAddress = "http://ubuntu:8080/stats";
const char* startupAddress = "http://ubuntu:8080/startup";

BH1750 lightMeter;
int sensorDelay = 10000;

int dataPin = 2; // PIN D4 = GPIO 2
DHT dht(dataPin, sensor);

//Webserver für migrate
ESP8266WebServer server(80);
String headers[20];

void setup() {
  Serial.begin(9600);
  Serial.setTimeout(2000);
  //clearEEPROM();
  while(!Serial) { }

  String ssid = readEEPROM(0,32,wifi_ssid_private);
  if (ssid == "") {
    migrateMode = true;
  }
  
  if (migrateMode) {
    migration();
  } else {
    connectWifi();
    setupTempSensor();
    setupLightSensor();
    sendStartupPing();
  }
}

void loop() {
  if (!migrateMode) {
    delay(sensorDelay);
    float lux = lightMeter.readLightLevel();
    float temperature = dht.readTemperature(); // Gets the values of the temperature
    float humidity = dht.readHumidity(); // Gets the values of the humidity
    float heatindex = dht.computeHeatIndex(temperature, humidity, false);
    sendData(lux, temperature, heatindex, humidity);
  } else {
    server.handleClient();
  }
}

void setupTempSensor() {
  digitalWrite(0, LOW); // sets output to gnd
  pinMode(0, OUTPUT); // switches power to DHT on
  delay(1000); // delay necessary after power up for DHT to stabilize
  pinMode(dataPin, INPUT);
  dht.begin();
  Serial.println("dht begin");
}

void setupLightSensor() {
  Wire.begin(D2, D1);
  if (lightMeter.begin(BH1750::CONTINUOUS_HIGH_RES_MODE)) {
    Serial.println(F("BH1750 Advanced begin"));
  } else {
    Serial.println(F("Error initialising BH1750"));
  }
}

//startAdr: offset (bytes), writeString: String to be written to EEPROM
void writeEEPROM(int startAdr, String writeString) {
  EEPROM.begin(512); //Max bytes of eeprom to use
  yield();
  //write to eeprom
  int charLength=writeString.length();
  for (int i = 0; i < charLength; ++i) {
    EEPROM.write(startAdr + i, writeString[i]);
  }
  EEPROM.commit();
  EEPROM.end();
}

String readEEPROM(int startAdr, int maxLength, char* dest) {
  EEPROM.begin(512);
  delay(10);
  for (int i = 0; i < maxLength; ++i) {
    dest[i] = char(EEPROM.read(startAdr + i));
  }
  EEPROM.end();
  return dest;
}

void clearEEPROM() {
  EEPROM.begin(512);
  // write a 0 to all 512 bytes of the EEPROM
  for (int i = 0; i < 512; i++) {
    EEPROM.write(i, 0);
  }
  EEPROM.end();
}

void connectWifi() {
  String ssid = readEEPROM(0,32,wifi_ssid_private);
  String pw = readEEPROM(32,64,wifi_password_private);
  Serial.println("Connecting to ");
  Serial.println(ssid);

  //connect to your local wi-fi network
  WiFi.begin(ssid, pw);

  //check wi-fi is connected to wi-fi network
  while (WiFi.status() != WL_CONNECTED) {
    delay(1000);
    Serial.print(".");
  }
  Serial.println("");
  Serial.println("WiFi connected..!");
  Serial.print("Got IP: ");
  Serial.println(WiFi.localIP());
}

// The network established by softAP will have default IP address of 192.168.4.1.
void startAP() {
  boolean result = WiFi.softAP("MeiSo", "", 1, false, 1);  
  Serial.print("creating AccessPoint was ");
  if(result == false){
    Serial.println("NOT ");
  }
  Serial.print("successfully!");
  Serial.println();
}

void migration() {
  Serial.println("Start migration..!");
  startAP();
  
  server.on("/health", handleHealthCheck);
  server.on("/wifi", handleWifi);
  server.begin(); //Start the server
  Serial.println("Server listening");
}

void sendData (float lux, float temp, float fTemp, float humidity) {
  WiFiClient client;
  if (!client.connect(host, httpPort))
  {
    Serial.println("connection failed");
  }
  String postData = "{\"temp\": "+String(temp)+",\"fTemp\": "+String(fTemp)+",\"humidity\": "+String(humidity)+",\"lux\": "+String(lux)+"}";
  HTTPClient http;
  http.begin(client, statAddress);
  http.addHeader("Content-Type", "application/json");
  Serial.println("send stat data: " + postData);
  auto httpCode = http.POST(postData);
}


void sendStartupPing () {
  WiFiClient client;
  if (!client.connect(host, httpPort))
  {
    Serial.println("connection failed");
  }
  HTTPClient http;
  http.begin(client, startupAddress);
  http.addHeader("Content-Type", "text/plain");
  Serial.println("send startup ping");
  auto httpCode = http.GET();
}

void handleHealthCheck() {
  server.send(200, "text/plain", "OK");
  return;
}

void handleWifi() {
  if (server.hasArg("plain")== false){ //Check if body received
    server.send(200, "text/plain", "Body not received");
    return;
  }
  
  String ssid = split(server.arg("plain"), '\n', 0);
  String pw = split(server.arg("plain"), '\n', 1);
  ssid.toCharArray(wifi_ssid_private, 32);
  pw.toCharArray(wifi_password_private, 32);
  
  writeEEPROM(0,wifi_ssid_private);//32 byte max length
  writeEEPROM(32,wifi_password_private);//32 byte max length
  
  server.send(200, "text/plain", "");
  delay(2000);
  ESP.restart();
}

String split(String s, char parser, int index) {
  String rs="";
  int parserIndex = index;
  int parserCnt=0;
  int rFromIndex=0, rToIndex=-1;
  while (index >= parserCnt) {
    rFromIndex = rToIndex+1;
    rToIndex = s.indexOf(parser,rFromIndex);
    if (index == parserCnt) {
      if (rToIndex == 0 || rToIndex == -1) return "";
      return s.substring(rFromIndex,rToIndex-1);
    } else parserCnt++;
  }
  return rs;
}
