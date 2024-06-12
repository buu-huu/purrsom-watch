# 🐱 PurrsomWatch: Early Ransomware Detection for Windows

Welcome to **PurrsomWatch**, a Golang-based software designed to detect ransomware activity on your Windows systems by using decoy files and advanced logging techniques. This project aims to provide early detection of ransomware attacks, allowing cybersecurity teams to take immediate action.

### Disclaimer
This repository is work in progress. Development is ongoing but will take some time. Releases will be made and tagged, when a certain set of features is implemented and tested. 

## 📜 Table of Contents
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Roadmap](#roadmap)
- [Contributing](#contributing)
- [License](#license)

## 🌟 Features <a name="features"></a>
- **Decoy File Detection:** Places a decoy file in specified directories to detect ransomware activity.
- **Event Logging:** Logs ransomware detection events into the Windows Event Log with custom logs for easy SIEM integration.
- **Real-Time Monitoring:** Monitors decoy files for read and modification (encryption) activities.
- **Future Enhancements:** Planned entropy-based detection to enhance ransomware detection capabilities.

## 🛠️ Installation
To install PurrsomWatch, follow these steps:

1. Clone the repository:
    ```sh
    git clone https://github.com/buu-huu/purrsom-watch.git
    cd purrsom-watch
    ```

2. Build the project:
    ```sh
    go build -o purrsomwatch.exe
    ```

3. Run the executable:
    ```sh
    ./purrsomwatch.exe <CONFIG_FILE>
    ```

## 🚀 Usage
Edit the configuration file to specify directories to monitor. The provided [config file template](configs/config_template.json) gets updated continuously.

``` json
{
  "purrEngine": {
    "purrInterval": "10",
    "decoyFile": {
      "fileName": "purrguard",
      "fileExtension": "docx",
      "location": {
        "fileDir": "%userdir%/Documents/purr",
        "username": "user"
      }
    }
  },
  "winEventProvider": {
    "eventId": "7700"
  }
}
```

PurrsomWatch will begin monitoring the specified directories and log any ransomware activity detected.

## 🛤️ Roadmap
- **Windows Event Logging:** Add custom windows event logging for SIEM based use cases
- **Entropy-Based Detection:** Implement entropy-based detection for enhanced ransomware identification.
- **Improved Logging:** Add more detailed logging and reporting features.
- **User Interface:** Optional windows systray icon with notification system.

## 🤝 Contributing
Contributions from the community are appreciated! If you have ideas for improvements or want to help with development, please fork the repository and submit a pull request or open an issue.

## 📄 License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

For any questions or support, please open an issue.

---

Stay safe and secure! 🛡️