# go-cmd-client

### Консольный клиент для взаиммодействия с ОС Windows посредством выполнения команд прописанных в JSON скрипте

### Настройка

1. Клонировать репозиторий и открыть его VSCode
2. Установить переменные среды в `.env` файле директории с проектом как в 'env.example'
3. Установить утилиту make любым удобным способом

### Пример использования

Команды а так же условия сформированы в файле task_list.json

```
[
    {
        "name": "New-Item",
        "params": {
            "-Path": "./filename.txt"
        },
        "result": {}
    },

...

    {
        "name": "Test-Path",
        "params": {
            "-Path": "./valera.txt"
        },
        "result": {},
        "positiveOutcome": {
            "name": "Write-Output",
            "params": {
                "-InputObject": "'File is exist!'"
            },
            "result": {}
        },
        "negativeOutcome": {
            "name": "New-Item",
            "params": {
                "-Path": "./valera.txt"
            },
            "result": {}
        }
    }
]
```

Вы можете добавить или сконфигурировать этот файл как угодно сохраняя структуру действия или условия как в существующем task_list.json

Чтобы запустить процесс выполнения необходимо в директории с проектом выполнить команду
```
make start
```

В ходе выполнения вы можете видеть промежуточные результаты в консоли
```
time=11-12-2023 00:30:21 level=info msg=
[POINT]: Runtime [INFO]: Prepare Task List...
time=11-12-2023 00:30:21 level=info msg=
[POINT]: Runtime [INFO]: Tasks processing...
time=11-12-2023 00:30:21 level=info msg=
[POINT]: Task execution [INFO]: Task [New-Item] start...
time=11-12-2023 00:30:21 level=info msg=
[POINT]: Task execution [INFO]: Task [New-Item] finished Successfully!

...

time=11-12-2023 00:30:22 level=info msg=
[POINT]: Runtime [INFO]: All tasks have been done! Saving results...
time=11-12-2023 00:30:22 level=info msg=
[POINT]: Runtime [INFO]: Process complite for 1.6539323s - find results in task_list.json!
```

По окончанию выполнения все результаты будут занесены в исходный файл task_list.json
Однако так как файл будет перезаписан - то удобнее всего просмотреть его будет в каком-нибудь ридере например jsonviewer.stack.hu

Существующий go-cmd-client тестировался на ОС Windows используя терминал Powershell как средство выполнения задач из скрипта
