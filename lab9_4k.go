package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

var token string

// Функция для авторизации пользователя
func login(username, password string) error {
	user := User{Username: username, Password: password}
	body, err := json.Marshal(user)
	if err != nil {
		return err
	}

	resp, err := http.Post("http://localhost:8080/login", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("Ошибка авторизации: %s", bodyBytes)
	}

	var tokenResponse TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResponse); err != nil {
		return err
	}

	token = tokenResponse.Token
	fmt.Println("Успешная авторизация! Токен:", token)
	return nil
}

// Функция для добавления пользователя
func addUser(username, password string) error {
	user := User{Username: username, Password: password}
	body, err := json.Marshal(user)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", "http://localhost:8080/adduser", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("Ошибка при добавлении пользователя: %s", bodyBytes)
	}

	var response struct {
		User  User   `json:"user"`
		Token string `json:"token"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return err
	}

	token = response.Token // Сохраняем токен для нового пользователя
	fmt.Println("Пользователь добавлен:", response.User)
	fmt.Println("Токен:", token) // Выводим токен для подтверждения
	return nil
}

// Функция для получения всех пользователей
func getAllUsers() error {
	req, err := http.NewRequest("GET", "http://localhost:8080/users", nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("Ошибка при получении пользователей: %s", bodyBytes)
	}

	var users []User
	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		return err
	}

	fmt.Println("Список пользователей:")
	for _, user := range users {
		fmt.Printf("- ID: %d, Username: %s\n", user.ID, user.Username)
	}
	return nil
}

// Основная функция клиента с пользовательским интерфейсом в консоли
func main() {
	var choice int

	for {
		fmt.Println("\nВыберите действие:")
		fmt.Println("1. Войти")
		fmt.Println("2. Добавить пользователя")
		fmt.Println("3. Получить всех пользователей")
		fmt.Println("4. Выход")

		_, _ = fmt.Scanf("%d\n", &choice)

		switch choice {
		case 1:
			var username, password string
			fmt.Print("Введите имя пользователя: ")
			_, _ = fmt.Scanf("%s\n", &username)
			fmt.Print("Введите пароль: ")
			_, _ = fmt.Scanf("%s\n", &password)

			if err := login(username, password); err != nil {
				fmt.Println(err)
			}
		case 2:
			if token == "" {
				fmt.Println("Сначала выполните вход.")
				continue
			}
			var username, password string
			fmt.Print("Введите имя нового пользователя: ")
			_, _ = fmt.Scanf("%s\n", &username)
			fmt.Print("Введите пароль нового пользователя: ")
			_, _ = fmt.Scanf("%s\n", &password)

			if err := addUser(username, password); err != nil {
				fmt.Println(err)
			}
		case 3:
			if token == "" {
				fmt.Println("Сначала выполните вход.")
				continue
			}

			if err := getAllUsers(); err != nil {
				fmt.Println(err)
			}
		case 4:
			os.Exit(0)
		default:
			fmt.Println("Неверный выбор. Пожалуйста, попробуйте снова.")
		}
	}
}
