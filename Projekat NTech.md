# NTech Go i Vue.js

## Postavke foldera i kreiranje projekta

<details>
<summary>Struktura foldera</summary>

```bash
NTech/
│
├── backend/               # Go server i logika
│   ├── main.go            # Glavni Go fajl
│   ├── handlers/          # HTTP handleri (rute)
│   ├── models/            # Modeli za bazu podataka
│   ├── database/          # Povezivanje sa bazom
│   └── config/            # Konfiguracija (npr. SQLite putanja)
│
├── frontend/              # Vue.js frontend aplikacija
│   ├── public/            # Public fajlovi (index.html, slike, itd.)
│   ├── src/               # Vue.js izvorni kod
│   │   ├── components/    # Vue komponente
│   │   ├── views/         # Različite stranice ili prikazi
│   │   └── App.vue        # Glavna Vue komponenta
│   └── package.json       # Node.js zavisnosti
│
└── .gitignore             # Git ignoriše fajlove
```
</details>

<details>
<summary>Skripta za kreiranje foldera</summary>

```bash
# Kreiranje foldera
mkdir -p backend/{handlers,models,database,config}
mkdir -p frontend/public
mkdir -p frontend/src/{components,views}

# Kreiranje fajlova
touch backend/main.go
touch frontend/src/App.vue
touch frontend/package.json

echo "Struktura projekta NTech je uspešno kreirana!"
```

</details>

<details>
<summary>Sadržaj .gitignore</summary>

```bash
# Go build fajlovi
/backend/NTech
/backend/NTech.exe
/backend/*.out
/backend/*.exe
/backend/*.test
/backend/*.o
/backend/*.a

# Keš fajlovi i privremeni fajlovi
/backend/tmp/
/backend/*.log
/backend/*.tmp

# SQLite baza i privremeni fajlovi
/backend/database/*.db
/backend/database/*.db-shm
/backend/database/*.db-wal

# Node.js i Vue build foldere
/frontend/node_modules/
/frontend/dist/
/frontend/.vite/

# IDE/editor konfiguracije i sistemski fajlovi
.vscode/
.idea/
*.swp
.DS_Store
Thumbs.db
```

</details>

<details>
<summary>Sadržaj .gitattributes</summary>

```bash
# Osnovno EOL podešavanje – koristi LF za sve fajlove (Unix-style)
* text=auto eol=lf

# Go fajlovi
*.go text diff=go

# Vue i JS fajlovi
*.vue text
*.js   text
*.ts   text
*.json text
*.css  text
*.html text

# Node fajlovi koji su binarni ili mašinski generisani
package-lock.json text

# SQLite fajlovi – binarni
*.db binary
*.db-shm binary
*.db-wal binary

# Images – binarni
*.png binary
*.jpg binary
*.jpeg binary
*.gif binary
*.svg text

# Ignore compiled binaries
*.exe binary
*.out binary
```

</details>
