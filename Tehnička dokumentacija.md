# GO-NTech

GO-NTech je web aplikacija za vođenje prodaje i servisa računara. Backend je razvijen u Go jeziku, frontend koristi Vue.js, a podaci se čuvaju u lokalnoj SQLite bazi. Aplikacija je podeljena u dva dela: backend i frontend.

---

## 📁 Struktura projekta

```
Go-NTech/
├── backend/             # Backend servis (Go)
│   ├── main.go
│   ├── routes/
│   ├── handlers/
│   ├── models/
│   ├── database/
│   └── config/
├── frontend/            # Frontend aplikacija (Vue.js)
│   ├── public/          # Static root (favicon, robots.txt…)
│   ├── src/             # Vue komponente i logika
│   └── assets/          # Slike, CSS fajlovi, JS skripte
└── README.md            # Tehnička dokumentacija
```

---

## 🔧 Funkcionalnosti

* Upravljanje artiklima i komponentama
* Evidencija servisa računara i laptopova
* Praćenje ulaza/izlaza robe
* Fakturisanje
* Upravljanje dobavljačima
* Kanban zadaci za organizaciju posla

---

## 🗃 Baza podataka (SQLite)

Primer tabele `artikli`:

| ID | Naziv   | Cena | Količina | Lokacija | Status |
| -- | ------- | ---- | -------- | -------- | ------ |
| 1  | RAM 8GB | 4000 | 5        | POL-1    | 1      |

* `Status INTEGER DEFAULT 1` — aktivan artikl
* Ostale tabele: `servisi`, `klijenti`, `fakture`, `dobavljaci`, `kanban`

---

## 🔌 API rute (primer)

| Metod  | Ruta          | Opis                      |
| ------ | ------------- | ------------------------- |
| GET    | /artikli      | Lista svih artikala       |
| POST   | /artikli      | Dodavanje novog artikla   |
| PUT    | /artikli/\:id | Izmena postojećeg artikla |
| DELETE | /artikli/\:id | Brisanje artikla          |

(Sve rute su JSON REST)

---

## 🚀 Pokretanje aplikacije

### Backend

```bash
cd Go-NTech/backend
go run main.go
```

### Frontend (ako koristi Vue.js)

```bash
cd Go-NTech/frontend
npm install
npm run dev
```

Ako koristiš Go templates, frontend se servira direktno iz backend-a.

---

## ⚙ Konfiguracija

* Konfiguracija baze: `database/db.go`
* Port: 8080 (može se menjati u `.env` ili `config.go`)
* SQLite fajl: `NTech.db` u korenu projekta

---

## 🔐 Autentifikacija (opciono)

* JWT tokeni ili sesije
* Korisnici: admin, serviser, prodaja (ako bude više nivoa pristupa)

---

## 🛠 Assets

U `frontend/assets/` direktorijumu nalaze se slike, CSS fajlovi i JS skripte.

### Primer korišćenja asset-a u Vue komponenti:

```vue
<template>
  <img :src="require('@/assets/logo.png')" alt="Logo" />
</template>
```

---

## 📝 TODO / Planirano

* Logovanje aktivnosti
* PDF export faktura
* Automatski backup baze
* Verzionisanje i Gitea integracija

---

## 📄 Licenca

MIT licenca (ili navedi drugu ukoliko koristiš).

---

© Dalibor Marković – GO-NTech
