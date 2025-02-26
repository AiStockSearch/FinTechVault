#!/bin/bash

# Проверка наличия Node.js и Go
if ! command -v node &> /dev/null; then
    echo "Node.js не установлен. Пожалуйста, установите Node.js и повторите попытку."
    exit 1
fi

if ! command -v go &> /dev/null; then
    echo "Go не установлен. Пожалуйста, установите Go и повторите попытку."
    exit 1
fi

# Переход в директорию frontend
cd frontend || { echo "Директория frontend не найдена."; exit 1; }

# Инициализация npm (если еще не инициализировано)
if [ ! -f "package.json" ]; then
    echo "Инициализация npm..."
    npm init -y
fi

# Установка зависимостей для тестирования фронтенда
echo "Установка Jest, @testing-library/react и Playwright..."
npm install --save-dev jest @testing-library/react @testing-library/jest-dom playwright playwright-chromium

# Создание базового файла конфигурации для Jest
echo "Создание конфигурации Jest..."
cat <<EOL > jest.config.js
module.exports = {
  testEnvironment: 'jsdom',
  setupFilesAfterEnv: ['<rootDir>/setupTests.js'],
};
EOL

# Создание файла setupTests.js для Jest
cat <<EOL > setupTests.js
import '@testing-library/jest-dom/extend-expect';
EOL

# Настройка Playwright
echo "Настройка Playwright..."
npx playwright install chromium

# Добавление скриптов в package.json
echo "Добавление скриптов для тестирования..."
node -e "const pkg = require('./package.json'); pkg.scripts = { ...pkg.scripts, 'test': 'jest', 'test-e2e': 'npx playwright test' }; require('fs').writeFileSync('./package.json', JSON.stringify(pkg, null, 2));"

# Настройка тестов для Go
echo "Установка testify для Go..."
go get github.com/stretchr/testify

# Создание примера unit-теста для Go
echo "Создание примера unit-теста для Go..."
if [ ! -d "../cmd/tests" ]; then
    mkdir ../cmd/tests
fi

cat <<EOL > ../cmd/tests/example_test.go
package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
    result := add(2, 3)
    assert.Equal(t, 5, result, "Сумма должна быть равна 5")
}

func add(a, b int) int {
    return a + b
}
EOL

# Завершение
echo "Интеграция тестовых библиотек завершена!"
echo "Для запуска тестов фронтенда используйте: npm run test"
echo "Для запуска E2E-тестов используйте: npm run test-e2e"
echo "Для запуска тестов Go используйте: go test ./..."