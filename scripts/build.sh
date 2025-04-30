#!/bin/bash

APP_NAME="qwenty"
BASE_DIR="$(pwd)/.."

BUILDS_DIR="$BASE_DIR/builds"
CMD_DIR="$BASE_DIR/cmd"

mkdir -p "$BUILDS_DIR"

dirs=("$CMD_DIR"/*)

if [ ${#dirs[@]} -eq 0 ]; then
  echo "Ошибка состояния директории cmd"
  exit 1
fi

for dir in "${dirs[@]}"; do
  subdir_name=$(basename "$dir")
  main_file="$dir/main.go"

  if [ ! -f "$main_file" ]; then
    echo "Не найден main.go файл в $main_file"
    continue
  fi

  output_name="$APP_NAME-$subdir_name"
  output_path="$BUILDS_DIR/$output_name"

  echo "Билд $subdir_name..."
  GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o "$output_path" "$main_file"

  if [ $? -eq 0 ]; then
    echo "Успешно!"
  else
    echo "Неудачно."
  fi
done

echo "Цикл завершен"
