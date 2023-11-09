package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	Red    = "\033[31m"
	Green  = "\033[32m"
	Reset  = "\033[0m"
	Yellow = "\033[33m"
)

func main() {
	fmt.Println("Lütfen giriş türünüzü seçin:")
	fmt.Println("0 - Admin Girişi")
	fmt.Println("1 - Öğrenci Girişi")

	var userType int
	fmt.Print("Seçiminiz: ")
	fmt.Scan(&userType)

	if userType == 0 {
		adminLogin()
	} else if userType == 1 {
		userLogin()
	} else {
		fmt.Println("Geçersiz seçenek. Program sonlandırılıyor.")
	}
}

func adminLogin() {
	adminUsername := "admin"
	adminPassword := "admin"
	var username string
	var password string

	for i := 0; i < 5; i++ {
		fmt.Print("Kullanıcı Adı: ")
		fmt.Scan(&username)
		fmt.Print("Şifre: ")
		fmt.Scan(&password)

		if username == adminUsername && password == adminPassword {
			fmt.Println(Green + "Başarılı giriş!" + Reset)
			adminMenu()
			return
		} else {
			fmt.Println(Yellow+"Hatalı kullanıcı adı veya şifre. Kalan deneme hakkı:"+Reset, 4-i)
			if i < 6 {
				logLogin(username, false)
			}
		}
	}

	logLogin(username, false)
	fmt.Println("5 başarısız deneme. Program Sonlandırılıyor")
}

func userLogin() {
	userUsername := "user"
	userPassword := "user"
	var username string
	var password string

	for i := 0; i < 5; i++ {
		fmt.Print("Kullanıcı Adı: ")
		fmt.Scan(&username)
		fmt.Print("Şifre: ")
		fmt.Scan(&password)

		if username == userUsername && password == userPassword {
			fmt.Println(Green + "Başarılı giriş!")
			logLogin(username, true)
			return
		} else {
			fmt.Println(Yellow+"Hatalı kullanıcı adı veya şifre. Kalan deneme hakkı:"+Reset, 4-i)
			if i < 6 {
				logLogin(username, false)
			}
		}
	}

	logLogin(username, false)
	fmt.Println("5 başarısız deneme. Program Sonlandırılıyor")
}

func logLogin(username string, success bool) {
	currentTime := time.Now()
	logFile, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// Append:  Dosyaya yazma işlemi yaparken veriyi dosyanın en sonuna ekler
	// Create:  Dosya yoksa oluşturur
	// Wronly:  Dosyaya yazma işlemi yapar
	// 0644:    Dosya izinleri
	if err != nil {
		fmt.Println("Log dosyasına erişilemiyor:", err)
		return
	}
	defer logFile.Close()

	logStatus := "Başarılı"
	if !success {
		logStatus = "Başarısız"
	}

	logEntry := fmt.Sprintf("Kullanıcı Adı: %s\nGiriş Tarihi ve Saati: %s\nGiriş Durumu: %s\n\n", username, currentTime.Format("2006-01-02 15:04:05"), logStatus)
	//fmt.Sprintf: Verilen format ile string oluşturur
	_, err = logFile.WriteString(logEntry)
	// WriteString: Dosyaya string yazar
	if err != nil {
		fmt.Println("Log dosyasına kayıt eklenemedi:", err)
	}
}

func adminMenu() {
	for {
		fmt.Println("Lütfen bir işlem seçin:")
		fmt.Println("0 - Logları Görüntüle")
		fmt.Println("1 - Çıkış Yap")

		var choice int
		fmt.Print("Seçenek: ")
		fmt.Scan(&choice)

		if choice == 0 {
			viewLogs()
		} else if choice == 1 {
			fmt.Println("Çıkış yapılıyor.")
			os.Exit(0)
		} else {
			fmt.Println("Geçersiz seçenek. Lütfen tekrar deneyin.")
		}
	}
}

func viewLogs() {
	logFile, err := os.Open("logs.txt")
	if err != nil {
		fmt.Println("Log dosyasına erişilemiyor:", err)
		return
	}
	defer logFile.Close()

	scanner := bufio.NewScanner(logFile)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "Giriş Durumu: Başarılı") {
			// Contains: Verilen string içerisinde arama yapar
			fmt.Print(Green)
		} else if strings.Contains(line, "Giriş Durumu: Başarısız") {
			fmt.Print(Red)
		}

		fmt.Println(line)
		if strings.Contains(line, "Giriş Durumu:") {
			fmt.Println(strings.Repeat("-", 35))
		}

		fmt.Print(Reset)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Log dosyası okunurken bir hata oluştu:", err)
	}
}
