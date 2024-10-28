TR
# LogParserCLI
LogParserCLI, stealer malware tarafından oluşturulan log dosyalarını otomatik olarak analiz eden bir komut satırı aracıdır. Birden fazla log formatını destekleyerek, kullanıcılara farklı kaynaklardan veri çıkarma ve analiz etme imkanı sağlar. Bu araç, log analiz sürecini kolaylaştırarak güvenlik uzmanlarının tehditleri tanımlamalarına ve olay müdahale verimliliğini artırmalarına yardımcı olur.

- Kullanım:
  * Bu proje, çeşitli dosya formatlarındaki (JSON, XML, CSV, TXT) verileri işleyerek, belirli anahtar kelimelere göre arama yapabilen bir log 
  analiz aracıdır. Dosya içinde belirli bir anahtar kelime aramak ve anahtar kelime içeren JSON nesnelerini veya diğer veri parçalarını almak 
  için şu adımları izleyin:

- 1. Kurulum:
  * Bu proje Go programlama dili kullanılarak yazılmıştır. Çalıştırmadan önce, Go kurulu olduğundan emin olun. Gerekli bağımlılıkları kurmak 
  için terminalde proje klasöründe şu komutu çalıştırın:
  **go mod tidy**

- 2. Çalıştırma:
  * Ana çalışma dosyasını (main.go) terminal üzerinden çalıştırarak projeyi başlatın. Örnek komut:
  **go run main.go -file <dosya_yolu> -search <anahtar_kelime>**
  * Argüman	Açıklama:
  -file	İşlenmesini istediğiniz dosyanın yolu
  -search	Aranacak anahtar kelime veya filtreleme kriteri

- 3. Örnek:
  * Örneğin, data.json dosyasında "warning" anahtar kelimesini aramak için aşağıdaki komut kullanılabilir:
  **go run main.go -file test_files/data.json -search warning**
  * Bu komut, data.json dosyasındaki tüm "warning" içeren JSON nesnelerini terminalde gösterecektir.

-------------------------------------------------------------------------------------------------------------------------------------------------
ENG
# LogParserCLI
LogParserCLI is a command-line tool for automatically parsing logs generated by stealer malware. It supports multiple log formats, allowing users to extract and analyze data from different sources. This tool streamlines the log parsing process, assisting security professionals in identifying threats and enhancing incident response efficiency.

- Usage:
  * This project is a log analysis tool that processes various file formats (JSON, XML, CSV, TXT) and allows users to search for specific 
  keywords within the files. It retrieves JSON objects or other relevant data containing the specified keyword. Follow these steps to use the 
  tool:

- 1. Installation:
  * This project is written in Go. Ensure Go is installed on your system. To install dependencies, run the following command in the project 
  directory:
  **go mod tidy**

- 2. Running the Program:
  * Run the main file (main.go) in the terminal to start the project. Example command:
  **go run main.go -file <dosya_yolu> -search <anahtar_kelime>**
  * Argument	Description:
  -file	Path to the file you want to process
  -search	The keyword or filter criteria to search for

- 3. Example:
  * To search for the keyword "warning" within the data.json file, use the following command:
  **go run main.go -file test_files/data.json -search warning**
  * This command will display all JSON objects in data.json that contain "warning" in the terminal.






