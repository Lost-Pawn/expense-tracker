# expense-cli

A simple command-line expense tracker written in Go. Track expenses, filter and summarize them by month, and set monthly budgets with overspend warnings.

Built as part of the [roadmap.sh backend path](https://roadmap.sh/projects/expense-tracker) project list.

## Features

- Add, update, delete, and list expenses
- Filter expenses by description
- View spending summaries (overall or by month)
- Set monthly budgets and get warned when you exceed them
- Data persisted locally as JSON (`expenses.json`, `budgets.json`)

## Installation

Clone the repo and build the binary:

```bash
git clone <your-repo-url>
cd expense-tracker
go build -o expense-cli ./cmd/expense-cli
```

Or install it to your Go bin path so it's available system-wide:

```bash
go install ./cmd/expense-cli
```

Make sure `$(go env GOPATH)/bin` is on your `PATH` (see [Verifying installation](#verifying-installation) below).

## Usage

```bash
expense-cli add --description "Lunch" --amount 20
expense-cli list
expense-cli list --filter "lunch"
expense-cli delete --id 1
expense-cli update --id 1 --description "Dinner" --amount 30
expense-cli summary
expense-cli summary --month 8
expense-cli set-budget --month 8 --amount 50
expense-cli help
```

### Commands

| Command       | Flags                              | Description                                      |
|---------------|-------------------------------------|--------------------------------------------------|
| `add`         | `--description`, `--amount`         | Add a new expense                                |
| `list`        | `--filter` (optional)               | List all expenses, optionally filtered by description |
| `delete`      | `--id`                              | Delete an expense by ID                          |
| `update`      | `--id`, `--description`, `--amount` | Update an existing expense's description and/or amount |
| `summary`     | `--month` (optional)                | Show total spending, overall or for a specific month |
| `set-budget`  | `--month`, `--amount`               | Set (or update) a budget for a given month        |
| `help`        | —                                    | Show usage information                            |

## Data storage

Data is stored as JSON files (`expenses.json`, `budgets.json`) in the directory you run the command from. Each directory you run `expense-cli` in gets its own independent set of data files.

## Project structure

```
expense-tracker/
├── cmd/
│   └── expense-cli/
│       └── main.go        # CLI entrypoint, command dispatch
├── internal/
│   ├── storage/
│   │   └── storage.go     # Load/save expenses and budgets to JSON
│   └── task/
│       └── task.go        # Expense and Budget types, ID generation, budget checks
├── go.mod
└── README.md
```

## License

See [LICENSE](./LICENSE).